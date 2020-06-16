package server

import (
	"testing"

	gooserv1 "github.com/rbicker/gooser/api/proto/v1"
	"github.com/stretchr/testify/assert"
)

func TestUserHasAnyOfRoles(t *testing.T) {
	tests := []struct {
		name  string
		user  *gooserv1.User
		roles []string
		want  bool
	}{
		{
			name: "having role",
			user: &gooserv1.User{Roles: []string{
				"worker",
				"tester",
				"helper",
			}},
			roles: []string{
				"admin",
				"tester",
			},
			want: true,
		},
		{
			name: "not having role",
			user: &gooserv1.User{Roles: []string{
				"worker",
				"tester",
				"helper",
			}},
			roles: []string{
				"validator",
				"editor",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			got := userHasAnyOfRoles(tt.user, tt.roles...)
			assert.Equal(tt.want, got)
		})
	}
}
