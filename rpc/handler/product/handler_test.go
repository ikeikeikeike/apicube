package product

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/ikeikeikeike/apicube/rpc/mockpb/mock_product"
	pb "github.com/ikeikeikeike/apicube/rpc/pb/apicube/product"
)

func TestListSimilars(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_product.NewMockProductServiceClient(ctrl)
	products := []*pb.Product{{Name: "1"}, {Name: "2"}}

	client.EXPECT().ListSimilars(
		gomock.Any(),
		&pb.ListSimilarsRequest{
			Parent: "products",
			Name:   "name-of-product",
		},
	).Return(&pb.ListSimilarsResponse{
		Products: products,
	}, nil)

	t.Helper()

	resp, err := client.ListSimilars(
		context.Background(),
		&pb.ListSimilarsRequest{
			Parent: "products",
			Name:   "name-of-product",
		},
	)
	if err != nil && resp.Products != nil && len(resp.Products) == 2 {
		t.Errorf("mocking failed")
	}

	t.Log("Reply : ", resp)
}
