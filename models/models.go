package models

// DB models

import (
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Username      string `bun:"username,unique"`
	Password      string `bun:"password"`
	Phone         string `bun:"phone"`
	Email         string `bun:"email"`
	Firstname     string `bun:"firstname"`
	Lastname      string `bun:"lastname"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %s %s>", u.ID, u.Username, u.Email, u.Password)
}

type Turno struct {
	bun.BaseModel `bun:"table:turnos,alias:t"`
	ID            int64     `bun:"id,pk,autoincrement"`
	Day           string    `bun:"day"`
	Hour          int       `bun:"hour"`
	Field         string    `bun:"field"`
	Status        string    `bun:"status"`
	Fijo          bool      `bun:"fijo,default:false"`
	OwnerId       int64     `bun:"owner"`
	Date          time.Time `bun:"date"`
}

func (t Turno) String() string {
	return fmt.Sprintf("-------------------- \n"+
		"[Turno] id: %d | day: %s | hour: %d | field: %s \n"+
		"status: | %s | owner: %d | date: %s \n", t.ID, t.Day, t.Hour, t.Field, t.Status, t.OwnerId, t.Date)
}

const (
	RESERVERD string = "RESERVADO"
	AVAILABLE string = "DISPONIBLE"
)
