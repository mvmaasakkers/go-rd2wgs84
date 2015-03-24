# go-rd2wgs84

RD coordinates to WGS84 (latitude longitude)conversion function
Rijks driehoeks coördinaten naar WGS84 (latitude longitude) conversie

This calculates the correct Latitude and Longitude based on the (Dutch) RD system.

Usage:

```golang
package main

import (
	"fmt"
	"github.com/mvmaasakkers/go-rd2wgs84"
)

func main() {
	wgs84 := rd2wgs84.Convert(163835.370083, 446830.763585)
	fmt.Printf("%+v\n", wgs84)
}

```

Will print
```
&{Latitude:52.00977421758342 Longitude:5.515894213047998}
```
