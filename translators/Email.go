package translators

import "net/mail"

type (
	// A translator that checks that the supplied string is a valid email,
	// returning a [mail.Address].
	Email struct{}

	// A translator that checks that the supplied string is a valid email,
	// returning the supplied string.
	StrEmail struct{}
)

func (_ Email) Translate(arg string) (mail.Address, error) {
	a, err := mail.ParseAddress(arg)
	if err != nil {
		return mail.Address{}, err
	}
	return *a, err
}
func (_ Email) Reset() {}
func (_ Email) HelpAddendum() string {
	return "Parses a single RFC 5322 email address."
}

func (_ StrEmail) Translate(arg string) (string, error) {
	_, err := mail.ParseAddress(arg)
	return arg, err
}
func (_ StrEmail) Reset() {}
func (_ StrEmail) HelpAddendum() string {
	return "Parses a single RFC 5322 email address."
}
