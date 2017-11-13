package countries

func PadLeftZero(s string, length int) string {
	for len(s) < length {
		s = "0" + s
	}

	return s
}
