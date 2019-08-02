package scraper

import (
	"github.com/pkg/errors"
)

func baseError(err error, message string) error {
	if err == nil {
		err = errors.New(message)
	} else {
		err = errors.Wrap(err, message)
	}

	return err
}

func (EmptyTarget) RenderingError() error {
	return baseError(nil, "Can't render an empty target")
}

func ContentMissingError() error {
	return baseError(nil, "Target has no content")
}

func MarshallingError(err error) error {
	return baseError(err, "Failed deserializing page content")
}

func RenderingError(err error) error {
	return baseError(err, "failed rendering the target hierarchy to text")
}

func (target htmlTarget) ambiguousTargetError() error {
	return baseError(nil, "cannot locate root for given target")
}
