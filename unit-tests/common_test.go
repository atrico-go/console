package unit_tests

import "github.com/atrico-go/testing/random"

var randomValues = random.NewValueGeneratorBuilder().
	WithDefaultStringLength(5).
	Build()

