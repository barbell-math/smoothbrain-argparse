package sbargp

import (
	"flag"

	"golang.org/x/exp/constraints"
)

// Sets two flags: -verbose and -v. They both increment the same underlying
// value by one every time they are supplied. Both -v and -verbose can be
// supplied multiple times.
func Verbosity[T constraints.Signed](
	fs *flag.FlagSet,
	val *T,
	_default T,
) {
	fs.Func(
		"verbose",
		"An integer value that controls how much information to print to the console. Higher number=more information",
		Int(val, _default, 1),
	)
	fs.Func(
		"v",
		"An integer value that controls how much information to print to the console. Higher number=more information",
		Int(val, _default, 1),
	)
}
