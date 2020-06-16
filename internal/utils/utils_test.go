package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want string
	}{
		{
			name: "simple",
			msg:  "I am secret",
		},
	}
	h := md5.New()
	plainKey := "thisisthesecretkeythatweareusingwhiletestingifthatisok?"
	if _, err := io.WriteString(h, plainKey); err != nil {
		t.Fatalf("unable to hash secret key: %s", err)
	}
	key := fmt.Sprintf("%x", h.Sum(nil))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			encrypted, _ := Encrypt(key, tt.msg)
			assert.NotEqual(tt.msg, encrypted)
			decrypted, _ := Decrypt(key, encrypted)
			assert.Equal(tt.msg, decrypted)
		})
	}
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		name       string
		n          int
		wantLength int
	}{
		{
			name:       "0",
			n:          0,
			wantLength: 0,
		},
		{
			name:       "1",
			n:          1,
			wantLength: 1,
		},
		{
			name:       "5",
			n:          5,
			wantLength: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			got := RandomString(tt.n)
			assert.Equal(tt.wantLength, len(got))
		})
	}
}

func TestIsMailAddress(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "test@example.com",
			want: true,
		},
		{
			s:    "test@example_abc.com",
			want: false,
		},
		{
			s:    "test@example-abc.com",
			want: true,
		},
		{
			s:    "test@example",
			want: true,
		},
		{
			s:    "test@example@example.com",
			want: false,
		},
	}
	for _, tt := range tests {
		assert := assert.New(t)
		got := IsMailAddress(tt.s)
		assert.Equal(tt.want, got, tt.s)
	}
}

func TestIsPhoneNumber(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "076 123 45 67",
			want: true,
		},
		{
			s:    "+41 76 123 45 67",
			want: true,
		},
		{
			s:    "+41 76a 123 45 67",
			want: false,
		},
		{
			s:    "+41 76 123 +45 67",
			want: false,
		},
	}
	for _, tt := range tests {
		assert := assert.New(t)
		got := IsPhoneNumber(tt.s)
		assert.Equal(tt.want, got, tt.s)
	}
}

func TestRequiredFields(t *testing.T) {
	tests := []struct {
		name    string
		obj     interface{}
		fields  []string
		wantErr bool
	}{
		{
			name: "valid",
			obj: struct {
				Name    string
				Counter int
			}{
				Name:    "Peter",
				Counter: 3,
			},
			fields: []string{
				"Name",
				"Counter",
			},
		},
		{
			name: "invalid",
			obj: struct {
				Name    string
				Counter int
			}{
				Counter: 3,
			},
			fields: []string{
				"Name",
				"Counter",
				"Length",
			},
			wantErr: true,
		},
	}
	p := message.NewPrinter(language.English)
	for _, tt := range tests {
		assert := assert.New(t)
		err := RequiredFields(p, tt.obj, tt.fields...)
		if tt.wantErr {
			assert.NotNil(err, "expected error")
		} else {
			assert.Nil(err, "did not expect error")
		}
	}
}
