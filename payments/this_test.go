package payments

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/calce/sgo/tests"
//	"github.com/calce/sgo"
)

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
//		Convey("")
	})
}