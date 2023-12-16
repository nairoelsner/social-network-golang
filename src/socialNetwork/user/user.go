package user

import "fmt"

type User struct {
	username       string
	password       string
	profilePicture string
	name           string
	bio            string
	mural          []map[string]string
}

func NewUser(username string, password string, name string) *User {
	user := &User{
		username:       username,
		password:       password,
		profilePicture: "url",
		name:           name,
		bio:            "",
		mural:          []map[string]string{},
	}

	return user
}

func (u *User) String() string {
	return fmt.Sprintf(u.username, u.name, u.bio)
}

func (u *User) GetInfo() map[string]interface{} {
	info := map[string]interface{}{
		"username":       u.username,
		"profilePicture": u.profilePicture,
		"name":           u.name,
		"bio":            u.bio,
		"mural":          u.mural,
	}

	return info
}

func (u *User) UpdateInfo(newInfo map[string]string) error {
	for key, value := range newInfo {
		switch key {
		case "name":
			u.name = value
		case "bio":
			u.bio = value
		case "profilePicture":
			u.profilePicture = value
		default:
			return fmt.Errorf("Invalid key: %s", key)
		}
	}

	return nil

}

func (u *User) GetAuthInfo() map[string]string {
	authInfo := map[string]string{"username": u.username, "password": u.password}
	return authInfo
}

func (u *User) GetSearchableInfo() string {
	return u.username + " " + u.name + " " + u.bio
}

func (u *User) CreatePost(username string, text string) {
	u.mural = append([]map[string]string{{"username": username, "text": text}}, u.mural...)
}
