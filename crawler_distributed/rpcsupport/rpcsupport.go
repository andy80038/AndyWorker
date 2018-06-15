package rpcsupport

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	fmt.Printf("listening on port:%s", host)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("accept error:%v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)

	}
	return nil
}
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil

}
