package payments

import (
	"github.com/calce/sgo"
	"errors"
)

const (
	endpoint = "payments"
)

var defaultClient *Payments = nil

func New(backend *sgo.Backend) (*Payments, error) {
	if backend == nil { return nil, errors.New("Backend needed") }
	return &Payments {
		backend: backend,
	}, nil
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
		defaultClient, _ = New(defaultBackend)
	}
	return defaultClient
}