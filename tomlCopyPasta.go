package sbargp

import (
	"strings"
	"time"
)

var (
	localOffset   = func() int { _, o := time.Now().Zone(); return o }()
	localDatetime = time.FixedZone("datetime-local", localOffset)
	localDate     = time.FixedZone("date-local", localOffset)
	localTime     = time.FixedZone("time-local", localOffset)

	dtTypes = []struct {
		fmt  string
		zone *time.Location
	}{
		{time.RFC3339Nano, time.Local},
		{"2006-01-02T15:04:05.999999999", localDatetime},
		{"2006-01-02", localDate},
		{"15:04:05.999999999", localTime},

		// tomlNext
		{"2006-01-02T15:04Z07:00", time.Local},
		{"2006-01-02T15:04", localDatetime},
		{"15:04", localTime},
	}

	datetimeRepl = strings.NewReplacer(
		"z", "Z",
		"t", "T",
		" ", "T",
	)
)

// Go's time.Parse() will accept numbers without a leading zero; there isn't any
// way to require it. https://github.com/golang/go/issues/29911
//
// Depend on the fact that the separators (- and :) should always be at the same
// location.
func missingLeadingZero(d, l string) bool {
	for i, c := range []byte(l) {
		if c == '.' || c == 'Z' {
			return false
		}
		if (c < '0' || c > '9') && d[i] != c {
			return true
		}
	}
	return false
}
