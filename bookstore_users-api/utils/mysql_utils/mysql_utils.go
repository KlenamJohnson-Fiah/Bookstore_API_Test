package mysqlutils

import (
	"bookstore_users-api/utils/errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRow = "no row in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRow) {
			return errors.NewNotFoundError("no record matching given ID")
		}
		return errors.NewNotFoundError("Error parsing Database Response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing data")
}
