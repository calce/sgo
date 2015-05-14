// forward-only iterator base class, used for requested lists
// Iter is not threadsafe
package sgo

import (
	"errors"
//	"log"
)

type NextPageFunc func() (error)

/*type NextPageLoader interface {
	loadNextPage(iter *Iter) *Error
}*/

type Iter struct {
	backend *Backend
	query *Query
	index int
	list *[]interface{}
	currentPage int

	nextPageQuery *Query
}

func (this *Iter) HasNext() bool {
	l := len(*this.list)
	if l == 0 { return false }
	if this.index < l-1 { return true }
	if this.nextPageQuery != nil { 
		return AutoLoadNextPage
	}
	return false
}

// Return next item, load next page if there is and when AutoLoadNextPage is true
func (this *Iter) Next() (interface{}, *Error) {

	if this.index > len(*this.list)-2 {
		if (this.nextPageQuery == nil) { return nil, MakeError(ErrorParse, nil, errors.New("At last item")) }
		if err := this.nextPage(); err != nil { return nil, err }
		this.index = -1
	}
	
	this.index++
	item := (*this.list)[this.index]
	
	return item, nil
}

func (this *Iter) HasNextPage() bool {
	return this.nextPageQuery != nil
}

func (this *Iter) nextPage() *Error {
	return this.backend.loadNextPage(this)
}

func (this *Iter) GetCurrentPage() int {
	return this.currentPage
}

// For testing purpose only
func (this *Iter) _SetList(l *[]interface{}) {
	this.list = l
	this.index = -1
}
// For testing purpose only
func (this *Iter) _SetNextQuery(query *Query) {
	this.nextPageQuery = query
}