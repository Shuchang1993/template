/*
@Time : 2023/3/31 13:35
@Author : sc-52766
@File : rpc_client
@Software: GoLand
*/
package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// Client is a struct that holds the grpc client connection and options
type Client struct {
	conn   *grpc.ClientConn
	opts   []grpc.DialOption
	ctx    context.Context
	cancel context.CancelFunc
}

// NewClient creates a new grpc client with the given options
func NewClient(address string, opts ...grpc.DialOption) (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		opts:   opts,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// Close closes the grpc client connection
func (c *Client) Close() error {
	c.cancel()
	return c.conn.Close()
}

// WithTimeout sets the timeout for the grpc client
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.ctx, c.cancel = context.WithTimeout(c.ctx, timeout)
	return c
}

// WithAddress sets the address for the grpc client
func (c *Client) WithAddress(address string) *Client {
	c.opts = append(c.opts, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	c.conn, _ = grpc.Dial(address, c.opts...)
	return c
}

// WithOptions sets the options for the grpc client
func (c *Client) WithOptions(opts ...grpc.DialOption) *Client {
	c.opts = append(c.opts, opts...)
	c.conn, _ = grpc.Dial(c.conn.Target(), c.opts...)
	return c
}

// Call invokes the specified method on the grpc server
func (c *Client) Call(ctx context.Context, method string, request interface{}, response interface{}) error {
	return c.conn.Invoke(ctx, method, request, response)
}

// SetContext sets the context for the grpc client
func (c *Client) SetContext(ctx context.Context) *Client {
	c.ctx, c.cancel = context.WithCancel(ctx)
	return c
} 

// The above code defines a grpc client framework that can set timeout, control the sending address, and modify rpc parameters.