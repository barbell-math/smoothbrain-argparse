package translators

type (
	// Used to represent a flag that can only be supplied once, returning a
	// boolean value indicating the flags presence.
	Flag struct{}
)

func (_ Flag) Translate(arg string) (bool, error) {
	return true, nil
}
func (_ Flag) Reset()               {}
func (_ Flag) HelpAddendum() string { return "" }
