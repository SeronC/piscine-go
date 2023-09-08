package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}
	totalLength := 0
	for _, arg := range args {
		totalLength += len(arg) + 1
	}
	concatenated := make([]byte, 0, totalLength)
	for i, arg := range args {
		concatenated = append(concatenated, []byte(arg)...)
		if i != len(args)-1 {
			concatenated = append(concatenated, '\n')
		}
	}
	return string(concatenated)
}
