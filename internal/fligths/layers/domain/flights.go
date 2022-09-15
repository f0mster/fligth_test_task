package domain

import (
	"fmt"

	"github.com/f0mster/flights/gen/golang/proto"
)

type Flights struct {
}

var ErrBadPath = fmt.Errorf("bad path")
var ErrEmptyPath = fmt.Errorf("empty path")

func (f *Flights) CalculateRoute(req *proto.GetEndpointReq) (*proto.GetEndpointResp, error) {
	pathFT := map[string]string{}
	pathTF := map[string]string{}
	if req == nil || len(req.Flights) == 0 {
		return nil, ErrEmptyPath
	}
	for _, flight := range req.Flights {
		if pathFT[flight.From] != "" {
			return nil, ErrBadPath
		}
		if pathTF[flight.To] != "" {
			return nil, ErrBadPath
		}

		pathFT[flight.From] = flight.To
		pathTF[flight.To] = flight.From
	}
	start := ""
	end := ""
	for from, to := range pathFT {
		if pathFT[to] == "" {
			end = to
		}
		if pathTF[from] == "" {
			start = from
		}
	}

	return &proto.GetEndpointResp{
		From: start,
		To:   end,
	}, nil
}
