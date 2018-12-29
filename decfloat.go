/*******************************************************************************
The MIT License (MIT)

Copyright (c) 2019 Hajime Nakagami

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*******************************************************************************/

package firebirdsql

import (
	"github.com/shopspring/decimal"
	"math/big"
)

func dpdBitToInt(dpd uint, mask uint) int {
	if (dpd & mask) != 0 {
		return 1
	} else {
		return 0
	}
}

func dpdToInt(dpd uint) int {
	// Convert DPD encodined value to int (0-999)
	// dpd: DPD encoded value. 10bit unsigned int

	b := make([]int, 10)
	b[9] = dpdBitToInt(dpd, 0x0200)
	b[8] = dpdBitToInt(dpd, 0x0100)
	b[7] = dpdBitToInt(dpd, 0x0080)
	b[6] = dpdBitToInt(dpd, 0x0040)
	b[5] = dpdBitToInt(dpd, 0x0020)
	b[4] = dpdBitToInt(dpd, 0x0010)
	b[3] = dpdBitToInt(dpd, 0x0008)
	b[2] = dpdBitToInt(dpd, 0x0004)
	b[1] = dpdBitToInt(dpd, 0x0002)
	b[0] = dpdBitToInt(dpd, 0x0001)

	d := make([]int, 3)
	if b[3] == 0 {
		d[2] = b[9]*4 + b[8]*2 + b[7]
		d[1] = b[6]*4 + b[5]*2 + b[4]
		d[0] = b[2]*4 + b[1]*2 + b[0]
	} else if b[3] == 1 && b[2] == 0 && b[1] == 0 {
		d[2] = b[9]*4 + b[8]*2 + b[7]
		d[1] = b[6]*4 + b[5]*2 + b[4]
		d[0] = 8 + b[0]
	} else if b[3] == 1 && b[2] == 0 && b[1] == 1 {
		d[2] = b[9]*4 + b[8]*2 + b[7]
		d[1] = 8 + b[4]
		d[0] = b[6]*4 + b[5]*2 + b[0]
	} else if b[3] == 1 && b[2] == 1 && b[1] == 0 {
		d[2] = 8 + b[7]
		d[1] = b[6]*4 + b[5]*2 + b[4]
		d[0] = b[9]*4 + b[8]*2 + b[0]
	} else if b[6] == 0 && b[5] == 0 && b[3] == 1 && b[2] == 1 && b[1] == 1 {
		d[2] = 8 + b[7]
		d[1] = 8 + b[4]
		d[0] = b[9]*4 + b[8]*2 + b[0]
	} else if b[6] == 0 && b[5] == 1 && b[3] == 1 && b[2] == 1 && b[1] == 1 {
		d[2] = 8 + b[7]
		d[1] = b[9]*4 + b[8]*2 + b[4]
		d[0] = 8 + b[0]
	} else if b[6] == 1 && b[5] == 0 && b[3] == 1 && b[2] == 1 && b[1] == 1 {
		d[2] = b[9]*4 + b[8]*2 + b[7]
		d[1] = 8 + b[4]
		d[0] = 8 + b[0]
	} else if b[6] == 1 && b[5] == 1 && b[3] == 1 && b[2] == 1 && b[1] == 1 {
		d[2] = 8 + b[7]
		d[1] = 8 + b[4]
		d[0] = 8 + b[0]
	} else {
		panic("Invalid DPD encoding")
	}

	return d[2]*100 + d[1]*10 + d[0]
}

func calcSignificand(prefix int64, dpdBits big.Int, numBits int) *big.Int {
	// prefix: High bits integer value
	// dpdBits: dpd encoded bits
	// numBits: bit length of dpd_bits
	// https://en.wikipedia.org/wiki/Decimal128_floating-point_format#Densely_packed_decimal_significand_field
	numSegments := numBits / 10
	segments := make([]uint, numSegments)
	bi1024 := big.NewInt(1024)

	for i := 0; i < numSegments; i++ {
		work := dpdBits
		segments[numSegments-i-1] = uint(dpdBits.Mod(&work, bi1024).Int64())
		dpdBits.Rsh(&dpdBits, 10)
	}

	v := big.NewInt(prefix)
	bi1000 := big.NewInt(1000)
	for _, dpd := range segments {
		v.Mul(v, bi1000)
		v.Add(v, big.NewInt(int64(dpd)))
	}

	return v
}

func decimal128ToSignDigitsExponent(b []byte) (v decimal.Decimal, sign int, digits int, exponent int) {
	// https://en.wikipedia.org/wiki/Decimal128_floating-point_format
	// TODO:
	v = decimal.Zero
	return
}

func decimalFixedToDecimal(b []byte, scale int) decimal.Decimal {
	// TODO:
	return decimal.Zero
}

func decimal128ToDecimal(b []byte) decimal.Decimal {
	// https://en.wikipedia.org/wiki/Decimal64_floating-point_format
	// TODO:
	return decimal.Zero
}
