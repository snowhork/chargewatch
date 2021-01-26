package vals

type UserID string

func NewUserID(id string) UserID {
	return UserID(id)
}
