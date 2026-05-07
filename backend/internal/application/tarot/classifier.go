package tarot

import "strings"

const (
	topicGeneral   = "general"
	topicLove      = "love"
	topicCareer    = "career"
	topicMoney     = "money"
	topicInnerWork = "inner_work"

	intentClarity  = "clarity"
	intentDecision = "decision"
	intentTimeline = "timeline"
	intentAdvice   = "advice"
)

type questionAnalyzer struct{}

func (questionAnalyzer) Analyze(question string) QuestionAnalysis {
	q := strings.TrimSpace(question)
	normalized := strings.ToLower(q)

	analysis := QuestionAnalysis{
		Topic:             topicGeneral,
		Intent:            intentClarity,
		RewrittenQuestion: rewriteQuestion(q, topicGeneral, intentClarity),
		Confidence:        0.52,
		Source:            "rules",
	}

	if hits := collectHits(normalized, []string{"yêu", "tình cảm", "người yêu", "crush", "người cũ", "nyc", "quay lại", "chia tay", "mối quan hệ", "hôn nhân", "anh ấy", "cô ấy", "bạn trai", "bạn gái", "chồng", "vợ"}); len(hits) > 0 {
		analysis.Topic = topicLove
		analysis.Signals = append(analysis.Signals, hits...)
		analysis.Confidence += 0.22
	}

	if hits := collectHits(normalized, []string{"công việc", "sự nghiệp", "việc làm", "dự án", "sếp", "đồng nghiệp", "phỏng vấn", "nghỉ việc", "chuyển việc", "lương", "job", "career"}); len(hits) > 0 {
		if analysis.Topic == topicGeneral {
			analysis.Topic = topicCareer
		}
		analysis.Signals = append(analysis.Signals, hits...)
		analysis.Confidence += 0.18
	}

	if hits := collectHits(normalized, []string{"tiền", "tài chính", "thu nhập", "kinh doanh", "đầu tư", "nợ", "mua nhà", "cổ phiếu", "crypto"}); len(hits) > 0 {
		if analysis.Topic == topicGeneral {
			analysis.Topic = topicMoney
		}
		analysis.Signals = append(analysis.Signals, hits...)
		analysis.Confidence += 0.16
	}

	if hits := collectHits(normalized, []string{"bản thân", "nội tâm", "chữa lành", "mất phương hướng", "lo âu", "căng thẳng", "tự tin", "sợ hãi"}); len(hits) > 0 {
		if analysis.Topic == topicGeneral {
			analysis.Topic = topicInnerWork
		}
		analysis.Signals = append(analysis.Signals, hits...)
		analysis.Confidence += 0.14
	}

	if hits := collectHits(normalized, []string{"có nên", "nên", "hay không", "lựa chọn", "chọn", "quyết định", "phân vân", "tiếp tục", "dừng lại", "từ bỏ"}); len(hits) > 0 {
		analysis.Intent = intentDecision
		analysis.Signals = append(analysis.Signals, hits...)
		analysis.Confidence += 0.16
	}

	if hits := collectHits(normalized, []string{"sắp tới", "tương lai", "thời gian tới", "diễn biến", "kết quả", "đi về đâu", "trở lại"}); len(hits) > 0 {
		if analysis.Intent != intentDecision {
			analysis.Intent = intentTimeline
		}
		analysis.Signals = append(analysis.Signals, hits...)
		analysis.Confidence += 0.12
	}

	if hits := collectHits(normalized, []string{"làm gì", "hành động", "bước tiếp", "cách nào", "giải quyết", "nên làm"}); len(hits) > 0 {
		if analysis.Intent == intentClarity {
			analysis.Intent = intentAdvice
		}
		analysis.Signals = append(analysis.Signals, hits...)
		analysis.Confidence += 0.12
	}

	analysis.EmotionalTone = detectTone(normalized)
	analysis.SafetyFlags = detectSafetyFlags(normalized)
	analysis.RewrittenQuestion = rewriteQuestion(q, analysis.Topic, analysis.Intent)
	if analysis.Confidence > 0.95 {
		analysis.Confidence = 0.95
	}

	return analysis
}

func collectHits(text string, keywords []string) []string {
	var hits []string
	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			hits = append(hits, keyword)
		}
	}
	return hits
}

func detectTone(text string) []string {
	var tones []string
	for _, item := range []struct {
		label    string
		keywords []string
	}{
		{"confused", []string{"rối", "không biết", "mơ hồ", "phân vân", "bối rối"}},
		{"anxious", []string{"lo", "sợ", "áp lực", "căng thẳng", "bất an"}},
		{"attached", []string{"chờ", "nhớ", "còn thương", "khó buông", "quay lại"}},
		{"blocked", []string{"bế tắc", "kẹt", "không tiến", "cản", "điểm nghẽn"}},
	} {
		if len(collectHits(text, item.keywords)) > 0 {
			tones = append(tones, item.label)
		}
	}
	return tones
}

func detectSafetyFlags(text string) []string {
	var flags []string
	checks := []struct {
		flag     string
		keywords []string
	}{
		{"medical", []string{"bệnh", "đau", "mang thai", "thuốc", "phẫu thuật", "bác sĩ", "sức khỏe"}},
		{"legal", []string{"kiện", "luật", "pháp lý", "hợp đồng", "tòa", "ly hôn"}},
		{"financial_investment", []string{"cổ phiếu", "crypto", "coin", "forex", "đầu tư mã", "all in"}},
		{"self_harm", []string{"tự tử", "tự hại", "không muốn sống", "chết đi"}},
	}
	for _, check := range checks {
		if len(collectHits(text, check.keywords)) > 0 {
			flags = append(flags, check.flag)
		}
	}
	return flags
}

func rewriteQuestion(question, topic, intent string) string {
	q := strings.TrimSpace(question)
	if q == "" {
		return "Tôi cần nhìn rõ điều gì trong tình huống hiện tại, và bước tiếp theo nên là gì?"
	}

	if intent == intentDecision {
		switch topic {
		case topicLove:
			return "Điều gì đang ảnh hưởng đến khả năng tiếp tục kết nối này, và tôi nên hành động thế nào để rõ ràng hơn?"
		case topicCareer, topicMoney:
			return "Bản chất lựa chọn này là gì, rủi ro nào cần nhìn rõ, và bước thực tế tiếp theo nên là gì?"
		default:
			return "Tôi đang đứng trước lựa chọn nào, điều gì đang cản trở sự rõ ràng, và bước tiếp theo nên là gì?"
		}
	}

	if intent == intentTimeline {
		return "Tình huống này đang đi theo xu hướng nào nếu mọi thứ tiếp tục như hiện tại?"
	}

	if intent == intentAdvice {
		return "Tôi cần hiểu điều gì trong tình huống này, và nên chọn hành động nào tiếp theo?"
	}

	return q
}
