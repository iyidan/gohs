package hyperscan_test

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/iyidan/gohs/hyperscan"
)

func TestMatch(t *testing.T) {
	Convey("Given a compatible API for regexp package", t, func() {
		Convey("When match a simple expression", func() {
			matched, err := hyperscan.MatchString("test", "abctestdef")

			So(matched, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("When match a invalid expression", func() {
			matched, err := hyperscan.MatchString(`\R`, "abctestdef")

			So(matched, ShouldBeFalse)
			So(err, ShouldNotBeNil)
		})

		Convey("When match a simple expression with io.Reader", func() {
			var buf bytes.Buffer

			buf.Write(make([]byte, 1024*1024))
			buf.WriteString("abctestdef")

			matched, err := hyperscan.MatchReader("test", &buf)

			So(matched, ShouldBeTrue)
			So(err, ShouldBeNil)
		})
	})
}
