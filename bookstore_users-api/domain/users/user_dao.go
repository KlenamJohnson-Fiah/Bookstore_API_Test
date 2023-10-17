package users

import (
	users_db "bookstore_users-api/datasources/mysql/user_db"
	"bookstore_users-api/utils/errors"
	"fmt"
	"strings"

	"bookstore_users-api/logger"
	mysqlutils "bookstore_users-api/utils/mysql_utils"
)

const (
	//indexUniqueEmail = "users.email_UNIQUE"
	errorNoRow                  = "no row in result set"
	queryInsertUser             = "INSERT into users(first_name,last_name,email,date_created,status,password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name,last_name,email,date_created,status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id =?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id,first_name,last_name,email,date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE email=? AND password=? AND status=?;"
)

// var (
// 	userDB = make(map[int64]*User)
// )

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error when trying to prepare GET user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("Error when trying to GET user by id", err)
		return errors.NewInternalServerError("database error")
	}

	/*
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()
		if err := users_db.Client.Ping(); err != nil {
			panic(err)
			//log.Printf("Here is the error : %s\n", err)
		}
		result := userDB[user.Id]
		if result == nil {
			errors.NewNotFoundError(fmt.Sprintf("User %d not found\n", user.Id))
		}
		user.Id = result.Id
		user.FirstName = result.FirstName
		user.LastName = result.LastName
		user.Email = result.Email
		user.DateCreated = result.DateCreated
		fmt.Println(user)
		return nil
	*/
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	// we use the Prepare to validate the query. We can then use the
	// the stmt in other places

	if err != nil {
		logger.Error("Error when trying to prepare SAVE user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	inserResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("Error when trying to CREATE user by id", err)
		return mysqlutils.ParseError(err)

		// if strings.Contains(err.Error(), indexUniqueEmail) {
		// 	return errors.NewBadRequestError(fmt.Sprintf("user email: %s, already exists", user.Email))
		// }

	}

	// Could have written a short statement like
	/*
		 result ,err := users_db.Client.Exec(queryInsertUser,user.FirstName, user.LastName, user.Email, user.DateCreated)
		 if err != nil{
			 return errors.errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s/n", err.Error())))

		}
	*/
	// Pretty much the same to the  method above
	userId, err := inserResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get LastInsertID from DB", err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	/*
		current := userDB[user.Id]
		if current != nil {
			if current.Email == user.Email {
				return errors.NewBadRequestError(fmt.Sprintf("user %s already registered", user.Email))
			}
			return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.Id))
		}

		user.DateCreated = date_utils.GetNowString()

		userDB[user.Id] = user
	*/
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error when trying to prepare UPDATE statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	UpdateResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	_ = UpdateResult
	if err != nil {
		logger.Error("Error when trying to execute UPDATE user", err)
		return errors.NewInternalServerError("database error")

	}
	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare DELETE user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, deleteErr := stmt.Exec(user.Id)
	if deleteErr != nil {
		logger.Error("Error when trying to execute DELETE user", err)
		return errors.NewInternalServerError("database error")

	}
	return nil

}

func (user *User) FindByStatus(status string) (Users, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare FIND user statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error when trying to execute FIND user", err)
		return nil, errors.NewInternalServerError("database error")
	}

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error when trying to scan for users who meet a specified status", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)

	}
	rows.Close()

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil

}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("Error when trying to prepare GET user by Email and Password statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, user.Status == StatusActive)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		if strings.Contains(err.Error(), mysqlutils.ErrorNoRow) {
			return errors.NewNotFoundError("invalid user credentials")
		}

		logger.Error("Error when trying to GET user by email and password", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
