package sgo

import (
	json "github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"net/http"
	"regexp"
	"bytes"
	"time"
//	"log"
	"io"
)

var linkParser = regexp.MustCompile(`http[^>]+`)

type Backend struct {
	token string
	baseURL string
	merchantId string
	client *http.Client
}

func NewBackend(token string, baseURL string, merchantId string, timeout time.Duration, client *http.Client) (*Backend, error) {
	backend := &Backend{
		token: token,
		merchantId: merchantId,
		baseURL: baseURL + "/" + merchantId + "/",		
	}
	
	if client == nil { client = &http.Client{Timeout: timeout} }
	client.Timeout = timeout
	backend.client = client
	
	return backend, nil
}

func (this *Backend) makeRequest(query *Query) (*http.Request, error) {

	var body io.Reader
	var url string = query.URL

	if query.Params != "" {
		if query.Method == GET {
			url += "?" + query.Params
		} else {
			body = bytes.NewBufferString(query.Params)
		}
	}
	
	req, err := http.NewRequest(query.Method, url, body)
	if err != nil { return nil, err }

	req.Header.Add("Authorization", "Bearer " + this.token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

// Parses api errors if there is, otherwise parse results
func (this *Backend) parseResponse(res *http.Response, query *Query) (interface{}, *Error) {

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil { return nil, MakeError(ErrorParse, nil, err) }
	
	if res.StatusCode != http.StatusOK { 
		var apiErr ApiError
		
		err = json.Unmarshal(b, &apiErr)
		if err != nil { return nil, MakeError(ErrorParse, nil, err) }
		
		// api error
		apiErr.ErrorCode = res.StatusCode
		return nil, MakeError(ErrorApi, &apiErr, nil)
	}
	
	v := query.GetObject()
	
	err = json.Unmarshal(b, v)
	if (err != nil) { return nil, MakeError(ErrorParse, nil, err) }
	
	return v, nil
}


// Creates an iterator object based on the returned list type
func (this *Backend) createIterator(l interface{}, res *http.Response, query*Query) (interface{}) {

	list := query.GetList(l)
	var nextPageQuery *Query = nil
	nextPageURL := this.getNextPageURL(res)
	if nextPageURL != "" {
		nextPageQuery = query.CreateNextPageQuery(nextPageURL)
	}

	iter := Iter{
		nextPageQuery: nextPageQuery,
		query: query,
		backend: this,
		index: -1,
		currentPage: 1,
		list: list,
	}
	
	return &iter
}

func (this *Backend) Dial(query *Query) (interface{}, *Error) {
	
	req, e := this.makeRequest(query);
	if e != nil { return nil, MakeError(ErrorParse, nil, e) }
	
	res, e := this.client.Do(req);
	if e != nil { return nil, MakeError(ErrorParse, nil, e) }
		
	v, err :=  this.parseResponse(res, query)
	if err != nil { return nil, err }
	
	if query.IsList {
		return this.createIterator(v, res, query), nil
	}
	return v, nil
}

// Updates an iterator with the next page's data
func (this *Backend) dialNextPage(iter *Iter) (*Error) {

	req, e := this.makeRequest(iter.nextPageQuery);
	if e != nil { MakeError(ErrorParse, nil, e) }
	
	res, e := this.client.Do(req);
	if e != nil { MakeError(ErrorParse, nil, e) }
		
	v, err :=  this.parseResponse(res, iter.nextPageQuery)
	if err != nil { return err }

	nextPageURL := this.getNextPageURL(res)
	if nextPageURL != "" {
		iter.nextPageQuery = iter.nextPageQuery.CreateNextPageQuery(nextPageURL)
	} else { 
		iter.nextPageQuery = nil
	}
	iter.SetList(iter.query.GetList(v))
	iter.index = -1
	iter.currentPage++
	
	return nil
}

// return true if res's header has the next page's url
func (this *Backend) getNextPageURL(res *http.Response) string {
	r := linkParser.FindStringSubmatch(res.Header.Get("Link"))
	if len(r) > 0 {
		return r[0]
	}
	return ""
}

func (this *Backend) GetBaseURL() string {
	return this.baseURL
}