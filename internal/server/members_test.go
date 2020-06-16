package server

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/genproto/protobuf/field_mask"

	gooserv1 "github.com/rbicker/gooser/api/proto/v1"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/golang/protobuf/ptypes"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	clubv1 "github.com/rbicker/club/api/proto/v1"

	"github.com/rbicker/club/internal/mocks"
	"github.com/rbicker/club/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (suite *Suite) TestListMembers() {
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
		name              string
		prepare           func(db *mocks.Store)
		accessToken       string
		req               *clubv1.ListRequest
		wantCode          codes.Code
		wantLen           int
		wantNextPageToken string
	}{
		{
			name:        "unauthenticated",
			accessToken: "",
			req: &clubv1.ListRequest{
				PageSize: 1,
			},
			wantCode: codes.Unauthenticated,
		},
		{
			name:        "list members",
			accessToken: "user",
			req: &clubv1.ListRequest{
				PageSize: 1,
			},
			prepare: func(db *mocks.Store) {
				db.On("ListMembers", mock.Anything, mock.Anything, "", "", "", int32(1)).Return(
					&[]store.Member{
						{
							Id: "tester",
						},
					},
					int32(5),
					"token",
					nil,
				).Once()
			},
			wantCode:          codes.OK,
			wantLen:           1,
			wantNextPageToken: "token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			// prepare mock
			db := new(mocks.Store)
			if tt.prepare != nil {
				tt.prepare(db)
			}
			suite.srv.store = db
			// prepare context with access token
			ctx := context.Background()
			if tt.accessToken != "" {
				ctx = context.WithValue(ctx, "access_token", tt.accessToken)
			}
			// run function
			res, err := client.ListMembers(ctx, tt.req)
			// check status code
			code, _ := status.FromError(err)
			assert.Equal(tt.wantCode, code.Code(), "response statuscode mismatch")
			db.AssertExpectations(t)
			if code.Code() != codes.OK {
				// if status code is not ok, response should be nil
				assert.Nil(res)
				return
			}
			// check result
			assert.Equal(tt.wantLen, len(res.Members), "length mismatch")
			assert.Equal(tt.wantNextPageToken, res.NextPageToken, "token mismatch")
		})
	}
}

