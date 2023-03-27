package gokey

func isEmpty(key string) bool {
	return key == ""
}

func sizeLimit(slimit int) int {
	if slimit < 1 {
		return getLimitPairsSet()
	}
	return slimit
}
