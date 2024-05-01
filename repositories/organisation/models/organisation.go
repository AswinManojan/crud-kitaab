package models

type Organization struct {
	ID        uint   `json:"id" bun:"id,pk,autoincrement"`
	LegalName string `json:"legalname"`
	Alias     string `json:"alias"`
	Country   string `json:"country"`
	Currency  string `json:"currency"`
	Gstreg    string `json:"gstreg"`
	Gstin     string `json:"gstin"`
	State     string `json:"state"`
	Pan       string `json:"pan"`
}
