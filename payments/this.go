package payments

import (
	"github.com/calce/sgo"
	"github.com/calce/sgo/url"	
)

type Payments struct {
	backend *sgo.Backend
}

func (this *Payments) List(params *sgo.PaymentListParams) (*Iter, *sgo.Error) {

	encodedParams := ""	
	if params != nil {
		encodedParams = url.ParamsEncode(params) 
	}

	var query = sgo.Query{
		Method: sgo.GET,
		URL: this.backend.GetBaseURL() + endpoint,
		Params: encodedParams,
		IsList: true,
	}
	query.GetObject = getListObject
	query.GetList = getList
	
	v, err := this.backend.Dial(&query)
	if err != nil { return nil, err }

	iter := v.(*sgo.Iter)
	return &Iter{
		iter,
	}, nil

}

func (this *Payments) Retrieve(paymentId string) (*sgo.Payment, *sgo.Error) {

	var query = sgo.Query{
		Method: sgo.GET,
		URL: this.backend.GetBaseURL() + endpoint + "/" + paymentId,
		Params: "",
		IsList: false,
	}
	query.GetObject = getObject
	
	payment, err := this.backend.Dial(&query)
	if err != nil { return nil, err }
	
	return payment.(*sgo.Payment), nil
	
}