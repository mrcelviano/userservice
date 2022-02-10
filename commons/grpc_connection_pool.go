package commons

import (
	"context"
	pool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
	"time"
)

const maxBodySize = 1024 * 1024 * 1024

type Pool struct {
	p *pool.Pool
}

type CustClientConn struct {
	conn  *grpc.ClientConn
	close func() error
}

func (cc CustClientConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	err := cc.conn.Invoke(ctx, method, args, reply, opts...)
	return err
}

func (cc CustClientConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return cc.conn.NewStream(ctx, desc, method, opts...)
}

//InitGRPCConnectionPool создает пул соединений с конкретным сервисом
func InitGRPCConnectionPool(serviceName string, capacity int) (Pool, error) {
	factory := func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(GetServiceAddr(serviceName),
			grpc.WithInsecure(),
			grpc.WithDefaultCallOptions(
				grpc.WaitForReady(true),
				grpc.MaxCallRecvMsgSize(maxBodySize),
			),
		)
		if err != nil {
			return nil, err
		}
		return conn, err
	}
	pl, err := pool.New(factory, 0, capacity*10, time.Minute, time.Minute)
	if err != nil {
		return Pool{}, err
	}
	return Pool{p: pl}, nil
}

func NewCustomClientConn(clientConn *pool.ClientConn, close func() error) grpc.ClientConnInterface {
	return CustClientConn{
		conn:  clientConn.ClientConn,
		close: close,
	}
}

// Соединение нужно обязательно использовать, иначе будет утечка
func (p Pool) Get(ctx context.Context) (grpc.ClientConnInterface, error) {
	conn, err := p.p.Get(ctx)
	if err != nil {
		return nil, err
	}
	return NewCustomClientConn(conn, conn.Close), err
}
