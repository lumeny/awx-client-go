// This file contains data structures that are used by other multiple data structures.

package data

type ListGetResponse struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}
