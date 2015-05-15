package settlements

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/calce/sgo/tests"
// 	"github.com/calce/sgo"	
)

func TestList(t *testing.T) {

	backend := tests.GetTestBackend()
	Convey("When calling List() with good params", t, func(){
		Convey("It should return an iterator to a list of settlements", func(){
			settlements, _ := New(backend)
			iter, err := settlements.List(goodParams)
			
			So(err, ShouldBeNil)
			So(iter, ShouldHaveSameTypeAs, &Iter{})
			
			Convey("We should be able to iterate the list", func(){
				for iter.HasNext() { 
					_, err := iter.Next()
					So(err, ShouldBeNil)
				}
			})
		})
		
	})
	
}