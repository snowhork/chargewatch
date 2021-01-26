package repo

import "golang.org/x/xerrors"

var ErrNotFound = xerrors.New("entry not found")
