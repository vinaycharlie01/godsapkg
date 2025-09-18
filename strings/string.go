package strings

func ReverseString(s []byte) {
	L := 0
	R := len(s) - 1
	for L < R {
		temp := s[L]
		s[L] = s[R]
		s[R] = temp
		L++
		R--
	}
}
