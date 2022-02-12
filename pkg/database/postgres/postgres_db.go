package postgres

import (
	"errors"
	"fmt"
	goCraft "github.com/gocraft/dbr"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"strings"
)

const (
	maxOpenConnections = 6
	maxIdleConnections = 2

	sslMode = "disable"

	defaultLimit  = 10
	defaultOffset = 0
)

var ErrDuplicateValue = errors.New("duplicate key value violates unique constraint")

func NewGoCraftDBConnectionPG(host, port, user, password, dbName string) (*goCraft.Connection, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		host, port, user, dbName, sslMode, password)

	conn, err := goCraft.Open("postgres", connectionString, nil)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(maxOpenConnections)
	conn.SetMaxIdleConns(maxIdleConnections)

	return conn, nil
}

func SetLimit(limit uint64, stmt *goCraft.SelectStmt) {
	if limit != 0 {
		stmt.Limit(limit)
		return
	}

	stmt.Limit(defaultLimit)
}

func SetOffset(offset uint64, stmt *goCraft.SelectStmt) {
	if offset != 0 {
		stmt.Offset(offset)
		return
	}

	stmt.Offset(defaultOffset)
}

func SetSortKeyAndSortOrder(sortKey, sortOrder string, stmt *goCraft.SelectStmt) {
	if sortKey == "" {
		sortKey = "id"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	stmt.OrderDir(sortKey, strings.ToLower(sortOrder) == "asc")
}

func GetError(err error) error {
	if strings.Contains(err.(*pq.Error).Message, ErrDuplicateValue.Error()) {
		return ErrDuplicateValue
	}
	return err
}
