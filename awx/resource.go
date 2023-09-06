// This file contains the basic implementation shared by all the resources.

package awx

import (
	"net/url"
)

type Resource struct {
	connection *Connection
	path       string
}

func (r *Resource) get(query url.Values, output interface{}) error {
	return r.connection.authenticatedGet(r.path, query, output)
}

func (r *Resource) post(query url.Values, input interface{}, output interface{}) error {
	return r.connection.authenticatedPost(r.path, query, input, output)
}

func (r *Resource) String() string {
	return r.path
}
