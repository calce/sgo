package sgo

type GetObjectFunc func() interface{}
type GetListFunc func(l interface{}) *[]interface{}

type Query struct {
	Method string
	URL string
	Params string
	IsList bool
	GetObject GetObjectFunc
	GetList GetListFunc
}

// creates a copy of query with the next page's URL and empty params
func (this *Query) CreateNextPageQuery(URL string) *Query {
	p := Query{
		Method: this.Method,
		URL: URL,
		Params: "",
		IsList: this.IsList,
		GetObject: this.GetObject,
		GetList: this.GetList,
	}
	return &p
}