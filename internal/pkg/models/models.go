package models

type Link struct {
	ID       int64  `json:"id"`
	ShortURL string `json:"shortURL"`
	LongURL  string `json:"longURL"`
	Counter  int    `json:"counter"`
}
