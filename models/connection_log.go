package models

import (
	"time"

	"github.com/google/uuid"
)

type ConnectionLog struct {
	Id             string
	UserId         string
	LastConnection time.Time
}

func (c *ConnectionLog) SetDefaultValues() { c.Id = uuid.NewString() }
