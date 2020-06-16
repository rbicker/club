package server

import (
	"context"
	"net"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"

	gooserv1 "github.com/rbicker/gooser/api/proto/v1"

	"github.com/rbicker/club/internal/mocks"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
)

// Suite is the test suite
type Suite struct {
	suite.Suite
	srv       *Server
	listener  *bufconn.Listener
	mockStore *mocks.Store
}

// GooserProviderStub implements the GooserProvider interface for testing.
type GooserProviderStub struct {
	gooserClient GooserClient
}

// NewGooserClient
func (g GooserProviderStub) NewGooserClient(target string) (GooserClient, error) {
	return g.gooserClient, nil
}

// Close does nothing.
func (g GooserProviderStub) Close() error {
	return nil
}

// ensure GooserProviderStub implements the GooserProvider interface.
var _ GooserProvider = &GooserProviderStub{}

// SetupSuite runs once before all tests
func (suite *Suite) SetupSuite() {
	t := suite.T()
	contextUserReceiver = func(ctx context.Context, gooserTarget string) (*gooserv1.User, error) {
		accessToken, ok := ctx.Value("access_token").(string)
		if !ok || accessToken == "" {
			return nil, nil
		}
		plainPassword := "password"
		hashed, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
		ss := strings.Split(accessToken, ",")
		return &gooserv1.User{
			Id:        ss[0],
			Username:  ss[0],
			Password:  string(hashed),
			Mail:      ss[0],
			Roles:     ss,
			Confirmed: true,
		}, nil
	}
	// store
	db := new(mocks.Store)
	// messenger
	messenger := new(mocks.Messenger)
	// create test grpc server
	srv, err := NewServer(db, messenger)
	if err != nil {
		t.Fatalf("unable to create server: %s", err)
	}
	suite.listener = bufconn.Listen(1024 * 1024)
	srv.listener = suite.listener
	suite.srv = srv
	// run server
	go func() {
		if err := suite.srv.Serve(); err != nil {
			t.Fatalf("grpc server failed: %s", err)
		}
	}()
}

// NewClientConnection creates a new grpc client connection.
func (suite *Suite) NewClientConnection() (*grpc.ClientConn, error) {
	// create dialer for client
	bufDialer := func(context.Context, string) (net.Conn, error) {
		return suite.listener.Dial()
	}
	unaryInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		meta := make(map[string]string)
		accessToken, ok := ctx.Value("access_token").(string)
		if ok && accessToken != "" {
			meta["access_token"] = accessToken
		}
		md := metadata.New(meta)
		ctx = metadata.NewOutgoingContext(context.TODO(), md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
	return grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithUnaryInterceptor(unaryInterceptor),
		grpc.WithInsecure(),
	)
}

// TearDownSuite runs once after all tests.
func (suite *Suite) TearDownSuite() {
	suite.srv.Stop()
}

// TestSuite is needed in order for 'go test' to run this suite.
func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
