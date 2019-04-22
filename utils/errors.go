/*
package utils provides non-specific functionality to the library
*/
package utils

import "github.com/pkg/errors"

func BaseError(err error, message string) {
	panic(errors.Wrap(err, message+"\n"))
}
