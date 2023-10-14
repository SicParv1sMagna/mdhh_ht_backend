package model

type Talon struct {
	ID          int  `json:"id,omitempty"`
	LegalEntity bool `json:"legalEntity"`
	UserID      int  `json:"userId"`
	BranchID    int  `json:"branchId"`
}
