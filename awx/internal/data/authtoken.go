// This file contains the data structures used for requesting authentication tokens.

package data

import (
	"time"
)

type AuthTokenPostRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthTokenPostResponse struct {
	Token   string    `json:"token,omitempty"`
	Expires time.Time `json:"expires,omitempty"`
}
