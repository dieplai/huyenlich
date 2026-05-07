package tarot

import (
	"fmt"
	"sort"
	"strings"
)

type selectedCard struct {
	info CardInfo
	raw  rawCard
}

func (s *Service) buildReading(question string, analysis QuestionAnalysis, spread SpreadInfo, cardIDs []string, orientations []string) (FreeResult, string, error) {
	cards := make([]selectedCard, len(cardIDs))
	for i, id := range cardIDs {
		raw, ok := s.cards[id]
		if !ok {
			return FreeResult{}, "", validationError(fmt.Sprintf("unknown tarot card: %s", id))
		}

		orientation := normalizeOrientation("")
		if len(orientations) > i {
			orientation = normalizeOrientation(orientations[i])
		}
		pos := spread.Positions[i]
		meaning := raw.UprightVI
		if orientation == "reversed" && raw.ReversedVI != "" {
			meaning = raw.ReversedVI
		}

		cards[i] = selectedCard{
			raw: raw,
			info: CardInfo{
				ID:               raw.ID,
				NameVI:           raw.NameVI,
				Position:         pos.NameVI,
				PositionID:       pos.ID,
				PositionFunction: pos.Function,
				Orientation:      orientation,
				Keywords:         raw.KeywordsVI,
				UprightVI:        raw.UprightVI,
				ReversedVI:       raw.ReversedVI,
				MeaningVI:        interpretCard(question, analysis, pos, raw, meaning, orientation),
				ImagePath:        "/" + strings.TrimPrefix(raw.Image.LocalPath, "public/"),
			},
		}
	}

	patterns := analyzePatterns(cards)
	cardInfos := make([]CardInfo, len(cards))
	for i := range cards {
		cardInfos[i] = cards[i].info
	}

	free := FreeResult{
		Question:          question,
		RewrittenQuestion: analysis.RewrittenQuestion,
		SpreadID:          spread.ID,
		SpreadName:        spread.NameVI,
		Analysis:          analysis,
		Cards:             cardInfos,
		Patterns:          patterns,
		Summary:           buildSummary(question, analysis, spread, cards, patterns),
		Synthesis:         buildSynthesis(spread, cards, patterns),
		Advice:            buildAdvice(spread, cards, analysis),
		Cautions:          buildCautions(analysis, spread),
	}

	return free, buildReadingContent(free), nil
}

func normalizeOrientation(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "reversed", "reverse", "nguoc", "ngược":
		return "reversed"
	default:
		return "upright"
	}
}

func interpretCard(question string, analysis QuestionAnalysis, pos SpreadPosition, card rawCard, baseMeaning, orientation string) string {
	orientationText := "thuận"
	if orientation == "reversed" {
		orientationText = "ngược"
	}

	intro := fmt.Sprintf("%s (%s) ở vị trí %s", card.NameVI, orientationText, strings.ToLower(pos.NameVI))
	switch pos.Function {
	case "advice":
		return fmt.Sprintf("%s không chỉ mô tả tình huống, mà gợi ý cách đáp lại: %s", intro, softenAbsolute(baseMeaning))
	case "challenge", "shadow":
		return fmt.Sprintf("%s chỉ ra điểm cần nhìn thẳng trước khi hành động: %s", intro, softenAbsolute(baseMeaning))
	case "forecast", "outcome":
		return fmt.Sprintf("%s nói về xu hướng gần nếu quỹ đạo hiện tại tiếp tục, không phải một kết luận cố định: %s", intro, softenAbsolute(baseMeaning))
	case "external":
		return fmt.Sprintf("%s nên được đọc như tín hiệu thể hiện qua tương tác và hoàn cảnh, không phải lời khẳng định chắc về suy nghĩ của người khác: %s", intro, softenAbsolute(baseMeaning))
	case "relationship_dynamic":
		return fmt.Sprintf("%s mô tả cách năng lượng giữa hai bên đang vận hành trong câu hỏi này: %s", intro, softenAbsolute(baseMeaning))
	case "resource", "opportunity":
		return fmt.Sprintf("%s làm rõ nguồn lực hoặc hướng mở có thể tận dụng: %s", intro, softenAbsolute(baseMeaning))
	case "background":
		return fmt.Sprintf("%s cho thấy một ảnh hưởng đã đặt nền cho câu chuyện hiện tại: %s", intro, softenAbsolute(baseMeaning))
	default:
		return fmt.Sprintf("%s là trọng tâm cần quan sát trong câu hỏi này: %s", intro, softenAbsolute(baseMeaning))
	}
}

