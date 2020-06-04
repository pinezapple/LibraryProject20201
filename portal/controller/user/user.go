package user

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/controller"
)

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
	ID       uint64 `json:"id_user"`
	Username string `json:"username"`
}

func SelectAllUser(c echo.Context) (erro error) {
	fmt.Println("Select all user")
	return controller.ExecHandler(c, nil, selectAllUser)
}

func SelectUserByID(c echo.Context) (erro error) {
	fmt.Println("Select user by id")
	return controller.ExecHandler(c, &reqUserID{}, selectUserByID)
}

func SaveUser(c echo.Context) (erro error) {
	fmt.Println("Save user")
	return controller.ExecHandler(c, &reqUser{}, saveUser)
}

func UpdateUser(c echo.Context) (erro error) {
	fmt.Println("Update user")
	return controller.ExecHandler(c, &reqUser{}, updateUser)
}

func DeleteUser(c echo.Context) (erro error) {
	fmt.Println("Delete user")
	return controller.ExecHandler(c, &reqUserID{}, deleteUser)
}
