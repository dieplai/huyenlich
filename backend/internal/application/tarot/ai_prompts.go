package tarot

import (
	"encoding/json"
	"fmt"
	"strings"
)

type aiQuestionAnalysis struct {
	Topic             string   `json:"topic"`
	Intent            string   `json:"intent"`
	EmotionalTone     []string `json:"emotional_tone"`
	Signals           []string `json:"signals"`
	SafetyFlags       []string `json:"safety_flags"`
	RewrittenQuestion string   `json:"rewritten_question"`
	Confidence        float64  `json:"confidence"`
	PreferredDepth    string   `json:"preferred_depth"`
}

func buildQuestionAnalysisPrompt(question string) string {
	return fmt.Sprintf(`Bạn là tarot reader chuyên nghiệp. Phân tích câu hỏi thật nhanh để backend chọn trải bài.

Chỉ trả về JSON hợp lệ, không markdown, không giải thích.

Schema:
{
  "topic": "general|love|career|money|inner_work",
  "intent": "clarity|decision|timeline|advice",
  "emotional_tone": ["confused|anxious|attached|blocked"],
  "signals": ["cụm từ quan trọng trong câu hỏi"],
  "safety_flags": ["medical|legal|financial_investment|self_harm"],
  "rewritten_question": "câu hỏi mở, đúng đạo đức tarot, không yes/no tuyệt đối",
  "confidence": 0.0,
  "preferred_depth": "three|five"
}

Quy tắc:
- User không cần chọn chủ đề; tự hiểu như người xem bài thật.
- "có thích tôi không", "crush", "người con gái", "người đó có tình cảm không" là love + clarity.
- Nếu câu hỏi về người khác, không khẳng định chắc suy nghĩ riêng tư của họ.
- Nếu có y tế/pháp lý/đầu tư/tự hại, thêm safety flag.
- preferred_depth=five nếu câu hỏi có nhiều lớp hoặc cần so sánh; three nếu câu hỏi cần đọc nhanh.

Câu hỏi: %q`, question)
}

func parseAIQuestionAnalysis(raw string, fallback QuestionAnalysis) (QuestionAnalysis, int, error) {
	cleaned := cleanJSONText(raw)
	var parsed aiQuestionAnalysis
	if err := json.Unmarshal([]byte(cleaned), &parsed); err != nil {
		return fallback, 0, err
	}

	analysis := fallback
	if isAllowed(parsed.Topic, []string{topicGeneral, topicLove, topicCareer, topicMoney, topicInnerWork}) {
		analysis.Topic = parsed.Topic
	}
	if isAllowed(parsed.Intent, []string{intentClarity, intentDecision, intentTimeline, intentAdvice}) {
		analysis.Intent = parsed.Intent
	}
	analysis.EmotionalTone = filterAllowed(parsed.EmotionalTone, []string{"confused", "anxious", "attached", "blocked"})
	analysis.Signals = uniqueStrings(parsed.Signals)
	analysis.SafetyFlags = filterAllowed(parsed.SafetyFlags, []string{"medical", "legal", "financial_investment", "self_harm"})
	if strings.TrimSpace(parsed.RewrittenQuestion) != "" {
		analysis.RewrittenQuestion = strings.TrimSpace(parsed.RewrittenQuestion)
	}
	if parsed.Confidence > 0 {
		analysis.Confidence = parsed.Confidence
		if analysis.Confidence > 0.98 {
			analysis.Confidence = 0.98
		}
	}
	analysis.Source = "ai"

	cardCount := 0
	if parsed.PreferredDepth == "three" {
		cardCount = 3
	}
	if parsed.PreferredDepth == "five" {
		cardCount = 5
	}

	return analysis, cardCount, nil
}

