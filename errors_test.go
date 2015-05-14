package sgo
		
import (
	"testing"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {

	Convey("Errors", t, func(){

		Convey("should be able to create parse error", func(){
			err := MakeError(ErrorParse, nil, errors.New("Error!"))
			So(err, ShouldNotBeNil)
			So(err.ErrorType, ShouldEqual, ErrorParse)
			So(err.ParseError, ShouldNotBeNil)
		})		

		Convey("should be able to create api error", func(){
			err := MakeApiError(404, "not-found",	"Not Found")
			So(err, ShouldNotBeNil)
			So(err.ErrorType, ShouldEqual, ErrorApi)
			So(err.ApiError, ShouldNotBeNil)
		})		

	})
	
}