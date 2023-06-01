package converter

func ArrayNotInArray(original []string, search []string) []string {
	var result []string
	originalMap := make(map[string]bool)
	for _, v := range original {
		originalMap[v] = true
	}
	for _, v := range search {
		if _, ok := originalMap[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}
