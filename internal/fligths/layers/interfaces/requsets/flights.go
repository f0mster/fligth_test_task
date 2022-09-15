package requsets

import (
	"context"
	"fmt"

	"github.com/f0mster/fligth_test_task/gen/golang/proto"
	"github.com/f0mster/fligth_test_task/internal/fligths/common/services"
)

type FlightsController struct {
	fnd services.Finder
}

func NewFlightsController(fnd services.Finder) *FlightsController {
	return &FlightsController{fnd: fnd}
}

func (f *FlightsController) GetEndpoint(
	ctx context.Context,
	req *proto.GetEndpointReq,
) (*proto.GetEndpointResp, error) {
	resp, err := f.fnd.CalculateRoute(req)
	if err != nil {
		err = fmt.Errorf("failed to calculate route: %w", err)
		return nil, err
	}

	return resp, nil
}