func softenAbsolute(text string) string {
	replacements := map[string]string{
		"chắc chắn": "có xu hướng",
		"phải":      "nên cân nhắc",
	}
	result := text
	for old, replacement := range replacements {
		result = strings.ReplaceAll(result, old, replacement)
	}
	return result
}

func analyzePatterns(cards []selectedCard) ReadingPattern {
	pattern := ReadingPattern{
		SuitCounts: make(map[string]int),
	}

	for _, card := range cards {
		if card.raw.Arcana == "major" {
			pattern.MajorCount++
		} else {
			pattern.MinorCount++
		}
		if card.info.Orientation == "reversed" {
			pattern.ReversalCount++
		}
		if suit := stringPtrValue(card.raw.Suit); suit != "" {
			pattern.SuitCounts[suit]++
		}
	}

	pattern.DominantSuit = dominantSuit(pattern.SuitCounts)
	if pattern.MajorCount >= 2 {
		pattern.Notes = append(pattern.Notes, "Nhiều lá Major Arcana cho thấy câu hỏi chạm vào bài học lớn hoặc bước chuyển đáng chú ý, không chỉ là một chi tiết nhỏ.")
	}
	if pattern.ReversalCount >= 2 {
		pattern.Notes = append(pattern.Notes, "Nhiều lá ngược cho thấy năng lượng đang bị chặn, quay vào bên trong hoặc cần thêm thời gian để rõ.")
	}
	if pattern.DominantSuit != "" {
		pattern.Notes = append(pattern.Notes, suitMeaning(pattern.DominantSuit))
	}

	for i := 0; i < len(cards)-1; i++ {
		pattern.PairRelations = append(pattern.PairRelations, relateCards(cards[i], cards[i+1]))
	}

	return pattern
}

func dominantSuit(counts map[string]int) string {
	type suitCount struct {
		suit  string
		count int
	}
	var items []suitCount
	for suit, count := range counts {
		items = append(items, suitCount{suit: suit, count: count})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].count == items[j].count {
			return items[i].suit < items[j].suit
		}
		return items[i].count > items[j].count
	})
	if len(items) == 0 || items[0].count < 2 {
		return ""
	}
	return items[0].suit
}

func relateCards(left, right selectedCard) PairRelation {
	leftSuit := stringPtrValue(left.raw.Suit)
	rightSuit := stringPtrValue(right.raw.Suit)
	leftElement := elementForCard(left.raw)
	rightElement := elementForCard(right.raw)

	relation := "clarifies"
	message := fmt.Sprintf("%s làm rõ thêm cho %s: lá sau bổ sung lớp nghĩa mới cho lá trước.", right.raw.NameVI, left.raw.NameVI)

	switch {
	case left.raw.Arcana == "major" || right.raw.Arcana == "major":
		relation = "anchors"
		message = fmt.Sprintf("%s và %s đặt một lớp bài học lớn lên tình huống, nên đọc chúng như trục chính thay vì chi tiết rời.", left.raw.NameVI, right.raw.NameVI)
	case leftSuit != "" && leftSuit == rightSuit:
		relation = "amplifies"
		message = fmt.Sprintf("%s và %s cùng chất %s, vì vậy chủ đề này được nhấn mạnh trong bài.", left.raw.NameVI, right.raw.NameVI, suitNameVI(leftSuit))
	case elementsSupport(leftElement, rightElement):
		relation = "supports"
		message = fmt.Sprintf("%s hỗ trợ %s: hai lá tạo lực đẩy tương đối thuận cho câu chuyện.", right.raw.NameVI, left.raw.NameVI)
	case elementsChallenge(leftElement, rightElement):
		relation = "challenges"
		message = fmt.Sprintf("%s tạo lực căng với %s, cho thấy có mâu thuẫn cần điều hòa thay vì đi thẳng một chiều.", right.raw.NameVI, left.raw.NameVI)
	}

	return PairRelation{
		FromCardID: left.raw.ID,
		ToCardID:   right.raw.ID,
		Relation:   relation,
		MessageVI:  message,
	}
}

