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

type square struct {
	length int
	height int
}

type triangle struct {
	base   int
	height int
}

func (s *square) getArea() float64 {
	return float64(s.length) * float64(s.height)
}

func (t *triangle) getArea() float64 {
	return float64(t.base) * float64(t.height) / 2
}

type shape interface {
	getArea() float64
}

func draw(s shape) {
	//draw
	s.getArea()
}

func (s *square) Read(p []byte) (n int, err error) {
	return 0, nil
}
