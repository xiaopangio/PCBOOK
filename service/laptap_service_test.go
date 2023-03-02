// Package service  @Author xiaobaiio 2023/2/21 13:30:00
package service_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/xiaopangio/pcbook/pb"
	"github.com/xiaopangio/pcbook/sample"
	"github.com/xiaopangio/pcbook/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestLaptapServer_CreateLaptap(t *testing.T) {
	t.Parallel()
	laptapNoID := sample.NewLaptap()
	laptapNoID.Id = ""
	laptapInvalidID := sample.NewLaptap()
	laptapInvalidID.Id = "invalid-uuid"

	laptapDuplicateID := sample.NewLaptap()
	storeDuplicateID := service.NewInMemoryLaptapStore()
	err := storeDuplicateID.Save(laptapDuplicateID)
	require.Nil(t, err)
	testCases := []struct {
		name        string
		laptap      *pb.Laptap
		laptapStore service.LaptapStore
		code        codes.Code
	}{
		{
			name:        "success_with_id",
			laptap:      sample.NewLaptap(),
			laptapStore: service.NewInMemoryLaptapStore(),
			code:        codes.OK,
		},
		{
			name:        "success_no_id",
			laptap:      laptapNoID,
			laptapStore: service.NewInMemoryLaptapStore(),
			code:        codes.OK,
		},
		{
			name:        "failure_invalid_id",
			laptap:      laptapInvalidID,
			laptapStore: service.NewInMemoryLaptapStore(),
			code:        codes.InvalidArgument,
		},
		{
			name:        "failure_duplicate_id",
			laptap:      laptapDuplicateID,
			laptapStore: storeDuplicateID,
			code:        codes.AlreadyExists,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			request := &pb.CreateLaptapRequest{Laptap: tc.laptap}
			server := service.NewLaptapServer(tc.laptapStore, nil, nil)
			response, err := server.CreateLaptap(context.Background(), request)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.NotEmpty(t, response.Id)
				if len(tc.laptap.Id) > 0 {
					require.Equal(t, tc.laptap.Id, response.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, response)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
