package coalesce_test

import (
	"testing"

	"github.com/ddollar/coalesce"
	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	assert.Equal(t, 0, coalesce.Any(0))
	assert.Equal(t, 0, coalesce.Any(0, 0))
	assert.Equal(t, 3, coalesce.Any(3))
	assert.Equal(t, 3, coalesce.Any(3, 0))
	assert.Equal(t, 3, coalesce.Any(0, 3))
	assert.Equal(t, 3, coalesce.Any(0, 3, 0))
	assert.Equal(t, 3, coalesce.Any(0, 3, 5))

	assert.Equal(t, "", coalesce.Any(""))
	assert.Equal(t, "", coalesce.Any("", ""))
	assert.Equal(t, "foo", coalesce.Any("foo"))
	assert.Equal(t, "foo", coalesce.Any("foo", ""))
	assert.Equal(t, "foo", coalesce.Any("", "foo"))
	assert.Equal(t, "foo", coalesce.Any("", "foo", ""))
	assert.Equal(t, "foo", coalesce.Any("", "foo", "bar"))

	type s1 struct{ foo int }
	assert.Equal(t, s1{}, coalesce.Any(s1{}))
	assert.Equal(t, s1{}, coalesce.Any(s1{0}))
	assert.Equal(t, s1{}, coalesce.Any(s1{}, s1{}))
	assert.Equal(t, s1{3}, coalesce.Any(s1{3}))
	assert.Equal(t, s1{3}, coalesce.Any(s1{3}, s1{}))
	assert.Equal(t, s1{3}, coalesce.Any(s1{3}, s1{}))
	assert.Equal(t, s1{3}, coalesce.Any(s1{}, s1{3}))
	assert.Equal(t, s1{3}, coalesce.Any(s1{}, s1{3}, s1{}))
	assert.Equal(t, s1{3}, coalesce.Any(s1{}, s1{3}, s1{5}))

	type s2 struct {
		foo int
		bar map[string]string
	}
	assert.Equal(t, s2{}, coalesce.Any(s2{}))
	assert.Equal(t, s2{}, coalesce.Any(s2{0, nil}))
	assert.Equal(t, s2{}, coalesce.Any(s2{}, s2{}))
	assert.Equal(t, s2{3, map[string]string{"baz": "qux"}}, coalesce.Any(s2{3, map[string]string{"baz": "qux"}}))
	assert.Equal(t, s2{3, map[string]string{"baz": "qux"}}, coalesce.Any(s2{3, map[string]string{"baz": "qux"}}, s2{}))
	assert.Equal(t, s2{3, map[string]string{"baz": "qux"}}, coalesce.Any(s2{3, map[string]string{"baz": "qux"}}, s2{}))
	assert.Equal(t, s2{3, map[string]string{"baz": "qux"}}, coalesce.Any(s2{}, s2{3, map[string]string{"baz": "qux"}}))
	assert.Equal(t, s2{3, map[string]string{"baz": "qux"}}, coalesce.Any(s2{}, s2{3, map[string]string{"baz": "qux"}}, s2{}))
	assert.Equal(t, s2{3, map[string]string{"baz": "qux"}}, coalesce.Any(s2{}, s2{3, map[string]string{"baz": "qux"}}, s2{5, nil}))
}

func TestPtr(t *testing.T) {
	value := 12

	assert.Equal(t, value, coalesce.Ptr(&value, 22))
	assert.Equal(t, 22, coalesce.Ptr(nil, 22))
}

func TestString(t *testing.T) {
	assert.Equal(t, "", coalesce.String(""))
	assert.Equal(t, "", coalesce.String("", ""))
	assert.Equal(t, "foo", coalesce.String("foo"))
	assert.Equal(t, "foo", coalesce.String("foo", ""))
	assert.Equal(t, "foo", coalesce.String("", "foo"))
	assert.Equal(t, "foo", coalesce.String("", "foo", ""))
	assert.Equal(t, "foo", coalesce.String("", "foo", "bar"))
}
