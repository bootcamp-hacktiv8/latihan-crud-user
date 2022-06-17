package entity

import "time"

type User struct {
	ID         int
	Username   string
	Email      string
	Password   string
	Age        int
	Created_at time.Time
	Update_at  time.Time
}
