package main

import (
    "fmt"
    "encoding/json"
    "log"
    "time"
)

func main() {
    var period = `["2014-03-01 13:00:00+00","2015-05-11 15:30:00+00"]`
    var iso8601Period = ToIso8601Period(period) 
    fmt.Printf("%#v\n", iso8601Period)
}
