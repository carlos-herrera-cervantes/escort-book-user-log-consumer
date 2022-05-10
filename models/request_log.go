package models

import (
	"time"

	"github.com/google/uuid"
)

type RequestLog struct {
	Id        string
	UserId    string
	Component string
	Path      string
	Method    string
	Payload   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *RequestLog) SetDefaultValues() {
	r.Id = uuid.NewString()
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}
