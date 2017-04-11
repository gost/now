package main

import (
	"fmt"

	"github.com/gost/now"
)

func main() {
	var postgresperiod = `["2014-03-01 13:00:00+00","2015-05-11 15:30:00+00"]`
	var iso8601Period = now.PostgresToIso8601Period(postgresperiod)
	fmt.Println("OSI8601 period:" + iso8601Period)
}
