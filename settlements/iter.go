package settlements

import "github.com/calce/sgo"

type Iter struct {
	*sgo.Iter
}

// takes the next object from sgo.Iter.Next() and convert into *Settlement
func (i *Iter) Next() (*sgo.Settlement, *sgo.Error) {
	iter, err := i.Iter.Next()
	if err != nil { return nil, err }
	return iter.(*sgo.Settlement), nil
}