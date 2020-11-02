package offer

func permutation(s string) []string {
	if len(s) == 0 {
		return nil
	}
	return permutationHelper([]byte(s), []byte{})
}

func permutationHelper(remain []byte, last []byte) []string {
	if len(remain) == 1 {
		return []string{string(append(last, remain[0]))}
	}

	var ret []string
	var remainBytes []byte
	m := make(map[byte]bool)

	for i := 0; i < len(remain); i++ {
		if m[remain[i]] {
			continue
		}
		m[remain[i]] = true
		remainBytes = append(remainBytes, remain[:i]...)
		remainBytes = append(remainBytes, remain[i+1:]...)
		ret = append(ret, permutationHelper(remainBytes, append(last, remain[i]))...)
		remainBytes = remainBytes[0:0]
	}
	return ret
}
