package domain

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/f0mster/fligth_test_task/gen/golang/proto"
)

func TestFlights_CalculateEndpoints(t *testing.T) {
	tests := []struct {
		name string
		req  *proto.GetEndpointReq
		resp *proto.GetEndpointResp
		err  error
	}{
		{
			name: "positive. simple path",
			req: &proto.GetEndpointReq{Flights: []*proto.Flight{
				{
					From: "AAA",
					To:   "BBB",
				},
			}},
			resp: &proto.GetEndpointResp{
				From: "AAA",
				To:   "BBB",
			},
			err: nil,
		},
		{
			name: "positive. harder path",
			req: &proto.GetEndpointReq{Flights: []*proto.Flight{
				{
					From: "AAA",
					To:   "BBB",
				},
				{
					From: "BBB",
					To:   "CCC",
				},
			}},
			resp: &proto.GetEndpointResp{
				From: "AAA",
				To:   "CCC",
			},
			err: nil,
		},
		{
			name: "negative. double start points",
			req: &proto.GetEndpointReq{Flights: []*proto.Flight{
				{
					From: "AAA",
					To:   "BBB",
				},
				{
					From: "EEE",
					To:   "BBB",
				},
				{
					From: "BBB",
					To:   "CCC",
				},
			}},
			resp: nil,
			err:  ErrBadPath,
		},
		{
			name: "negative. split route in the middle",
			req: &proto.GetEndpointReq{Flights: []*proto.Flight{
				{
					From: "AAA",
					To:   "BBB",
				},
				{
					From: "BBB",
					To:   "DDD",
				},
				{
					From: "BBB",
					To:   "CCC",
				},
				{
					From: "DDD",
					To:   "EEE",
				},
				{
					From: "CCC",
					To:   "EEE",
				},
			}},
			resp: nil,
			err:  ErrBadPath,
		},

		{
			name: "negative. double end points",
			req: &proto.GetEndpointReq{Flights: []*proto.Flight{
				{
					From: "AAA",
					To:   "BBB",
				},
				{
					From: "BBB",
					To:   "CCC",
				},
				{
					From: "BBB",
					To:   "DDD",
				},
			}},
			resp: nil,
			err:  ErrBadPath,
		},

		{
			name: "negative. empty route",
			req:  &proto.GetEndpointReq{Flights: []*proto.Flight{}},
			resp: nil,
			err:  ErrEmptyPath,
		},
		{
			name: "negative. nil",
			req:  nil,
			resp: nil,
			err:  ErrEmptyPath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := &Flights{}
			got, err := f.CalculateRoute(tt.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.resp, got)
		})
	}
}
