package utils

import (
	"github.com/jackc/pgerrcode"
	"github.com/uptrace/bun/driver/pgdriver"
)

func PgError(err error) error {
	if err, ok := err.(pgdriver.Error); ok && err.IntegrityViolation() {
		return err
	} else if err.Field('C') == pgerrcode.InvalidTransactionState {
		return err
	}
	return nil
}
