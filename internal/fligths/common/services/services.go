package services

import "github.com/f0mster/fligth_test_task/gen/golang/proto"

type Finder interface {
	CalculateRoute(req *proto.GetEndpointReq) (*proto.GetEndpointResp, error)
}
