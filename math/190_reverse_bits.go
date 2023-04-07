package math

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res * 2 + num % 2
		num = num / 2
	}
	return res
}

// 同理，10进制数翻转
//for x != 0 {
//	res = res * 10 + x % 10
//	x /= 10
//}