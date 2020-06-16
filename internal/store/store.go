package store

import (
	"context"

	"golang.org/x/text/message"
)

// Store abstracts saving and receiving data.
type Store interface {
	ListMembers(ctx context.Context, printer *message.Printer, filterString, orderBy, token string, size int32) (members *[]Member, totalSize int32, nextToken string, err error)
	CountMembers(ctx context.Context, printer *message.Printer, filterString string) (int32, error)
	GetMember(ctx context.Context, printer *message.Printer, id string) (*Member, error)
	GetMemberByUserId(ctx context.Context, printer *message.Printer, userId string) (*Member, error)
	SaveMember(ctx context.Context, printer *message.Printer, member *Member) (*Member, error)
	DeleteMember(ctx context.Context, printer *message.Printer, id string) error
}
