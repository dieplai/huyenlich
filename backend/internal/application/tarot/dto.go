package tarot

type PrepareRequest struct {
	Question           string `json:"question" binding:"required"`
	PreferredCardCount int    `json:"preferred_card_count,omitempty"`
}

type PrepareResponse struct {
	Analysis QuestionAnalysis `json:"analysis"`
	Spread   SpreadInfo       `json:"spread"`
}

type ReadRequest struct {
	Question     string            `json:"question" binding:"required"`
	SpreadID     string            `json:"spread_id"`
	CardIDs      []string          `json:"card_ids" binding:"required,min=1,max=5"`
	Orientations []string          `json:"orientations,omitempty"`
	Analysis     *QuestionAnalysis `json:"analysis,omitempty"`
}

type QuestionAnalysis struct {
	Topic             string   `json:"topic"`
	Intent            string   `json:"intent"`
	EmotionalTone     []string `json:"emotional_tone,omitempty"`
	Signals           []string `json:"signals,omitempty"`
	SafetyFlags       []string `json:"safety_flags,omitempty"`
	RewrittenQuestion string   `json:"rewritten_question"`
	Confidence        float64  `json:"confidence"`
	Source            string   `json:"source"`
}

type SpreadPosition struct {
	Index               int    `json:"index"`
	ID                  string `json:"id"`
	NameVI              string `json:"name_vi"`
	Function            string `json:"function"`
	PromptInstructionVI string `json:"prompt_instruction_vi"`
}

type SpreadInfo struct {
	ID        string           `json:"id"`
	NameVI    string           `json:"name_vi"`
	CardCount int              `json:"card_count"`
	Layout    string           `json:"layout"`
	Positions []SpreadPosition `json:"positions"`
}

type CardInfo struct {
	ID               string   `json:"id"`
	NameVI           string   `json:"name_vi"`
	Position         string   `json:"position"`
	PositionID       string   `json:"position_id"`
	PositionFunction string   `json:"position_function"`
	Orientation      string   `json:"orientation"`
	Keywords         []string `json:"keywords_vi"`
	UprightVI        string   `json:"upright_vi,omitempty"`
	ReversedVI       string   `json:"reversed_vi,omitempty"`
	MeaningVI        string   `json:"meaning_vi"`
	ImagePath        string   `json:"image_path,omitempty"`
}

type PairRelation struct {
	FromCardID string `json:"from_card_id"`
	ToCardID   string `json:"to_card_id"`
	Relation   string `json:"relation"`
	MessageVI  string `json:"message_vi"`
}

type ReadingPattern struct {
	MajorCount    int            `json:"major_count"`
	MinorCount    int            `json:"minor_count"`
	ReversalCount int            `json:"reversal_count"`
	DominantSuit  string         `json:"dominant_suit,omitempty"`
	SuitCounts    map[string]int `json:"suit_counts,omitempty"`
	Notes         []string       `json:"notes,omitempty"`
	PairRelations []PairRelation `json:"pair_relations,omitempty"`
}

type FreeResult struct {
	Question          string           `json:"question"`
	RewrittenQuestion string           `json:"rewritten_question,omitempty"`
	SpreadID          string           `json:"spread_id"`
	SpreadName        string           `json:"spread_name"`
	Analysis          QuestionAnalysis `json:"analysis"`
	Cards             []CardInfo       `json:"cards"`
	Patterns          ReadingPattern   `json:"patterns"`
	Summary           string           `json:"summary"`
	Synthesis         string           `json:"synthesis"`
	Advice            []string         `json:"advice,omitempty"`
	Cautions          []string         `json:"cautions,omitempty"`
}

type ReadResponse struct {
	ReadingID  string     `json:"reading_id"`
	Spread     SpreadInfo `json:"spread"`
	Free       FreeResult `json:"free"`
	AIContent  *string    `json:"ai_content,omitempty"`
	AIStatus   string     `json:"ai_status"`
	IsUnlocked bool       `json:"is_unlocked"`
}

type DeckCard struct {
	ID         string   `json:"id"`
	NameVI     string   `json:"name_vi"`
	Arcana     string   `json:"arcana,omitempty"`
	Suit       string   `json:"suit,omitempty"`
	KeywordsVI []string `json:"keywords_vi,omitempty"`
	UprightVI  string   `json:"upright_vi,omitempty"`
	ReversedVI string   `json:"reversed_vi,omitempty"`
	ImagePath  string   `json:"image_path,omitempty"`
}

type DeckResponse struct {
	Cards   []DeckCard   `json:"cards"`
	Spreads []SpreadInfo `json:"spreads"`
}
