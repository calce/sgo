package settlements

import (
	"github.com/calce/sgo"
	"github.com/calce/sgo/url"	
)

type Settlements struct {
	backend *sgo.Backend
}

func (this *Settlements) List(params *sgo.SettlementListParams) (*Iter, *sgo.Error) {

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
	
	v, err := this.backend.Call(&query)
	if err != nil { return nil, err }

	iter := v.(*sgo.Iter)
	return &Iter{
		iter,
	}, nil

}

func (this *Settlements) Retrieve(settlementId string) (*sgo.Settlement, *sgo.Error) {

	var query = sgo.Query{
		Method: sgo.GET,
		URL: this.backend.GetBaseURL() + endpoint + "/" + settlementId,
		Params: "",
		IsList: false,
	}
	query.GetObject = getObject
	
	settlement, err := this.backend.Call(&query)
	if err != nil { return nil, err }
	
	return settlement.(*sgo.Settlement), nil
	
}