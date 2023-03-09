package delivery

type ClassResponse struct {
	Id           uint         `json:"id"`
	ClassName    string       `json:"class_name"`
	PicUserId    uint         `json:"pic_user_id"`
	DateStart    string       `json:"date_start"`
	DateGraduate string       `json:"date_graduate"`
	User         UserResponse `json:"user,omitempty"`
}

type UserResponse struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
