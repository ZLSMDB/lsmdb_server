package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tsandl/skvdb/leveldb"
)

type LevelDBRepo interface {
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

// LevelDBUsecase is a leveldb usecase.
type LevelDBUsecase struct {
	repo LevelDBRepo
	log  *log.Helper
}

// NewLevelDBUsecase new a leveldb usecase.
func NewLevelDBUsecase(repo LevelDBRepo, logger log.Logger) *LevelDBUsecase {
	return &LevelDBUsecase{repo: repo, log: log.NewHelper(logger)}
}

// TODO 提供给用户的接口
func (s *LevelDBUsecase) NewLevelDBCli(dbname string) error {
	return s.repo.NewLevelDBCli(dbname)
}

func (s *LevelDBUsecase) Set(key string, value []byte) error {
	return s.repo.Set(key, value)
}

func (s *LevelDBUsecase) Get(key string) ([]byte, error) {
	return s.repo.Get(key)
}

func (s *LevelDBUsecase) Delete(key string) error {
	return s.repo.Del(key)
}

func (s *LevelDBUsecase) State(value string) (string, error) {
	return s.repo.State(value)
}

func (s *LevelDBUsecase) Iterator(prefix string) (map[string]string, error) {
	return s.repo.Iterator(prefix)
}

func (s *LevelDBUsecase) IteratorOnlyKey(prefix string) ([]string, error) {
	return s.repo.IteratorOnlyKey(prefix)
}

func (s *LevelDBUsecase) OpenDB(bucketName string) (*leveldb.DB, error) {
	return s.repo.OpenDB(bucketName)
}

func (s *LevelDBUsecase) CloseDB() error {
	return s.repo.CloseDB()
}
