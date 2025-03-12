package main

import (
	"RAG1/internal/store"
	"context"
	"errors"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

type userRegisterPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,max=100,email"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (app *application) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload userRegisterPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	ctx := context.Background()

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		app.internalServerError(w, r, errors.New("failed to hash password"))
		return
	}

	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hash,
	}

	_, err = app.store.User.CreateUser(ctx, user)
	if err != nil {
		app.CustomErrorResponse(w, r, err.Error())
		return
	}
	app.jsonResponse(w, http.StatusCreated, "Generated User Sucessfully")
}

type loginPayload struct {
	Email    string `json:"email" validate:"required,max=100,email"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (app *application) LoginUser(w http.ResponseWriter, r *http.Request) {
	//get the data from request
	var payload loginPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//validate the data
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//get the user with the email id
	ctx := context.Background()
	user,err:=app.store.User.GetUserWithEmail(ctx,payload.Email)
	if err!=nil{
		app.notFoundRequest(w,r,err)
		return
	}
	//validate the password
	//compare the entered password with the db user password
	err=bcrypt.CompareHashAndPassword(user.Password,[]byte(payload.Password))
	if err!=nil{
		app.CustomErrorResponse(w,r,"Invalid email or password")
		return
	}
	//send success response
	app.jsonResponse(w,http.StatusFound,"dummytoken")
	return
}
