package settlements

import (
	"github.com/calce/sgo"
	"errors"
)

const (
	endpoint = "settlements"
)

var defaultClient *Settlements = nil

func New(backend *sgo.Backend) (*Settlements, error) {
	if backend == nil { return nil, errors.New("Backend needed") }
	return &Settlements {
		backend: backend,
	}, nil
}

// getObject creates a new instance of Settlement
// objects created with getObject() are to be used with Unmarshal() by sgo.Backend
func getObject() interface{} {
	return &sgo.Settlement{}
}

// getObject creates a new instance of []Settlement
// to be used with Unmarshal() by sgo.Backend
func getListObject() interface{} {
	return &[]sgo.Settlement{}
}

// Convert a slice of *Settlement into a slice of interface{}
// to be stored in an sgo.Iter object
func getList(l interface{}) *[]interface{} {
	
	pl := l.(*[]sgo.Settlement)	
	size := len(*pl)
	il := make([]interface{}, size, cap(*pl))
	
	for i:=0; i<size; i++ {
		il[i] = &(*pl)[i]
	}
	return &il

}

// Square API List
func List(params *sgo.SettlementListParams) (*Iter, *sgo.Error) {
	return getClient().List(params)
}

// Square API Retrieve
func Retrieve(settlementId string) (*sgo.Settlement, *sgo.Error) {
	return getClient().Retrieve(settlementId)
}

// getClient() returns the default Settlements object
func getClient() *Settlements {
	if defaultClient == nil {
		defaultBackend := sgo.GetDefaultBackend()
		defaultClient, _ = New(defaultBackend)
	}
	return defaultClient
}