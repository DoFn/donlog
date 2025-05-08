package backend

import (
	"errors"
	"regexp"
	"github.com/google/uuid"
)

// This function checks if the details for creating a blog are valid
func ValidBlogCreate(
	userId string,
	title string,
	description string,
	data *data,
) (error) {
	// Ensure blog title length is valid
	if len(title) < 6 {
		return errors.New("blog title cannot be less than 6 characters")
	}

	// Check that title is alphanumeric
	byteName := ([]byte)(title)
	alphanumeric, _ := regexp.Match(`^[a-zA-Z0-9 ]*$`, byteName)
	if !alphanumeric {
		return errors.New("title must be alphanumeric")
	}

	// Check that blog description isn't too long
	if len(description) > 1000 {
		return errors.New("blog description cannot exceed 1000 characters")
	}

	return nil
}

// This function creates a blog, then returns its blogid
func DoBlogCreate(
	userId string,
	title string,
	description string,
	data *data,
) (string) {
	// Create blog, and append it to profile
	blog := blogInfo{
		title: title,
		description: description,
		blogId: uuid.NewString(),
		pages: []pageInfo{},
	}

	for _, profile := range data.users {
		if (profile.userId == userId) {
			profile.blogs = append(profile.blogs, blog)
		}
	}

	return blog.blogId
}

// This function checks if a blog list call is valid
func ValidBlogList(
	userId string,
	targetNameUser string,
	data *data,
) (error) {
	// Ensure userId exists
	idFound := false
	userFound := false
	for _, profile := range data.users {
		if profile.userId == userId {
			idFound = true
		}
		if profile.username == targetNameUser {
			userFound = true
		}
	}
	if !idFound {
		return errors.New("user searching is invalid")
	} 
	
	if !userFound {
		return errors.New("target user is invalid")
	}

	return nil
}

func DoBlogList(
	targetNameUser string,
	data *data,
) ([]blogDetails) {
	var profile userInfo
	for _, prof := range data.users {
		if prof.username == targetNameUser {
			profile = prof
		}
	}

	list := []blogDetails{}
	for _, blog := range profile.blogs {
		blogSummary := blogDetails{title: blog.title, description: blog.description, blogId: blog.blogId}
		list = append(list, blogSummary)
	}

	return list
}