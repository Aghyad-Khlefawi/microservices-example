package users

import (
	"context"
	"errors"

	"github.com/aghyad-khlefawi/identity/pkg/servicecollection"
	"github.com/aghyad-khlefawi/identity/utils"
	passwordHelper "github.com/sethvargo/go-password/password"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
  Id primitive.ObjectID `bson:"_id"`
	Email string `bson:"email"`
	Password string `bson:"password"`
}

func CreateUser(email string, password string, generatePassword bool) (*User, error) {
	if generatePassword{
		res, err := passwordHelper.Generate(14, 3, 2, false, true)
		if err != nil {	
			return nil,err
		}
		password = res

	} else {
		if len(password) < 10 {
			return nil,ErrPasswordNotMeetingRequirement
		}
	}

	// Hash the password
	hashedPassword, err := utils.HashPasswword(password)
	if err != nil {
		return nil,err
	}

	// Create the user record
	user:= User{
		Email:    email,
		Password: hashedPassword,
	}
	servicecollection.Default().MongoClient.Database("identity").Collection("users").InsertOne(context.TODO(), user)

	return &user,nil

}

var ErrPasswordNotMeetingRequirement = errors.New("The password must be at least 10 charaters long")
