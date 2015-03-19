package rd2wgs84

import (
	"testing"
)

var parseTests = []struct {
	in  RD
	out *WGS84
}{
	{RD{163835.370083, 446830.763585}, &WGS84{52.00977421758342, 5.515894213047998}},
}

func TestConvert(t *testing.T) {
	for i, tt := range parseTests {
		wgs := Convert(tt.in.X, tt.in.Y)
		if wgs.Latitude != tt.out.Latitude || wgs.Longitude != tt.out.Longitude {
			t.Errorf("%d. Convert(%f, %f) => %+v returned, expected %+v", i, tt.in.X, tt.in.Y, wgs, tt.out)
		}
	}
}
