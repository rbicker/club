package store

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/rbicker/go-rsql"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBsonDocFromSortString(t *testing.T) {
	tests := []struct {
		name    string
		sort    string
		want    bson.D
		wantErr bool
	}{
		{
			name: "empty",
			sort: "",
			want: bson.D{},
		},
		{
			name: "simple",
			sort: "userId, -age",
			want: bson.D{
				{"userId", 1},
				{"age", -1},
			},
		},
		{
			name:    "invalid",
			sort:    ",-age",
			wantErr: true,
		},
	}
	p := message.NewPrinter(language.English)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			got, err := bsonDocFromOrderByString(p, tt.sort)
			if !tt.wantErr {
				assert.Nil(err, "error is not nil")
				assert.Equal(tt.want, got)
			} else {
				assert.Error(err)
			}
		})
	}
}

func TestPaginatedFilterBuilder(t *testing.T) {
	// idFilter is a filter that we expect to see
	// in some results, where it checks that the document
	// with the next id is being returned.
	idFilter := bson.E{
		Key: "_id",
		Value: bson.E{
			Key: "$oid",
			Value: bson.E{
				Key:   "$gt",
				Value: "x",
			},
		},
	}
	// inputFilter is a filter which we will send
	// to the function, to filter for the "count"-field.
	inputFilter := bson.D{
		{
			Key:   "Count",
			Value: 5,
		},
	}
	// sortPageFilter is the filter we expect to see in the
	// result when we sort for the "count"-field.
	sortPageFilter := bson.D{
		bson.E{
			Key: "$or",
			Value: bson.A{
				bson.D{
					bson.E{
						// next match
						Key: "count",
						Value: bson.E{
							Key:   "$gt",
							Value: 5,
						},
					},
				},
				bson.D{
					// exact match + greater id
					bson.E{
						Key:   "count",
						Value: 5,
					},
					idFilter,
				},
			},
		},
	}
	// inputObj is the struct on which
	// the filter creation will depend.
	inputObj := struct {
		Id    string `json:"_id,omitempty"`
		Count int    `json:"count"`
	}{
		Id:    "x",
		Count: 5,
	}
	tests := []struct {
		name    string
		filter  bson.D
		orderBy string
		obj     interface{}
		want    bson.D
		wantErr bool
	}{
		{
			name:    "no struct object",
			obj:     "x",
			wantErr: true,
		},
		{
			name:    "object without id",
			obj:     struct{}{},
			wantErr: true,
		},
		{
			name: "valid object",
			obj:  inputObj,
			want: bson.D{
				idFilter,
			},
		},
		{
			name:    "only id ascending",
			obj:     inputObj,
			orderBy: "id",
			want: bson.D{
				idFilter,
			},
		},
		{
			name:    "id and others",
			obj:     inputObj,
			orderBy: "id,-count",
			want: bson.D{
				bson.E{
					Key: "$or",
					Value: bson.A{
						// next filter
						bson.D{
							idFilter,
							bson.E{
								Key: "count",
								Value: bson.E{
									Key:   "$lt",
									Value: 5,
								},
							},
						},
						// exact filter
						bson.D{
							bson.E{
								Key:   "count",
								Value: 5,
							},
							idFilter,
						},
					},
				},
			},
		},
		{
			name:    "only id descending",
			obj:     inputObj,
			orderBy: "-id",
			want: bson.D{
				bson.E{
					Key: "_id",
					Value: bson.E{
						Key: "$oid",
						Value: bson.E{
							Key:   "$lt",
							Value: "x",
						},
					},
				},
			},
		},
		{
			name:   "filter, no sorting",
			obj:    inputObj,
			filter: inputFilter,
			want: bson.D{
				bson.E{
					Key: "$and",
					Value: bson.A{
						inputFilter,
						bson.D{
							idFilter,
						},
					},
				},
			},
		},
		{
			name:    "sorting, no filter",
			obj:     inputObj,
			orderBy: "Count",
			want:    sortPageFilter,
		},
		{
			name:    "sorting and filter",
			obj:     inputObj,
			filter:  inputFilter,
			orderBy: "Count",
			want: bson.D{
				bson.E{
					Key: "$and",
					Value: bson.A{
						inputFilter,
						sortPageFilter,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			parser, err := rsql.NewParser(rsql.Mongo())
			if err != nil {
				t.Fatalf("error while creating rsql parser: %s", err)
			}
			m := &MGO{
				rsqlParser:  parser,
				errorLogger: log.New(ioutil.Discard, "", 0),
			}
			p := message.NewPrinter(language.English)
			got, err := m.paginatedFilterBuilder(p, tt.filter, tt.orderBy, tt.obj)
			if !tt.wantErr {
				assert.Nil(err, "error is not nil")
				assert.Equal(tt.want, got)
			} else {
				assert.Error(err)
			}
		})
	}
}
