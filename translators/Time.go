package translators

import (
	"fmt"
	stdtime "time"
)

type (
	// Represents a cmd line argument that will be parsed to a time value with
	// a date and time component. Uses [stdtime.Parse] internally.
	time struct {
		format string
	}

	// Represents a cmd line argument that will be parsed as a duration. All
	// the rules that [stdtime.ParseDuration] use will be used here.
	Duration struct{}
)

func NewTime(format string) time {
	return time{format: format}
}
func (t time) Translate(arg string) (stdtime.Time, error) {
	return stdtime.Parse(t.format, arg)
}
func (_ time) Reset() {}
func (t time) HelpAddendum() string {
	return fmt.Sprintf("Expected format: %s", t.format)
}

func (_ Duration) Translate(arg string) (stdtime.Duration, error) {
	return stdtime.ParseDuration(arg)
}
func (_ Duration) Reset() {}
func (_ Duration) HelpAddendum() string {
	return `Example formats: "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`
}
