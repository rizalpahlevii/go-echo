package controller

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"go-echo/helper"
	"go-echo/model"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	DB *sql.DB
}

func NewUserController(DB *sql.DB) *UserControllerImpl {
	return &UserControllerImpl{DB: DB}
}

func (controller *UserControllerImpl) Create(e echo.Context) error {
	tx, err := controller.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := model.User{
		Name:     e.FormValue("name"),
		Email:    e.FormValue("email"),
		Password: e.FormValue("password"),
	}

	helper.PanicIfError(err)

	SQL := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(e.Request().Context(), SQL, user.Name, user.Email, user.Password)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	resource := helper.ToUserResponse(user)
	response := helper.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   resource,
	}
	return e.JSON(http.StatusOK, response)
}

func (controller *UserControllerImpl) Update(e echo.Context) error {
	tx, err := controller.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := model.User{
		Name:     e.FormValue("name"),
		Email:    e.FormValue("email"),
		Password: e.FormValue("password"),
	}

	SQL := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err = tx.ExecContext(e.Request().Context(), SQL, user.Name, user.Email, user.Password, e.Param("userId"))
	helper.PanicIfError(err)

	user.Id, err = strconv.Atoi(e.Param("userId"))

	response := helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	}
	return e.JSON(http.StatusOK, response)

}

func (controller *UserControllerImpl) Delete(e echo.Context) error {
	tx, err := controller.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "DELETE FROM users WHERE id = ?"
	_, err = tx.ExecContext(e.Request().Context(), SQL, e.Param("userId"))
	helper.PanicIfError(err)
	response := helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	return e.JSON(http.StatusOK, response)
}

func (controller *UserControllerImpl) FindById(e echo.Context) error {
	tx, err := controller.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name, email, password FROM users WHERE id = ?"
	rows, err := tx.QueryContext(e.Request().Context(), SQL, e.Param("userId"))
	helper.PanicIfError(err)
	defer rows.Close()

	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		helper.PanicIfError(err)
		resources := helper.ToUserResponse(user)
		response := helper.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   resources,
		}
		return e.JSON(http.StatusOK, response)
	} else {
		response := helper.WebResponse{
			Code:   404,
			Status: "Not Found",
			Data:   nil,
		}
		return e.JSON(http.StatusNotFound, response)
	}

}

func (controller *UserControllerImpl) FindAll(e echo.Context) error {
	tx, err := controller.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name, email, password FROM users"
	rows, err := tx.QueryContext(e.Request().Context(), SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	resources := helper.ToUserResponses(users)
	response := helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   resources,
	}
	return e.JSON(http.StatusOK, response)
}
