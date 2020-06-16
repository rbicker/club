package server

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/golang/protobuf/ptypes"
	gooserv1 "github.com/rbicker/gooser/api/proto/v1"

	"github.com/golang/protobuf/protoc-gen-go/generator"
	fieldmaskUtils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/genproto/protobuf/field_mask"

	"github.com/rbicker/club/internal/store"

	"github.com/rbicker/club/internal/utils"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	clubv1 "github.com/rbicker/club/api/proto/v1"
)

// ListMembers lists the club members.
func (srv *Server) ListMembers(ctx context.Context, request *clubv1.ListRequest) (*clubv1.ListMembersResponse, error) {
	u, err := srv.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	printer := message.NewPrinter(language.Make(u.Language))
	filter := request.GetFilter()
	members, totalSize, token, err := srv.store.ListMembers(ctx, printer, filter, request.GetOrderBy(), request.GetPageToken(), request.GetPageSize())
	if err != nil {
		return nil, err
	}
	var pbMembers []*clubv1.Member
	var pageSize int32
	if members != nil {
		pageSize = int32(len(*members))
		for _, m := range *members {
			pbMembers = append(pbMembers, m.ToPb())
		}
	}
	return &clubv1.ListMembersResponse{
		Members:       pbMembers,
		NextPageToken: token,
		PageSize:      pageSize,
		TotalSize:     totalSize,
	}, nil
}

// GetMember queries the club member with the given id.
// If the id is empty, the own member will be returned.
func (srv *Server) GetMember(ctx context.Context, request *clubv1.IdRequest) (*clubv1.Member, error) {
	u, err := srv.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	printer := message.NewPrinter(language.Make(u.Language))
	id := request.GetId()
	if id == "" {
		// return own member by default
		member, err := srv.store.GetMemberByUserId(ctx, printer, u.Id)
		if err != nil {
			return nil, err
		}
		return member.ToPb(), nil
	}
	member, err := srv.store.GetMember(ctx, printer, id)
	if err != nil {
		return nil, err
	}
	return member.ToPb(), nil
}

// ValidateMember validates the given member's fields. This function should be run
// before saving a member to the store.
func (srv *Server) ValidateMember(ctx context.Context, printer *message.Printer, member *clubv1.Member) error {
	if !utils.IsMailAddress(member.GetMail()) {
		return status.Errorf(codes.InvalidArgument, printer.Sprintf("%s is not a valid mail address", member.GetMail()))
	}
	if !utils.IsPhoneNumber(member.GetPhone()) {
		return status.Errorf(codes.InvalidArgument, printer.Sprintf("%s is not a valid phone number", member.GetPhone()))
	}
	if _, err := language.Parse(member.GetLanguage()); err != nil {
		return status.Errorf(codes.InvalidArgument, printer.Sprintf("invalid language '%s'", member.GetLanguage()))
	}
	// required for juristic & individual members
	if err := utils.RequiredFields(printer, member, "FirstName", "LastName", "Address", "PostalCode", "City"); err != nil {
		return err
	}
	if member.GetJuristic() {
		// required for juristic members
		if err := utils.RequiredFields(printer, member, "Organisation", "OrganisationType"); err != nil {
			return err
		}
		reTypes := regexp.MustCompile(`^(club|other)$`)
		if !reTypes.Match([]byte(member.GetOrganisationType())) {
			return status.Errorf(codes.InvalidArgument, "invalid organisation type: %s", member.GetOrganisationType())
		}
	} else {
		// required for individual members
		if err := utils.RequiredFields(printer, member, "DateOfBirth"); err != nil {
			return err
		}
		// date of birth
		ts, _ := ptypes.Timestamp(member.GetDateOfBirth())
		if time.Now().Before(ts) {
			return status.Errorf(codes.InvalidArgument, printer.Sprintf("date of birth cannot be in the future"))
		}
	}
	// check for duplicate values
	filter := fmt.Sprintf(`userId=="%s",username=="%s",mail=="%s"`, member.GetUserId(), member.GetUsername(), member.GetMail())
	if member.GetId() != "" {
		filter = fmt.Sprintf(`(%s);_id!=oid="%s"`, filter, member.GetId())
	}
	size, err := srv.store.CountMembers(ctx, printer, filter)
	if err != nil {
		return err
	}
	if size > 0 {
		return status.Errorf(codes.InvalidArgument, printer.Sprintf("either userId, username or mail address are already taken"))
	}
	if code, _ := status.FromError(err); err != nil && code.Code() != codes.NotFound {
		srv.errorLogger.Printf("error while looking for duplicate members: %s", err)
		return status.Errorf(codes.Internal, printer.Sprintf("error while looking for duplicates"))
	}
	return nil
}

// CreateMember creates a new member.
func (srv *Server) CreateMember(ctx context.Context, member *clubv1.Member) (*clubv1.Member, error) {
	printer := message.NewPrinter(language.Make(utils.LookupEnv("CLUB_DEFAULT_LANGUAGE", "en")))
	member.Id = ""
	member.UserId = ""
	if len(member.Password) < 7 {
		return nil, status.Errorf(codes.InvalidArgument, printer.Sprintf("password must have a length of at least 7"))
	}
	err := srv.ValidateMember(ctx, printer, member)
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, printer.Sprintf("the request was canceled by the client"))
	}
	c, err := srv.gooserProvider.NewGooserClient(srv.gooserTarget)
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, printer.Sprintf("gooser connection failed"))
	}
	defer srv.gooserProvider.Close()
	user, err := c.CreateUser(ctx, &gooserv1.User{
		Username: member.GetUsername(),
		Mail:     member.GetMail(),
		Language: member.GetLanguage(),
		Password: member.GetPassword(),
	})
	if err != nil {
		// pass through error
		return nil, err
	}
	member.UserId = user.GetId()
	created, err := srv.store.SaveMember(ctx, printer, store.PbToMember(member))
	if err != nil {
		return nil, err
	}
	return created.ToPb(), nil
}

