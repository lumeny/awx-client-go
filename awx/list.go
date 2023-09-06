// This file contains the implementation of the properties shared by all lists.

package awx

type ListGetResponse struct {
	count    int
	next     string
	previous string
}

func (r *ListGetResponse) Count() int {
	return r.count
}
