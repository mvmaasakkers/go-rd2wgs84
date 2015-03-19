package rd2wgs84

import (
	"math"
)

// RD struct defines the X and Y coordinates
type RD struct {
	X float64
	Y float64
}

// WGS84 struct defines the Latitude and Longitude coordinates
type WGS84 struct {
	Latitude  float64
	Longitude float64
}

const (
	// Base RD coordinates Amersfoort
	RD_X_BASE int = 155000
	RD_Y_BASE int = 463000

	// Same base, but as wgs84 coordinates
	WGS84_LATITUDE_BASE  float64 = 52.15517440
	WGS84_LONGITUDE_BASE float64 = 5.38720621
)

type K map[int]LI
type LI map[int]float64

var Kpq K = K{
	0: LI{1: 3235.65389, 2: -0.2475, 3: -0.0655},
	1: LI{0: -0.00738, 1: -0.00012},
	2: LI{0: -32.58297, 1: -0.84978, 2: -0.01709, 3: -0.00039},
	4: LI{0: 0.0053, 1: 0.00033},
}

var Lpq K = K{
	0: LI{1: 0.01199, 2: 0.00022},
	1: LI{0: 5260.52916, 1: 105.94684, 2: 2.45656, 3: 0.05594, 4: 0.00128},
	2: LI{0: -0.00022},
	3: LI{0: -0.81885, 1: -0.05607, 2: -0.00256},
	5: LI{0: 0.00026},
}

// Convert converts the dutch RD coordinates to WGS84
func Convert(x, y float64) *WGS84 {
	wgs := &WGS84{}

	calc_lat := float64(0)
	calc_long := float64(0)

	d_latitude := (x - float64(RD_X_BASE)) * 0.00001
	d_longitude := (y - float64(RD_Y_BASE)) * 0.00001

	pmax := 5
	qmax := 4

	for p := 0; p <= pmax; p++ {
		for q := 0; q <= qmax; q++ {
			if _, ok := Kpq[p]; ok {
				if val2, ok2 := Kpq[p][q]; ok2 {
					calc_lat += val2 * math.Pow(d_latitude, float64(p)) * math.Pow(d_longitude, float64(q))
				}
			}

			if _, ok := Lpq[p]; ok {
				if val2, ok2 := Lpq[p][q]; ok2 {
					calc_long += val2 * math.Pow(d_latitude, float64(p)) * math.Pow(d_longitude, float64(q))
				}
			}
		}
	}

	wgs.Latitude = WGS84_LATITUDE_BASE + (calc_lat / 3600)
	wgs.Longitude = WGS84_LONGITUDE_BASE + (calc_long / 3600)

	return wgs
}
