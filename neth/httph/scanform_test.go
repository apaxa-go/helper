package httph

import (
	"errors"
	"github.com/apaxa-go/helper/mathh"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestScanError_Error(t *testing.T) {
	tests := []struct {
		err ScanError
		str string
	}{
		{ScanError{FieldNum: 0, FieldName: "f1", Type: ScanErrorTypeNoSuchField, SubError: nil}, "Scan error in #0 field with name 'f1': no field with such name"},
		{ScanError{FieldNum: 1, FieldName: "f2", Type: ScanErrorTypeMultipleValues, SubError: nil}, "Scan error in #1 field with name 'f2': there is more than 1 field with such name"},
		{ScanError{FieldNum: 2, FieldName: "f3", Type: ScanErrorTypeIncompatibleValue, SubError: nil}, "Scan error in #2 field with name 'f3': unable to parse string to required type"},
		{ScanError{FieldNum: 3, FieldName: "f4", Type: ScanErrorTypeIncompatibleValue, SubError: errors.New("sub error")}, "Scan error in #3 field with name 'f4': sub error"},
		{ScanError{FieldNum: 4, FieldName: "f5", Type: ScanErrorTypeIncompatibleType, SubError: nil}, "Scan error in #4 field with name 'f5': type of this field is imcompatible with this function type"},
		{ScanError{FieldNum: 5, FieldName: "f6", Type: ScanErrorType(mathh.MaxUint8), SubError: nil}, "Scan error in #5 field with name 'f6': unknown error"},
	}

	for _, test := range tests {
		if r := test.err.Error(); r != test.str {
			t.Errorf("Expected error string is '%v', but got '%v'", test.str, r)
		}
	}
}

func TestScanFormData(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.Form = url.Values{}
	req.Form.Add("f1", "1")
	req.Form.Add("f2", "2")
	req.Form.Add("f3", "3")
	req.Form.Add("f4", "4")
	req.Form.Add("f5", "5")
	req.Form.Add("f6", "6")
	req.Form.Add("f7", "7")
	req.Form.Add("f8", "8")
	req.Form.Add("f9", "9")
	req.Form.Add("f10", "10")
	req.Form.Add("f20", "String")
	req.Form.Add("f30", "on")
	req.Form.Add("f31", "off")
	req.Form.Add("fm", "v1")
	req.Form.Add("fm", "v2")

	var (
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64
		u   uint
		u8  uint8
		u16 uint16
		u32 uint32
		u64 uint64
		s   string
		b1  bool
		b2  bool
	)

	// Test all types in positive case
	fields := []ScanField{
		{Name: "f1", Value: &i},
		{Name: "f2", Value: &i8},
		{Name: "f3", Value: &i16},
		{Name: "f4", Value: &i32},
		{Name: "f5", Value: &i64},
		{Name: "f6", Value: &u},
		{Name: "f7", Value: &u8},
		{Name: "f8", Value: &u16},
		{Name: "f9", Value: &u32},
		{Name: "f10", Value: &u64},
		{Name: "f20", Value: &s},
		{Name: "f30", Value: &b1},
		{Name: "f31", Value: &b2},
	}
	if err := ScanFormData(req, fields...); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if i != 1 || i8 != 2 || i16 != 3 || i32 != 4 || i64 != 5 || u != 6 || u8 != 7 || u16 != 8 || u32 != 9 || u64 != 10 || s != "String" || !b1 || b2 {
		t.Error("Invalid values:", i, i8, i16, i32, i64, u, u8, u16, u32, u64, s, b1, b2)
	}

	// Test bad value for bool
	if err := ScanFormData(req, ScanField{Name: "f1", Value: &b1}); err == nil {
		t.Error("Expected error, but got nil")
	} else if err1, ok := err.(ScanError); !ok {
		t.Errorf("Unxepected type of error: %v", err1)
	} else if err1.Type != ScanErrorTypeIncompatibleValue || err1.FieldName != "f1" || err1.FieldNum != 0 || err1.SubError == nil {
		t.Errorf("Invalid fields in error: %v", err1)
	}

	// Test bad value for int
	for _, v := range []interface{}{&i, &i8, &i16, &i32, &i64, &u, &u8, &u16, &u32, &u64} {
		if err := ScanFormData(req, ScanField{Name: "f20", Value: v}); err == nil {
			t.Error("Expected error, but got nil")
		} else if err1, ok := err.(ScanError); !ok {
			t.Errorf("Unxepected type of error: %v", err1)
		} else if err1.Type != ScanErrorTypeIncompatibleValue || err1.FieldName != "f20" || err1.FieldNum != 0 || err1.SubError == nil {
			t.Errorf("Invalid fields in error: %v", err1)
		}
	}

	// Test unknown type
	var tmp os.File
	if err := ScanFormData(req, ScanField{Name: "f1", Value: &tmp}); err == nil {
		t.Error("Expected error, but got nil")
	} else if err1, ok := err.(ScanError); !ok {
		t.Errorf("Unxepected type of error: %v", err1)
	} else if err1.Type != ScanErrorTypeIncompatibleType || err1.FieldName != "f1" || err1.FieldNum != 0 || err1.SubError != nil {
		t.Errorf("Invalid fields in error: %v", err1)
	}

	// Test non-exists field
	if err := ScanFormData(req, ScanField{Name: "f-1", Value: &i}); err == nil {
		t.Error("Expected error, but got nil")
	} else if err1, ok := err.(ScanError); !ok {
		t.Errorf("Unxepected type of error: %v", err1)
	} else if err1.Type != ScanErrorTypeNoSuchField || err1.FieldName != "f-1" || err1.FieldNum != 0 || err1.SubError != nil {
		t.Errorf("Invalid fields in error: %v", err1)
	}

	// Test field with multiple values
	if err := ScanFormData(req, ScanField{Name: "fm", Value: &s}); err == nil {
		t.Error("Expected error, but got nil")
	} else if err1, ok := err.(ScanError); !ok {
		t.Errorf("Unxepected type of error: %v", err1)
	} else if err1.Type != ScanErrorTypeMultipleValues || err1.FieldName != "fm" || err1.FieldNum != 0 || err1.SubError != nil {
		t.Errorf("Invalid fields in error: %v", err1)
	}
}
