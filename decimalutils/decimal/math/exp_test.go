package math

import (
	"fmt"
	"testing"

	"g.newcoretech.com/mobile/gocommon/decimalutils/decimal"
)

var exp_X, _ = new(decimal.Big).SetString("123.456")

func BenchmarkExp(b *testing.B) {
	for _, prec := range benchPrecs {
		b.Run(fmt.Sprintf("%d", prec), func(b *testing.B) {
			b.ReportAllocs()
			z := decimal.WithPrecision(prec)
			for j := 0; j < b.N; j++ {
				Exp(z, exp_X)
			}
			gB = z
		})
	}
}

func TestBrokenJobs_Exp(t *testing.T) {
	for i, s := range [...]struct {
		x, r string
	}{
		{
			x: "-4.196711681127197916094391539123262189586909347963506502543424520204269305640664305277347577002702737250370072340050961484385104884969242076870376232111486905959065396493009164151561622858914473431085133053988002190096068967537402110957786645990448175717096753891271854515878515736050199594390165820346179495731423460807010421108300654567720490182E-11",
			r: "0.9999999999580328831896086402857692391751165335212832970499070407439918336957012664619593253323412937604524594000218133755918787968163830746272347910313962376243766200053425149074918367705985619799635171168309589413777382356581724053283138880135649239454288402003937458350598982662476784676965172757399371508486621844945763917842573387863934924333",
		},
	} {
		x, _ := new(decimal.Big).SetString(s.x)
		r, _ := new(decimal.Big).SetString(s.r)

		z := Exp(decimal.WithPrecision(r.Precision()), x)
		if z.Cmp(r) != 0 {
			t.Fatalf(`#%d: Exp(%s)
wanted: %s
got   : %s
`, i, x, r, z)
		}
	}
}