// UpdateMember updates the given member in the store.
func (srv *Server) UpdateMember(ctx context.Context, request *clubv1.UpdateMemberRequest) (*clubv1.Member, error) {
	u, err := srv.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	printer := message.NewPrinter(language.Make(u.GetLanguage()))
	isAdmin := userHasAnyOfRoles(u, "admin")
	isMemberAdmin := userHasAnyOfRoles(u, "admin", "member-admin")
	member := request.GetMember()
	id := member.GetId()
	var existing *store.Member
	if id == "" {
		// use own member by default
		existing, err = srv.store.GetMemberByUserId(ctx, printer, u.Id)
		if err != nil {
			return nil, err
		}
	} else {
		existing, err = srv.store.GetMember(ctx, printer, id)
		if err != nil {
			return nil, err
		}
	}
	// if changing member other than own as non-admin
	if existing.UserId != u.Id && !isMemberAdmin {
		return nil, status.Errorf(codes.PermissionDenied, printer.Sprintf("permission denied"))
	}
	mask, err := fieldmaskUtils.MaskFromProtoFieldMask(request.GetFieldMask(), generator.CamelCase)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, printer.Sprintf("invalid field mask: %s", err))
	}

	// copy request document to existing document with field mask applied
	merged := existing.ToPb()
	err = fieldmaskUtils.StructToStruct(mask, member, merged)
	if err != nil {
		srv.errorLogger.Printf("unable to merge documents: %s", err)
		return nil, status.Errorf(codes.Internal, printer.Sprintf("unable to merge"))
	}
	// check if userId was changed
	if _, ok := mask.Get("UserId"); ok && !isAdmin && existing.UserId != merged.UserId {
		return nil, status.Errorf(codes.PermissionDenied, printer.Sprintf("you are not allowed to change the user id"))
	}
	// if password was changed
	if _, ok := mask.Get("Password"); ok {
		return nil, status.Errorf(codes.InvalidArgument, printer.Sprintf("password cannot be changed using the UpdateMember function, use ChangePassword instead"))
	}
	// check if mail was changed
	// by a non-admin for another user
	// this would be blocked by gooser
	if _, ok := mask.Get("Mail"); ok &&
		existing.Mail != merged.Mail &&
		!isAdmin &&
		existing.UserId != u.Id {
		return nil, status.Errorf(codes.PermissionDenied, printer.Sprintf("you are not allowed to change the mail address for other users"))
	}
	// validate
	if err := srv.ValidateMember(ctx, printer, merged); err != nil {
		return nil, err
	}
	var paths []string
	if _, ok := mask.Get("Language"); ok {
		paths = append(paths, "Language")
	}
	if _, ok := mask.Get("Mail"); ok {
		paths = append(paths, "Mail")
	}
	if _, ok := mask.Get("Username"); ok {
		paths = append(paths, "Username")
	}
	// update user if necessary
	if len(paths) > 0 {
		if ctx.Err() == context.Canceled {
			return nil, status.Errorf(codes.Canceled, printer.Sprintf("the request was canceled by the client"))
		}
		c, err := srv.gooserProvider.NewGooserClient(srv.gooserTarget)
		if err != nil {
			srv.errorLogger.Printf("unable to connect to gooser: %s", err)
			return nil, status.Errorf(codes.Internal, printer.Sprintf("gooser connection failed"))
		}
		defer srv.gooserProvider.Close()
		user, err := c.GetUser(ctx, &gooserv1.IdRequest{Id: member.UserId})
		if err != nil {
			return nil, err
		}
		user.Mail = merged.Mail
		user.Language = merged.Language
		user.Username = merged.Username
		// update gooser user
		if _, err := c.UpdateUser(ctx, &gooserv1.UpdateUserRequest{
			User:      user,
			FieldMask: &field_mask.FieldMask{Paths: paths},
		}); err != nil {
			// if anything goes wrong, pass error to requester
			return nil, err
		}
	}
	updated, err := srv.store.SaveMember(ctx, printer, store.PbToMember(merged))
	if err != nil {
		return nil, err
	}
	return updated.ToPb(), nil
}

// DeleteMember deletes the given member.
func (srv *Server) DeleteMember(ctx context.Context, request *clubv1.IdRequest) (*empty.Empty, error) {
	u, err := srv.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	printer := message.NewPrinter(language.Make(u.GetLanguage()))
	if !userHasAnyOfRoles(u, "admin") {
		return nil, status.Errorf(codes.PermissionDenied, printer.Sprintf("permission denied"))
	}
	member, err := srv.store.GetMember(ctx, printer, request.GetId())
	if err != nil {
		return nil, err
	}
	c, err := srv.gooserProvider.NewGooserClient(srv.gooserTarget)
	if err != nil {
		srv.errorLogger.Printf("unable to connect to gooser: %s", err)
		return nil, status.Errorf(codes.Internal, printer.Sprintf("gooser connection failed"))
	}
	defer srv.gooserProvider.Close()
	if _, err := c.DeleteUser(ctx, &gooserv1.IdRequest{Id: member.UserId}); err != nil {
		return nil, err
	}
	if err := srv.store.DeleteMember(ctx, printer, request.GetId()); err != nil {
		srv.errorLogger.Printf("user with id %s was deleted but an error occurred while deleting member %s with id %s", member.UserId, member.Username, member.Id)
		return nil, err
	}
	return &empty.Empty{}, nil
}
