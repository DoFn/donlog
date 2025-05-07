package backend

type data struct {
	users []userInfo
}

type userInfo struct {
	name string
	email string
	username string
	userId string
	blogs []blogInfo
}

type blogInfo struct {
	title string
	description string
	blogId int
	pages []pageInfo
}

// NOTE: page text is stored in files, and thus not included in this struct
type pageInfo struct {
	title string
	pageId int
}

type userDetails struct {
	name string
	email string
	username string
}