func (suite *Suite) TestGetMember() {
	t := suite.T()
	now := time.Now()
	protoNow, _ := ptypes.TimestampProto(now)
	// client connection
	conn, err := suite.NewClientConnection()
	if err != nil {
		t.Fatalf("unable to create client connection: %s", err)
	}
	defer conn.Close()
	client := clubv1.NewClubClient(conn)
	// tests
	tests := []struct {
		name                 string
		prepare              func(db *mocks.Store)
		accessToken          string
		req                  *clubv1.IdRequest
		wantCode             codes.Code
		wantId               string
		wantCreatedAt        *timestamppb.Timestamp
		wantUpdatedAt        *timestamppb.Timestamp
		wantUserId           string
		wantUsername         string
		wantMail             string
		wantLanguage         string
		wantFirstName        string
		wantLastName         string
		wantDateOfBirth      *timestamppb.Timestamp
		wantPhone            string
		wantAddress          string
		wantAddress2         string
		wantPostalCode       string
		wantCity             string
		wantJuristic         bool
		wantOrganisation     string
		wantWebsite          string
		wantOrganisationType string
		wantInactive         bool
	}{
		{
			name:     "unauthenticated",
			req:      &clubv1.IdRequest{Id: "member"},
			wantCode: codes.Unauthenticated,
		},
		{
			name:        "get member",
			accessToken: "user",
			prepare: func(db *mocks.Store) {
				db.On("GetMember", mock.Anything, mock.Anything, "member").Return(
					&store.Member{
						Id:               "member",
						CreatedAt:        now,
						UpdatedAt:        now,
						UserId:           "member",
						Username:         "member",
						Mail:             "member@test.com",
						Language:         "en",
						FirstName:        "member",
						LastName:         "member",
						DateOfBirth:      now,
						Phone:            "0",
						Address:          "test",
						Address2:         "test",
						PostalCode:       "1234",
						City:             "test",
						Juristic:         true,
						Organisation:     "test",
						Website:          "test.com",
						OrganisationType: "club",
						Inactive:         true,
					},
					nil,
				)
			},
			req:                  &clubv1.IdRequest{Id: "member"},
			wantCode:             codes.OK,
			wantId:               "member",
			wantCreatedAt:        protoNow,
			wantUpdatedAt:        protoNow,
			wantUserId:           "member",
			wantUsername:         "member",
			wantMail:             "member@test.com",
			wantLanguage:         "en",
			wantFirstName:        "member",
			wantLastName:         "member",
			wantDateOfBirth:      protoNow,
			wantPhone:            "0",
			wantAddress:          "test",
			wantAddress2:         "test",
			wantPostalCode:       "1234",
			wantCity:             "test",
			wantJuristic:         true,
			wantOrganisation:     "test",
			wantWebsite:          "test.com",
			wantOrganisationType: "club",
			wantInactive:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			// prepare mock
			db := new(mocks.Store)
			if tt.prepare != nil {
				tt.prepare(db)
			}
			suite.srv.store = db
			// prepare context with access token
			ctx := context.Background()
			if tt.accessToken != "" {
				ctx = context.WithValue(ctx, "access_token", tt.accessToken)
			}
			// run function
			res, err := client.GetMember(ctx, tt.req)
			// check status code
			code, _ := status.FromError(err)
			assert.Equal(tt.wantCode, code.Code(), "response statuscode mismatch")
			db.AssertExpectations(t)
			if code.Code() != codes.OK {
				// if status code is not ok, response should be nil
				assert.Nil(res)
				return
			}

			// check result
			assert.Equal(tt.wantId, res.Id, "id mismatch")
			assert.Equal(tt.wantCreatedAt, res.CreatedAt, "createdAt mismatch")
			assert.Equal(tt.wantUpdatedAt, res.UpdatedAt, "updatedAt mismatch")
			assert.Equal(tt.wantUserId, res.UserId, "userId mismatch")
			assert.Equal(tt.wantUsername, res.Username, "username mismatch")
			assert.Equal(tt.wantMail, res.Mail, "mail mismatch")
			assert.Equal(tt.wantLanguage, res.Language, "language mismatch")
			assert.Equal(tt.wantFirstName, res.FirstName, "firstName mismatch")
			assert.Equal(tt.wantLastName, res.LastName, "lastName mismatch")
			assert.Equal(tt.wantDateOfBirth, res.DateOfBirth, "dateOfBirth mismatch")
			assert.Equal(tt.wantPhone, res.Phone, "phone mismatch")
			assert.Equal(tt.wantAddress, res.Address, "address mismatch")
			assert.Equal(tt.wantAddress2, res.Address2, "address2 mismatch")
			assert.Equal(tt.wantPostalCode, res.PostalCode, "postalCode mismatch")
			assert.Equal(tt.wantCity, res.City, "city mismatch")
			assert.Equal(tt.wantJuristic, res.Juristic, "juristic mismatch")
			assert.Equal(tt.wantOrganisation, res.Organisation, "organisation mismatch")
			assert.Equal(tt.wantWebsite, res.Website, "website mismatch")
			assert.Equal(tt.wantOrganisationType, res.OrganisationType, "organisationType mismatch")
			assert.Equal(tt.wantInactive, res.Inactive, "inactive mismatch")
		})
	}
}

