package util

import (
	"github.com/google/uuid"
)

type UUIDGen struct {
}

func (u *UUIDGen) New() string {
	return uuid.New().String()
}
