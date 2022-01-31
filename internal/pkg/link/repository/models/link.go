package models

type Link struct {
	ID       int64  `db:"id"`
	ShortURL string `db:"short_url"`
	LongURL  string `db:"long_url"`
	Counter  int    `db:"counter"`
}
