package structs

type RequestUpdateVersioning struct {
	ID                int64  `json:"id"`
	Version           int64  `json:"version" validate:"required"`
	Status            int8   `json:"status" validate:"required"`
}