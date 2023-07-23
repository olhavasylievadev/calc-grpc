package calculator

import (
	"context"
	"github.com/olhavasylievadev/calc-grpc/pkg/api"
	"testing"
)

func TestGRPCServer_CalculateAdd(t *testing.T) {
	s := GRPCServer{}

	testAdd := api.CalculateRequest{First: 5, Second: 5, Operation: "+"}
	exp := api.CalculateResponse{Res: 10}

	resp, err := s.Calculate(context.Background(), &testAdd)
	if err != nil {
		t.Error("Got unexpected error while testing addition")
	}
	if resp.Res != exp.Res {
		t.Errorf("Got %v, expected %v", resp.Res, exp.Res)
	}
}

func TestGRPCServer_CalculateSub(t *testing.T) {
	s := GRPCServer{}

	testSub := api.CalculateRequest{First: 10, Second: 5, Operation: "-"}
	exp := api.CalculateResponse{Res: 5}

	resp, err := s.Calculate(context.Background(), &testSub)
	if err != nil {
		t.Error("Got unexpected error while testing subtraction")
	}
	if resp.Res != exp.Res {
		t.Errorf("Got %v, expected %v", resp.Res, exp.Res)
	}
}

func TestGRPCServer_CalculateMul(t *testing.T) {
	s := GRPCServer{}

	testMul := api.CalculateRequest{First: 10, Second: 5, Operation: "*"}
	exp := api.CalculateResponse{Res: 50}

	resp, err := s.Calculate(context.Background(), &testMul)
	if err != nil {
		t.Error("Got unexpected error while testing multiplication")
	}
	if resp.Res != exp.Res {
		t.Errorf("Got %v, expected %v", resp.Res, exp.Res)
	}
}

func TestGRPCServer_CalculateDiv(t *testing.T) {
	s := GRPCServer{}

	testDiv := api.CalculateRequest{First: 10, Second: 5, Operation: "/"}
	exp := api.CalculateResponse{Res: 2}

	resp, err := s.Calculate(context.Background(), &testDiv)
	if err != nil {
		t.Error("Got unexpected error while testing division")
	}
	if resp.Res != exp.Res {
		t.Errorf("Got %v, expected %v", resp.Res, exp.Res)
	}
}
