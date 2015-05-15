package payments

import "github.com/calce/sgo"

type Iter struct {
	*sgo.Iter
}

// takes the next object from sgo.Iter.Next() and convert into *Payment
func (i *Iter) Next() (*sgo.Payment, *sgo.Error) {
	iter, err := i.Iter.Next()
	if err != nil { return nil, err }
	return iter.(*sgo.Payment), nil
}