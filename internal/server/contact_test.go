package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"

	clubv1 "github.com/rbicker/club/api/proto/v1"
	"github.com/rbicker/club/internal/mocks"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (suite *Suite) TestContact() {
	t := suite.T()
	// client connection
	conn, err := suite.NewClientConnection()
	if err != nil {
		t.Fatalf("unable to create client connection: %s", err)
	}
	defer conn.Close()
	client := clubv1.NewClubClient(conn)
	// tests
	tests := []struct {
		name     string
		prepare  func(db *mocks.Store, messenger *mocks.Messenger)
		req      *clubv1.ContactRequest
		wantCode codes.Code
	}{
		{
			name: "valid request",
			req: &clubv1.ContactRequest{
				Name:     "Hans Muster",
				Mail:     "test@example.ch",
				Subject:  "hello world",
				Message:  "Just wanted to say hi.",
				Language: "en",
			},
			prepare: func(db *mocks.Store, messenger *mocks.Messenger) {
				messenger.On("Contact", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
					nil,
				)
			},
			wantCode: codes.OK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			// prepare mocks
			db := new(mocks.Store)
			messenger := new(mocks.Messenger)
			if tt.prepare != nil {
				tt.prepare(db, messenger)
			}
			suite.srv.store = db
			suite.srv.messenger = messenger
			// run function
			res, err := client.Contact(context.Background(), tt.req)
			// check status code
			code, _ := status.FromError(err)
			assert.Equal(tt.wantCode, code.Code(), "response statuscode mismatch")
			db.AssertExpectations(t)
			if code.Code() != codes.OK {
				// if status code is not ok, response should be nil
				assert.Nil(res)
				return
			}
		})
	}
}
