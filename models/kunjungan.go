package models

type Kunjungans []Kunjungan

type Kunjungan struct {
	ID          int    `json:"id"`
	Akuisisi    int    `json:"akuisisi"`
	AgenJawara  int    `json:"agen_jawara"`
	AgenJuragan int    `json:"agen_juragan"`
	AgenBEP     int    `json:"agen_bep"`
	Status      string `json:"status"`
	UserId      int    `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
