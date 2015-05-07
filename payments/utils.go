package payments

import (
	"github.com/calce/sgo"
)

const (
	endpoint = "payments"
)

var defaultClient *Payments = nil

type Iter struct {
	*sgo.Iter
}

func New(backend *sgo.Backend) *Payments {
	return &Payments {
		backend: backend,
	}
}

// getObject creates a new instance of Payment
// objects created with getObject() are to be used with Unmarshal() by sgo.Backend
func getObject() interface{} {
	return &sgo.Payment{}
}

// getObject creates a new instance of []Payment
// to be used with Unmarshal() by sgo.Backend
func getListObject() interface{} {
	return &[]sgo.Payment{}
}

// Convert a slice of *Payment into a slice of interface{}
// to be stored in an sgo.Iter object
func getList(l interface{}) *[]interface{} {
	
	pl := l.(*[]sgo.Payment)	
	size := len(*pl)
	il := make([]interface{}, size, cap(*pl))
	
	for i:=0; i<size; i++ {
		il[i] = &(*pl)[i]
	}
	return &il

}

// Square API List
func List(params *sgo.PaymentListParams) (*Iter, *sgo.Error) {
	return getClient().List(params)
}

// Square API Retrieve
func Retrieve(paymentId string) (*sgo.Payment, *sgo.Error) {
	return getClient().Retrieve(paymentId)
}

// getClient() returns the default Payments object
func getClient() *Payments {
	if defaultClient == nil {
		defaultBackend, err := sgo.GetDefaultBackend()
		if err != nil { return nil }
		defaultClient = New(defaultBackend)
	}
	return defaultClient
}

// takes the next object from sgo.Iter.Next() and convert into *Payment
func (i *Iter) Next() (*sgo.Payment, *sgo.Error) {
	iter, err := i.Iter.Next()
	if err != nil { return nil, err }
	return iter.(*sgo.Payment), nil
}