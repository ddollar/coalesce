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

	type s struct{ foo int }
	assert.Equal(t, s{}, coalesce.Any(s{}))
	assert.Equal(t, s{}, coalesce.Any(s{0}))
	assert.Equal(t, s{}, coalesce.Any(s{}, s{}))
	assert.Equal(t, s{3}, coalesce.Any(s{3}))
	assert.Equal(t, s{3}, coalesce.Any(s{3}, s{}))
	assert.Equal(t, s{3}, coalesce.Any(s{3}, s{}))
	assert.Equal(t, s{3}, coalesce.Any(s{}, s{3}))
	assert.Equal(t, s{3}, coalesce.Any(s{}, s{3}, s{}))
	assert.Equal(t, s{3}, coalesce.Any(s{}, s{3}, s{5}))
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
