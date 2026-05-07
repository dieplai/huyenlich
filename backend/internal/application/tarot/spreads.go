package tarot

const (
	spreadSituationActionOutcome = "situation_action_outcome"
	spreadPastPresentFuture      = "past_present_future"
	spreadRelationshipThree      = "relationship_three"
	spreadLoveFive               = "love_five"
	spreadCareerFive             = "career_five"
	spreadDecisionFive           = "decision_five"
	spreadFiveCardCross          = "five_card_cross"
)

func defaultSpreadRegistry() map[string]SpreadInfo {
	spreads := []SpreadInfo{
		newSpread(
			spreadSituationActionOutcome,
			"Tình Huống - Điểm Nghẽn - Hành Động",
			"line",
			position(1, "situation", "Tình huống hiện tại", "diagnostic", "Đọc như bản chất thật của tình huống, không chỉ bề mặt."),
			position(2, "block", "Điểm nghẽn cần nhìn rõ", "challenge", "Đọc như điều đang cản tiến trình, có thể là hoàn cảnh bên ngoài hoặc mô thức bên trong."),
			position(3, "next_step", "Bước nên làm tiếp theo", "advice", "Chuyển nghĩa lá thành một hành động thực tế, mềm nhưng cụ thể."),
		),
		newSpread(
			spreadPastPresentFuture,
			"Ảnh Hưởng - Hiện Tại - Xu Hướng Gần",
			"line",
			position(1, "past_influence", "Ảnh hưởng đã qua", "background", "Đọc như nguồn gốc, mô thức cũ hoặc điều đã đặt nền cho tình huống."),
			position(2, "present", "Điều đang diễn ra", "diagnostic", "Đọc như năng lượng chính đang vận hành ở hiện tại."),
			position(3, "near_trend", "Xu hướng gần", "forecast", "Đọc như xu hướng nếu quỹ đạo hiện tại tiếp tục; không diễn đạt như định mệnh."),
		),
		newSpread(
			spreadRelationshipThree,
			"Bạn - Tín Hiệu Đối Phương - Kết Nối",
			"line",
			position(1, "user_energy", "Vai trò/năng lượng của bạn", "diagnostic", "Đọc như cảm xúc, nhu cầu hoặc vai trò hiện tại của querent trong mối quan hệ."),
			position(2, "other_signal", "Tín hiệu từ phía người kia", "external", "Đọc như dấu hiệu thể hiện trong tương tác; tránh khẳng định chắc tâm trí người khác."),
			position(3, "connection_dynamic", "Động lực kết nối", "relationship_dynamic", "Đọc như điều đang kéo hai bên lại gần, đẩy ra xa hoặc làm mối quan hệ vận hành."),
		),
		newSpread(
			spreadLoveFive,
			"Tình Cảm 5 Lá",
			"relationship_cross",
			position(1, "user_energy", "Năng lượng của bạn", "diagnostic", "Đọc như cảm xúc, nhu cầu hoặc vai trò hiện tại của user trong mối quan hệ."),
			position(2, "other_signal", "Tín hiệu từ phía người kia", "external", "Đọc như năng lượng/khả năng thể hiện của người kia qua tình huống; tránh đọc chắc suy nghĩ riêng tư."),
			position(3, "connection_dynamic", "Động lực kết nối", "relationship_dynamic", "Đọc như điều đang làm mối quan hệ vận hành."),
			position(4, "challenge", "Thử thách chính", "challenge", "Đọc như điểm nghẽn, nỗi sợ, lệch pha hoặc điều cần đối diện."),
			position(5, "next_step", "Bước tiếp theo", "advice", "Đưa lời khuyên thực tế để user hành động hoặc quan sát rõ hơn."),
		),
		newSpread(
			spreadCareerFive,
			"Công Việc - Nguồn Lực - Cơ Hội",
			"line",
			position(1, "current_state", "Hiện trạng", "diagnostic", "Đọc như bức tranh thật của công việc/dự án hiện tại."),
			position(2, "resource", "Thế mạnh đang có", "resource", "Đọc như nguồn lực, kỹ năng hoặc lợi thế user đang có."),
			position(3, "block", "Điểm nghẽn", "challenge", "Đọc như rào cản, thói quen hoặc điều kiện đang làm chậm tiến trình."),
			position(4, "opportunity", "Cơ hội gần nhất", "opportunity", "Đọc như hướng mở hoặc điều nên chú ý trong thời gian tới."),
			position(5, "next_step", "Hành động ưu tiên", "advice", "Đưa hành động thực tế, tránh hứa chắc kết quả tài chính."),
		),
		newSpread(
			spreadDecisionFive,
			"Ra Quyết Định 5 Lá",
			"decision_split",
			position(1, "option_a_nature", "Bản chất lựa chọn A", "diagnostic", "Đọc như điều lựa chọn đầu tiên thật sự đại diện."),
			position(2, "option_a_trend", "Xu hướng nếu chọn A", "forecast", "Đọc như xu hướng gần nếu đi theo lựa chọn A."),
			position(3, "option_b_nature", "Bản chất lựa chọn B", "diagnostic", "Đọc như điều lựa chọn thứ hai thật sự đại diện."),
			position(4, "option_b_trend", "Xu hướng nếu chọn B", "forecast", "Đọc như xu hướng gần nếu đi theo lựa chọn B."),
			position(5, "decision_lens", "Tiêu chí để quyết định", "advice", "Không ra lệnh chọn thay user; chỉ làm rõ tiêu chí và hành động tiếp theo."),
		),
		newSpread(
			spreadFiveCardCross,
			"Trải Bài 5 Lá",
			"cross",
			position(1, "current_state", "Hiện trạng", "diagnostic", "Đọc như vấn đề thật đang nằm ở đâu trong câu hỏi của user."),
			position(2, "hidden_layer", "Điều ẩn bên dưới", "shadow", "Đọc như nỗi sợ, động lực, kỳ vọng hoặc phần user chưa nhìn rõ."),
			position(3, "external_influence", "Tác động bên ngoài", "external", "Đọc như người khác, hoàn cảnh, môi trường hoặc tín hiệu ngoài đời đang tác động vào tình huống."),
			position(4, "turning_point", "Điểm chuyển hướng", "forecast", "Đọc như điều có thể làm tình huống rẽ hướng, tốt hoặc xấu, nếu năng lượng hiện tại tiếp diễn."),
			position(5, "next_direction", "Hướng đi tiếp", "advice", "Đọc như hướng xử lý hoặc hành động gần nhất user nên cân nhắc."),
		),
	}

	registry := make(map[string]SpreadInfo, len(spreads))
	for _, spread := range spreads {
		registry[spread.ID] = spread
	}
	return registry
}

func newSpread(id, nameVI, layout string, positions ...SpreadPosition) SpreadInfo {
	return SpreadInfo{
		ID:        id,
		NameVI:    nameVI,
		CardCount: len(positions),
		Layout:    layout,
		Positions: positions,
	}
}

func position(index int, id, nameVI, function, instruction string) SpreadPosition {
	return SpreadPosition{
		Index:               index,
		ID:                  id,
		NameVI:              nameVI,
		Function:            function,
		PromptInstructionVI: instruction,
	}
}
