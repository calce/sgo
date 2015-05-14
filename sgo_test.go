package sgo

import (
	"time"
	"testing"
	"net/http"
	"github.com/calce/sgo/mocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBackend(t *testing.T) {

	Convey("General", t, func(){

		// live client
		Convey("should be able to create http client when not supplied", func(){			
			liveend, e := NewBackend("", "", "", 0, nil)
			So(e, ShouldBeNil)
			So(liveend.client, ShouldNotBeNil)
		})

		mock := mocks.New("mocks")
		backend, e := NewBackend("yay", "http://connect.squareup.com/v1", "me", time.Second * 30, mock.Client)

		Convey("should create new backend properly", func(){
			So(e, ShouldBeNil)

			Convey("baseURL should not be blank", func(){
				So(backend.GetBaseURL(), ShouldNotBeBlank)
			})
	
			query := Query{
				Method: GET,
				URL: backend.GetBaseURL() + "payments",
				Params: "?",
				IsList: true,
				GetObject: getListObject,
				GetList: getList,
			}
			
			req, e := backend.makeRequest(&query);
			Convey("should create proper request", func(){
				So(e, ShouldBeNil)
				So(req, ShouldNotBeNil)
	
				res, e := backend.client.Do(req);
				Convey("should return a good response after calling an endpoint", func(){
					So(e, ShouldBeNil)
					So(res, ShouldNotBeNil)
		
					list, err := backend.parseResponse(res, &query)
					Convey("should be able to parse a good response", func(){
						ESo(err)
						So(list, ShouldHaveSameTypeAs, &[]Payment{})
			
						var iter *Iter = backend.createIterator(list, res, &query)
						Convey("should be able to create iterators out of requested lists", func(){
							So(iter, ShouldNotBeNil)
							So(iter.backend, ShouldEqual, backend)
							
						})
					})
				})	
			})
		})		
	})
}

func TestCall(t *testing.T) {

	mock := mocks.New("mocks")
	backend, _ := NewBackend("yay", "http://connect.squareup.com/v1", "me", time.Second * 30, mock.Client)
		
	Convey("Calling", t, func(){


		Convey("Listing", func(){
			query := Query{
				Method: GET,
				URL: backend.GetBaseURL() + "payments",
				Params: "?",
				IsList: true,
				GetObject: getListObject,
				GetList: getList,
			}
			
			v, err := backend.Call(&query)
			Convey("should be able to call properly", func(){
				So(err, ShouldBeNil)
		
				iter := v.(*Iter)
				err = backend.loadNextPage(iter)
				Convey("should be able to load an iterator's next page", func(){
					So(err, ShouldBeNil)
					So(iter.list, ShouldNotBeNil)
					So(iter.currentPage, ShouldEqual, 2)
				})
			})	
		})
		
		Convey("Retrieving", func(){
			
			Convey("should retrieve existing object", func(){
				
				query := Query{
					Method: GET,
					URL: backend.GetBaseURL() + "payments/yes",
					Params: "?",
					IsList: false,
					GetObject: getObject,
					GetList: getList,
				}
				
				_, err := backend.Call(&query)
				So(err, ShouldBeNil)
			})

			Convey("should return 404 on non-existing object", func(){
				
				query := Query{
					Method: GET,
					URL: backend.GetBaseURL() + "payments/YAY",
					Params: "?",
					IsList: false,
					GetObject: getObject,
					GetList: getList,
				}
				
				_, err := backend.Call(&query)
				So(err, ShouldNotBeNil)
				So(err.ApiError.ErrorCode, ShouldEqual, http.StatusNotFound)
			})
			
			
		})

	})	
}

func TestDefaultBackend(t *testing.T) {
	Convey("Default backend", t, func(){

		Convey("should be able to get/set default backend", func(){
			backend, _ := GetDefaultBackend()
			So(backend, ShouldNotBeNil)
			newBackend := Backend{}
			SetDefaultBackend(&newBackend)
			newDefaultBackend, _ := GetDefaultBackend()
			So(newDefaultBackend, ShouldPointTo, &newBackend)
		})
		
	})
}

func TestQuery(t *testing.T) {
	query := Query{}
	Convey("Query", t, func(){
		So(query.CreateNextPageQuery(""), ShouldNotBeNil)
	})	
}

func getObject() interface{} {
	return &Payment{}
}

func getListObject() interface{} {
	return &[]Payment{}
}

func getList(l interface{}) *[]interface{} {
	
	pl := l.(*[]Payment)	
	size := len(*pl)
	il := make([]interface{}, size, cap(*pl))
	
	for i:=0; i<size; i++ {
		il[i] = &(*pl)[i]
	}
	return &il

}

func ESo(err *Error) {
	if err != nil {
		if err.ParseError != nil { Print("( Parse: ", err) }
		if err.ApiError != nil { Print("(", err.ApiError.ErrorCode, " - ", err.ApiError.Message, ") ") }
	}
	So(err, ShouldBeNil)
}