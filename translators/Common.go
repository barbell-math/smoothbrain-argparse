package translators

type (
	// Used to control the way the parser treats values.
	ArgType int

	// An interface that defines what actions can be performed when translating
	// a string argument to a typed value. The translator is expected to perform
	// all validation required to ensure a correct value is returned. It is also
	// expected to return an error if a value is found to be invalid.
	Translator[T any] interface {
		Translate(arg string) (T, error)
		Reset()
		HelpAddendum() string
		ArgType() ArgType
	}
)

const (
	UnknownArgType ArgType = iota
	// Represents a flag type that must accept a value as an argument and must
	// only be supplied once.
	ValueArgType
	// Represents a flag type that can accept many values as an argument and
	// can be supplied many times.
	MultiValueArgType
	// Represents a flag type that must not accept a value and must only be
	// supplied once.
	FlagArgType
	// Represents a flag type that must not accept a value and may be supplied
	// many times.
	MultiFlagArgType
)

var (
	singleSpecificationArgTypes = map[ArgType]struct{}{
		ValueArgType: struct{}{}, FlagArgType: struct{}{},
	}

	multiSpecificationArgTypes = map[ArgType]struct{}{
		MultiValueArgType: struct{}{}, MultiFlagArgType: struct{}{},
	}
)
