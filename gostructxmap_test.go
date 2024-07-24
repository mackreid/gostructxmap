package gostructxmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	st := struct {
		Name   string `sxm:"name"`
		Age    int    `sxm:"age"`
		Gender string `sxm:"-"`
	}{
		Name:   "TestName",
		Age:    30,
		Gender: "TestGender",
	}
	exp := map[string]any{
		"name": "TestName",
		"age":  30,
	}
	ma, err := Make(st)
	assert.NoError(t, err)
	assert.Equal(t, exp, ma)
}

func TestDoesSkip(t *testing.T) {
	tag := "NoSkip"
	emptyTag := ""
	skipTag := "-"
	assert.Equal(t, false, doesSkip(tag))
	assert.Equal(t, true, doesSkip(emptyTag))
	assert.Equal(t, true, doesSkip(skipTag))
}

func TestWithOmit(t *testing.T) {
	f := WithOmit()

	contain := f("", "test", "omitempty")
	assert.Equal(t, false, contain)
}

func TestWithMask(t *testing.T) {
	mask := []string{"name", "age"}
	f := WithMask(mask...)
	contain := f("test", "name", "mask")
	assert.Equal(t, true, contain)
}

type FullTestStruct struct {
	ID     string `sxm:"-"`
	Name   string `sxm:"name,omitempty,mask"`
	Title  string `sxm:"title,omitempty"`
	Author string `sxm:"author,mask"`
	Pages  int    `sxm:"pages,omitempty"`
}

func TestFullWithOmit(t *testing.T) {
	st := &FullTestStruct{
		ID:    "123",
		Title: "TestTitle",
	}
	exp := map[string]any{
		"title":  "TestTitle",
		"author": "",
	}
	ma, err := Make(st, WithOmit())
	assert.NoError(t, err)
	assert.Equal(t, exp, ma)
}
