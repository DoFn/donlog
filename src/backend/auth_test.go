package backend

import (
	"testing"
)

// UserRegistration Tests
func TestSuccessfulRegistration(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error != nil {
		t.Errorf("Successful registration failed")
	}
}

func TestFailedRegistrationEmailUsed(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}
	_, error2 := UserRegister("Donald", "Danilo", "DoFnow", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error2 == nil {
		t.Errorf("Duplicate email registration succeeded")
	}
}

func TestFailedRegistrationEmailInvalid(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error := UserRegister("Donald", "Danilo", "DoFn", "realemailguys", "P1#rfcJMN1", &data);
	if error == nil {
		t.Errorf("Invalid email registration succeeded")
	}
}

func TestFailedRegistrationUsernameUsed(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}
	_, error2 := UserRegister("Donald", "Danilo", "DoFn", "doonowyt@gmail.com", "P1#rfcJMN1", &data);
	if error2 == nil {
		t.Errorf("Duplicate username registration succeeded")
	}
}

func TestFailedRegistrationNameFirstInvalid(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error := UserRegister("Donald!!!", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error == nil {
		t.Errorf("Invalid NameFirst registration succeeded")
	}
}

func TestFailedRegistrationNameLastInvalid(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error := UserRegister("Donald", "Danilo!!!", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error == nil {
		t.Errorf("Invalid NameLast registration succeeded")
	}
}

func TestFailedRegistrationNameFirstEmpty(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error := UserRegister("", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error == nil {
		t.Errorf("Empty NameFirst registration succeeded")
	}
}

func TestFailedRegistrationNameLastEmpty(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error := UserRegister("Donald", "", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error == nil {
		t.Errorf("Empty NameLast registration succeeded")
	}
}

func TestFailedRegistrationPasswordLength(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error := UserRegister("Donald", "", "DoFn", "thedonenzo@gmail.com", "P1#a", &data);
	if error == nil {
		t.Errorf("Short password registration succeeded")
	}
}

func TestFailedRegistrationPasswordCharacters(t *testing.T) {
	data := data{users: []userInfo{}}
	_, error := UserRegister("Donald", "", "DoFn", "thedonenzo@gmail.com", "aaaa#janigbei1", &data);
	if error == nil {
		t.Errorf("Weak password registration succeeded")
	}
}

// UserDelete Tests
func TestSuccessfulRegistrationDeletion(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	error2 := UserDelete(userId, "P1#rfcJMN1");
	if error2 != nil {
		t.Errorf("Successful deletion failed")
	}
}

func TestFailedRegistrationDeletionUserId(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	error2 := UserDelete(userId + "1", "P1#rfcJMN1");
	if error2 == nil {
		t.Errorf("Deletion of invalid userId succeeded")
	}
}

func TestFailedRegistrationDeletionPassword(t *testing.T) {
	data := data{users: []userInfo{}}
	userId, error1 := UserRegister("Donald", "Danilo", "DoFn", "thedonenzo@gmail.com", "P1#rfcJMN1", &data);
	if error1 != nil {
		t.Errorf("Successful registration failed")
	}

	error2 := UserDelete(userId, "P1#rfcJMN12");
	if error2 == nil {
		t.Errorf("Deletion with invalid password succeeded")
	}
}