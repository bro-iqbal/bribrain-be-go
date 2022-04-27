package models

type Users []User

type User struct {
	ID                   int64  `json:"id"`
	Name                 string `json:"name"`
	InputStatusKunjungan string `json:"input_status_kunjungan"`
	InputStatusRKA       string `json:"input_status_rka"`
	Status               string `json:"status"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	DeletedAt            string `json:"deleted_at"`
}
