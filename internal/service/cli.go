package service

import (
	"fmt"

	pblsmdb "github.com/ZLSMDB/lsmdb_server/api/lsmdb/v1"
	g "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	MAX_MESSAGE_LENGTH = 256 * 1024 * 1024 // 可根据具体需求设置，此处设为256M
)

func ConnectDB(addr string) (pblsmdb.LsmdbClient, *g.ClientConn) {
	diaOpt := g.WithDefaultCallOptions(g.MaxCallRecvMsgSize(MAX_MESSAGE_LENGTH), g.MaxCallSendMsgSize(MAX_MESSAGE_LENGTH))
	conn, err := g.Dial(addr, g.WithTransportCredentials(insecure.NewCredentials()), diaOpt)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, nil
	}
	// defer conn.Close()
	client := pblsmdb.NewLsmdbClient(conn)
	return client, conn
}
