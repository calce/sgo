package url

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

type Params struct {
	local string
	BeginTime string						`param:"begin_time"`
	EndTime string							`param:"end_time"`
	Ninja string								`param:"_"`
	Yes string									`param:"yes"`
	Number int									`param:"number"`
	Money float64								`param:"money"`	
}

type BadParams struct {
	M map[string]string
}

func TestParams(t *testing.T) {

	good := Params{
		local: "things",
		BeginTime: "2013-01-15T00:00:00Z",
		EndTime: "2013-01-31T00:00:00Z",
		Ninja: "things",
		Number: 999,
		Money: 999.999,
	}
	
	bad := BadParams{
		M: make(map[string]string),
	}
	
	Convey("Params", t, func(){

		s, err := Encode(&good)
		Convey("Should be able to encode good structs", func(){
			So(err, ShouldBeNil)	
		})
		Convey("Should generate correct value", func(){			
			So(s, ShouldEqual, "begin_time=2013-01-15T00%3A00%3A00Z&end_time=2013-01-31T00%3A00%3A00Z&money=999.999&number=999")			
		})
		
		Convey("Should return an error from structs with unsupported types", func(){
			_, err = Encode(&bad)
			So(err, ShouldNotBeNil)
		})		

	})
	
}