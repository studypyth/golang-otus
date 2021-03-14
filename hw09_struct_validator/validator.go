package hw09

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const validateTag = "validate"

const ruleStrLen = "len"
const ruleStrRegexp = "regexp"
const ruleStrIn = "in"

const ruleIntMin = "min"
const ruleIntMax = "max"
const ruleIntIn = "in"

// ErrNotStruct is raise when trying to validate not a struct.
var ErrNotStruct = errors.New("type of given value is not Struct")

// ErrNotSupported is raise when trying to validate value of not supported type.
var ErrNotSupported = errors.New("validation is not supported to value of given type")

// ErrBadValidationTag is raise when incorrect validation tag found.
var ErrBadValidationTag = errors.New("incorrect validation tag")

// ErrNotApplyable is raise when trying to validate value with not applyable tag.
var ErrNotApplyable = errors.New("validation rule not applyable to value of given type")

// ErrBadLength is raise when string has incorrect length.
var ErrBadLength = errors.New("bad string length")

// ErrNotMatchRegexp is raise when string not match to regexp.
var ErrNotMatchRegexp = errors.New("do not match with regexp")

// ErrNotMatchIn is raise when string not match with given list.
var ErrNotMatchIn = errors.New("do not match with list")

// ErrMinVal is raise when int less than min value.
var ErrMinVal = errors.New("int less than min")

// ErrMaxVal is raise when int lesgreaters than max value.
var ErrMaxVal = errors.New("int greater than min")

// ValidationError represents validation error.
type ValidationError struct {
	Field string
	Err   error
}

// ValidationErrors represents bunch of validation errors.
type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	return fmt.Sprintf("validation failed: %d errors", len(v))
}

// Validate validate struct based on field's tags.
func Validate(v interface{}) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("validation value of type %T failed: %w", v, ErrNotStruct)
	}

	var validationErrors ValidationErrors
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		rules, err := parseRules(field.Tag)
		if err != nil {
			return fmt.Errorf("parsing validation rules failed: %w", err)
		}

		if len(rules) == 0 {
			continue
		}

		fValue := rv.Field(i)
		switch fValue.Kind() { // nolint: exhaustive
		case reflect.Slice:
			for i := 0; i < fValue.Len(); i++ {
				el := fValue.Index(i)
				err := validateScalar(el, rules)
				if err != nil {
					if isProgramError(err) {
						return fmt.Errorf("validation of field %s failed: %w", field.Name, err)
					}
					validationErrors = append(
						validationErrors,
						ValidationError{Field: field.Name, Err: fmt.Errorf("validation for element %d failed: %w", i, err)},
					)
				}
			}
		case reflect.String, reflect.Int:
			err := validateScalar(fValue, rules)
			if err != nil {
				if isProgramError(err) {
					return fmt.Errorf("validation of field %s failed: %w", field.Name, err)
				}
				validationErrors = append(
					validationErrors,
					ValidationError{Field: field.Name, Err: fmt.Errorf("validation failed: %w", err)},
				)
			}
		default:
			return fmt.Errorf("validation for value of type %s failed: %w", fValue.Kind(), ErrNotSupported)
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

type rule struct {
	name        string
	constraints string
}

func parseRules(sTag reflect.StructTag) ([]rule, error) {
	rules := make([]rule, 0)

	vTag := sTag.Get(validateTag)
	if vTag == "" {
		return rules, nil
	}

	for _, t := range strings.Split(vTag, "|") {
		v := strings.Split(t, ":")
		if len(v) != 2 {
			return nil, fmt.Errorf("parsing validation tag %s failed: %w", t, ErrBadValidationTag)
		}

		rules = append(rules, rule{name: v[0], constraints: v[1]})
	}

	return rules, nil
}

func isProgramError(e error) bool {
	if errors.Is(e, ErrNotStruct) || errors.Is(e, ErrNotSupported) || errors.Is(e, ErrNotApplyable) {
		return true
	}
	return false
}

func validateScalar(v reflect.Value, rules []rule) error {
	var err error
	switch v.Kind() { // nolint: exhaustive
	case reflect.String:
		err = validateString(v.String(), rules)
	case reflect.Int:
		err = validateInt(v.Int(), rules)
	default:
		err = fmt.Errorf("validation for value of type %s failed: %w", v.Kind(), ErrNotSupported)
	}
	return err
}

func validateString(v string, rules []rule) error {
	for _, r := range rules {
		switch r.name {
		case ruleStrLen:
			l, err := strconv.Atoi(r.constraints)
			if err != nil {
				return fmt.Errorf("validation of %s failed: %w", v, ErrBadValidationTag)
			}

			if len(v) != l {
				return fmt.Errorf("validation of %s failed: %w", v, ErrBadLength)
			}
		case ruleStrRegexp:
			re, err := regexp.Compile(r.constraints)
			if err != nil {
				return fmt.Errorf("validation of %s failed: %w", v, ErrBadValidationTag)
			}

			if !re.MatchString(v) {
				return fmt.Errorf("validation of %s failed: %w", v, ErrNotMatchRegexp)
			}
		case ruleStrIn:
			in := strings.Split(r.constraints, ",")

			flag := false
			for _, s := range in {
				if s == v {
					flag = true
					break
				}
			}

			if !flag {
				return fmt.Errorf("validation of %s failed: %w", v, ErrNotMatchIn)
			}
		default:
			return fmt.Errorf("validation of %s failed: %w", v, ErrNotApplyable)
		}
	}
	return nil
}

func validateInt(v int64, rules []rule) error {
	for _, r := range rules {
		switch r.name {
		case ruleIntMin:
			min, err := strconv.Atoi(r.constraints)
			if err != nil {
				return fmt.Errorf("validation of %d failed: %w", v, ErrBadValidationTag)
			}

			if v < int64(min) {
				return fmt.Errorf("validation of %d failed: %w", v, ErrMinVal)
			}
		case ruleIntMax:
			max, err := strconv.Atoi(r.constraints)
			if err != nil {
				return fmt.Errorf("validation of %d failed: %w", v, ErrBadValidationTag)
			}

			if v > int64(max) {
				return fmt.Errorf("validation of %d failed: %w", v, ErrMaxVal)
			}
		case ruleIntIn:
			var in []int
			for _, s := range strings.Split(r.constraints, ",") {
				val, err := strconv.Atoi(s)
				if err != nil {
					return fmt.Errorf("validation of %d failed: %w", v, ErrNotMatchIn)
				}

				in = append(in, val)
			}

			for _, s := range in {
				if int64(s) == v {
					break
				}
			}

			return fmt.Errorf("validation of %d failed: %w", v, ErrNotMatchIn)
		default:
			return fmt.Errorf("validation of %d failed: %w", v, ErrNotApplyable)
		}
	}
	return nil
}
