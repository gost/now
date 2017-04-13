package main

import (
	"fmt"

	"github.com/gost/now"
)

func main() {
	var postgresperiod = `["2014-03-01 13:00:00+00","2015-05-11 15:30:00+00"]`
	var iso8601Period = now.PostgresToIso8601Period(postgresperiod)
	var postgresPeriodAfter = now.Iso8601ToPostgresPeriodFormat(iso8601Period)
	fmt.Println("Input:" + postgresperiod)
	fmt.Println("ISO8601 period:" + iso8601Period)
	fmt.Println("Postgres period after:" + postgresPeriodAfter)

}
