# now
Now is a time toolkit for GOST

Functions:

- Convert Time to/from ISO8601 and Postgres format

- Convert Period to/from ISO8601 and Postgres format

### PostgresToIso8601Period:

Converts from ["2014-03-01 13:00:00+00","2015-05-11 15:30:00+00"] to 2014-03-01T13:00:00.000Z/2015-05-11T15:30:00.000Z

Sample:

var iso8601Period = PostgresToIso8601Period(postgresperiod)

### Iso8601ToPostgresPeriodFormat:

Converts from 2014-03-01T13:00:00.000Z/2015-05-11T15:30:00.000Z to ["2014-03-01 13:00:00+00","2015-05-11 15:30:00+00"]

Sample:

var postgresPeriod = Iso8601ToPostgresPeriod(period)




