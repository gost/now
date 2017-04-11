package now

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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

// PostgresToIso8601Period converts Postgres period format to Iso8601 format
func PostgresToIso8601Period(period string) string {
	var p = GetPeriodFromPostgresString(period)
	var startTime = ParsePostgresTime(p.Start)
	var endTime = ParsePostgresTime(p.End)
	var iso8601Period = TimeToIso8601(startTime) + "/" + TimeToIso8601(endTime)
	return iso8601Period
}
