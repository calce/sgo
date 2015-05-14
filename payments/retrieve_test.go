package payments

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/calce/sgo/tests"
//	"github.com/calce/sgo"
	"testing"	
)

func testRetrieve(payments *Payments, t *testing.T) {
	
	Convey("When calling Retrieve() with a good payment id", t, func(){
		Convey("The returned payment should be in good form", func(){

			payment, err := payments.Retrieve("yes")
			So(err, ShouldBeNil)
			So(payment.Id, ShouldNotBeBlank)
			
		})
	})
	
	Convey("When calling Retrieve() with a bad payment id", t, func(){
		Convey("It should return a 404 Not Found error", func(){

			payment, err := payments.Retrieve("no")
			So(err, ShouldNotBeNil)
			So(payment, ShouldBeNil)

		})
	})

}

func TestRetrieve(t *testing.T) {

	backend := tests.GetTestBackend()	
	payments, _ := New(backend)
	testRetrieve(payments, t)
	
}