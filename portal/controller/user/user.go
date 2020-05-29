package user

import "github.com/labstack/echo"

type reqUser struct {
	ID          uint64 `json:"id_user" db:"id_user"`
	Username    string `json:"username" db:"username"`
	Password    string `json:"password"`
	Name        string `json:"name" db:"name"`
	Role        string `json:"role" db:"role"`
	Dob         string `json:"dob" db:"dob"`
	Sex         string `json:"sex" db:"sex"`
	PhoneNumber string `json:"phonenumber" db:"phonenumber"`
	Status      byte   `json:"status" db:"status"`
}

type reqUserID struct {
	ID uint64 `json:"id"`
}

func SelectAllUser(c echo.Context) {
	return
}

func SelectUserByID(c echo.Context) {
	return
}

func SaveUser(c echo.Context) {
	return
}

func UpdateUser(c echo.Context) {
	return
}

func DeleteUser(c echo.Context) {
	return
}
