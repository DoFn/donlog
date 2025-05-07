package backend

import (
	"github.com/AfterShip/email-verifier"
	"errors"
	"regexp"
	"github.com/google/uuid"
)

// This function checks if a registration has correct details
func ValidRegistration(
	nameFirst string, 
	nameLast string,
	nameUser string,
	email string,
	password string,
	data *data,
	) (error) {
	// Check that email is valid
	if !emailverifier.IsAddressValid(email) {
		return errors.New("email Invalid")
	}

	// Check that email and username aren't being used already
	for _, val := range data.users {
		if val.email == email || val.username == nameUser {
			return errors.New("username or email in use")
		}
	}

	// Check that username and names aren't empty
	if nameUser == "" || nameFirst == "" || nameLast == "" {
		return errors.New("name fields cannot be empty")
	}

	// Check that names are alphabetic
	byteName := ([]byte)(nameFirst + nameLast)
	alphabetic, _ := regexp.Match(`^[a-zA-Z]*$`, byteName)
	if !alphabetic {
		return errors.New("names can only be alphabetic")
	}

	// Check password is at least 10 characters
	if len(password) < 10 {
		return errors.New("password must be at least 10 characters long")
	}

	// Check password has at least 1 of each necessary character type
	bytePassword := ([]byte)(password)
	noLower, _ := regexp.Match(`/^[^a-z]+$/`, bytePassword)
	noUpper, _ := regexp.Match(`/^[^A-Z]+$/`, bytePassword)
	noNumber, _ := regexp.Match(`/^[^0-9]+$/`, bytePassword)
	noSpecial, _ := regexp.Match(`/^[A-Za-z '-]+$/`, bytePassword)
	if noLower || noUpper || noNumber || noSpecial {
		return errors.New("password must have at least 1 lowercase, uppercase, number and special character")
	}

	// In case of no error, return nil
	return nil
}

// This function appends a user registration into the data struct
func DoRegistration(
	nameFirst string, 
	nameLast string,
	nameUser string,
	email string,
	password string,
	data *data,
	) (string) {
	userId := uuid.NewString()
	profile := userInfo{
		name: nameFirst + " " + nameLast, 
		email: email,
		username: nameUser,
		userId: userId,
		blogs: []blogInfo{},
	}

	data.users = append(data.users, profile);
	return userId
}