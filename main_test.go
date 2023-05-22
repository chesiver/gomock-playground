package main

import (
	"testing"

	mock_clients "github.com/chesiver/gomock-playground/mocks"
	"github.com/golang/mock/gomock"
)

func TestInvokeClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := mock_clients.NewMockClient(ctrl)
	mockClient.EXPECT().GetData().Times(1)
	InvokeClient(mockClient)
}
