package services

import "github.com/f0mster/flights/gen/golang/proto"

type Finder interface {
	CalculateRoute(req *proto.GetEndpointReq) (*proto.GetEndpointResp, error)
}