func buildSummary(question string, analysis QuestionAnalysis, spread SpreadInfo, cards []selectedCard, pattern ReadingPattern) string {
	names := make([]string, len(cards))
	for i, card := range cards {
		names[i] = fmt.Sprintf("%s: %s", card.info.Position, card.info.NameVI)
	}

	topic := topicNameVI(analysis.Topic)
	if analysis.RewrittenQuestion != "" && analysis.RewrittenQuestion != question {
		return fmt.Sprintf("Câu hỏi được đọc theo hướng %s: %s. Spread %s mở ra qua %s.", topic, analysis.RewrittenQuestion, spread.NameVI, strings.Join(names, "; "))
	}
	return fmt.Sprintf("Spread %s đang đọc câu hỏi theo hướng %s. Các lá chính là %s.", spread.NameVI, topic, strings.Join(names, "; "))
}

func buildSynthesis(spread SpreadInfo, cards []selectedCard, pattern ReadingPattern) string {
	var parts []string
	if len(cards) > 0 {
		parts = append(parts, fmt.Sprintf("Trọng tâm nằm ở %s: %s.", cards[0].info.Position, cards[0].info.NameVI))
	}
	for _, relation := range pattern.PairRelations {
		parts = append(parts, relation.MessageVI)
	}
	if len(pattern.Notes) > 0 {
		parts = append(parts, pattern.Notes...)
	}
	if len(parts) == 0 {
		return "Bài đọc này nên được xem như một chuỗi tín hiệu liên kết, không phải ba hay năm định nghĩa rời rạc."
	}
	return strings.Join(parts, " ")
}

func buildAdvice(spread SpreadInfo, cards []selectedCard, analysis QuestionAnalysis) []string {
	var advice []string
	for _, card := range cards {
		if card.info.PositionFunction == "advice" {
			advice = append(advice, fmt.Sprintf("Bắt đầu từ vị trí %s: chọn một hành động nhỏ phản ánh năng lượng của %s, thay vì chờ mọi thứ tự rõ.", strings.ToLower(card.info.Position), card.info.NameVI))
		}
	}
	if len(advice) == 0 && len(cards) > 0 {
		last := cards[len(cards)-1]
		advice = append(advice, fmt.Sprintf("Dùng lá cuối %s như tiêu chí hành động: hỏi bản thân bước nào khiến tình huống rõ hơn trong 24-72 giờ tới.", last.info.NameVI))
	}
	if analysis.Intent == intentDecision {
		advice = append(advice, "Không cần ép bài thành câu trả lời có/không; hãy dùng bài này để xác định tiêu chí quyết định và bước thử nhỏ ít rủi ro.")
	}
	if analysis.Topic == topicLove {
		advice = append(advice, "Với câu hỏi tình cảm, ưu tiên quan sát hành động và mức độ rõ ràng trong giao tiếp, không chỉ bám vào suy đoán cảm xúc của người kia.")
	}
	return uniqueStrings(advice)
}

