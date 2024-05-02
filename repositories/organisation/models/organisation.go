package models

type Organization struct {
	ID        uint   `json:"ID" bun:"id,pk,autoincrement"`
	LegalName string `json:"legalName"`
	Alias     string `json:"alias"`
	Country   string `json:"country"`
	Currency  string `json:"currency"`
	Gstreg    string `json:"gstreg"`
	Gstin     string `json:"gstin"`
	State     string `json:"state"`
	Pan       string `json:"pan"`
	UserID    uint   `json:"userID"`
	ParentID  uint   `json:"parentID"`
}

type ParentData struct {
	ParentID uint `json:"PID"`
	Children uint `json:"ChildID"`
}