func (suite *Suite) TestValidateMember() {
	t := suite.T()
	dateOfBirth, _ := time.Parse("2006-01-02", "1987-02-25")
	protoDateOfBirth, _ := ptypes.TimestampProto(dateOfBirth)
	// tests
	tests := []struct {
		name        string
		prepare     func(db *mocks.Store)
		accessToken string
		member      *clubv1.Member
		wantErr     bool
	}{
		{
			name: "valid individual",
			prepare: func(db *mocks.Store) {
				db.On("CountMembers",
					mock.Anything,
					mock.Anything,
					`(userId=="member",username=="member",mail=="member@test.com");_id!=oid="member"`,
				).Return(
					int32(0),
					nil,
				)
			},
			member: &clubv1.Member{
				Id:          "member",
				UserId:      "member",
				Username:    "member",
				Mail:        "member@test.com",
				FirstName:   "Member",
				LastName:    "Member",
				DateOfBirth: protoDateOfBirth,
				Phone:       "+41 123 45 67",
				Address:     "Street 1",
				Address2:    "At the roundabout",
				PostalCode:  "1234",
				City:        "Testing",
				Language:    "en",
				Juristic:    false,
			},
			wantErr: false,
		},
		{
			name: "valid juristic",
			prepare: func(db *mocks.Store) {
				db.On("CountMembers",
					mock.Anything,
					mock.Anything,
					`(userId=="member",username=="member",mail=="member@test.com");_id!=oid="member"`,
				).Return(
					int32(0),
					nil,
				)
			},
			member: &clubv1.Member{
				Id:               "member",
				UserId:           "member",
				Username:         "member",
				Mail:             "member@test.com",
				FirstName:        "Member",
				LastName:         "Member",
				Phone:            "+41 123 45 67",
				Address:          "Street 1",
				Address2:         "At the roundabout",
				PostalCode:       "1234",
				City:             "Testing",
				Language:         "en",
				Juristic:         true,
				Organisation:     "test",
				OrganisationType: "club",
			},
			wantErr: false,
		},
		{
			name: "missing field",
			member: &clubv1.Member{
				Id:          "member",
				UserId:      "member",
				Username:    "member",
				Mail:        "member@test.com",
				FirstName:   "",
				LastName:    "Member",
				DateOfBirth: protoDateOfBirth,
				Phone:       "+41 123 45 67",
				Address:     "Street 1",
				Address2:    "At the roundabout",
				PostalCode:  "1234",
				City:        "Testing",
				Language:    "de",
				Juristic:    false,
			},
			wantErr: true,
		},
		{
			name: "invalid mail",
			member: &clubv1.Member{
				Id:          "member",
				UserId:      "member",
				Username:    "member",
				Mail:        "memb@er@test.com",
				FirstName:   "Member",
				LastName:    "Member",
				DateOfBirth: protoDateOfBirth,
				Phone:       "+41 123 45 67",
				Address:     "Street 1",
				Address2:    "At the roundabout",
				PostalCode:  "1234",
				City:        "Testing",
				Language:    "en",
				Juristic:    false,
			},
			wantErr: true,
		},
		{
			name: "invalid phone number",
			member: &clubv1.Member{
				Id:          "member",
				UserId:      "member",
				Username:    "member",
				Mail:        "member@test.com",
				FirstName:   "Member",
				LastName:    "Member",
				DateOfBirth: protoDateOfBirth,
				Phone:       "+41 +123 45 67",
				Address:     "Street 1",
				Address2:    "At the roundabout",
				PostalCode:  "1234",
				City:        "Testing",
				Language:    "en",
				Juristic:    false,
			},
			wantErr: true,
		},
		{
			name: "invalid language",
			member: &clubv1.Member{
				Id:          "member",
				UserId:      "member",
				Username:    "member",
				Mail:        "member@test.com",
				FirstName:   "Member",
				LastName:    "Member",
				DateOfBirth: protoDateOfBirth,
				Phone:       "+41 123 45 67",
				Address:     "Street 1",
				Address2:    "At the roundabout",
				PostalCode:  "1234",
				City:        "Testing",
				Language:    "invalid",
				Juristic:    false,
			},
			wantErr: true,
		},
		{
			name: "duplicate user",
			prepare: func(db *mocks.Store) {
				db.On("CountMembers",
					mock.Anything,
					mock.Anything,
					`(userId=="member",username=="member",mail=="member@test.com");_id!=oid="member"`,
				).Return(
					int32(1),
					nil,
				)
			},
			member: &clubv1.Member{
				Id:          "member",
				UserId:      "member",
				Username:    "member",
				Mail:        "member@test.com",
				FirstName:   "Member",
				LastName:    "Member",
				DateOfBirth: protoDateOfBirth,
				Phone:       "+41 123 45 67",
				Address:     "Street 1",
				Address2:    "At the roundabout",
				PostalCode:  "1234",
				City:        "Testing",
				Language:    "en",
				Juristic:    false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			// prepare mock
			db := new(mocks.Store)
			if tt.prepare != nil {
				tt.prepare(db)
			}
			suite.srv.store = db
			// prepare context with access token
			ctx := context.Background()
			if tt.accessToken != "" {
				ctx = context.WithValue(ctx, "access_token", tt.accessToken)
			}
			// run function
			err := suite.srv.ValidateMember(ctx, message.NewPrinter(language.English), tt.member)
			if tt.wantErr {
				assert.Error(err, "expected error")
			} else {
				assert.Nil(err, "expected error to be nil")
			}
		})
	}
}

