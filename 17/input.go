package day17

type Input struct {
	registerA int
	registerB int
	registerC int
	program   []int
}

var TEST_INPUT = Input{
	registerA: 729,
	registerB: 0,
	registerC: 0,
	program:   []int{0, 1, 5, 4, 3, 0},
}

var INPUT = Input{
	registerA: 202356708354602, // registerA: 32916674,
	registerB: 0,
	registerC: 0,
	program: []int{
		2, 4, // registerB = registerA % 8
		1, 1, // registerB = registerB XOR 001
		7, 5, // registerC = registerA / 2^registerB
		0, 3, // registerA = registerA / 8
		1, 4, // registerB = registerB XOR 100
		4, 0, // registerB = registerB XOR registerC
		5, 5, // out: registerB % 8
		3, 0, // jump to start!!!
		// if A = 101111011110010101100111101010010110001000011111

		// B = 111
		// B = 110 = 111 XOR 001
		// C = 1011110111100101011001111010100101100010000
		// A = 101111011110010101100111101010010110001000011111
		// B = 000 = 100 XOR 100
		// B = 1 = 001 = 000 XOR 001

		// B = 011
		//

		// out: registerB % 8
		// out: (registerB XOR registerC) % 8
		// out: ((registerB XOR 100) XOR registerC) % 8
		// out: ((registerB XOR 100) XOR (registerA / 2^registerB)) % 8
		// out: (((registerB XOR 001) XOR 100) XOR (registerA / 2^registerB)) % 8
		// out: (((registerA % 8 XOR 001) XOR 100) XOR (registerA / 2^(registerA % 8 XOR 001))) % 8

		// out = ((registerA % 8 XOR 001) XOR 100) XOR registerA / 2^(registerA % 8 XOR 001)) % 8
		// out = ((registerA XOR 001) XOR 100) XOR registerA / 2^(registerA % 8 XOR 001)

		// Egyenletrendszer (A = registerA)
		// out = A XOR 101 XOR (A / 2^(A XOR 001))

		// 111 = A XOR (A >> (A XOR 001))
		// A XOR 111 = A >> (A XOR 001); A % 8 = 010
		// 111 = 010 XOR A >> 3;

		// legyen 3 a shift? ^^^^^^^^^
		// lehet, hogy minimalizálni kell a shiftet?
		// ;
		// 110 = 011 XOR A >> 2;

		// x = A; y = (A / 2^(A XOR 001))

		// 010 XOR 101 = 111 = 010 XOR 101		111 = ? XOR ? A >> ?
		// 100 XOR 101 = 001 = 010 XOR 011		001 = ? XOR ? A >> ?
		// 001 XOR 101 = 100 = 010 XOR 110		100 = ? XOR ? A >> ?
		// 001 XOR 101 = 100 = 010 XOR 110		100 = ? XOR ? A >> ?
		// 111 XOR 101 = 010 = 010 XOR 000		010 = 111 XOR  A >> 5
		// 101 XOR 101 = 000 = 010 XOR 010		000 = 110 XOR 110 A >> 7
		// 000 XOR 101 = 101 = 010 XOR 111		101 = 110 XOR 011 A >> 7
		// 011 XOR 101 = 110 = 010 XOR 100		110 = 101 XOR 011 A >> 4
		// 001 XOR 101 = 100 = 010 XOR 110		100 = 111 XOR 011 A >> 6
		// 100 XOR 101 = 001 = 010 XOR 011		001 = 010 XOR 011 A >> 3
		// 100 XOR 101 = 001 = 010 XOR 011		001 = 011 XOR 010 A >> 2
		// 000 XOR 101 = 101 = 010 XOR 111		101 = 101 XOR 000 A >> 4
		// 101 XOR 101 = 000 = 010 XOR 010		000 = 000 XOR 000 A >> 1
		// 101 XOR 101 = 000 = 010 XOR 010		000 = 000 XOR 000 A >> 1
		// 011 XOR 101 = 110 = 010 XOR 100		110 = 110 XOR 000 A >> 7

		// 000 XOR 101 = 101 = 101 XOR 000		101 = 011 XOR 110 A >> 2
		// 000 XOR 101 = 101 = 101 XOR 000		101 = 010 XOR 111 A >> 3
		// 000 XOR 101 = 101 = 101 XOR 000		101 = 101 XOR 000 A >> 4
		// 000 XOR 101 = 101 = 101 XOR 000		101 = 100 XOR 001 A >> 5
		// 000 XOR 101 = 101 = 101 XOR 000		101 = 111 XOR 010 A >> 6
		// 000 XOR 101 = 101 = 101 XOR 000		101 = 110 XOR 011 A >> 7
		// ami itt van, az     ^^^    tuti a 45-ik számjegy után van

		// ..101 11000 00001 01011 01011 11011 10110 ..... ..... .....
		//
		// y 2-vel a meglévő után, x folytonosan

		// x = (A);        y = (A / 2^5)
		// x = (A / 2^3);  y = (A / 2^8)
		// x = (A / 2^6);  y = (A / 2^11)
		// x = (A / 2^9);  y = (A / 2^14)
		// x = (A / 2^12); y = (A / 2^17)
		// x = (A / 2^15); y = (A / 2^20)
		// x = (A / 2^18); y = (A / 2^23)
		// x = (A / 2^21); y = (A / 2^26)
		// x = (A / 2^24); y = (A / 2^29)
		// x = (A / 2^27); y = (A / 2^32)
		// x = (A / 2^30); y = (A / 2^35)
		// x = (A / 2^33); y = (A / 2^38)
		// x = (A / 2^36); y = (A / 2^41)
		// x = (A / 2^39); y = (A / 2^44)
		// x = (A / 2^42); y = (A / 2^47)
		// x = (A / 2^45); y = (A / 2^50)
	},
}
