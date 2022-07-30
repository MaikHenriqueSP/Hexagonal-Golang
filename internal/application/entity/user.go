package entity

type User struct {
	email string
	fullName string
}

func NewUser(email, fullName string) (u *User) {
	return &User{email: email, fullName: fullName}
}