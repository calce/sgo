package payments

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/calce/sgo/tests"
	"github.com/calce/sgo"
)

var goodParams = &sgo.PaymentListParams{
	BeginTime: "",
	EndTime: "",
	Order: "DESC",
	Limit: 0,
}

func TestNew(t *testing.T) {
	
	Convey("Payments", t, func(){		
		Convey("Given a good backend", func(){			
			Convey("Payments object should be created properly", func(){
				backend := tests.GetTestBackend()
				payments, _ := New(backend)
				So(payments, ShouldNotBeNil)
			})
		})
		Convey("Given a nil backend", func(){
			Convey("creating new payments object should raise error", func(){
				_, err := New(nil)
				So(err, ShouldNotBeNil)
			})
		})		
	})
	
}

func TestDefaults(t *testing.T) {

	Convey("Given the default backend", t, func(){
		Convey("The default payments client should be functional", func(){

			backend := tests.GetTestBackend()
			sgo.SetDefaultBackend(backend)
		
			payments := getClient()
			So(payments, ShouldNotBeNil)
			
			_, err := List(goodParams)
			So(err, ShouldBeNil)
			
			_, err = Retrieve("yes")
			So(err, ShouldBeNil)
			
		})
	})
}

func TestTypeCasts(t *testing.T) {

	Convey("getObject() should return a pointer to sgo.Payment", t, func(){
		So(getObject(), ShouldHaveSameTypeAs, &sgo.Payment{})
	})

	Convey("getListObject() should return a pointer to []sgo.Payment", t, func(){
		So(getListObject(), ShouldHaveSameTypeAs, &[]sgo.Payment{})
	})
	
	Convey("Giving a list of payments", t, func(){
		Convey("Calling getList() should return a correct list of interface{}", func(){
			pl := []sgo.Payment{
				sgo.Payment{},
				sgo.Payment{},
				sgo.Payment{},
			}
			
			list := getList(&pl)
			l := new([]interface{})
			So(list, ShouldHaveSameTypeAs, l)
			So(len(*list), ShouldEqual, 3)
		})
	})
}