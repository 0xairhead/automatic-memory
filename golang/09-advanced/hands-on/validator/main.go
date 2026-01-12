package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ValidationError represents an error in validation
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation failed for '%s': %s", e.Field, e.Reason)
}

// Validator validates structs based on "validate" tags
// Supported: "required", "min=X", "max=X" (for strings/ints)
func Validate(s interface{}) []error {
	var errs []error

	val := reflect.ValueOf(s) // Value
	typ := reflect.TypeOf(s)  // Type

	if val.Kind() != reflect.Struct {
		return []error{fmt.Errorf("input must be a struct")}
	}

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")

		if tag == "" {
			continue
		}

		rules := strings.Split(tag, ",")
		for _, rule := range rules {
			if rule == "required" {
				if isZero(fieldVal) {
					errs = append(errs, &ValidationError{Field: fieldType.Name, Reason: "is required"})
				}
			} else if strings.HasPrefix(rule, "min=") {
				minVal, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
				if !checkMin(fieldVal, minVal) {
					errs = append(errs, &ValidationError{Field: fieldType.Name, Reason: fmt.Sprintf("min value/len is %d", minVal)})
				}
			}
		}
	}
	return errs
}

func isZero(v reflect.Value) bool {
	return v.IsZero()
}

func checkMin(v reflect.Value, min int) bool {
	switch v.Kind() {
	case reflect.String:
		return len(v.String()) >= min
	case reflect.Int, reflect.Int64:
		return int(v.Int()) >= min
	default:
		return true // Skip unsupported
	}
}

type UserRequest struct {
	Username string `validate:"required,min=5"`
	Age      int    `validate:"min=18"`
	Email    string `validate:"required"`
}

func main() {
	// Case 1: Invalid User
	req1 := UserRequest{
		Username: "bob", // Too short
		Age:      16,    // Too young
		Email:    "",    // Missing
	}

	fmt.Println("--- Validating Req 1 ---")
	errs := Validate(req1)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println("❌", err)
		}
	} else {
		fmt.Println("✅ Valid")
	}

	// Case 2: Valid User
	req2 := UserRequest{
		Username: "phoenix_user",
		Age:      20,
		Email:    "phoenix@example.com",
	}

	fmt.Println("\n--- Validating Req 2 ---")
	errs2 := Validate(req2)
	if len(errs2) > 0 {
		for _, err := range errs2 {
			fmt.Println("❌", err)
		}
	} else {
		fmt.Println("✅ Valid")
	}
}
