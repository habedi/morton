package morton

import "fmt"

// Convert an uint32 to a bit array
func uInt32ToBitArray(x uint32) []uint8 {
	var result []uint8
	for i := 31; i >= 0; i-- {
		result = append(result, uint8(x>>i&1))
	}
	return result
}

// Convert an uint64 to a bit array
func uInt64ToBitArray(x uint64) []uint8 {
	var result []uint8
	for i := 63; i >= 0; i-- {
		result = append(result, uint8(x>>i&1))
	}
	return result
}

// Interleave two bit arrays
func interleave2D(x, y []uint8) []uint8 {
	var result []uint8
	for i := 0; i < 32; i++ {
		result = append(result, x[i])
		result = append(result, y[i])
	}
	return result
}

// Interleave three bit arrays
func interleave3D(x, y, z []uint8) []uint8 {
	var result []uint8

	// Add one 0 bit to the start of result to make it 64-bit long
	result = append(result, 0)

	// Currently only supports using the last 21 bits of each input array because the output should fit into a uint64
	for i := 11; i < 32; i++ {
		result = append(result, x[i])
		result = append(result, y[i])
		result = append(result, z[i])
	}
	return result
}

// Convert a bit array to an unit64
func fromBitArray(x []uint8) uint64 {
	var l = len(x)
	// Check if the bit array has correct length
	if (l != 64) && (l != 32) {
		panic("Bit array must have length 64 or 32 not " + fmt.Sprint(l))
	}

	// Convert the bit array to an uint64
	var result uint64 = 0
	for i := 0; i < l; i++ {
		result = result<<1 | uint64(x[i])
	}
	return result
}

// Encode2D Encodes a 2D point (x, y) to a Morton code
func Encode2D(x, y uint32) uint64 {

	// Interleave the bit arrays of x and y and convert the result to an uint64
	return fromBitArray(interleave2D(uInt32ToBitArray(x), uInt32ToBitArray(y)))
}

// Decode2D Decodes a Morton code to a 2D point (x, y)
func Decode2D(code uint64) (uint32, uint32) {

	// Convert the Morton code to a bit array
	var bits = uInt64ToBitArray(code)

	// Split the bit array into two separate bit arrays
	var x, y []uint8
	for i := 0; i < 64; i++ {
		if i%2 == 0 {
			x = append(x, bits[i])
		} else {
			y = append(y, bits[i])
		}
	}

	// Convert the bit arrays to uint32
	return uint32(fromBitArray(x)), uint32(fromBitArray(y))
}

// Encode3D Encodes a 3D point (x, y, z) to a Morton code
func Encode3D(x, y, z uint32) uint64 {

	// Interleave the bit arrays of x, y, and z and convert the result to an uint64
	return fromBitArray(interleave3D(uInt32ToBitArray(x), uInt32ToBitArray(y), uInt32ToBitArray(z)))
}

// Decode3D Decodes a Morton code to a 3D point (x, y, z)
func Decode3D(code uint64) (uint32, uint32, uint32) {

	// Convert the Morton code to a bit array
	var bits = uInt64ToBitArray(code)

	// Split the bit array into three separate bit arrays
	var x, y, z []uint8

	// Add 11 0 bits to the start of each bit array to make them 32-bit long
	for i := 0; i < 11; i++ {
		x = append(x, 0)
		y = append(y, 0)
		z = append(z, 0)
	}

	// Deinterleave the bit array into three smaller bit arrays
	for i := 1; i < 64; i = i + 3 {
		x = append(x, bits[i])
		y = append(y, bits[i+1])
		z = append(z, bits[i+2])
	}

	// Convert the bit arrays to uint32
	return uint32(fromBitArray(x)), uint32(fromBitArray(y)), uint32(fromBitArray(z))
}
