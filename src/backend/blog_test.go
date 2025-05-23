package backend

import (
	"testing"
)

// BlogCreate Tests
func TestSuccessfulBlogCreation(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	_, error2 := BlogCreate(userId, "Amazing", "A cool and interesting blog", &data)
	if error2 != nil {
		t.Errorf("Successful blog creation failed")
	}
}

func TestFailedBlogCreationShortTitle(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	_, error2 := BlogCreate(userId, "Cool", "A cool and interesting blog", &data)
	if error2 == nil {
		t.Errorf("Blog creation with short title succeeded")
	}
}

func TestFailedBlogCreationInvalidTitle(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	_, error2 := BlogCreate(userId, "Cool!!!!!!", "A cool and interesting blog", &data)
	if error2 == nil {
		t.Errorf("Blog creation with invalid title characters succeeded")
	}
}

func TestFailedBlogCreationInvalidDescription(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	description := ""
	for i := 0; i < 1001; i++ {
		description = description + "a"
	}
	_, error2 := BlogCreate(userId, "Amazing", description, &data)
	if error2 == nil {
		t.Errorf("Blog creation with long description succeeded")
	}
}

// BlogList Tests
func TestSuccessfulBlogList(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	_, error2 := BlogCreate(userId, "Amazing", "A cool and interesting blog", &data)
	if error2 != nil {
		t.Errorf("Successful blog creation failed")
	}

	_, error3 := BlogList(userId, "DoFn", &data);
	if error3 != nil {
		t.Errorf("Successful blog list failed")
	}
}

func TestFailedBlogListUserId(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	_, error2 := BlogCreate(userId, "Amazing", "A cool and interesting blog", &data)
	if error2 != nil {
		t.Errorf("Successful blog creation failed")
	}

	_, error3 := BlogList(userId + "1", "DoFn", &data);
	if error3 == nil {
		t.Errorf("Blog list with invalid userId succeeded")
	}
}

func TestFailedBlogListUsername(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	_, error2 := BlogCreate(userId, "Amazing", "A cool and interesting blog", &data)
	if error2 != nil {
		t.Errorf("Successful blog creation failed")
	}

	_, error3 := BlogList(userId, "DoFnow", &data);
	if error3 == nil {
		t.Errorf("Blog list with invalid username succeeded")
	}
}
