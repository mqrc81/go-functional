package check

import (
	"fmt"
	"testing"
)

var strict bool

var nilChan chan string
var nilFunc func()
var nilInterface interface{}
var nilMap map[bool]struct{}
var nilSlice []int

type isValidTests []struct {
	name  string
	value any
	valid bool
}

func Test_isValid_Default(t *testing.T) {
	strict = false
	tests := isValidTests{
		{
			name:  "int: zero",
			value: 0,
			valid: false,
		},
		{
			name:  "int: positive",
			value: 100,
			valid: true,
		},
		{
			name:  "bool: zero",
			value: false,
			valid: false,
		},
		{
			name:  "bool: true",
			value: true,
			valid: true,
		},
		{
			name:  "int64: zero",
			value: int64(0),
			valid: false,
		},
		{
			name:  "string: zero",
			value: "",
			valid: false,
		},
		{
			name:  "string: text",
			value: "sadasdasdasdasd",
			valid: true,
		},
		{
			name:  "uint: zero",
			value: uint(0),
			valid: false,
		},
		{
			name:  "uint32: positive",
			value: uint32(10000),
			valid: true,
		},
		{
			name:  "uintptr: zero",
			value: uintptr(0),
			valid: false,
		},
		{
			name:  "complex64: zero",
			value: complex(0, 0),
			valid: false,
		},
		{
			name:  "array: zero empty",
			value: [0]int{},
			valid: false,
		},
		{
			name:  "array: zero big",
			value: [1000]int{},
			valid: false,
		},
		{
			name:  "array: all zeroes",
			value: [2]int{0, 0},
			valid: false,
		},
		{
			name:  "array: normal",
			value: [2]bool{true, false},
			valid: true,
		},
		{
			name:  "chan: make empty",
			value: make(chan int),
			valid: true,
		},
		{
			name:  "chan: zero nil",
			value: nilChan,
			valid: false,
		},
		{
			name:  "chan: make empty size",
			value: make(chan int, 13),
			valid: true,
		},
		{
			name:  "func: zero",
			value: *new(func()),
			valid: false,
		},
		{
			name:  "func: zero nil",
			value: nilFunc,
			valid: false,
		},
		{
			name:  "func: normal",
			value: func(val int) int { return val * 3 },
			valid: true,
		},
		{
			name:  "interface: zero",
			value: *new(interface{}),
			valid: false,
		},
		{
			name:  "interface: nil",
			value: nilInterface,
			valid: false,
		},
		{
			name:  "nil",
			value: nil,
			valid: false,
		},
		{
			name:  "interface: empty",
			value: check[int]{strict: false},
			valid: false,
		},
		{
			name:  "interface: normal",
			value: check[struct{}]{strict: true},
			valid: true,
		},
		{
			name:  "map: normal empty",
			value: map[string]int{},
			valid: true,
		},
		{
			name:  "map: zero nil",
			value: nilMap,
			valid: false,
		},
		{
			name:  "map: make empty",
			value: make(map[uint]struct{}, 0),
			valid: true,
		},
		{
			name:  "pointer: new",
			value: new(int),
			valid: true,
		},
		{
			name:  "pointer: zero",
			value: *new(*int),
			valid: false,
		},
		{
			name:  "pointer: normal nil",
			value: &nilFunc,
			valid: true,
		},
		{
			name:  "slice: nil",
			value: nilSlice,
			valid: false,
		},
		{
			name:  "slice: normal empty",
			value: []string{},
			valid: true,
		},
		{
			name:  "slice: normal zeroes",
			value: []string{"", "", ""},
			valid: true,
		},
		{
			name:  "struct: zero empty",
			value: struct{}{},
			valid: false,
		},
		{
			name: "struct: zero empty complex",
			value: struct {
				name string
				m    map[int]bool
			}{},
			valid: false,
		},
		{
			name: "struct: normal",
			value: struct {
				name string
				m    map[int]bool
			}{name: "no"},
			valid: true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%v - %v", i, tt.name), func(t *testing.T) {
			if got := isValid(tt.value, strict); got != tt.valid {
				t.Errorf("isValid() = %v, want %v", got, tt.valid)
			}
		})
	}
}