func (suite *Suite) TestCreateMember() {
	t := suite.T()
	dateOfBirth, _ := time.Parse("2006-01-02", "1987-02-25")
	protoDateOfBirth, _ := ptypes.TimestampProto(dateOfBirth)
	// client connection
	conn, err := suite.NewClientConnection()
	if err != nil {
		t.Fatalf("unable to create client connection: %s", err)
	}
	defer conn.Close()
	client := clubv1.NewClubClient(conn)
	// tests
	tests := []struct {
		name        string
		prepare     func(db *mocks.Store, gooser *mocks.GooserClient)
		accessToken string
		member      *clubv1.Member
		wantCode    codes.Code
	}{
		{
			name: "valid member",
			prepare: func(db *mocks.Store, gooser *mocks.GooserClient) {
				db.On("CountMembers",
					mock.Anything,
					mock.Anything,
					`userId=="",username=="member",mail=="member@test.com"`,
				).Return(
					int32(0),
					nil,
				)
				db.On("SaveMember", mock.Anything, mock.Anything, mock.MatchedBy(func(member *store.Member) bool {
					return member.Username == "member" &&
						member.Mail == "member@test.com" &&
						member.Language == "en"
				})).Return(
					&store.Member{
						Id:        "member",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						UserId:    "member",
						Username:  "member",
						Mail:      "member@test.com",
						Language:  "en",
					},
					nil,
				)
				gooser.On("CreateUser", mock.Anything, mock.MatchedBy(func(user *gooserv1.User) bool {
					return user.Username == "member" &&
						user.Password == "Member1234!" &&
						user.Mail == "member@test.com" &&
						user.Language == "en"
				})).Return(
					&gooserv1.User{
						Id:        "member",
						CreatedAt: ptypes.TimestampNow(),
						UpdatedAt: ptypes.TimestampNow(),
						Username:  "member",
						Mail:      "member@test.co¶o",
						Language:  "en",
					},
					nil,
				)
			},
			member: &clubv1.Member{
				Username:    "member",
				Password:    "Member1234!",
				Mail:        "member@test.com",
				FirstName:   "Member",
				LastName:    "Member",
				DateOfBirth: protoDateOfBirth,
				Phone:       "+41 123 45 56",
				Address:     "Street1",
				PostalCode:  "1234",
				City:        "Testing",
				Language:    "en",
			},
			wantCode: codes.OK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			// prepare mocks
			db := new(mocks.Store)
			gooser := new(mocks.GooserClient)
			suite.srv.gooserProvider = GooserProviderStub{gooserClient: gooser}
			if tt.prepare != nil {
				tt.prepare(db, gooser)
			}
			suite.srv.store = db
			// prepare context with access token
			ctx := context.Background()
			if tt.accessToken != "" {
				ctx = context.WithValue(ctx, "access_token", tt.accessToken)
			}
			// run function
			res, err := client.CreateMember(ctx, tt.member)
			// check status code
			code, _ := status.FromError(err)
			assert.Equal(tt.wantCode, code.Code(), "response statuscode mismatch")
			db.AssertExpectations(t)
			if code.Code() != codes.OK {
				// if status code is not ok, response should be nil
				assert.Nil(res)
				return
			}
			// check result
			assert.Equal(tt.member.Username, res.GetUsername())
			assert.Equal(tt.member.Mail, res.GetMail())
		})
	}
}

