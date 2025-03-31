package flags

import (
	"reflect"
)

// Arg represents a positional argument on the command line.
type Arg struct {
	// The name of the positional argument (used in the help)
	Name string

	// A description of the positional argument (used in the help)
	Description string

	// If non-empty, only a certain set of values is allowed for the positional
	// argument. This can only be enforced if the positional argument is not a
	// slice type (i.e., it does not collect all remaining arguments).
	Choices []string

	// The minimal number of required positional arguments
	Required int

	// The maximum number of required positional arguments
	RequiredMaximum int

	field reflect.StructField
	value reflect.Value
	tag   multiTag
}

func (a *Arg) isRemaining() bool {
	return a.value.Type().Kind() == reflect.Slice
}
