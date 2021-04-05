package applestrings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Test Data 1
 * Quoted keys
 * No line break at the end
 */

var reader1 = strings.NewReader(`/* Insert menu item */
"Insert" = "Insert";
/* Error string used for unknown error types. */
"ErrorString_1" = "An unknown error occurred.";`)

type struct1 struct {
	Insert        string
	ErrorString_1 string
}

var expectedStruct1 = struct1{
	"Insert", "An unknown error occurred.",
}

func TestDecoder_Decode_1(t *testing.T) {
	decoder := NewDecoder(reader1)
	var target struct1
	err := decoder.Decode(&target)
	assert.Nil(t, err)
	assert.Equal(t, expectedStruct1, target)
}

/*
 * Test Data 2
 * Unquoted keys
 * Linebreak at the end
 */

// .strings file supported unquoted keys, exemplified by an Apple's documentation at
// https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/LocalizingYourApp/LocalizingYourApp.html
var reader2 = strings.NewReader(`CFBundleDisplayName = "Maisons";
NSHumanReadableCopyright = "Copyright © 2014 My Great Company Tous droits réservés.";
`)

type struct2 struct {
	CFBundleDisplayName      string
	NSHumanReadableCopyright string
}

var expectedStruct2 = struct2{
	"Maisons", "Copyright © 2014 My Great Company Tous droits réservés.",
}

func TestDecoder_Decode_2(t *testing.T) {
	decoder := NewDecoder(reader2)
	var target struct2
	err := decoder.Decode(&target)
	assert.Nil(t, err)
	assert.Equal(t, expectedStruct2, target)
}

/*
 * Test Data 3
 * Unquoted keys
 * No linebreak at the end
 */

var reader3 = strings.NewReader(`CFBundleDisplayName = "Maisons";
NSHumanReadableCopyright = "Copyright © 2014 My Great Company Tous droits réservés.";`)

type struct3 struct {
	CFBundleDisplayName      string
	NSHumanReadableCopyright string
}

var expectedStruct3 = struct3{
	"Maisons", "Copyright © 2014 My Great Company Tous droits réservés.",
}

func TestDecoder_Decode_3(t *testing.T) {
	decoder := NewDecoder(reader3)
	var target struct3
	err := decoder.Decode(&target)
	assert.Nil(t, err)
	assert.Equal(t, expectedStruct3, target)
}

/*
 * Test Data 4
 * Unquoted keys
 * No linebreak at the end
 * With empty lines & lines with only spaces
 */

var reader4 = strings.NewReader(`/* Comment */
CFBundleDisplayName = "Maisons";

NSHumanReadableCopyright = "Copyright © 2014 My Great Company Tous droits réservés.";
 
`)

type struct4 struct {
	CFBundleDisplayName      string
	NSHumanReadableCopyright string
}

var expectedStruct4 = struct4{
	"Maisons", "Copyright © 2014 My Great Company Tous droits réservés.",
}

func TestDecoder_Decode_4(t *testing.T) {
	decoder := NewDecoder(reader4)
	var target struct4
	err := decoder.Decode(&target)
	assert.Nil(t, err)
	assert.Equal(t, expectedStruct4, target)
}
