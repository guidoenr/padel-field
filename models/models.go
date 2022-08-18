package models

import "time"

type User struct {
	Id        int64  `db:"ID" json:"id"`
	Username  string `db:"Username" json:"username"`
	Password  string `db:"Password" json:"password"`
	Firstname string `db:"Firstname" json:"firstname"`
	Lastname  string `db:"Lastname" json:"lastname"`
}

type Turno struct {
	Id     int64     `db:"ID" json:"id"`
	Day    string    `db:"Day" json:"day"`
	Field  string    `db:"Field" json:"field"`
	Date   time.Time `db:"Date" json:"date"`
	Status string    `db:"Status" json:"status"`
	Owner  *User     `db:"User" json:"user" pg:"rel:has-one"`
}

const (
	Reserved  string = "RESERVADO"
	Available string = "DISPONIBLE"
)
