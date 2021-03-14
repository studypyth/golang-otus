package hw09

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Header struct {
		Name string `json:"header"`
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	BadValidationTag struct {
		Flag bool `validate:"1:2:3"`
	}

	NotSupportedType struct {
		Flag bool `validate:"len:5"`
	}

	NotSupportedSliceType struct {
		Flag []bool `validate:"len:5"`
	}

	ErrNotApplyableToStringTag struct {
		Name string `validate:"min:5"`
	}

	ErrNotApplyableToIntTag struct {
		Code int `validate:"regexp:\\s+"`
	}

	SomePhone struct {
		Phone string `validate:"len:11"`
	}

	AnotherPhones struct {
		Phones []string `validate:"len:11"`
	}

	SomeCode struct {
		Code string `validate:"regexp:^\\d+$"`
	}

	AnotherCodes struct {
		Codes []string `validate:"regexp:^\\d+$"`
	}

	SomeRole struct {
		Role string `validate:"in:admin,moderator"`
	}

	AnotherRoles struct {
		Roles []string `validate:"in:admin,moderator"`
	}

	SomeMinInt struct {
		Number int `validate:"min:1"`
	}

	AnotherMinInts struct {
		Numbers []int `validate:"min:1"`
	}

	SomeMaxInt struct {
		Number int `validate:"max:5"`
	}

	AnotherMaxInts struct {
		Numbers []int `validate:"max:5"`
	}

	SomeInInt struct {
		Number int `validate:"in:1,2,3"`
	}

	AnotherInInts struct {
		Numbers []int `validate:"in:1,2,3"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{UserRole("admin"), ErrNotStruct},
		{Token{}, nil},
		{Header{}, nil},
		{BadValidationTag{}, ErrBadValidationTag},
		{NotSupportedType{}, ErrNotSupported},
		{NotSupportedSliceType{Flag: []bool{true}}, ErrNotSupported},
		{ErrNotApplyableToStringTag{Name: "test"}, ErrNotApplyable},
		{
			SomePhone{Phone: "1234567890"},
			ValidationErrors{
				{Field: "Phone", Err: ErrBadLength},
			},
		},
		{
			AnotherPhones{Phones: []string{"1234567890"}},
			ValidationErrors{
				{Field: "Phones", Err: ErrBadLength},
			},
		},
		{
			SomeCode{Code: "123abc"},
			ValidationErrors{
				{Field: "Code", Err: ErrNotMatchRegexp},
			},
		},
		{
			AnotherCodes{Codes: []string{"123abc"}},
			ValidationErrors{
				{Field: "Codes", Err: ErrNotMatchRegexp},
			},
		},
		{
			SomeRole{Role: "staff"},
			ValidationErrors{
				{Field: "Role", Err: ErrNotMatchIn},
			},
		},
		{
			AnotherRoles{Roles: []string{"staff"}},
			ValidationErrors{
				{Field: "Roles", Err: ErrNotMatchIn},
			},
		},
		{
			ErrNotApplyableToIntTag{Code: 123},
			ErrNotApplyable,
		},
		{
			SomeMinInt{Number: 0},
			ValidationErrors{
				{Field: "Number", Err: ErrMinVal},
			},
		},
		{
			AnotherMinInts{Numbers: []int{0}},
			ValidationErrors{
				{Field: "Numbers", Err: ErrMinVal},
			},
		},
		{
			SomeMaxInt{Number: 10},
			ValidationErrors{
				{Field: "Number", Err: ErrMaxVal},
			},
		},
		{
			AnotherMaxInts{Numbers: []int{10}},
			ValidationErrors{
				{Field: "Numbers", Err: ErrMaxVal},
			},
		},
		{
			SomeInInt{Number: 5},
			ValidationErrors{
				{Field: "Number", Err: ErrNotMatchIn},
			},
		},
		{
			AnotherInInts{Numbers: []int{5}},
			ValidationErrors{
				{Field: "Numbers", Err: ErrNotMatchIn},
			},
		},
		{
			User{
				ID:     "123456789012345678901234567890123456",
				Age:    50,
				Email:  "nobody@nowhere.org",
				Role:   UserRole("stuff"),
				Phones: []string{"12345678901"},
			},
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := Validate(tt.in)
			if err == nil {
				assert.Equal(t, err, tt.expectedErr, "success")
			} else {
				if ve, ok := err.(ValidationErrors); ok {
					expVe := tt.expectedErr.(ValidationErrors)
					assert.True(t, len(ve) == len(expVe), "right amount of errors")
					for i, actual := range ve {
						expected := expVe[i]
						assert.Equal(t, actual.Field, expected.Field, "right field")
						assert.True(t, errors.Is(actual.Err, expected.Err), "right error")
					}
				} else {
					assert.True(t, errors.Is(err, tt.expectedErr), "right error")
				}
			}
		})
	}
}