func buildCautions(analysis QuestionAnalysis, spread SpreadInfo) []string {
	var cautions []string
	if len(analysis.SafetyFlags) > 0 {
		cautions = append(cautions, "Bài Tarot này chỉ nên dùng để soi lại cảm xúc, lựa chọn và cách phản ứng; với y tế, pháp lý hoặc đầu tư cụ thể, hãy hỏi chuyên gia phù hợp.")
	}
	cautions = append(cautions, "Không đọc lá xu hướng/kết quả như định mệnh cố định; đó là hướng có khả năng mở ra nếu năng lượng và hành động hiện tại không đổi.")
	if spread.ID == spreadLoveFive || spread.ID == spreadRelationshipThree {
		cautions = append(cautions, "Không dùng bài này để khẳng định chắc suy nghĩ riêng tư của người khác; hãy đọc nó như tín hiệu trong tương tác và điều bạn có thể chủ động làm rõ.")
	}
	return uniqueStrings(cautions)
}

func buildReadingContent(free FreeResult) string {
	var b strings.Builder
	b.WriteString("Mở bài\n")
	b.WriteString(free.Summary)
	b.WriteString("\n\nTừng lá\n")
	for _, card := range free.Cards {
		b.WriteString(fmt.Sprintf("- %s - %s: %s\n", card.Position, card.NameVI, card.MeaningVI))
	}
	b.WriteString("\nCâu chuyện chung\n")
	b.WriteString(free.Synthesis)
	if len(free.Advice) > 0 {
		b.WriteString("\n\nLời khuyên hành động\n")
		for _, advice := range free.Advice {
			b.WriteString("- ")
			b.WriteString(advice)
			b.WriteString("\n")
		}
	}
	if len(free.Cautions) > 0 {
		b.WriteString("\nĐiều không nên vội kết luận\n")
		for _, caution := range free.Cautions {
			b.WriteString("- ")
			b.WriteString(caution)
			b.WriteString("\n")
		}
	}
	return strings.TrimSpace(b.String())
}

func topicNameVI(topic string) string {
	switch topic {
	case topicLove:
		return "tình cảm / quan hệ"
	case topicCareer:
		return "công việc / sự nghiệp"
	case topicMoney:
		return "tiền bạc / nguồn lực"
	case topicInnerWork:
		return "nội tâm / phát triển cá nhân"
	default:
		return "tổng quan"
	}
}

func suitMeaning(suit string) string {
	switch suit {
	case "cups":
		return "Nhiều Cups cho thấy cảm xúc, kết nối và nhu cầu được thấu hiểu là trọng tâm."
	case "swords":
		return "Nhiều Swords cho thấy suy nghĩ, giao tiếp, sự thật và căng thẳng tinh thần là trọng tâm."
	case "wands":
		return "Nhiều Wands cho thấy động lực, ham muốn, hành động và nhịp tiến triển là trọng tâm."
	case "pentacles":
		return "Nhiều Pentacles cho thấy thực tế, tiền bạc, công việc, thân thể hoặc sự ổn định là trọng tâm."
	default:
		return ""
	}
}

func suitNameVI(suit string) string {
	switch suit {
	case "cups":
		return "Cups/cảm xúc"
	case "swords":
		return "Swords/lý trí"
	case "wands":
		return "Wands/hành động"
	case "pentacles":
		return "Pentacles/thực tế"
	default:
		return suit
	}
}

func elementForCard(card rawCard) string {
	if value := stringPtrValue(card.Element); value != "" {
		return value
	}
	switch stringPtrValue(card.Suit) {
	case "wands":
		return "fire"
	case "cups":
		return "water"
	case "swords":
		return "air"
	case "pentacles":
		return "earth"
	default:
		return ""
	}
}

func elementsSupport(left, right string) bool {
	pair := left + ":" + right
	return pair == "fire:air" || pair == "air:fire" || pair == "water:earth" || pair == "earth:water" || (left != "" && left == right)
}

func elementsChallenge(left, right string) bool {
	pair := left + ":" + right
	return pair == "fire:water" || pair == "water:fire" || pair == "air:earth" || pair == "earth:air"
}

func stringPtrValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func uniqueStrings(values []string) []string {
	seen := make(map[string]bool, len(values))
	var unique []string
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		unique = append(unique, value)
	}
	return unique
}
