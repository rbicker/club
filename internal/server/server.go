package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	clubv1 "github.com/rbicker/club/api/proto/v1"
	"github.com/rbicker/club/internal/mailer"
	"github.com/rbicker/club/internal/store"
	gooserv1 "github.com/rbicker/gooser/api/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Server implements the club server.
type Server struct {
	secret              string
	port                string
	store               store.Store
	messenger           mailer.Messenger
	grpcServer          *grpc.Server
	useReflection       bool
	listener            net.Listener
	errorLogger         *log.Logger
	infoLogger          *log.Logger
	contextUserReceiver func(ctx context.Context, db store.Store) (*gooserv1.User, error)
	gooserTarget        string
	gooserProvider      GooserProvider
}

// ensure MGO implements the store interface.
var _ clubv1.ClubServer = &Server{}

// PageToken represents a pagination token.
type PageToken struct {
	Filter string
	Sort   string
	Skip   int32
}

// contextUserReceiver gets the access token from the context
// and queries the corresponding user from gooser.
var contextUserReceiver = func(ctx context.Context, gooserTarget string) (*gooserv1.User, error) {
	accessToken, ok := ctx.Value("access_token").(string)
	if !ok || accessToken == "" {
		return nil, nil
	}
	// receive user from gooser
	conn, err := grpc.Dial(gooserTarget, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	client := gooserv1.NewGooserClient(conn)
	ctx = context.WithValue(ctx, "access_token", accessToken)
	return client.GetUser(ctx, &gooserv1.IdRequest{})
}

// NewServer returns a new gooser server.
func NewServer(db store.Store, messenger mailer.Messenger, opts ...func(*Server) error) (*Server, error) {
	// create server (with default values)
	var srv = Server{
		port:           "50051",
		gooserTarget:   "gooser:50051",
		store:          db,
		messenger:      messenger,
		gooserProvider: &Gooser{},
	}
	// run functional options
	for _, op := range opts {
		err := op(&srv)
		if err != nil {
			return nil, fmt.Errorf("setting option failed: %w", err)
		}
	}
	// default loggers
	if srv.infoLogger == nil {
		srv.infoLogger = log.New(os.Stdout, "INFO: ", log.Lmsgprefix+log.LstdFlags)
	}
	if srv.errorLogger == nil {
		srv.errorLogger = log.New(os.Stdout, "ERROR: ", log.Lmsgprefix+log.LstdFlags)
	}
	// unary server interceptor
	unaryInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Internal, "retrieving metadata failed")
		}
		if header, ok := md["access_token"]; ok {
			token := header[0]
			ctx = context.WithValue(ctx, "access_token", token)
		}
		return handler(ctx, req)
	}
	// register grpc server
	srv.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)
	// enable reflection
	if srv.useReflection {
		reflection.Register(srv.grpcServer)
	}
	clubv1.RegisterClubServer(srv.grpcServer, &srv)
	return &srv, nil
}

// Serve starts serving the gooser server.
func (srv *Server) Serve() error {
	var err error
	if srv.listener == nil {
		// use tcp listener by default
		srv.listener, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", srv.port))
	}
	if err != nil {
		return fmt.Errorf("gooser server is unable to server: %w", err)
	}
	return srv.grpcServer.Serve(srv.listener)
}

// GetUserFromContext returns the user corresponding
// to the access token in the given context.
func (srv *Server) GetUserFromContext(ctx context.Context) (*gooserv1.User, error) {
	return contextUserReceiver(ctx, srv.gooserTarget)
}

// EnableReflection instructs the grpc server to enable reflection.
func EnableReflection() func(*Server) error {
	return func(srv *Server) error {
		srv.useReflection = true
		return nil
	}
}

// Stop stops the gooser server.
func (srv *Server) Stop() error {
	stopped := make(chan struct{})
	go func() {
		srv.grpcServer.GracefulStop()
		close(stopped)
	}()
	t := time.NewTimer(10 * time.Second)
	select {
	case <-t.C:
		srv.grpcServer.Stop()
	case <-stopped:
		t.Stop()
	}
	return nil
}

// WithPort sets the club server port.
func WithPort(port string) func(*Server) error {
	return func(srv *Server) error {
		i, err := strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("unable to convert given port '%s' to number", port)
		}
		if i <= 0 {
			return fmt.Errorf("port number %s is invalid because it is less or equal 0", port)
		}
		srv.port = port
		return nil
	}
}

// WithInfoLogger sets the info logger for the server.
// By default, log entries will be written to stdout.
func WithInfoLogger(logger *log.Logger) func(*Server) error {
	return func(srv *Server) error {
		srv.infoLogger = logger
		return nil
	}
}

// WithErrorLogger sets the error logger for the server.
// By default, log entries will be written to stderr.
func WithErrorLogger(logger *log.Logger) func(*Server) error {
	return func(srv *Server) error {
		srv.infoLogger = logger
		return nil
	}
}

// WithGooserTarget sets the target (host:port) for the gooser connection.
func WithGooserTarget(target string) func(*Server) error {
	return func(srv *Server) error {
		srv.gooserTarget = target
		return nil
	}
}
