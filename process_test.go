package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"path/filepath"
)

func TestReadInput(t *testing.T) {
	actual, err := readInputs("foo", nil)
	assert.Nil(t, err)
	assert.Equal(t, "foo", actual[0])

	// stdin is "" because during tests it's given /dev/null
	actual, err = readInputs("", []string{"-"})
	assert.Nil(t, err)
	assert.Equal(t, "", actual[0])

	actual, err = readInputs("", []string{"main_test.go"})
	assert.Nil(t, err)
	thisFile, _ := os.Open("main_test.go")
	expected, _ := ioutil.ReadAll(thisFile)
	assert.Equal(t, string(expected), actual[0])
}

func TestInputDir(t *testing.T) {
	outDir, err := ioutil.TempDir("test/files/input-dir", "out-temp-")
	assert.Nil(t, err)
	defer os.RemoveAll(outDir)

	src, err := ParseSource("config=test/files/input-dir/config.yml")
	assert.Nil(t, err)

	data := &Data{
		Sources: map[string]*Source { "config" : src },
	}
	gomplate := NewGomplate(data, "{{", "}}")
	err = processDir(gomplate,"test/files/input-dir/in", outDir)
	assert.Nil(t, err)

	top, err := ioutil.ReadFile(filepath.Join(outDir, "top.txt"))
	assert.Nil(t, err)
	assert.Equal(t, "eins", string(top))

	inner, err := ioutil.ReadFile(filepath.Join(outDir, "inner/nested.txt"))
	assert.Nil(t, err)
	assert.Equal(t, "zwei", string(inner))
}
