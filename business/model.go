package business

import "time"

type Member struct {
	ID         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Permission int    `json:"permission" db:"permission"`
	Token      string `json:"token"`
}

type Food struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Quantity  int       `json:"quantity" db:"quantity"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
