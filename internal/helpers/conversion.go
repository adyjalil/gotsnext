package helpers

import "strconv"

func IntToString(n int) string {
	return strconv.Itoa(n)
}

func StringToInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return n, nil
}

var Difficulty = struct {
	Easy   string
	Medium string
	Hard   string
}{
	"easy",
	"medium",
	"hard",
}
