package webrpc

import "github.com/pkg/errors"

// TODO: lets get error code.

// TODO: add well defined errors like -- ErrXXX

// perhaps we take extra msg, like msg ...string
func Error(err error) error {
	return errors.Wrapf(err, "rpc error")
}
