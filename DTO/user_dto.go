package DTO

import "awesomeProject/model"

type UserDto struct {
	UserName  string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		UserName:  user.Name,
		Telephone: user.Phone,
	}

}
