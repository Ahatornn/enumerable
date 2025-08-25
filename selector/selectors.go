package selector

import "time"

// Int is a predefined selector function that returns the input int value unchanged.
// It is useful when working with operations that require a key selector function
// but the element type is already int and no transformation is needed.
//
// This selector is functionally equivalent to the identity function for int values
// and ensures zero overhead in key extraction.
var Int = func(x int) int { return x }

// Int64 is a predefined selector function that returns the input int64 value unchanged.
// It is useful when working with operations that require a key selector function
// but the element type is already int64 and no transformation is needed.
//
// This selector is functionally equivalent to the identity function for int64 values
// and ensures zero overhead in key extraction.
var Int64 = func(x int64) int64 { return x }

// Bool is a predefined selector function that returns the input bool value unchanged.
// It is useful when working with operations that require a key selector function
// but the element type is already bool and no transformation is needed.
//
// This selector is functionally equivalent to the identity function for bool values
// and ensures zero overhead in key extraction.
var Bool = func(x bool) bool { return x }

// Byte is a predefined selector function that returns the input byte value unchanged.
// It is useful when working with operations that require a key selector function
// but the element type is already byte (uint8) and no transformation is needed.
//
// This selector is functionally equivalent to the identity function for byte values
// and ensures zero overhead in key extraction.
var Byte = func(x byte) byte { return x }

// Float32 is a predefined selector function that returns the input float32 value unchanged.
// It is useful when working with operations that require a key selector function
// but the element type is already float32 and no transformation is needed.
//
// Note: When comparing floating-point values, consider the impact of NaN and precision.
// This selector does not perform any validation or normalization.
var Float32 = func(x float32) float32 { return x }

// Float64 is a predefined selector function that returns the input float64 value unchanged.
// It is useful when working with operations that require a key selector function
// but the element type is already float64 and no transformation is needed.
//
// Note: When comparing floating-point values, consider the impact of NaN and precision.
// This selector does not perform any validation or normalization.
var Float64 = func(x float64) float64 { return x }

// Rune is a predefined selector function that returns the input rune value unchanged.
// It is useful when working with operations that require a key selector function
// but the element type is already rune (int32) and no transformation is needed.
//
// This selector is functionally equivalent to the identity function for rune values
// and ensures zero overhead in key extraction.
var Rune = func(x rune) rune { return x }

// String is a predefined selector function that returns the input string value unchanged.
// It is useful when working with operations that require a key selector function
// but the element type is already string and no transformation is needed.
//
// This selector is functionally equivalent to the identity function for string values
// and ensures zero overhead in key extraction.
var String = func(x string) string { return x }

// Time is a predefined selector function that returns the input time.Time value unchanged.
// It is useful when working with operations that require a key selector function
// but the element type is already time.Time and no transformation is needed.
//
// This selector is functionally equivalent to the identity function for time values
// and ensures zero overhead in key extraction.
var Time = func(x time.Time) time.Time { return x }
