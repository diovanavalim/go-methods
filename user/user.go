package user

type user struct {
	name     string
	lastName string
	age      uint
	email    string
	password string
}

func CreateUser(name string, lastName string, age uint, email string, password string) user {
	user := user{
		name:     name,
		lastName: lastName,
		age:      age,
		email:    email,
		password: password,
	}

	return user
}

func (u user) GetName() string {
	return u.name
}

func (u user) GetLastName() string {
	return u.lastName
}

func (u user) GetAge() uint {
	return u.age
}

func (u user) GetEmail() string {
	return u.email
}

func (u *user) SetName(name string) {
	u.name = name
}

func (u *user) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *user) SetAge(age uint) {
	u.age = age
}

func (u *user) SetEmail(email string) {
	u.email = email
}

func (u *user) SetPassword(password string) {
	u.password = password
}
