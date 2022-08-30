package models

// DB models

import (
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

// turnos status
const (
	RESERVERD  string = "RESERVADO"
	AVAILABLE  string = "DISPONIBLE"
	TOURNAMENT string = "TORNEO"
)

// user roles
const (
	ADMIN  string = "admin"
	NORMAL string = "normal"
)

// -------------------------- USER

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Username      string `bun:"username,unique"`
	Password      string `bun:"password"`
	Role          string `bun:"role"`
	Phone         string `bun:"phone"`
	Email         string `bun:"email,unique"`
	Firstname     string `bun:"firstname"`
	Lastname      string `bun:"lastname"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %s %s>", u.ID, u.Username, u.Email, u.Password)
}

// -------------------------- TURNO

type Turno struct {
	bun.BaseModel `bun:"table:turnos,alias:t"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	Day           string    `bun:"day" json:"day"`
	Hour          int       `bun:"hour" json:"hour"`
	Field         string    `bun:"field" json:"field"`
	Status        string    `bun:"status" json:"status"`
	Fijo          bool      `bun:"fijo,default:false" json:"fijo"`
	OwnerId       int64     `bun:"owner" json:"owner"`
	Date          time.Time `bun:"date" json:"date"`
}

func (t Turno) String() string {
	return fmt.Sprintf("-------------------- \n"+
		"[Turno] id: %d | day: %s | hour: %d | field: %s \n"+
		"status: | %s | owner: %d | date: %s \n", t.ID, t.Day, t.Hour, t.Field, t.Status, t.OwnerId, t.Date)
}
