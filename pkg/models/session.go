package models

import "time"

type UserSession struct {
	ID                 string
	User               *User
	DigitalFingerPrint string
	SessionID          string
	CreatedAt          time.Time
	ExpiryTime         time.Time
}

func (s *UserSession) HasExpired() bool {
	return time.Now().After(s.ExpiryTime)
}
