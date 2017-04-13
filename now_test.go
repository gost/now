package now

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIso8601 tests the ToIso8601 function
func TestPostgresIoso8601(t *testing.T) {
	// arrange
	var postgresperiod = `["2014-03-01 13:00:00+00","2015-05-11 15:30:00+00"]`
	var expectedResult = "2014-03-01T13:00:00.000Z/2015-05-11T15:30:00.000Z"

	// act
	var iso8601Period = PostgresToIso8601Period(postgresperiod)

	// assert
	assert.Equal(t, expectedResult, iso8601Period, "Iso8601 error")
}


// TestToPostgresPeriodFormat tests the ToPostgresPeriodFormat function
func TestIso8601ToPostgresPeriodFormat(t *testing.T) {
	// arrange
	var period = "2014-03-01T13:00:00Z/2015-05-11T15:30:00Z"
	var expectedResult = `["2014-03-01 13:00:00+00","2015-05-11 15:30:00+00"]`

	// act
	var postgresPeriod = Iso8601ToPostgresPeriod(period)

	// assert
	assert.Equal(t, expectedResult, postgresPeriod, "ToPostgresPeriodFormat error")
}

// RoundTestIso8601 Test Iso8601 period -> Postgres period -> Iso8601 period
func RoundTestIso8601(t *testing.T){
	// arrange
	var period = "2014-03-01T13:00:00Z/2015-05-11T15:30:00Z"

	// act
	var postgresPeriod = Iso8601ToPostgresPeriod(period)
	var iso8601Period = PostgresToIso8601Period(postgresPeriod)

	// assert
	assert.Equal(t, iso8601Period, period, "Round iso8601 test error")
}

// RoundTestPostgres Test Postgres period -> Iso8601 period -> Postgres period
func RoundTestPostgres(t *testing.T){
	// arrange
	var postgresPeriod = `["2014-03-01 13:00:00+00","2015-05-11 15:30:00+00"]`

	// act
	var iso8601Period = PostgresToIso8601Period(postgresPeriod)
	var postgresPeriodAfter = Iso8601ToPostgresPeriod(iso8601Period)

	// assert
	assert.Equal(t, postgresPeriod, postgresPeriodAfter, "Round iso8601 test error")
}
