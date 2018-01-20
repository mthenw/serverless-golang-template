package users

type Users struct{}

func (u Users) Create() string {
	return "user has been created"
}

func (u Users) Delete() string {
	return "user has been deleted"
}
