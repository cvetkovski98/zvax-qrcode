package errutil

import (
	"database/sql"
	"errors"
)

func ParseError(err error) error {
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}
