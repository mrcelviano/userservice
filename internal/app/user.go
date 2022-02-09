package app

type User struct {
	ID    int64  `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
	Name  string `json:"name" db:"name"`
}

type UserLogic interface {
}

type UserRepository interface {
}
