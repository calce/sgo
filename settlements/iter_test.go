package settlements

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
//	"github.com/calce/sgo/tests"
	"github.com/calce/sgo"	
)

func TestIter(t *testing.T) {

	Convey("Given a good iterator, the settlements iterator should", t, func(){
		pl := []sgo.Settlement{
			sgo.Settlement{},
			sgo.Settlement{},
			sgo.Settlement{},
		}
		p := &sgo.Settlement{}
		
		list := getList(&pl)
		siter := &sgo.Iter{}
		siter.SetList(list)
		
		iter := Iter{
			siter,
		}
		
		Convey("be iterable and have correct items", func(){
			count := 0
			for iter.HasNext() {
				settlement, err := iter.Next()
				So(err, ShouldBeNil)
				So(settlement, ShouldHaveSameTypeAs, p)
				count++
			}
			So(count, ShouldEqual, 3)
		})
	})
}