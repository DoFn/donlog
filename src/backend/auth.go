package backend

// This function registers a user into DonLog, and returns their userId, or
// otherwise sets error to not be nil
func UserRegister(
	nameFirst string, 
	nameLast string,
	nameUser string,
	email string,
	password string,
	) (int, error) {
	return 0, nil
}

// This function deletes a registered user from DonLog, or sets error to not be
// nil if deletion fails
func UserDelete(
	userId int,
	password string,
) (error) {
	return nil
}

// This function returns the details of a registered DonLog user, or sets error
// if action could not be performed
func UserDetails(
	userId int,
) (userDetails, error) {
	details := userDetails{name: "Donald", email: "thedonenzo@gmail.com", username: "DoFn"}
	return details, nil
}

// This function logs in a DonLog user, and returns their userId, or sets error
// if action could not be performed
func UserLogin(
	userName string,
	password string,
) (int, error) {
	return 0, nil
}

// This function logs out a DonLog user, or sets error if action could not be
// performed
func UserLogout(
	userId int,
) (error) {
	return nil
}
