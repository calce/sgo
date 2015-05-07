package payments

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/calce/sgo/tests"
	"github.com/calce/sgo"
)

func inspectPayment(index int, payment *sgo.Payment) {	
	//Println("Payment", payment.Id)
}

func TestPayments(t *testing.T) {

	sgo.AutoLoadNextPage = true
	backend := tests.GetTestBackend()
	payments := New(backend)
	
	Convey("List", t, func(){
		
		payments, err := payments.List(nil)
		Convey("Should list correct payments", func(){
			So(err, ShouldBeNil)			
			var i = 0
			for payments.HasNext() {
				payment, err := payments.Next()				
				So(err, ShouldBeNil)						
				So(payment, ShouldNotBeNil)
				inspectPayment(i, payment)
				i++
			}
		})
		
	})
	
	Convey("Retrieve", t, func(){
		
		_, err := payments.Retrieve("yes")
		
		Convey("Should perfectly retrieve a payment", func(){
			So(err, ShouldBeNil)
		})		
		
	})		
}

func TestMinors(t *testing.T) {
	sgo.SetDefaultBackend(tests.GetMockBackend())

	Convey("Skips", t, func(){
		List(nil)
		List(&sgo.PaymentListParams{})
		Retrieve("")
		getObject()
		getListObject()
	})
	
	Convey("Should return default payment client", t, func(){
		So(getClient(), ShouldNotBeNil)
	})

	Convey("Should properly convert lists", t, func(){
		l := []sgo.Payment{
			sgo.Payment{}, sgo.Payment{}, sgo.Payment{},
		}
		list := getList(&l)
		So(len(*list), ShouldEqual, 3)
	})
}

func TestIterator(t *testing.T) {

	backend := tests.GetTestBackend()
	sgo.SetDefaultBackend(backend)
	list := []interface{}{
		&sgo.Payment{
			Id: "0",
			MerchantId: "me",
		},
		&sgo.Payment{
			Id: "1",
			MerchantId: "me",
		}, 
		&sgo.Payment{
			Id: "2",
			MerchantId: "me",
		},
	}

	siter := sgo.Iter{}
	siter.SetList(&list)

	iter := Iter{
		&siter,
	}
	
	Convey("Payment iterator should return Payment type", t, func(){
		for iter.HasNext() {
			payment, err := iter.Next()
			So(err, ShouldBeNil)
			So(payment.MerchantId, ShouldEqual, "me")
		}		
	})
}