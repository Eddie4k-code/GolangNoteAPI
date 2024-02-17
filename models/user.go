package user

import (
	"context"
	db "note_api/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Data Shape of a User */
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

// Temporary collection of users will use database to store these.
var Users []*User = []*User{}

/* Retrieve all Users from the Users Collection */
func GetAllUsers() ([]User, error) {
	var users []User

	userCursor, err := db.UserCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	defer userCursor.Close(context.TODO())

	for userCursor.Next(context.TODO()) {
		var user User

		err := userCursor.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	if len(users) == 0 {
		return []User{}, nil
	}

	return users, nil
}

/* Create a new User and Store it in the User Collection */
func (user *User) CreateUser() (*User, error) {
	_, err := db.UserCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return user, nil
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
