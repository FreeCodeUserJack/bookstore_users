package users

import "encoding/json"

type PublicUser struct {
	Id int64 `json:"id"`
	// Firstname   string `json:"first_name"`
	// Lastname    string `json:"last_name"`
	// Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`
}

func (u Users) Marshall(isPublic bool) []interface{} {
	res := make([]interface{}, len(u))

	for index, user := range u {
		res[index] = user.Marshall(isPublic)
	}

	return res
}

func (u *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          u.Id,
			DateCreated: u.DateCreated,
			Status:      u.Status,
		}
	}

	jsonUser, _ := json.Marshal(u)
	var privateUser PrivateUser
	json.Unmarshal(jsonUser, &privateUser)
	return privateUser
}