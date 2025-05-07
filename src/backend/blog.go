package backend

// This function creates a blog with the given details, then returns the
// blogid, or sets an error value
func BlogCreate(
	userId string,
	title string,
	description string,
	data *data,
) (string, error) {
	return "0", nil
}

// This function returns the details of a user's blogs
func BlogList(
	userId string,
	targetNameUser string,
	data *data,
) ([]blogDetails, error) {
	return []blogDetails{}, nil
}

// This function deletes a blog with the given details, or returns an
// error if applicable
func BlogDelete(
	userId string,
	blogId string,
	data *data,
) (error) {
	return nil
}

// This function creates a new blog page with the given details, and returns
// the new pageid, or returns an error if applicable
func BlogPageCreate(
	userId string,
	blogId string,
	title string,
	description string,
	data *data,
) (string, error) {
	return "0", nil
}

// This function returns information on a blogs pages
func BlogPageList(
	userId string,
	blogId string,
	targetNameUser string,
) ([]pageInfo, error) {
	return []pageInfo{}, nil
}

// This function deletes a blog page with the given details, or returns
// an error if failed
func BlogPageDelete(
	userId string,
	blogId string,
	pageId string,
	data *data,
) (error) {
	return nil
}

// This function returns information on a blog page
func BlogPageDetails(
	userId string,
	blogId string,
	pageId string,
	targetNameUser string,
) ([]pageDetails, error) {
	return []pageDetails{}, nil
}