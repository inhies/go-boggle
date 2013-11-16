package boggle

import (
	"bufio"
	"os"
	"strings"
)

type Dictionary interface {
	Exists(string) bool
	Add(string)
}

// A dictionary created from a map. Much faster than slices.
type MapDict map[string]struct{}

func NewMapDict() *MapDict {
	d := make(MapDict)
	return &d
}

func (d *MapDict) Exists(str string) bool {
	_, ok := (*d)[str]
	return ok
}

func (d *MapDict) Add(str string) {
	(*d)[str] = struct{}{}
}

// A dictionary created from slices. Only for testing.
type SliceDict []string

func NewSliceDict() *SliceDict {
	d := make(SliceDict, 0)
	return &d
}
func (s *SliceDict) Exists(str string) bool {
	for _, v := range *s {
		if v == str {
			return true
		}
	}
	return false
}

func (s *SliceDict) Add(str string) {
	*s = append(*s, str)
}

func Load(file string, dict Dictionary) (err error) {
	var f *os.File
	f, err = os.Open(file)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dict.Add(strings.ToUpper(scanner.Text()))
	}

	if err = scanner.Err(); err != nil {
		return
	}
	return
}
