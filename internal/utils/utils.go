package utils

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strings"

	"golang.org/x/text/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LookupEnv loads the setting with the given name
// from the environment. If no environment variable
// with the given name is set, the defaultValue is
// returned.
func LookupEnv(name, defaultValue string) string {
	if v, ok := os.LookupEnv(name); ok {
		return v
	}
	return defaultValue
}

// RandomString is used to generate a random string with n letters.
func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// IsMailAddress checks if the given string is a s address.
func IsMailAddress(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(s)
}

// IsPhoneNumber validates the given string as a phone number.
func IsPhoneNumber(s string) bool {
	re := regexp.MustCompile(`^\+?\d(\d|\s)+$`)
	return re.MatchString(s)
}

// RequiredFields checks if all the required fields are defined on the given object.
func RequiredFields(printer *message.Printer, obj interface{}, fields ...string) error {
	v := reflect.ValueOf(obj)
	if !v.IsValid() {
		return status.Errorf(codes.Internal, printer.Sprintf("unexpected value while looking for required fields"))
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return status.Errorf(codes.Internal, printer.Sprintf("unexpected value while looking for required fields"))
	}
	var invalid []string
	for _, s := range fields {
		// get field and check if it is valid / is not zero
		if f := v.FieldByName(s); !f.IsValid() || f.IsZero() {
			invalid = append(invalid, s)
		}
	}
	if len(invalid) > 0 {
		l := strings.Join(invalid, ", ")
		return status.Errorf(codes.InvalidArgument, printer.Sprintf("required field(s) missing: %s", l))
	}
	return nil
}

// Encrypt encrypts the given message using the given key.
func Encrypt(key, msg string) (string, error) {
	plain := []byte(msg)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	// iv needs to be unique, but doesn't have to be secure.
	// it's common to put it at the beginning of the cipher text.
	cypher := make([]byte, aes.BlockSize+len(plain))
	iv := cypher[:aes.BlockSize]
	if _, err = io.ReadFull(crand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cypher[aes.BlockSize:], plain)
	// return base64 encoded string
	res := base64.URLEncoding.EncodeToString(cypher)
	return res, nil
}

// Decrypt decrypts the given message the given key.
func Decrypt(key, msg string) (string, error) {
	cipherText, err := base64.URLEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("cipher text block size is to short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}
