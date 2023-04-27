package database

func (database *Database) CreateUserTable() error {
	_, err := database.Instance.Exec(`
			create table if not exists users
			(
			    id       INTEGER
			        constraint users_pk
			            primary key autoincrement,
			    username TEXT,
			    password TEXT
			);
    `)
	return err
}
