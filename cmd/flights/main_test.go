package main

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/f0mster/flights/gen/golang/proto"
)

func Test_Main(t *testing.T) {
	freePort, err := GetFreePort()
	require.NoError(t, err)
	host := "localhost:" + strconv.Itoa(freePort)
	go startServer(host)
	time.Sleep(2 * time.Second)
	cli := proto.NewFlightsFinderJSONClient("http://"+host, http.DefaultClient)
	resp, err := cli.GetEndpoint(context.Background(), &proto.GetEndpointReq{Flights: []*proto.Flight{{
		From: "AAA",
		To:   "CCC",
	}}})
	require.NoError(t, err)
	require.Equal(t, (&proto.GetEndpointResp{
		From: "AAA",
		To:   "CCC",
	}).String(), resp.String())
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
