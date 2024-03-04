package service

import (
	"context"

	pb "github.com/ZLSMDB/lsmdb_server/api/lsmdb/v1"
	"github.com/ZLSMDB/lsmdb_server/internal/biz"
	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
)

type LsmdbService struct {
	pb.UnimplementedLsmdbServer
	uc     *biz.LevelDBUsecase
	ucS3   *biz.OssUsecase
	dbName string
	conf   *conf.Bootstrap
	log    *log.Helper
}

func NewLsmdbService(uc *biz.LevelDBUsecase, logger log.Logger, oss *biz.OssUsecase, conf *conf.Bootstrap) *LsmdbService {
	svc := &LsmdbService{uc: uc, log: log.NewHelper(logger), ucS3: oss, conf: conf}
	// svc.OpenDB(context.Background(), &pb.OpenDBRequest{DbName: svc.conf.Data.NodeName})
	return svc
}

func (s *LsmdbService) OpenDB(ctx context.Context, req *pb.OpenDBRequest) (*pb.OpenDBReply, error) {
	if err := s.uc.NewLevelDBCli(req.DbName); err != nil {
		s.log.Errorf("open db %v fail because of %v", req.DbName, err)
		return &pb.OpenDBReply{Value: false}, nil
	}
	s.dbName = req.DbName
	return &pb.OpenDBReply{Value: true}, nil
}

func (s *LsmdbService) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutReply, error) {
	// 分流调用oss存储大文件，需要记录当前bucketname的名字, etcd中存储格式为{key: data{dbname, bucketname}}
	// if len(req.Value) > 64*data.MiB {
	// 	// put to oss and etcd.
	// 	err := s.ucS3.PutBytes(s.dbName, req.Key, req.Value)
	// 	if err != nil {
	// 		s.log.Errorf("put key: %s to s3 fail", req.Key)
	// 	}
	// 	// 转化成key-addr更好， key前需要加一个前缀证明是直接存在上
	// 	s.ucEtcd.Put(req.Key, s.dbName, nil)
	// 	// 中
	// 	return &pb.PutReply{Data: true}, nil
	// }

	err := s.uc.Set(req.Key, req.Value)
	if err != nil {
		s.log.Errorf("put key-value <%v> fail", req.Key)
		return &pb.PutReply{Data: false}, err
	}
	return &pb.PutReply{Data: true}, nil
}

/*
	func (s *LsmdbService) PutStr(ctx context.Context, req *pb.PutStrRequest) (*pb.PutStrReply, error) {
		// 分流调用oss存储大文件，需要记录当前bucketname的名字, etcd中存储格式为{key: data{dbname, bucketname}}
		if len(req.Value) > 64*data.MiB {
			// put to oss and etcd.
			err := s.ucS3.PutBytes(s.dbName, req.Key, []byte(req.Value))
			if err != nil {
				s.log.Errorf("put key: %s to s3 fail", req.Key)
			}
			// TODO 存在key同名问题, key + dbname作为key, 还得区分是在lsmdb中还是oss中
			err = s.ucEtcd.Put(fmt.Sprintf("index/%s/%s", s.dbName, req.Key), "oss", nil) // 转化成key-addr更好， key前需要加一个前缀证明是直接存在上
			// 中
			return &pb.PutStrReply{Data: true}, err
		}
		err := s.uc.Set(req.Key, []byte(req.Value))
		if err != nil {
			s.log.Errorf("put key-value <%v> fail", req.Key)
			return &pb.PutStrReply{Data: false}, err
		}
		return &pb.PutStrReply{Data: true}, nil
	}
*/
func (s *LsmdbService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	// if s.ucEtcd.Get(fmt.Sprintf("index/%s/%s", s.dbName, req.Key)) == "oss" {
	// 	value, err := s.ucS3.GetBytes(s.dbName, req.Key)
	// 	return &pb.GetReply{Value: value}, err
	// }
	value, err := s.uc.Get(req.Key)
	if err != nil {
		s.log.Errorf("Get key %s fail", req.Key)
		return &pb.GetReply{}, err
	}
	return &pb.GetReply{Value: value}, nil
}

func (s *LsmdbService) CloseDB(ctx context.Context, req *pb.CloseDBRequest) (*pb.CloseDBReply, error) {
	err := s.uc.CloseDB()
	return &pb.CloseDBReply{}, err
}
