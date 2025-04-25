package model

type User struct {
	Id          string
	Email       string
	Permissions map[string]bool
}

func (u *User) HasPermission(p string) bool {
	allowed, ok := u.Permissions[p]
	return ok && allowed
}

type UserRequest struct {
	UserID string `json:"user_id"`
}
