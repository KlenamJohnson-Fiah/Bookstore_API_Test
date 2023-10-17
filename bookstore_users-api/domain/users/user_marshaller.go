package users

import (
	"encoding/json"
)

type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// This is to Marshal a slice of User using the  Marshal method. This a method that works on the User type Users.
func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result

}

func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJSON, _ := json.Marshal(user)
	privatePresentation := new(PrivateUser)
	json.Unmarshal(userJSON, &privatePresentation)
	return privatePresentation

	/*
			If the JSON identity of the User struct is different from the  Presentation Structs like the PrivateUser and PublicUser

			----------------------------------------------------------------------------------------------------------
			User Struct   ----------------------------------------------    PrivateUser Struct
				||																		||
			ID string `json: "id"`                                            Id string `json: "user_id"`
		--------------------------------------------------------------------------------------------------------------
			we can map the fields using the example below.Eg.


					return PrivateUser{
					Id:          user.Id,
					FirstName:   user.FirstName,
					LastName:    user.LastName,
					Email:       user.Email,
					DateCreated: user.DateCreated,
					Status:      user.Status,
				}
	*/

}
