package main

import "core:strconv"
import "core:fmt"
import "core:crypto/hash"
import "core:strings"

INPUT :: #load("../input/day4.txt", string)

main :: proc() {
	input := strings.trim_space(INPUT)

	first_5_zeros: int
	first_6_zeros: int

	key_buf: [64]u8
	copy(key_buf[:], input)

	for i := 1; ; i += 1 {
		i_str := strconv.write_int(key_buf[len(input):], i64(i), 10)
		input_str := key_buf[:len(input) + len(i_str)]

		digest: [16]u8
		hash.hash_bytes_to_buffer(.Insecure_MD5, input_str, digest[:])

		// 5 hex zeros = 2 zero bytes + top 4 bits of the 3rd byte are zero
		if first_5_zeros == 0 && digest[0] == 0 && digest[1] == 0 && digest[2] & 0xF0 == 0 {
			first_5_zeros = i
		}

		// 6 hex zeros = 3 zero bytes
		if first_6_zeros == 0 && digest[0] == 0 && digest[1] == 0 && digest[2] == 0 {
			first_6_zeros = i
		}

		if first_5_zeros != 0 && first_6_zeros != 0 {
			break
		}
	}

	fmt.printf("First 5 zeros: %d\n", first_5_zeros)
	fmt.printf("First 6 zeros: %d\n", first_6_zeros)
}
