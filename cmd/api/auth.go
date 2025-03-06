package main

import "net/http"


type UserRegisterPayload struct{
	Username string `json:"username" validate:"required,max=100"`
	Email string `json:"email" validate:"required,max=100,email"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}


func (app *application)RegisterUser(w http.ResponseWriter,r *http.Request){
	var payload UserRegisterPayload
	if err:= readJSON(w,r,&payload);err!=nil{
		app.badRequestResponse(w,r,err)
		return
	}

	if err:=Validate.Struct(payload); err!=nil{
		app.badRequestResponse(w,r,err)
		return
	}

	
	
}