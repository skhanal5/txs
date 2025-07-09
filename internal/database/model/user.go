package model

type User struct {
	Email    string `db:"email"`
	Password []byte `db:"password"`
}
