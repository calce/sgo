package settlements

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/calce/sgo/tests"
//	"github.com/calce/sgo"
	"testing"	
)

func TestRetrieve(t *testing.T) {
	
	backend := tests.GetTestBackend()	

	Convey("When calling Retrieve() with a good settlement id", t, func(){
		Convey("The returned settlement should be in good form", func(){
		
			settlements, _ := New(backend)
			settlement, err := settlements.Retrieve("yes")
			So(err, ShouldBeNil)
			So(settlement.Id, ShouldNotBeBlank)
			
		})
	})
	
	Convey("When calling Retrieve() with a bad settlement id", t, func(){
		Convey("It should return a 404 Not Found error", func(){

			settlements, _ := New(backend)
			settlement, err := settlements.Retrieve("no")
			So(err, ShouldNotBeNil)
			So(settlement, ShouldBeNil)

		})
	})

}