func Test_isValid_Strict(t *testing.T) {
	strict = true
	tests := isValidTests{
		{
			name:  "int: zero",
			value: 0,
			valid: false,
		},
		{
			name:  "int: positive",
			value: 100,
			valid: true,
		},
		{
			name:  "bool: zero",
			value: false,
			valid: false,
		},
		{
			name:  "bool: true",
			value: true,
			valid: true,
		},
		{
			name:  "int64: zero",
			value: int64(0),
			valid: false,
		},
		{
			name:  "string: zero",
			value: "",
			valid: false,
		},
		{
			name:  "string: text",
			value: "sadasdasdasdasd",
			valid: true,
		},
		{
			name:  "uint: zero",
			value: uint(0),
			valid: false,
		},
		{
			name:  "uint32: positive",
			value: uint32(10000),
			valid: true,
		},
		{
			name:  "uintptr: zero",
			value: uintptr(0),
			valid: false,
		},
		{
			name:  "complex64: zero",
			value: complex(0, 0),
			valid: false,
		},
		{
			name:  "array: zero empty",
			value: [0]int{},
			valid: false,
		},
		{
			name:  "array: zero big",
			value: [1000]int{},
			valid: false,
		},
		{
			name:  "array: all zeroes",
			value: [2]int{0, 0},
			valid: false,
		},
		{
			name:  "array: normal",
			value: [2]bool{true, false},
			valid: true,
		},
		{
			name:  "chan: make empty",
			value: make(chan int),
			valid: false,
		},
		{
			name:  "chan: zero nil",
			value: nilChan,
			valid: false,
		},
		{
			name:  "chan: make empty size",
			value: make(chan int, 13),
			valid: false,
		},
		{
			name:  "func: zero",
			value: *new(func()),
			valid: false,
		},
		{
			name:  "func: zero nil",
			value: nilFunc,
			valid: false,
		},
		{
			name:  "func: normal",
			value: func(val int) int { return val * 3 },
			valid: true,
		},
		{
			name:  "interface: zero",
			value: *new(interface{}),
			valid: false,
		},
		{
			name:  "interface: nil",
			value: nilInterface,
			valid: false,
		},
		{
			name:  "nil",
			value: nil,
			valid: false,
		},
		{
			name:  "interface: empty",
			value: check[int]{strict: false},
			valid: false,
		},
		{
			name:  "interface: normal",
			value: check[struct{}]{strict: true},
			valid: true,
		},
		{
			name:  "map: normal empty",
			value: map[string]int{},
			valid: false,
		},
		{
			name:  "map: zero nil",
			value: nilMap,
			valid: false,
		},
		{
			name:  "map: make empty",
			value: make(map[uint]struct{}, 0),
			valid: false,
		},
		{
			name:  "pointer: new",
			value: new(int),
			valid: true,
		},
		{
			name:  "pointer: zero",
			value: *new(*int),
			valid: false,
		},
		{
			name:  "pointer: normal nil",
			value: &nilFunc,
			valid: true,
		},
		{
			name:  "slice: nil",
			value: nilSlice,
			valid: false,
		},
		{
			name:  "slice: normal empty",
			value: []string{},
			valid: false,
		},
		{
			name:  "slice: normal zeroes",
			value: []string{"", "", ""},
			valid: true,
		},
		{
			name:  "struct: zero empty",
			value: struct{}{},
			valid: false,
		},
		{
			name: "struct: zero empty complex",
			value: struct {
				name string
				m    map[int]bool
			}{},
			valid: false,
		},
		{
			name: "struct: normal",
			value: struct {
				name string
				m    map[int]bool
			}{name: "no"},
			valid: true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%v - %v", i, tt.name), func(t *testing.T) {
			if got := isValid(tt.value, strict); got != tt.valid {
				t.Errorf("isValid() = %v, want %v", got, tt.valid)
			}
		})
	}
}
