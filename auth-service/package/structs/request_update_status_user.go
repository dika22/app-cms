package structs

import "time"

type RequestUpdateStatusUser struct {
	ID        int64     `json:"id"`
	IsSeller  bool      `json:"is_seller"`
	Role      int64     `json:"role"`
	UpdatedAt time.Time `json:"updated_at"`
}