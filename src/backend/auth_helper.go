package backend

import (
	"errors"
	"regexp"
	"github.com/AfterShip/email-verifier"
	"github.com/google/uuid"
	"slices"
	"golang.org/x/crypto/bcrypt"
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

	// Check password is at least 10 characters, but not over 30
	if len(password) < 10 || len(password) > 30{
		return errors.New("password must be between 10 and 30 characters inclusive")
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

// This function hashes a password using bcrypt, then returns the hash of said password. Uses
// userId uuid as additional salt in hash
func HashPassword(
	password string,
	userId string,
) (string) {
	hashText := ([]byte)(password + userId)
	hashedPassword, _ := bcrypt.GenerateFromPassword(hashText, 12);
	return (string)(hashedPassword)
}

// This function checks if a password matches a hashed password
func PasswordIsEquivalent(
	password string,
	userId string,
	hashText string,
) (bool) {
	hashBytes := ([]byte)(hashText)
	passwordBytes := ([]byte)(password + userId)
	error := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)
	
	return error == nil
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
		password: HashPassword(password, userId),
	}

	data.users = append(data.users, profile);
	return userId
}

// This function checks if a user can be deleted
func ValidDeletion(
	userId string,
	password string,
	data *data,
) (error) {
	// Check user for deletion exists, return error if not found
	profileFound := false
	var profile userInfo
	for _, prof := range data.users {
		if prof.userId == userId {
			profileFound = true
			profile = prof
		}
	}

	if !profileFound {
		return errors.New("profile with userId does not exist")
	}

	// Check that password matches. Use hashed version of password
	if !PasswordIsEquivalent(password, profile.userId, profile.password) {
		return errors.New("password does not match logged in user")
	}

	// Return nil if deletion is valid
	return nil
}

// This function deletes a user from the system
func DoDelete(
	userId string,
	data *data,
) {
	var userIndex int
	for index, val := range data.users {
		if val.userId == userId {
			userIndex = index
		}
	}

	data.users = slices.Delete(data.users, userIndex, userIndex + 1)
}

// This function checks whether a userId can return details
func ValidDetails(
	userId string,
	data *data,
) (error) {
	// Check if user exists
	for _, prof := range data.users {
		if prof.userId == userId {
			return nil
		}
	}

	// Return error since user does not exist
	return errors.New("user does not exist")
}

// This function returns the details about a user
func DoDetails(
	userId string,
	data *data,
) (userDetails) {
	var prof userInfo
	for _, profile := range data.users {
		if prof.userId == userId {
			prof = profile
		}
	}

	return userDetails{name: prof.name, email: prof.email, username: prof.username}
}

// This function checks if a login is valid
func ValidLogin(
	username string,
	password string,
	data *data,
 ) (error) {
	// Check profile exists, and return userId if password matches and it does exist
	for _, profile := range data.users {
		if profile.username == username && 
		PasswordIsEquivalent(password, profile.userId, profile.password) {
			return nil
		}
	}

	return errors.New("profile with supplied details not found")
}

// This function returns a userId that corresponds to a users details
func DoLogin(
	username string,
	data *data,
) (string) {
	var profile userInfo
	for _, prof := range data.users {
		if (prof.username == username) {
			profile = prof
		}
	}

	return profile.userId
}