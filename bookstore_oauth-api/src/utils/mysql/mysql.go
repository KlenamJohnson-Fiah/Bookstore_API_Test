package mysql

import (
	"strings"

	errors "bookstore_oauth-api/src/utils/errors_utils"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewNotFoundError("No record matching given id")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewInternalServerError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
