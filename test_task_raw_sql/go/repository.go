package swagger

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Repository interface {
	CreateItem(item *Request) error
	GetItem(ID string) (*Request, error)
	DeleteItem(ID string) error
	UpdateItem(item *Request) error
	GetItemsList() ([]*Request, error)
}

type Storage struct {
	db *sql.DB
}

var (
	ErrNotFound = errors.New("not found")
)

func wrapError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return err
	}
}

func DBConnect() Repository {
	dbHost, ok := os.LookupEnv("DATABASE_HOST")
	if !ok {
		log.Fatal("DATABASE_HOST env is not set")
	}
	dbName, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		log.Fatal("DATABASE_NAME env is not set")
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgres://%s/%s?sslmode=disable&user=postgres&database=%s&password=postgres", dbHost, dbName, dbName),
	)
	if err != nil {
		log.Fatal(err)
	}
	// for db.Ping() != nil {
	// 	log.Println(db.Ping())
	// }

	defer db.Close()

	stor := &Storage{
		db: db,
	}

	return stor
}

func (s *Storage) CreateItem(item *Request) error {
	if _, err := s.db.Exec("INSERT INTO requests VALUES (NULL, $1, $2, $3)",
		item.Head, item.Body, item.Email); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetItem(ID string) (*Request, error) {
	res := &Request{}

	if err := s.db.QueryRow("SELECT * FROM requests WHERE id = $1", ID).
		Scan(&res.Id, &res.Head, &res.Body, &res.Email); err != nil {
		return nil, wrapError(err)
	}
	return res, nil
}

func (s *Storage) DeleteItem(ID string) error {
	_, err := s.db.Exec("DELETE FROM requests WHERE id = $1", ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateItem(item *Request) error {
	_, err := s.db.Exec("UPDATE requests SET head=$1, body=$2 WHERE id = $3", item.Head, item.Body, item.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetItemsList() ([]*Request, error) {
	res := make([]*Request, 0)

	rows, err := s.db.Query("select * from requests")
	if err != nil {
		return res, wrapError(err)
	}

	defer rows.Close()

	for rows.Next() {
		var item *Request

		err = rows.Scan(&item.Id, &item.Head, &item.Body, &item.Email)
		if err != nil {
			return res, wrapError(err)
		}

		res = append(res, item)
	}

	return res, nil
}
