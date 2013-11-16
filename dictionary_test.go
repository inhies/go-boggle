package boggle

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Package global variable to ensure the compiler doesn't optimize away our
// tests and benchmarks.
var saveError error

// Test data.
var WordList = []struct {
	Word   string
	Exists bool
}{
	{"AA", true},
	{"ZZZ", true},
	{"THISISNOTAWORD", false},
}

// Global map and slice dictionaries to be pre-filled
var (
	mapDict   *MapDict
	sliceDict *SliceDict
)

//Preload dictionaries for benchmarks.
func init() {
	mapDict = NewMapDict()
	sliceDict = NewSliceDict()
	file := "./dictionaries/OWL2.txt"
	Load(file, mapDict)
	Load(file, sliceDict)
}

func Test_LoadMap(t *testing.T) {
	Convey("Given a valid file", t, func() {
		dict := NewMapDict()
		file := "./dictionaries/OWL2.txt"
		So(Load(file, dict), ShouldBeNil)
		Convey("We should be able to add each line to the dictionary", func() {
			Convey("We should be able to check if a word exists", func() {
				for _, v := range WordList {
					So(dict.Exists(v.Word), ShouldEqual, v.Exists)
				}
			})
		})
	})
}

func Test_LoadSlice(t *testing.T) {
	Convey("Given a valid file", t, func() {
		dict := NewSliceDict()
		file := "./dictionaries/OWL2.txt"
		So(Load(file, dict), ShouldBeNil)
		Convey("We should be able to add each line to the dictionary", func() {
			Convey("We should be able to check if a word exists", func() {
				for _, v := range WordList {
					So(dict.Exists(v.Word), ShouldEqual, v.Exists)
				}
			})
		})
	})
}

func Benchmark_LoadMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dict := NewMapDict()
		// TODO(inhies): Load file in to memory to be read for consistency.
		file := "./dictionaries/OWL2.txt"

		err := Load(file, dict)
		for _, v := range WordList {
			dict.Exists(v.Word)
		}
		saveError = err
	}
}

func Benchmark_LoadSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dict := NewSliceDict()
		file := "./dictionaries/OWL2.txt"

		err := Load(file, dict)
		for _, v := range WordList {
			dict.Exists(v.Word)
		}
		saveError = err
	}
}

/*
func Benchmark_TestMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range WordList {
			mapDict.Exists(v.Word)
		}
	}
}

func Benchmark_TestSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range WordList {
			sliceDict.Exists(v.Word)
		}
	}
}
*/

func benchExists(word string, d Dictionary, b *testing.B) {
	for i := 0; i < b.N; i++ {
		d.Exists(word)
	}
}

func Benchmark_Map_FirstWord(b *testing.B)   { benchExists("AA", mapDict, b) }
func Benchmark_Slice_FirstWord(b *testing.B) { benchExists("AA", sliceDict, b) }
func Benchmark_Map_LastWord(b *testing.B)    { benchExists("ZZZ", mapDict, b) }
func Benchmark_Slice_LastWord(b *testing.B)  { benchExists("ZZZ", sliceDict, b) }
