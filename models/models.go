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
	Date          time.Time `bun:"date"`
	Status        string    `bun:"status"`
	OwnerId       int64     `bun:"owner"`
}

func (t Turno) String() string {
	return fmt.Sprintf("Turno<id:%d %s %s %s owner:%v | hour:%d:00 %d/%s>", t.ID, t.Day, t.Field, t.Status, t.OwnerId, t.Hour, t.Date.Day(), t.Date.Month())
}

const (
	Reserved  string = "RESERVADO"
	Available string = "DISPONIBLE"
)
