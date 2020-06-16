package server

import (
	"context"

	gooserv1 "github.com/rbicker/gooser/api/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	clubv1 "github.com/rbicker/club/api/proto/v1"
	"google.golang.org/grpc"
)

// GooserClient is identical to gooserv1.GooserClient.
type GooserClient interface {
	gooserv1.GooserClient
}

// GooserProvider provides a method to receive a new (connected) client and to close the connection.
type GooserProvider interface {
	NewGooserClient(target string) (GooserClient, error)
	Close() error
}

// Gooser implements the GooserProvider interface.
type Gooser struct {
	conn *grpc.ClientConn
}

// ensure Gooser implements the GooserProvider interface.
var _ GooserProvider = &Gooser{}

// NewGooserClient creates a grpc connection to the given target
// and returns a new (connected) gooser client.
func (g *Gooser) NewGooserClient(target string) (GooserClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	g.conn = conn
	return gooserv1.NewGooserClient(conn), nil
}

// Close closes the grpc connection.
func (g *Gooser) Close() error {
	return g.conn.Close()
}

// userHasAnyOfRoles checks if the given user has any of the given roles.
// If so, true is returned.
func userHasAnyOfRoles(user *gooserv1.User, roles ...string) bool {
	for _, ur := range user.GetRoles() {
		for _, r := range roles {
			if ur == r {
				return true
			}
		}
	}
	return false
}

// gooserToClubGroup converts the given gooser group to a club group.
func gooserToClubGroup(group *gooserv1.Group) *clubv1.Group {
	return &clubv1.Group{
		Id:        group.Id,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
		Name:      group.Name,
		Roles:     group.Roles,
		Members:   group.Members,
	}
}

// clubToGooserGroup converts the given club group to a gooser group.
func clubToGooserGroup(group *clubv1.Group) *gooserv1.Group {
	return &gooserv1.Group{
		Id:        group.Id,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
		Name:      group.Name,
		Roles:     group.Roles,
		Members:   group.Members,
	}
}

// ChangePassword changes the password using gooser.
func (srv *Server) ChangePassword(ctx context.Context, request *clubv1.ChangePasswordRequest) (*empty.Empty, error) {
	conn, err := grpc.Dial(srv.gooserTarget, grpc.WithInsecure())
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, "gooser connection failed")
	}
	defer conn.Close()
	c := gooserv1.NewGooserClient(conn)
	return c.ChangePassword(ctx, &gooserv1.ChangePasswordRequest{
		Id:          request.Id,
		OldPassword: request.OldPassword,
		NewPassword: request.NewPassword,
	})
}

// ConfirmMail confirms the mail address using gooser.
func (srv *Server) ConfirmMail(ctx context.Context, request *clubv1.ConfirmMailRequest) (*empty.Empty, error) {
	conn, err := grpc.Dial(srv.gooserTarget, grpc.WithInsecure())
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, "gooser connection failed")
	}
	defer conn.Close()
	c := gooserv1.NewGooserClient(conn)
	return c.ConfirmMail(ctx, &gooserv1.ConfirmMailRequest{
		Token: request.Token,
	})
}

// ForgotPassword starts the password reset process using gooser.
func (srv *Server) ForgotPassword(ctx context.Context, request *clubv1.ForgotPasswordRequest) (*empty.Empty, error) {
	conn, err := grpc.Dial(srv.gooserTarget, grpc.WithInsecure())
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, "gooser connection failed")
	}
	defer conn.Close()
	c := gooserv1.NewGooserClient(conn)
	return c.ForgotPassword(ctx, &gooserv1.ForgotPasswordRequest{
		Username: request.Username,
		Mail:     request.Mail,
	})
}

// ResetPassword resets the password for the user corresponding to the given token using gooser.
func (srv *Server) ResetPassword(ctx context.Context, request *clubv1.ResetPasswordRequest) (*empty.Empty, error) {
	conn, err := grpc.Dial(srv.gooserTarget, grpc.WithInsecure())
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, "gooser connection failed")
	}
	defer conn.Close()
	c := gooserv1.NewGooserClient(conn)
	return c.ResetPassword(ctx, &gooserv1.ResetPasswordRequest{
		Token:    request.Token,
		Password: request.Password,
	})
}

// ListGroups lists the groups using gooser.
func (srv *Server) ListGroups(ctx context.Context, request *clubv1.ListRequest) (*clubv1.ListGroupsResponse, error) {
	conn, err := grpc.Dial(srv.gooserTarget, grpc.WithInsecure())
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, "gooser connection failed")
	}
	defer conn.Close()
	c := gooserv1.NewGooserClient(conn)
	res, err := c.ListGroups(ctx, &gooserv1.ListRequest{
		PageSize:  0,
		PageToken: request.PageToken,
		Filter:    request.Filter,
	})
	if err != nil {
		return nil, err
	}
	var groups []*clubv1.Group
	for _, g := range res.Groups {
		groups = append(groups, gooserToClubGroup(g))
	}
	return &clubv1.ListGroupsResponse{
		Groups:        groups,
		NextPageToken: res.NextPageToken,
		PageSize:      res.PageSize,
		TotalSize:     res.TotalSize,
	}, nil
}

// GetGroup gets the group with the given id from gooser.
func (srv *Server) GetGroup(ctx context.Context, request *clubv1.IdRequest) (*clubv1.Group, error) {
	conn, err := grpc.Dial(srv.gooserTarget, grpc.WithInsecure())
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, "gooser connection failed")
	}
	defer conn.Close()
	c := gooserv1.NewGooserClient(conn)
	g, err := c.GetGroup(ctx, &gooserv1.IdRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}
	return gooserToClubGroup(g), nil
}

// CreateGroup creates the given group using gooser.
func (srv *Server) CreateGroup(ctx context.Context, group *clubv1.Group) (*clubv1.Group, error) {
	conn, err := grpc.Dial(srv.gooserTarget, grpc.WithInsecure())
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, "gooser connection failed")
	}
	defer conn.Close()
	c := gooserv1.NewGooserClient(conn)
	g, err := c.CreateGroup(ctx, clubToGooserGroup(group))
	if err != nil {
		return nil, err
	}
	return gooserToClubGroup(g), nil
}

// UpdateGroup updates the given group using gooser.
func (srv *Server) UpdateGroup(ctx context.Context, request *clubv1.UpdateGroupRequest) (*clubv1.Group, error) {
	conn, err := grpc.Dial(srv.gooserTarget, grpc.WithInsecure())
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, "gooser connection failed")
	}
	defer conn.Close()
	c := gooserv1.NewGooserClient(conn)
	g, err := c.UpdateGroup(ctx, &gooserv1.UpdateGroupRequest{
		Group:     clubToGooserGroup(request.Group),
		FieldMask: request.FieldMask,
	})
	if err != nil {
		return nil, err
	}
	return gooserToClubGroup(g), nil
}

// DeleteGroup deletes the given group.
func (srv *Server) DeleteGroup(ctx context.Context, request *clubv1.IdRequest) (*empty.Empty, error) {
	conn, err := grpc.Dial(srv.gooserTarget, grpc.WithInsecure())
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, "gooser connection failed")
	}
	defer conn.Close()
	c := gooserv1.NewGooserClient(conn)
	return c.DeleteGroup(ctx, &gooserv1.IdRequest{Id: request.Id})
}
