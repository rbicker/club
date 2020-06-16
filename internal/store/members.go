package store

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	clubv1 "github.com/rbicker/club/api/proto/v1"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/text/message"
)

// Member describes a club member.
type Member struct {
	Id               string    `bson:"_id,omitempty"`
	CreatedAt        time.Time `bson:"createdAt"`
	UpdatedAt        time.Time `bson:"updatedAt"`
	UserId           string    `bson:"userId"`
	Username         string    `bson:"username"`
	Mail             string    `bson:"mail"`
	Language         string    `bson:"language"`
	FirstName        string    `bson:"firstName"`
	LastName         string    `bson:"lastName"`
	DateOfBirth      time.Time `bson:"dateOfBirth"`
	Phone            string    `bson:"phone"`
	Address          string    `bson:"address"`
	Address2         string    `bson:"address2"`
	PostalCode       string    `bson:"postalCode"`
	City             string    `bson:"city"`
	Juristic         bool      `bson:"juristic"`
	Organisation     string    `bson:"organisation"`
	Website          string    `bson:"website"`
	OrganisationType string    `bson:"organisationType"`
	Inactive         bool      `bson:"inactive"`
}

// ToPb converts the store member to a protobuf member.
func (m *Member) ToPb() *clubv1.Member {
	createdAt, _ := ptypes.TimestampProto(m.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(m.UpdatedAt)
	dateOfBirth, _ := ptypes.TimestampProto(m.DateOfBirth)
	return &clubv1.Member{
		Id:               m.Id,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		UserId:           m.UserId,
		Username:         m.Username,
		Mail:             m.Mail,
		Language:         m.Language,
		FirstName:        m.FirstName,
		LastName:         m.LastName,
		DateOfBirth:      dateOfBirth,
		Phone:            m.Phone,
		Address:          m.Address,
		Address2:         m.Address2,
		PostalCode:       m.PostalCode,
		City:             m.City,
		Juristic:         m.Juristic,
		Organisation:     m.Organisation,
		Website:          m.Website,
		OrganisationType: m.OrganisationType,
		Inactive:         m.Inactive,
	}
}

// PbToMember converts the given protobuf member to a store member.
func PbToMember(m *clubv1.Member) *Member {
	createdAt, _ := ptypes.Timestamp(m.CreatedAt)
	updatedAt, _ := ptypes.Timestamp(m.UpdatedAt)
	dateOfBirth, _ := ptypes.Timestamp(m.DateOfBirth)
	return &Member{
		Id:               m.Id,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		UserId:           m.UserId,
		Username:         m.Username,
		Mail:             m.Mail,
		Language:         m.Language,
		FirstName:        m.FirstName,
		LastName:         m.LastName,
		DateOfBirth:      dateOfBirth,
		Phone:            m.Phone,
		Address:          m.Address,
		Address2:         m.Address2,
		PostalCode:       m.PostalCode,
		City:             m.City,
		Juristic:         m.Juristic,
		Organisation:     m.Organisation,
		Website:          m.Website,
		OrganisationType: m.OrganisationType,
		Inactive:         m.Inactive,
	}
}

// ListMembers lists members stored in mongodb.
func (m *MGO) ListMembers(ctx context.Context, printer *message.Printer, filterString, orderBy, token string, size int32) (members *[]Member, totalSize int32, nextToken string, err error) {
	cur, total, err := m.queryDocuments(
		ctx,
		printer,
		m.members,
		filterString,
		orderBy,
		token,
		size,
	)
	if err != nil {
		return nil, 0, "", err
	}
	defer cur.Close(ctx)
	var member Member
	for cur.Next(ctx) {
		err = cur.Decode(&member)
		if err != nil {
			m.errorLogger.Printf("unable to decode member: %s", err)
			return nil, 0, "", status.Errorf(codes.Internal, printer.Sprintf("unable to decode document"))
		}
		*members = append(*members, member)
	}
	// if there might be more results
	l := int32(len(*members))
	if size == l && totalSize > l {
		nextToken, err = m.NextPageToken(
			ctx,
			printer,
			m.members,
			filterString,
			orderBy,
			member,
		)
		if err != nil {
			return nil, 0, "", err
		}
	}
	return members, total, nextToken, err
}

// getMember is the internal function to query a member based on the given filter.
func (m *MGO) getMember(ctx context.Context, printer *message.Printer, filter *bson.M) (*Member, error) {
	member := &Member{}
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, printer.Sprintf("the request was canceled by the client"))
	}
	if err := m.members.FindOne(ctx, filter).Decode(member); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, printer.Sprintf("unable to find document"))
		}
		return nil, err
	}
	return member, nil
}

// CountMembers returns the number of user documents corresponding to the given filter.
func (m *MGO) CountMembers(ctx context.Context, printer *message.Printer, filterString string) (int32, error) {
	filter, err := m.bsonDocFromRsqlString(printer, filterString)
	if err != nil {
		return 0, err
	}
	count, err := m.members.CountDocuments(ctx, filter, nil)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		m.errorLogger.Printf("unable to count members: %s", err)
		return 0, status.Errorf(codes.Internal, printer.Sprintf("unable to count members"))
	}
	return int32(count), nil
}

// GetMember gets the member with the given id from mongodb.
func (m *MGO) GetMember(ctx context.Context, printer *message.Printer, id string) (*Member, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, printer.Sprintf("invalid id '%s'", id))
	}
	filter := bson.M{"_id": oid}
	return m.getMember(ctx, printer, &filter)
}

// GetMemberByUserId returns the member with the given user id.
func (m *MGO) GetMemberByUserId(ctx context.Context, printer *message.Printer, userId string) (*Member, error) {
	filter := bson.M{"userId": userId}
	return m.getMember(ctx, printer, &filter)
}

// SaveMember saves the given member to mongodb.
func (m *MGO) SaveMember(ctx context.Context, printer *message.Printer, member *Member) (*Member, error) {
	var err error
	var oid primitive.ObjectID
	member.UpdatedAt = time.Now()
	if member.Id != "" {
		oid, err = primitive.ObjectIDFromHex(member.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, printer.Sprintf("invalid id '%s'", member.Id))
		}
	} else {
		oid = primitive.NewObjectID()
		member.CreatedAt = member.UpdatedAt
	}
	opts := options.FindOneAndUpdate()
	opts.SetUpsert(true)
	opts.SetReturnDocument(options.After)
	filter := bson.M{"_id": oid}
	doc := bson.M{"$set": member}
	updated := &Member{}
	err = m.members.FindOneAndUpdate(ctx, filter, doc, opts).Decode(updated)
	if err != nil {
		m.errorLogger.Printf("error while saving member: %s", err)
		return nil, status.Errorf(codes.Internal, printer.Sprintf("error while saving"))
	}
	return updated, nil
}

// DeleteMember deletes the member with the given id from mongodb.
func (m *MGO) DeleteMember(ctx context.Context, printer *message.Printer, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, printer.Sprintf("invalid id"))
	}
	filter := bson.M{"_id": oid}
	if ctx.Err() == context.Canceled {
		return status.Errorf(codes.Canceled, printer.Sprintf("the request was canceled by the client"))
	}
	res, err := m.members.DeleteOne(ctx, filter)
	if err != nil {
		return status.Errorf(codes.Internal, "unable to delete document")
	}
	if res.DeletedCount != 1 {
		return status.Errorf(codes.InvalidArgument, printer.Sprintf("unable to find %s document with id %s", m.members.Name(), id))
	}
	return nil
}
