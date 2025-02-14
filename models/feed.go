package models

type Feed struct {
	Friends []User `json:"friends"`
	Quote   string `json:"quote"`
}
