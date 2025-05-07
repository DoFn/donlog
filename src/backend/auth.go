package backend

// This function registers a user into DonLog, and returns their userId, or
// otherwise sets error to not be nil
func UserRegister(
	nameFirst string, 
	nameLast string,
	nameUser string,
	email string,
	password string,
	data *data,
	) (string, error) {
	// Check that registration is valid, otherwise return an error
	error := ValidRegistration(nameFirst, nameLast, nameUser, email, password, data)
	if error != nil {
		return "", error
	}

	// Make requested registration and return id
	userId := DoRegistration(nameFirst, nameLast, nameUser, email, password, data)
	return userId, nil
}

// This function deletes a registered user from DonLog, or sets error to not be
// nil if deletion fails
func UserDelete(
	userId string,
	password string,
	data *data,
) (error) {
	error := ValidDeletion(userId, password, data)
	if error != nil {
		return error
	}
	
	DoDelete(userId, data)
	return nil
}

// This function returns the details of a registered DonLog user, or sets error
// if action could not be performed
func UserDetails(
	userId string,
	data *data,
) (userDetails, error) {
	error := ValidDetails(userId, data)
	if error != nil {
		return userDetails{name: "John", email: "Doe@gmail.com", username: "Jane"}, error
	}

	return DoDetails(userId, data), nil
}

// This function logs in a DonLog user, and returns their userId, or sets error
// if action could not be performed
func UserLogin(
	userName string,
	password string,
	data *data,
) (string, error) {
	error := ValidLogin(userName, password, data)
	if error != nil {
		return "", error
	}

	return DoLogin(userName, data), nil
}

