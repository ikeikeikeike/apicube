package lifecycle

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/ikeikeikeike/apicube/rpc/mockpb/mock_lifecycle"
	pb "github.com/ikeikeikeike/apicube/rpc/pb/apicube/lifecycle"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_lifecycle.NewMockPingServiceClient(ctrl)

	client.EXPECT().GetPing(
		gomock.Any(),
		&pb.Ping{Msg: "one-two-token"},
	).Return(&pb.Ping{
		Msg: "one-two-token",
	}, nil)

	t.Helper()

	resp, err := client.GetPing(
		context.Background(),
		&pb.Ping{Msg: "one-two-token"},
	)
	if err != nil && resp.Msg != "one-two-token" {
		t.Errorf("mocking failed")
	}

	t.Log("Reply : ", resp)
}
