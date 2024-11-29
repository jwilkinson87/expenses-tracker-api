package responses

import (
	"time"
)

type AuthenticatedUserResponse struct {
	Token      string    `json:"token"`
	ExpiryTime time.Time `json:"expiry_time"`
}
