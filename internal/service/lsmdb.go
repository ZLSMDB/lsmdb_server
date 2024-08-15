package service

import (
	"context"
	"strings"

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
	if err := svc.uc.NewLevelDBCli(svc.conf.Data.NodeName); err != nil {
		svc.log.Errorf("open db %v fail because of %v", svc.conf.Data.NodeName, err)
		return nil
	}
	svc.dbName = svc.conf.Data.NodeName
	// svc.OpenDB(context.Background(), &pb.OpenDBRequest{DbName: svc.conf.Data.NodeName})
	return svc
}

// use for simple node
func (s *LsmdbService) OpenDB(ctx context.Context, req *pb.OpenDBRequest) (*pb.OpenDBReply, error) {
	// if err := s.uc.NewLevelDBCli(s.conf.Data.NodeName); err != nil {
	// 	s.log.Errorf("Open db %v fail because of %v", req.DbName, err)
	// 	return &pb.OpenDBReply{Value: false}, nil
	// }
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

	// err := s.uc.Set(fmt.Sprintf("%s_%s", s.dbName, req.Key), req.Value)
	err := s.uc.Set(req.Key, req.Value)
	if err != nil {
		s.log.Errorf("put key-value <%v> fail", req.Key)
		return &pb.PutReply{Data: false}, err
	}
	return &pb.PutReply{Data: true}, nil
}

func (s *LsmdbService) BatchPut(ctx context.Context, req *pb.BatchPutRequest) (*pb.BatchPutReply, error) {

	err := s.uc.BatchSet(req.Keys, req.Values)
	if err != nil {
		s.log.Errorf("put key-value fail")
		return &pb.BatchPutReply{Data: false}, err
	}
	return &pb.BatchPutReply{Data: true}, nil
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
	// value, err := s.uc.Get(fmt.Sprintf("%s_%s", s.dbName, req.Key))

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

func (s *LsmdbService) OpenDBWeb(ctx context.Context, req *pb.OpenDBWebRequest) (*pb.OpenDBWebReply, error) {
	if err := s.uc.NewLevelDBCli(s.conf.Data.NodeName); err != nil {
		s.log.Errorf("Open db %v fail because of %v", req.Dbname, err)
		return &pb.OpenDBWebReply{Value: false}, nil
	}
	s.dbName = req.Dbname
	return &pb.OpenDBWebReply{Value: true}, nil
}

func (s *LsmdbService) CloseDBWeb(ctx context.Context, req *pb.CloseDBWebRequest) (*pb.CloseDBWebReply, error) {
	err := s.uc.CloseDB()
	return &pb.CloseDBWebReply{}, err
}

func (s *LsmdbService) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferReply, error) {
	reply, err := s.TransferKV(ctx, (*pb.TransferKVRequest)(req))
	return (*pb.TransferReply)(reply), err
}

func (s *LsmdbService) TransferKV(ctx context.Context, req *pb.TransferKVRequest) (*pb.TransferKVReply, error) {
	cli, conn := ConnectDB(req.NodeAddress) // TODO 优化连接速度
	conn.Connect()
	defer conn.Close()
	cli.OpenDB(ctx, &pb.OpenDBRequest{DbName: strings.Split(req.RegionName, "_")[0]}) // dbname
	values, err := s.getKVs(cli, req.Keys)
	if err != nil {
		return &pb.TransferKVReply{Success: false}, err
	}
	for i, value := range values {
		_, err := cli.Put(ctx, &pb.PutRequest{Key: req.Keys[i], Value: value})
		if err != nil {
			return &pb.TransferKVReply{Success: false}, err
		}
	}

	return &pb.TransferKVReply{Success: true}, nil
}

func (s *LsmdbService) getKVs(cli pb.LsmdbClient, keys []string) ([][]byte, error) {
	reply, err := cli.GetKVs(context.Background(), &pb.GetKVsRequest{Keys: keys})
	if err != nil {
		return nil, err
	}
	return reply.Values, nil
}

func (s *LsmdbService) GetKVs(ctx context.Context, req *pb.GetKVsRequest) (*pb.GetKVsReply, error) {
	var values [][]byte
	var err error
	for _, key := range req.Keys {
		reply, err := s.uc.Get(key)
		if err != nil {
			continue
		}
		values = append(values, reply)
	}
	return &pb.GetKVsReply{Values: values}, err
}
