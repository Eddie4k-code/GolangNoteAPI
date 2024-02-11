package user

type User struct {
	ID       int    `json: "string"`
	Username string `json: "string"`
	Password string `json: "string"`
}

// Temporary collection of users will use database to store these.
var Users []*User = []*User{}

// Create a new User
func (user *User) CreateUser() *User {

	Users = append(Users, user)

	return user
}

// Find a Specific User based on ID
func FindUser(user User) *User {
	for _, v := range Users {
		if v.ID == user.ID {
			return v
		}
	}

	return nil
}
