package gtime

import "time"

//var Location, _ = time.LoadLocation("Asia/Shanghai")

//主要针对windows系统
func Location() *time.Location {
	loca, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loca = time.FixedZone("CST", 8*3600)
	}
	return loca
}
