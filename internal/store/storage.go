package store

import (
	"context"
	"database/sql"
)

type Storage struct{
	User interface{
		CreateUser(context.Context,*User)error
	}
}


func NewStorage(db *sql.DB)Storage{
	return Storage{
		
	}
}