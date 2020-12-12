package php

import (
	"github.com/syyongx/php2go"
)

// Empty empty()
var Empty = php2go.Empty

// IsNumeric is_numeric()
// Numeric strings consist of optional sign, any number of digits, optional decimal part and optional exponential part.
// Thus +0123.45e6 is a valid numeric value.
// In PHP hexadecimal (e.g. 0xf4c3b00c) is not supported, but IsNumeric is supported.
var IsNumeric = php2go.IsNumeric
