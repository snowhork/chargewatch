package vals

import "golang.org/x/xerrors"

var ErrValidation = xerrors.New("validation failed")
