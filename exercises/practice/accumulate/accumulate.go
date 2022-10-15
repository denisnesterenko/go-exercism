package accumulate

func Accumulate(list []string, transform func(string) string) []string {
	var result []string
	for _, elem := range list {
		result = append(result, transform(elem))
	}
	return result
}
