package data

import (
	"errors"

	"github.com/tsandl/skvdb/leveldb"
)

type DBType int

type DBRepo interface {
	NewLevelDBCli(bucketName string) error
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Del(key string) error
	State(value string) (string, error)
	Iterator(prefix string) (map[string]string, error)
	IteratorOnlyKey(prefix string) ([]string, error)
	OpenDB(bucketName string) (*leveldb.DB, error)
	CloseDB() error
}

const (
	LevelDBs DBType = iota
	RocksDB
)

type RepoFactory struct {
	dbType DBType
}

func NewRepo(dbType DBType) (DBRepo, error) {
	switch dbType {
	case LevelDBs:
		return &leveldbRepo{}, nil
	case RocksDB:
		return &rocksdbRepo{}, nil
	default:
		return nil, errors.New("unsupported database type")
	}
}