func buildReadingWriterPrompt(free FreeResult, spread SpreadInfo) string {
	payload := struct {
		Question string     `json:"question"`
		Spread   SpreadInfo `json:"spread"`
		Reading  FreeResult `json:"reading_facts"`
	}{
		Question: free.Question,
		Spread:   spread,
		Reading:  free,
	}
	data, _ := json.MarshalIndent(payload, "", "  ")

	return fmt.Sprintf(`Bạn là một tarot reader chuyên nghiệp, thực tế, sắc bén và trung thành với lá bài. Viết bài luận giải tiếng Việt dựa hoàn toàn trên JSON facts bên dưới.

Mục tiêu sản phẩm:
- User phải cảm thấy reader đang đọc đúng câu hỏi của họ, không phải đọc nghĩa bài chung chung.
- Câu trả lời phải có phán đoán rõ: tốt thì nói tốt, xấu thì nói xấu, lạ/đột biến thì nói đúng sự lạ/đột biến đó.
- Không được làm mềm mọi lá bài thành lời an ủi đẹp đẽ. Lá khó như Mười Kiếm, Tòa Tháp, Quỷ, Năm Cốc, Bảy Kiếm, Mặt Trăng... phải được đọc đúng mặt khó của nó.
- Mỗi đoạn phải bám ít nhất một dữ kiện cụ thể: câu hỏi, lá bài, vị trí, orientation, pattern hoặc quan hệ giữa các lá.

Luật bắt buộc:
- Trả lời thẳng trọng tâm ngay trong 2-3 câu đầu.
- Nếu câu hỏi là "có thích tôi không", "người đó nghĩ gì", "có quay lại không": đưa ra kết luận mềm ngay, ví dụ "nghiêng về có tín hiệu", "chưa đủ dấu hiệu", hoặc "tín hiệu trái chiều"; sau đó giải thích bằng bài.
- Với tình cảm: chỉ đọc tín hiệu, động lực tương tác và hành vi có thể quan sát; không tuyên bố chắc tâm trí riêng tư của người kia.
- Không tự bịa thêm lá, không đổi spread, không nói ngược orientation.
- Không dùng văn rỗng kiểu "vũ trụ", "định mệnh đã an bài", "hãy tin vào trực giác" nếu không chỉ ra hành động cụ thể.
- Không viết Markdown trang trí: không dùng **, ###, ---, blockquote.
- Không viết câu có thể đúng với mọi trải bài. Nếu một câu không gắn với lá/vị trí/câu hỏi, hãy viết lại cụ thể hơn.
- Không né các kết luận khó nghe. Nhưng không hù dọa, không thao túng, không khiến user lệ thuộc vào bài.
- Không liệt kê sách giáo khoa nghĩa lá. Luôn nói "lá này trong vị trí này đang nói gì với câu hỏi này".
- Không để lộ ngôn ngữ nội bộ dạng khung phân tích. Tuyệt đối không dùng các tiêu đề: "Câu trả lời thẳng", "Điểm bài đánh trúng", "Điều không dễ nghe", "Bạn nên làm gì", "Dấu hiệu ngoài đời nên kiểm chứng", "Mạch bài", "Dấu hiệu cần quan sát", "Hướng đi tiếp", "Theo từng lá".

Cấu trúc hiển thị cho user:
Thông điệp chính
4-6 câu. Trả lời trực diện câu hỏi của user, nói rõ bài nghiêng về hướng nào, phần nào có cửa, phần nào chưa ổn, và mức độ chắc/chưa chắc. Phần này phải đủ rõ để user đọc xong đã nắm được kết luận chính.

Luận giải
2-4 đoạn ngắn. Nối các lá thành một câu chuyện có logic theo vị trí trải bài. Phải đọc được năng lượng riêng của từng lá quan trọng trong spread, nhưng không biến thành danh sách từng lá. Nếu bài có lá nặng, nói thẳng mặt khó bằng giọng chuyên nghiệp, không giật gân.

Lời khuyên
1-2 đoạn ngắn. Gộp hành vi nên quan sát và hướng xử lý tiếp theo. Với tình cảm, ưu tiên nhịp tương tác, mức độ chủ động của đối phương và một hành động nhỏ user có thể làm trong vài ngày tới. Không dạy đời, không ra lệnh.

Độ dài:
- 3 lá: 220-320 từ.
- 5 lá: 340-500 từ.

JSON facts:
%s`, string(data))
}

func cleanJSONText(raw string) string {
	text := strings.TrimSpace(raw)
	text = strings.TrimPrefix(text, "```json")
	text = strings.TrimPrefix(text, "```")
	text = strings.TrimSuffix(text, "```")
	return strings.TrimSpace(text)
}

func isAllowed(value string, allowed []string) bool {
	for _, item := range allowed {
		if value == item {
			return true
		}
	}
	return false
}

func filterAllowed(values, allowed []string) []string {
	var result []string
	for _, value := range values {
		if isAllowed(value, allowed) {
			result = append(result, value)
		}
	}
	return uniqueStrings(result)
}
