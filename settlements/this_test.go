package settlements

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/calce/sgo/tests"
	"github.com/calce/sgo"
)

var goodParams = &sgo.SettlementListParams{
	BeginTime: "",
	EndTime: "",
	Order: "DESC",
	Limit: 0,
}

func TestNew(t *testing.T) {
	
	Convey("Settlements", t, func(){		
		Convey("Given a good backend", func(){			
			Convey("Settlements object should be created properly", func(){
				backend := tests.GetTestBackend()
				settlements, _ := New(backend)
				So(settlements, ShouldNotBeNil)
			})
		})
		Convey("Given a nil backend", func(){
			Convey("creating new settlements object should raise error", func(){
				_, err := New(nil)
				So(err, ShouldNotBeNil)
			})
		})		
	})
	
}

func TestDefaults(t *testing.T) {

	Convey("Given the default backend", t, func(){
		Convey("The default settlements client should be functional", func(){

			backend := tests.GetTestBackend()
			sgo.SetDefaultBackend(backend)
		
			settlements := getClient()
			So(settlements, ShouldNotBeNil)
			
			_, err := List(goodParams)
			So(err, ShouldBeNil)
			
			_, err = Retrieve("yes")
			So(err, ShouldBeNil)
			
		})
	})
}

func TestTypeCasts(t *testing.T) {

	Convey("getObject() should return a pointer to sgo.Settlement", t, func(){
		So(getObject(), ShouldHaveSameTypeAs, &sgo.Settlement{})
	})

	Convey("getListObject() should return a pointer to []sgo.Settlement", t, func(){
		So(getListObject(), ShouldHaveSameTypeAs, &[]sgo.Settlement{})
	})
	
	Convey("Giving a list of settlements", t, func(){
		Convey("Calling getList() should return a correct list of interface{}", func(){
			pl := []sgo.Settlement{
				sgo.Settlement{},
				sgo.Settlement{},
				sgo.Settlement{},
			}
			
			list := getList(&pl)
			l := new([]interface{})
			So(list, ShouldHaveSameTypeAs, l)
			So(len(*list), ShouldEqual, 3)
		})
	})
}