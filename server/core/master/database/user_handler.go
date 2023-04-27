package database

import "errors"

type User struct {
	Id       int
	Name     string
	Password string
}

func (database *Database) CreateUser(user *User) error {
	_, err := database.Instance.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Name, user.Password)
	return err
}

func (database *Database) UserFromName(name string) (*User, error) {
	rows, err := database.Instance.Query("SELECT * FROM users WHERE username=?", name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Password)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, nil
}

func (database *Database) UserExists(name string) (bool, error) {
	result, err := database.Instance.Query("SELECT * FROM users WHERE username=?", name)
	if err != nil {
		return false, err
	}

	defer result.Close()

	return result.Next(), nil
}

func (database *Database) ComparePwd(name, password string) error {
	user, err := database.UserFromName(name)
	if err != nil {
		return err
	}

	if user.Password == password {
		return nil
	}

	return errors.New("invalid credentials")
}
