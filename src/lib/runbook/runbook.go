package runbook

import (
	"database/sql"
	"fmt"
)

type DB interface {
	GetRunBooks() ([]RunBook, error)
}

type DBClient struct {
	Dao *sql.DB
}

type RunBook struct {
	ID          string
	Time        string
	Description string
}

func (dbClient *DBClient) GetRunBooks() ([]RunBook, error) {
	rows, err := dbClient.Dao.Query("SELECT * FROM runbooks;")
	if err != nil {
		fmt.Println("Error on dbClient.db.Exec in GetRunBooks()")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []RunBook{}
	for rows.Next() {
		i := RunBook{}
		err = rows.Scan(&i.ID, &i.Time, &i.Description)
		if err != nil {
			return nil, err
		}
		fmt.Println(i)
		data = append(data, i)
	}
	return data, nil
}
