package main

import (
	"context"
	"errors"
	"io"
	"log"

	"google.golang.org/grpc"
)

func myStreamClientInteceptor1(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	// ストリームがopenされる前に行われる前処理
	log.Println("[pre] my stream client interceptor 1", method)

	stream, err := streamer(ctx, desc, cc, method, opts...)
	return &myClientStreamWrapper1{stream}, err
}

type myClientStreamWrapper1 struct {
	grpc.ClientStream
}

func (s *myClientStreamWrapper1) SendMsg(m any) error {
	log.Println("[pre message] my stream client interceptor 1: ", m)
	return s.ClientStream.SendMsg(m)
}

func (s *myClientStreamWrapper1) RecvMsg(m any) error {
	err := s.ClientStream.RecvMsg(m)

	if !errors.Is(err, io.EOF) {
		log.Println("[post message] my stream client interceptor 1: ", m)
	}
	return err
}

func (s *myClientStreamWrapper1) CliseSend() error {
	err := s.ClientStream.CloseSend()
	log.Println("[post] my stream client interceptor 1")
	return err
}
