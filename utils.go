package gutils

import "github.com/google/uuid"

func NewUid() string {
	u := uuid.New().String()
	return u
}
