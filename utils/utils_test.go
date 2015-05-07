package utils

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

//var d = Convey

func TestUtils(t *testing.T) {
	var goodTimes = []string{
		"",
		"2013-01-31T00:00:00Z",
		"2014-05-05T04:11:59Z",
		"2015-01-30T01:02:03Z",
		"2014-11-31T12:42:00Z",
		"2019-12-02T06:11:59Z",
	}
	
	var badTime = "2013-13-31T00:00:ZZ"
	
	Convey("Utils", t, func(){

		Convey("Should be able to parse good times", func(){
			for _, s := range goodTimes {				
				So(ValidTime(s, true), ShouldBeTrue)
			}			
		})
		
		Convey("Should fail on a blank string when not accepting blank strings", func(){
			So(ValidTime("", false), ShouldBeFalse)
		})
		
		Convey("Should not be able to parse bad time", func(){
			So(ValidTime(badTime, true), ShouldBeFalse)
		})

	})
}