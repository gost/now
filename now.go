package now

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"strings"
)

// Period struct (Start, End)
type Period struct {
	Start string
	End   string
}

// UnmarshalJSON fills period object
func (p *Period) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&p.Start, &p.End}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Period: %d != %d", g, e)
	}
	return nil
}

// GetPeriodFromPostgresString converts Postgres period string to object
func GetPeriodFromPostgresString(period string) Period {
	var p Period
	if err := json.Unmarshal([]byte(period), &p); err != nil {
		log.Fatal(err)
	}
	return p
}

// ParsePostgresTime parses postgres time to object
func ParsePostgresTime(t string) time.Time {
	var result, _ = time.Parse("2006-01-02 15:04:05-07", t)
	return result
}

// TimeToIso8601 converts time object to ISO8601 object
func TimeToIso8601(t time.Time) string {
	var result = t.Format("2006-01-02T15:04:05.000Z")
	return result
}

// ToTime parses a time string to RCX3339Nano format
func ToTime(input string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, input)
}

// ToIso8601 formats a time object to ISO8601 string
func ToIso8601(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}

// ParseTMPeriod parses a TM Period into time array (start,end)
func ParseTMPeriod(input string) [2]time.Time {
	parts := strings.Split(input, "/")
	startTime, _ := ToTime(parts[0])
	endTime, _ := ToTime(parts[1])
	var period [2]time.Time
	period[0] = startTime.UTC()
	period[1] = endTime.UTC()
	return period
}

// Iso8601ToPostgresPeriodFormat formats to Postgres format
func Iso8601ToPostgresPeriodFormat(input string) string {
	period := ParseTMPeriod(input)
	start := ToIso8601(period[0])
	end := ToIso8601(period[1])
	return "[" + start + "," + end + "]"
}

// PostgresToIso8601Period converts Postgres period format to Iso8601 format
func PostgresToIso8601Period(period string) string {
	var p = GetPeriodFromPostgresString(period)
	var startTime = ParsePostgresTime(p.Start)
	var endTime = ParsePostgresTime(p.End)
	var iso8601Period = TimeToIso8601(startTime) + "/" + TimeToIso8601(endTime)
	return iso8601Period
}
