#include <iostream>
#include <vector>
#include <go_library.h>

int main() {
	// Print all of the base number types being modified from Go
	std::cout << "ModifyBool:       " << (ModifyBool(true) ? "true" : "false") << std::endl;
	std::cout << "ModifyInt:        " << ModifyInt(43) << std::endl;
	std::cout << "ModifyInt8:       " << (int)ModifyInt8(68) << std::endl;
	std::cout << "ModifyInt16:      " << ModifyInt16(1453) << std::endl;
	std::cout << "ModifyInt32:      " << ModifyInt32(72431565) << std::endl;
	std::cout << "ModifyInt64:      " << ModifyInt64(23262353858424) << std::endl;
	std::cout << "ModifyUint:       " << ModifyUint(44) << std::endl;
	std::cout << "ModifyUint8:      " << (int)ModifyUint8(69) << std::endl;
	std::cout << "ModifyUint16:     " << ModifyUint16(1454) << std::endl;
	std::cout << "ModifyUint32:     " << ModifyUint32(72431566) << std::endl;
	std::cout << "ModifyUint64:     " << ModifyUint64(23262353858425) << std::endl;
	std::cout << "ModifyFloat32:    " << ModifyFloat32(573.42) << std::endl;
	std::cout << "ModifyFloat64:    " << ModifyFloat64(6935.34562) << std::endl;

	// Strings must be freed, as they're allocated by Go using malloc
	auto modifiedString = ModifyString(GoString{.p = "aBcDeF", .n = 6});
	std::cout << "ModifyString:     " << modifiedString << std::endl;
	free(modifiedString);

	// We can use a familiar vector to store out ints, which we can easily convert to a GoSlice
	std::vector<std::int32_t> inputInts;
	for (int i = 0; i < 5; i++) {
		inputInts.push_back(std::int32_t(i) + 10);
	}
	auto outputInts = ModifyInt32Slice(GoSlice{
		.data = inputInts.data(),
		.len = GoInt(inputInts.size()),
		.cap = GoInt(inputInts.size()),
	});
	std::cout << "ModifyInt32Slice: [";
	for (int i = 0; i < inputInts.size(); i++) {
		if (i > 0) {
			std::cout << ",";
		}
		std::cout << inputInts[i];
	}
	std::cout << "] [";
	for (int i = 0; i < outputInts.r1; i++) {
		if (i > 0) {
			std::cout << ",";
		}
		std::cout << outputInts.r0[i];
	}
	std::cout << "]" << std::endl;
	// Just like with strings, we must remember to free our integers
	free(outputInts.r0);

	// This just shows how we can modify a local struct by passing in a pointer, as well as using a struct returned by
	// value
	DemoStruct inputStruct{.A = 12, .B = 34};
	auto outputStruct = ModifyDemoStruct(&inputStruct, DemoStruct{.A = 56, .B = 78});
	std::cout << "ModifyDemoStruct: {A: " << (int)inputStruct.A << ", B: " << inputStruct.B << "} {A: ";
	std::cout << (int)outputStruct.A << ", B: " << outputStruct.B << "}" << std::endl;

	// The Go runtime manages its own threads
	StartChannel();
	std::cout << "GoManagedThreads: " << ReadChannel() << "," << ReadChannel() << "," << ReadChannel() << std::endl;

	return 0;
}
