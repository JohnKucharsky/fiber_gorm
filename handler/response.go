package handler

import "github.com/JohnKucharsky/real_world_fiber_gorm/model"

type userResponse struct {
	User struct {
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
		Token    string  `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.UserName
	r.User.Email = u.Email
	r.User.Bio = u.Bio
	r.User.Image = u.Image
	r.User.Token = u.Password
	return r
}
