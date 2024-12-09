package users

import (
	"context"
	"errors"

	"github.com/aghyad-khlefawi/identity/pkg/servicecollection"
	"github.com/aghyad-khlefawi/identity/utils"
	passwordHelper "github.com/sethvargo/go-password/password"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func CreateUser(email string, password string, generatePassword bool) (*User, error) {
	if generatePassword {
		res, err := passwordHelper.Generate(14, 3, 2, false, true)
		if err != nil {
			return nil, err
		}
		password = res

	} else {
		if len(password) < 10 {
			return nil, ErrPasswordNotMeetingRequirement
		}
	}

	// Hash the password
	hashedPassword, err := utils.HashPasswword(password)
	if err != nil {
		return nil, err
	}

	// Create the user record
	user := User{
		Email:    email,
		Password: hashedPassword,
	}
	servicecollection.Default().MongoClient.Database("identity").Collection("users").InsertOne(context.TODO(), user)

	mailContent := "Welcome to our hr-app!"
	if generatePassword {
		mailContent += "\r\nYour password is " + password
	}
	mail := gomail.NewMessage()
	mail.SetHeader("From", "mail@hrapp.com")
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Welcome!")
	mail.SetBody("text/plain", mailContent)
	dialer:= gomail.NewDialer("mail-server",1025,"","")
	if err= dialer.DialAndSend(mail); err!=nil{
		return &user,err
	}
	return &user, nil

}

var ErrPasswordNotMeetingRequirement = errors.New("The password must be at least 10 charaters long")
