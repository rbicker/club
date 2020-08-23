package server

import (
	"context"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/rbicker/club/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	clubv1 "github.com/rbicker/club/api/proto/v1"
)

// Contact is used by visitors to send a message to the website owner.
func (srv *Server) Contact(ctx context.Context, request *clubv1.ContactRequest) (*empty.Empty, error) {
	printer := message.NewPrinter(language.Make(request.GetLanguage()))
	if err := utils.RequiredFields(printer, request, "Name", "Subject", "Mail", "Message"); err != nil {
		return nil, err
	}
	if !utils.IsMailAddress(request.GetMail()) {
		return nil, status.Errorf(codes.InvalidArgument, printer.Sprintf("%s is not a valid mail address", request.GetMail()))
	}
	err := srv.messenger.Contact(request.Name, request.Subject, request.Mail, request.Message, request.Language)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
