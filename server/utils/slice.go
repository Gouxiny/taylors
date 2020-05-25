package utils

func SliceUniqueString(s []string) []string {
	uniquedSlice := []string{}
	m := make(map[string]bool)
	for _, v := range s {
		if _, exists := m[v]; !exists {
			m[v] = true
			uniquedSlice = append(uniquedSlice, v)
		}
	}
	return uniquedSlice
}
