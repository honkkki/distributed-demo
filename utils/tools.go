package utils

// 字符串切片去重
func sliceUnique(src []string) []string {
	var res []string
	m := make(map[string]struct{})
	for _, v := range src {
		if _, ok := m[v]; ok {
			continue
		} else {
			m[v] = struct{}{}
			res = append(res, v)
		}
	}

	return res
}
