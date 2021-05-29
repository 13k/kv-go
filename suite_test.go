package kv_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/stretchr/testify/suite"

	"github.com/13k/kv-go"
)

type Suite struct {
	suite.Suite
}

func (s *Suite) FixturePath(path string) string {
	return filepath.Join("testdata", path)
}

func (s *Suite) OpenFixture(path string) (*os.File, error) {
	return os.Open(s.FixturePath(path))
}

func (s *Suite) MustOpenFixture(path string) *os.File {
	f, err := s.OpenFixture(path)

	if err != nil {
		panic(err)
	}

	return f
}

func (s *Suite) ReadFixture(path string) ([]byte, error) {
	return ioutil.ReadFile(s.FixturePath(path))
}

func (s *Suite) MustReadFixture(path string) []byte {
	data, err := s.ReadFixture(path)

	if err != nil {
		panic(err)
	}

	return data
}

func (s *Suite) RequireEqualKeyValue(expected, actual kv.KeyValue) {
	r := s.Require()

	r.Equalf(expected.Type(), actual.Type(), "wrong KeyValue type")
	r.Equalf(expected.Key(), actual.Key(), "wrong KeyValue key")
	r.Equalf(expected.Value(), actual.Value(), "wrong KeyValue value")
	r.Lenf(actual.Children(), len(expected.Children()), "wrong KeyValue children count")

	for childIdx, expectedChild := range expected.Children() {
		actualChild := actual.Children()[childIdx]
		s.RequireEqualKeyValuef(expectedChild, actualChild, "child %d", childIdx)
	}
}

func (s *Suite) RequireEqualKeyValuef(expected, actual kv.KeyValue, format string, args ...interface{}) {
	r := s.Require()

	r.Equalf(expected.Type(), actual.Type(), format+": wrong KeyValue type", args...)
	r.Equalf(expected.Key(), actual.Key(), format+": wrong KeyValue key", args...)
	r.Equalf(expected.Value(), actual.Value(), format+": wrong KeyValue value", args...)
	r.Lenf(actual.Children(), len(expected.Children()), format+": wrong KeyValue children count", args...)

	for childIdx, expectedChild := range expected.Children() {
		actualChild := actual.Children()[childIdx]
		s.RequireEqualKeyValuef(expectedChild, actualChild, format+": child %d", append(args, childIdx)...)
	}
}
