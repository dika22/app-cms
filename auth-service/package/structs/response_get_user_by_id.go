package structs

type ResponseGetUserByID struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      int64     `json:"role"`
	CreatedAt string	`json:"created_at"`
	UpdatedAt string	`json:"updated_at"`
}

func (u User) NewGetUserByID() ResponseGetUserByID {
	return ResponseGetUserByID{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Role:      u.Role,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	
}