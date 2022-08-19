package models

// DB models

import (
	"fmt"
	"time"
)

type User struct {
	Id        int64  `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	Email     string `db:"email" json:"email"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %s %s>", u.Id, u.Username, u.Email, u.Password)
}

type Turno struct {
	Id     int64     `db:"id" json:"id"`
	Day    string    `db:"day" json:"day"`
	Field  string    `db:"field" json:"field"`
	Date   time.Time `db:"date" json:"date"`
	Status string    `db:"status" json:"status"`
	Owner  *User     `db:"user" json:"user" pg:"rel:has-one"`
}

func (t Turno) String() string {
	return fmt.Sprintf("Turno<%d %s %s %s %v %s>", t.Id, t.Day, t.Field, t.Status, t.Owner, t.Date)
}

const (
	Reserved  string = "RESERVADO"
	Available string = "DISPONIBLE"
)