func (suite *Suite) TestUpdateMember() {
	t := suite.T()
	dateOfBirth, _ := time.Parse("2006-01-02", "1987-02-25")
	// client connection
	conn, err := suite.NewClientConnection()
	if err != nil {
		t.Fatalf("unable to create client connection: %s", err)
	}
	defer conn.Close()
	client := clubv1.NewClubClient(conn)
	// tests
	tests := []struct {
		name          string
		prepare       func(db *mocks.Store, gooser *mocks.GooserClient)
		accessToken   string
		req           *clubv1.UpdateMemberRequest
		wantCode      codes.Code
		wantUsername  string
		wantMail      string
		wantFirstName string
	}{
		{
			name: "unauthenticated",
			req: &clubv1.UpdateMemberRequest{
				Member: &clubv1.Member{
					Id:   "user1",
					Mail: "new@testing.com",
				},
				FieldMask: &field_mask.FieldMask{
					Paths: []string{"mail"},
				},
			},
			wantCode: codes.Unauthenticated,
		},
		{
			name:        "change userId",
			accessToken: "user1",
			prepare: func(db *mocks.Store, gooser *mocks.GooserClient) {
				db.On("GetMember", mock.Anything, mock.Anything, "user1").Return(
					&store.Member{
						Id:     "user1",
						UserId: "user1",
					},
					nil,
				)
			},
			req: &clubv1.UpdateMemberRequest{
				Member: &clubv1.Member{
					Id:     "user1",
					UserId: "admin",
				},
				FieldMask: &field_mask.FieldMask{
					Paths: []string{"userId"},
				},
			},
			wantCode: codes.PermissionDenied,
		},
		{
			name:        "change password",
			accessToken: "user1",
			prepare: func(db *mocks.Store, gooser *mocks.GooserClient) {
				db.On("GetMemberByUserId", mock.Anything, mock.Anything, "user1").Return(
					&store.Member{
						Id:     "user1",
						UserId: "user1",
					},
					nil,
				)
			},
			req: &clubv1.UpdateMemberRequest{
				Member: &clubv1.Member{
					Password: "Changed1234!",
				},
				FieldMask: &field_mask.FieldMask{
					Paths: []string{"password"},
				},
			},
			wantCode: codes.InvalidArgument,
		},
		{
			name:        "valid update",
			accessToken: "user1",
			prepare: func(db *mocks.Store, gooser *mocks.GooserClient) {
				db.On("GetMember", mock.Anything, mock.Anything, "user1").Return(
					&store.Member{
						Id:          "user1",
						UserId:      "user1",
						Username:    "user1",
						Mail:        "user1@test.com",
						Language:    "en",
						FirstName:   "User1",
						LastName:    "User1",
						DateOfBirth: dateOfBirth,
						Phone:       "+11 111 11 11",
						Address:     "Test1",
						Address2:    "",
						PostalCode:  "1234",
						City:        "Test",
					},
					nil,
				)
				db.On("CountMembers", mock.Anything, mock.Anything, `(userId=="user1",username=="user2",mail=="user2@test.com");_id!=oid="user1"`).Return(int32(0), nil)
				db.On("SaveMember", mock.Anything, mock.Anything, mock.MatchedBy(func(member *store.Member) bool {
					return member.Username == "user2" &&
						member.Mail == "user2@test.com"
				})).Return(
					&store.Member{
						Id:        "user1",
						UserId:    "user1",
						Username:  "user2",
						Mail:      "user2@test.com",
						FirstName: "User2",
					},
					nil,
				)
				gooser.On("GetUser", mock.Anything, mock.MatchedBy(func(idRequest *gooserv1.IdRequest) bool {
					return idRequest.Id == "user1"
				})).Return(
					&gooserv1.User{
						Id:       "user1",
						Username: "user1",
						Mail:     "user1@test.com",
						Language: "en",
					},
					nil,
				)
				gooser.On("UpdateUser", mock.Anything, mock.MatchedBy(func(updateUserRequest *gooserv1.UpdateUserRequest) bool {
					return updateUserRequest.User.Username == "user2" && updateUserRequest.User.Mail == "user2@test.com"
				})).Return(
					&gooserv1.User{
						Id:       "user1",
						Username: "user2",
						Mail:     "user2@test.com",
						Language: "en",
					},
					nil,
				)
			},
			req: &clubv1.UpdateMemberRequest{
				Member: &clubv1.Member{
					Id:        "user1",
					UserId:    "user1",
					Username:  "user2",
					Mail:      "user2@test.com",
					FirstName: "User2",
				},
				FieldMask: &field_mask.FieldMask{
					Paths: []string{"Username", "Mail", "FirstName"},
				},
			},
			wantCode:      codes.OK,
			wantFirstName: "User2",
			wantUsername:  "user2",
			wantMail:      "user2@test.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			// prepare mock
			db := new(mocks.Store)
			gooser := new(mocks.GooserClient)
			suite.srv.gooserProvider = GooserProviderStub{gooserClient: gooser}
			if tt.prepare != nil {
				tt.prepare(db, gooser)
			}
			suite.srv.store = db
			// prepare context with access token
			ctx := context.Background()
			if tt.accessToken != "" {
				ctx = context.WithValue(ctx, "access_token", tt.accessToken)
			}
			// run function
			res, err := client.UpdateMember(ctx, tt.req)
			// check status code
			code, _ := status.FromError(err)
			assert.Equal(tt.wantCode, code.Code(), "response statuscode mismatch")
			db.AssertExpectations(t)
			if code.Code() != codes.OK {
				// if status code is not ok, response should be nil
				assert.Nil(res)
				return
			}
			// check result
			assert.Equal(tt.wantUsername, res.GetUsername())
			assert.Equal(tt.wantMail, res.GetMail())
			assert.Equal(tt.wantFirstName, res.GetFirstName())
		})
	}
}

