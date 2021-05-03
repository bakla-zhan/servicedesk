package swagger

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	CreateItem(item *Request)
	GetItem(ID int) *Request
	DeleteItem(ID int)
	UpdateItem(item *Request)
	GetItemsList() []*Request
}

type Storage struct {
	db *gorm.DB
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
	dbPort, ok := os.LookupEnv("DATABASE_PORT")
	if !ok {
		log.Fatal("DATABASE_NAME env is not set")
	}
	dbUser, ok := os.LookupEnv("DATABASE_USER")
	if !ok {
		log.Fatal("DATABASE_NAME env is not set")
	}
	dbPassw, ok := os.LookupEnv("DATABASE_PASSWORD")
	if !ok {
		log.Fatal("DATABASE_NAME env is not set")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", dbHost, dbUser, dbPassw, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Request{})
	if err != nil {
		log.Println("unable to migrate the schema")
	}

	stor := &Storage{
		db: db,
	}

	return stor
}

func (s *Storage) CreateItem(item *Request) {
	s.db.Create(&Request{
		Head:  item.Head,
		Body:  item.Body,
		Email: item.Email,
	})
}

func (s *Storage) GetItem(ID int) *Request {
	res := &Request{}

	s.db.First(&res, ID)

	return res
}

func (s *Storage) DeleteItem(ID int) {
	res := &Request{}

	s.db.First(&res, ID)
	s.db.Delete(&res, ID)
}

func (s *Storage) UpdateItem(item *Request) {
	res := &Request{}

	s.db.First(&res, item.Id)
	s.db.Model(&res).Updates(Request{
		Head: item.Head,
		Body: item.Body,
	})
}

func (s *Storage) GetItemsList() []*Request {
	res := make([]*Request, 0)

	s.db.Find(&res)

	return res
}
