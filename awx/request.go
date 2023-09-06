// This file contains the implementation of the capabilities that are common to any kind of
// request.

package awx

import (
	"fmt"
	"net/url"
)

type Request struct {
	resource *Resource
	query    url.Values
}

func (r *Request) addFilter(name string, value interface{}) {
	if r.query == nil {
		r.query = make(url.Values)
	}
	r.query.Add(name, fmt.Sprintf("%s", value))
}

func (r *Request) get(output interface{}) error {
	return r.resource.get(r.query, output)
}

func (r *Request) post(input interface{}, output interface{}) error {
	return r.resource.post(r.query, input, output)
}