func (suite *Suite) TestDeleteMember() {
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
		name        string
		prepare     func(db *mocks.Store, gooser *mocks.GooserClient)
		accessToken string
		req         *clubv1.IdRequest
		wantCode    codes.Code
	}{
		{
			name: "unauthenticated",
			req: &clubv1.IdRequest{
				Id: "user1",
			},
			wantCode: codes.Unauthenticated,
		},
		{
			name:        "not admin",
			accessToken: "user1",
			req: &clubv1.IdRequest{
				Id: "member1",
			},
			wantCode: codes.PermissionDenied,
		},
		{
			name:        "valid",
			accessToken: "admin",
			prepare: func(db *mocks.Store, gooser *mocks.GooserClient) {
				db.On("GetMember", mock.Anything, mock.Anything, "member1").Return(
					&store.Member{
						Id:     "member1",
						UserId: "user1",
					},
					nil,
				)
				db.On("DeleteMember", mock.Anything, mock.Anything, "member1").Return(
					nil,
				)
				gooser.On("DeleteUser", mock.Anything, mock.MatchedBy(func(request *gooserv1.IdRequest) bool {
					return request.Id == "user1"
				})).Return(
					&empty.Empty{},
					nil,
				)
			},
			req: &clubv1.IdRequest{
				Id: "member1",
			},
			wantCode: codes.OK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			// prepare mocks
			db := new(mocks.Store)
			gooser := new(mocks.GooserClient)
			suite.srv.gooserProvider = GooserProviderStub{gooserClient: gooser}
			if tt.prepare != nil {
				tt.prepare(db, gooser)
			}
			suite.srv.store = db
			// prepare context with access token
			ctx := context.Background()
			if tt.accessToken != "" {
				ctx = context.WithValue(ctx, "access_token", tt.accessToken)
			}
			// run function
			res, err := client.DeleteMember(ctx, tt.req)
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
