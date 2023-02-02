package main

import (
	"C"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// ToCSlice converts a slice of floats or integers into the pointer + length combo that C operates on.
func ToCSlice[T constraints.Integer | constraints.Float](input []T) (*T, int) {
	var element T
	elementSize := int(unsafe.Sizeof(element))
	allocatedMemory := C.malloc(C.ulonglong(len(input) * elementSize))
	for i := 0; i < len(input); i++ {
		allocatedMemoryLocation := (*T)(unsafe.Add(allocatedMemory, i * elementSize))
		*allocatedMemoryLocation = input[i]
	}
	return (*T)(allocatedMemory), len(input)
}
