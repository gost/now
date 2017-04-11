package now

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Period struct {
	Start string
	End   string
}

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

func GetPeriod(period string) Period {
	var p Period
	if err := json.Unmarshal([]byte(period), &p); err != nil {
		log.Fatal(err)
	}
	return p
}

func ParsePostgresTime(t string) time.Time {
	var result, _ = time.Parse("2006-01-02 15:04:05-07", t)
	return result
}

func TimeToIso8601(t time.Time) string {
	var result = t.Format("2006-01-02T15:04:05.000Z")
	return result
}

func PostgresToIso8601Period(period string) string {
	var p = GetPeriod(period)
	var startTime = ParsePostgresTime(p.Start)
	var endTime = ParsePostgresTime(p.End)
	var iso8601Period = TimeToIso8601(startTime) + "/" + TimeToIso8601(endTime)
	return iso8601Period
}
