package applestrings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Test Data - Invalid comment
 */

var errorReader1 = strings.NewReader(`/* Insert menu item
 "Insert" = "Insert";
 /* Error string used for unknown error types. */
 "ErrorString_1" = "An unknown error occurred.";`)

type emptyStruct struct{}

func TestDecoder_Decode_Error_InvalidComment(t *testing.T) {
	decoder := NewDecoder(errorReader1)
	var target emptyStruct
	err := decoder.Decode(&target)
	assert.Equal(t, ErrInvalidSyntax, err)
	assert.Empty(t, target)
}

/*
 * Test Data - Invalid key-value syntax
 */

var errorReader2 = strings.NewReader(`"Insert" = "Insert"`)
var errorReader3 = strings.NewReader(`"Insert" = "Insert" = "Insert";`)

func TestDecoder_Decode_Error_InvalidKeyValueSyntax(t *testing.T) {
	decoder := NewDecoder(errorReader2)
	var target emptyStruct
	err := decoder.Decode(&target)
	assert.Equal(t, ErrInvalidSyntax, err)
	assert.Empty(t, target)

	decoder = NewDecoder(errorReader3)
	err = decoder.Decode(&target)
	assert.Equal(t, ErrInvalidSyntax, err)
	assert.Empty(t, target)
}

var successReader1 = strings.NewReader(`"Insert" = "Insert";`)

type successStruct1 struct {
	Insert string
}

func TestDecoder_Decode_NonPointer(t *testing.T) {
	successReader1.Seek(0, 0)
	decoder := NewDecoder(successReader1)
	var target successStruct1
	err := decoder.Decode(target)
	assert.Nil(t, err)
	assert.Empty(t, target)

	successReader1.Seek(0, 0)
	err = decoder.Decode(&target)
	assert.Nil(t, err)
	assert.NotEmpty(t, target)
}

func TestDecoder_Decode_NonStruct(t *testing.T) {
	successReader1.Seek(0, 0)
	decoder := NewDecoder(successReader1)
	var target int
	err := decoder.Decode(&target)
	assert.Nil(t, err)
	assert.Empty(t, target)
}

type errorStruct1 struct {
	Insert int
}

func TestDecoder_Decode_StructFieldNotString(t *testing.T) {
	successReader1.Seek(0, 0)
	decoder := NewDecoder(successReader1)
	var target errorStruct1
	err := decoder.Decode(&target)
	assert.Nil(t, err)
	assert.Empty(t, target)
}
