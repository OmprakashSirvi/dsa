package bits

import "math"

// Check if the ith bit is set in the binary representation of num
func IsBitSet(num int, i int) bool {
	return (num & (1 << i)) != 0
}

func IsPow2(num int) bool {
	return num > 0 && num&(num-1) == 0
}

func CountSetBits(num int) int {
	count := 0
	for {
		if num <= 0 {
			break
		}
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}

func Divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negative := (dividend > 0) != (divisor > 0)

	x := dividend
	y := divisor

	if x > 0 {
		x = -x
	}
	if y > 0 {
		y = -y
	}
	quotient := 0
	for x <= y {
		power := 0
		for {
			if (y * (1 << power)) < x {
				power--
				break
			}
			power++
		}

		quotient += 1 << power
		x -= (y * (1 << power))
	}
	if negative {
		return -quotient
	}

	return quotient
}

// func Divide(dividend int, divisor int) int {
// 	// We will figure out the largest number which can be removed from the dividend.
// 	if dividend == divisor {
// 		return 1
// 	}

// 	n := int(math.Abs(float64(dividend)))
// 	quotient := 0
// 	for {
// 		// exit condition
// 		if n < divisor {
// 			break
// 		}

// 		count := 0
// 		for {
// 			if (divisor * int(math.Pow(2, float64(count)))) > n {
// 				count--
// 				break
// 			}
// 			count++
// 		}

// 		n -= divisor * int(math.Pow(2, float64(count)))
// 		quotient += int(math.Pow(2, float64(count)))
// 	}

// 	neg := (dividend > 0 && divisor < 0) || (dividend < 0 && dividend > 0)
// 	if neg {
// 		return quotient * -1
// 	}

// 	return quotient
// }
