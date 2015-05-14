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
	Changes float32							`param:"changes"`
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
		Changes: 4.2,
	}
	
	bad := BadParams{
		M: make(map[string]string),
	}
	
	Convey("Params", t, func(){
		
		Convey("should be able to encode good structs", func(){
			_, err := Encode(&good)
			So(err, ShouldBeNil)
		})

		Convey("should be able to encode good non-pointer structs", func(){
			_, err := Encode(good)
			So(err, ShouldBeNil)
		})
		
		Convey("should generate correct value", func(){
			s, _ := Encode(&good)
			So(s, ShouldEqual, "begin_time=2013-01-15T00%3A00%3A00Z&changes=4.2&end_time=2013-01-31T00%3A00%3A00Z&money=999.999&number=999")			
		})
		
		Convey("should return an error from structs with unsupported types", func(){
			_, err := Encode(&bad)
			So(err, ShouldNotBeNil)
		})

		Convey("should not encode non-struct", func(){
			_, err := Encode("yay")
			So(err, ShouldNotBeNil)
		})

		Convey("ParamsEncode()", func(){			
			Convey("should generate correct value", func(){
				s := ParamsEncode(&good)
				So(s, ShouldEqual, "begin_time=2013-01-15T00%3A00%3A00Z&changes=4.2&end_time=2013-01-31T00%3A00%3A00Z&money=999.999&number=999")
			})
			Convey("should return blank on bad value", func(){
				s := ParamsEncode(&bad)
				So(s, ShouldBeBlank)				
			})
		})
		
	})
	
	
}