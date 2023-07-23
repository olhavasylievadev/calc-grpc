//package calculator implements GRPC server struct and Calculate method
//supported operations are Addition '+' Subtraction '-',
//Multiplication '*' and Division '/'

package calculator

import (
	"context"
	"errors"
	"github.com/olhavasylievadev/calc-grpc/pkg/api"
)

type GRPCServer struct {
	api.UnimplementedCalculatingServiceServer
}

// Calculate takes in context and a pointer to CalculateRequest
// It calculates the sum of the operation and sends back a pointer to CalculateResponse.
// In case of unsupported operation or division by 0 error occurs
func (srv *GRPCServer) Calculate(ctx context.Context, r *api.CalculateRequest) (*api.CalculateResponse, error) {
	first := r.GetFirst()
	second := r.GetSecond()
	operation := r.GetOperation()

	var res api.CalculateResponse

	switch operation {
	case "+":
		res.Res = first + second
	case "-":
		res.Res = first - second
	case "*":
		res.Res = first * second
	case "/":
		if second == 0 {
			return nil, errors.New("you can't divide by 0")
		}
		res.Res = first / second
	default:
		return nil, errors.New("this operation isn't supported.\n supported operations are: '+', '-', '*', '/'")
	}

	return &res, nil
}

