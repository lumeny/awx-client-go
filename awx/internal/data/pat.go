// This file contains the data structures used for requesting authentication tokens.

package data

import (
	"time"
)

// Personal Access Token, user token in OAuth2
type PATPostRequest struct {
	Description string  `json:"description,omitempty"`
	Application *string `json:"application"` // Must be "null" in a PAT request
	Scope       string  `json:"scope,omitempty"`
}

type PATPostResponse struct {
	Token   string    `json:"token,omitempty"`
	Expires time.Time `json:"expires,omitempty"`
}
