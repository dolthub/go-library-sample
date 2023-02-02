package main

/*
#include <stdint.h>
struct DemoStruct {
    uint8_t A;
    int32_t B;
};
*/
import "C"
import (
	"math"
	"strings"
)

func main() {}

//export ModifyBool
func ModifyBool(input bool) bool {
	// This flips the input
	// For reference, the C code will return an "uint8"
	return !input
}

//export ModifyInt
func ModifyInt(input int) int {
	// This applies a bitwise NOT to the input
	return ^input
}

//export ModifyInt8
func ModifyInt8(input int8) int8 {
	// This applies a bitwise NOT to the input
	return ^input
}

//export ModifyInt16
func ModifyInt16(input int16) int16 {
	// This applies a bitwise NOT to the input
	return ^input
}

//export ModifyInt32
func ModifyInt32(input int32) int32 {
	// This applies a bitwise NOT to the input
	// For reference, "rune" is an alias for "int32"
	return ^input
}

//export ModifyInt64
func ModifyInt64(input int64) int64 {
	// This applies a bitwise NOT to the input
	return ^input
}

//export ModifyUint
func ModifyUint(input uint) uint {
	// This applies a bitwise NOT to the input
	return ^input
}

//export ModifyUint8
func ModifyUint8(input uint8) uint8 {
	// This applies a bitwise NOT to the input
	// For reference, "byte" is an alias for "uint8"
	return ^input
}

//export ModifyUint16
func ModifyUint16(input uint16) uint16 {
	// This applies a bitwise NOT to the input
	return ^input
}

//export ModifyUint32
func ModifyUint32(input uint32) uint32 {
	// This applies a bitwise NOT to the input
	return ^input
}

//export ModifyUint64
func ModifyUint64(input uint64) uint64 {
	// This applies a bitwise NOT to the input
	return ^input
}

//export ModifyFloat32
func ModifyFloat32(input float32) float32 {
	// This applies a bitwise NOT to the bits of the input
	return math.Float32frombits(^math.Float32bits(input))
}

//export ModifyFloat64
func ModifyFloat64(input float64) float64 {
	// This applies a bitwise NOT to the bits of the input
	return math.Float64frombits(^math.Float64bits(input))
}

//export ModifyString
func ModifyString(input string) *C.char {
	// This lowercases and uppercases each half of the input depending on the length of the input. The new string must
	// be freed in the calling C code.
	inputLen := len(input)
	var output string
	if inputLen % 2 == 0 {
		output = strings.ToLower(input[:inputLen/2]) + strings.ToUpper(input[inputLen/2:])
	} else {
		output = strings.ToUpper(input[:inputLen/2]) + strings.ToLower(input[inputLen/2:])
	}
	return C.CString(output)
}

//export ModifyInt32Slice
func ModifyInt32Slice(input []int32) (*int32, int) {
	// This modifies the input while returning a new C-compatible array. The new array must be freed in the calling C
	// code.
	output := make([]int32, len(input))
	for i := range input {
		input[i] = ^input[i]
		output[i] = int32(i + 100)
	}
	return ToCSlice(output)
}

//export ModifyDemoStruct
func ModifyDemoStruct(input1 *C.struct_DemoStruct, input2 C.struct_DemoStruct) C.struct_DemoStruct {
	// This modifies the first input while returning a new DemoStruct based on the second input. The struct type is
	// declared in C, as there is not yet a way to pass standard Go structs.
	input1.A = ^input1.A
	input1.B = ^input1.B
	return C.struct_DemoStruct{
		A: (C.uint8_t)(^input2.A),
		B: (C.int32_t)(^input2.B),
	}
}

// Remember to close channels once you're done with them. This is being ignored for the sake of demonstration.
var demoChannel = make(chan int32)

//export StartChannel
func StartChannel() {
	// This starts sending integers over the channel
	go func() {
		count := int32(1)
		for {
			demoChannel <- count * 10
			count++
		}
	}()
}

//export ReadChannel
func ReadChannel() int32 {
	// This returns the next integer from the channel
	return <- demoChannel
}
