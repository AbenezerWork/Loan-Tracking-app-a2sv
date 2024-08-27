package domain

type Loan struct {
	ID        string  `json:"id"`
	Amount    float64 `json:"amount"`
	Approved  bool    `json:"approved"`
	Interest  float64 `json:"interest"`
	Duration  int     `json:"duration"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"createdAt"`
}
