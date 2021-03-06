package main

import (
	"context"
	"net/rpc/jsonrpc"
	"testing"

	_authuc "github.com/yeqown/gRBAC-server/internal-modules/auth/usecase"
	"github.com/yeqown/gRBAC-server/pkg/logger"
)

func Test_StartRPC(t *testing.T) {
	// init logger
	logger.InitLogger("./logs")

	// start rpc
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go StartRPC(ctx, 9090)

	client, err := jsonrpc.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		t.Errorf("could not dail rpc server: %v", err)
		return
	}
	var (
		args = &_authuc.Args{
			UID:     "18302889215",
			ResDesc: "admin",
			Action:  "login",
		}
		reply bool
	)
	if err := client.Call("Auth.IsPermitted", args, &reply); err != nil {
		t.Errorf("could not call rpc method: %v", err)
	}

	if reply != true {
		t.Errorf("want true, get %v", reply)
	}
}
