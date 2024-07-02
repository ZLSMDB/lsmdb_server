package data

import (
	"errors"

	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	gorocksdb "github.com/linxGnu/grocksdb"
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
	LevelDBs = "leveldb"
	RocksDB  = "rocksdb"
)

func NewRepo(conf *conf.Data, leveldb *leveldb.DB, rocksdb *gorocksdb.DB, logger log.Logger) (DBRepo, error) {
	switch conf.Dbtype {
	case LevelDBs:
		return NewLevelDBRepo(conf, leveldb, logger), nil
	case RocksDB:
		return NewRocksdbRepo(conf, rocksdb, logger), nil
	default:
		return nil, errors.New("unsupported database type")
	}
}
