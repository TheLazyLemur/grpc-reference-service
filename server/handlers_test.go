package main

import (
	"context"
	"prototut/pb"
	"testing"
)

func TestSum(t *testing.T) {
	tc := []struct {
		name string
		in   *pb.NumbersRequest
		out  *pb.CalculationResponse
		err  error
	}{
		{
			name: "Sum - result is 420",
			in: &pb.NumbersRequest{
				Numbers: []int64{210, 210},
			},
			out: &pb.CalculationResponse{
				Result: 420,
			},
			err: nil,
		},
		{
			name: "Sum - result is 69",
			in: &pb.NumbersRequest{
				Numbers: []int64{12, 57},
			},
			out: &pb.CalculationResponse{
				Result: 69,
			},
			err: nil,
		},
	}

	s := NewServer(":8083")
	defer s.Stop()

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			out, err := s.Sum(context.Background(), tt.in)
			if err != tt.err {
				t.Errorf("Sum(%v) = %v, want %v", tt.in, err, tt.err)
			}

			if out.GetResult() != tt.out.GetResult() {
				t.Errorf("Sum(%v) = %v, want %v", tt.in, out, tt.out)
			}
		})
	}
}
