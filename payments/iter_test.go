package payments

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
//	"github.com/calce/sgo/tests"
	"github.com/calce/sgo"	
)

func TestIter(t *testing.T) {

	Convey("Given a good iterator, the payments iterator should", t, func(){
		pl := []sgo.Payment{
			sgo.Payment{},
			sgo.Payment{},
			sgo.Payment{},
		}
		p := &sgo.Payment{}
		
		list := getList(&pl)
		siter := &sgo.Iter{}
		siter.SetList(list)
		
		iter := Iter{
			siter,
		}
		
		Convey("be iterable and have correct items", func(){
			count := 0
			for iter.HasNext() {
				payment, err := iter.Next()
				So(err, ShouldBeNil)
				So(payment, ShouldHaveSameTypeAs, p)
				count++
			}
			So(count, ShouldEqual, 3)
		})
	})
}