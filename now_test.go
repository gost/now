package now

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestIso8601 tests the ToIso8601 function
func TestPostgresIoso8601(t *testing.T) {
	// arrange
	var postgresperiod = `["2014-03-01 13:00:00+00","2015-05-11 15:30:00+00"]`
	var expectedResult = "2014-03-01T13:00:00.000Z/2015-05-11T15:30:00.000Z"

	// act
	var iso8601Period = PostgresToIso8601Period(postgresperiod)

	// assert
	assert.Equal(t, expectedResult, iso8601Period, "Iso8601 error")
}
