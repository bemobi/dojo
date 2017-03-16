package codec

import (
	"testing"

	"github.com/bemobi/goconvey/convey"
	"github.com/xor-gate/bjf"
)

func TestEncode(t *testing.T) {
	convey.Convey("Dado que eu passo o numero 1", t, func() {
		number := "1"
		convey.Convey("Quando eu chamo o Encode", func() {
			got := bjf.Encode(number)
			convey.Convey("Entao eu obtenho b", func() {
				convey.So(got, convey.ShouldEqual, "b")
			})
		})
	})
}
