package cdb_mmap_test

import (
	"testing"

	"github.com/bsergean/cdb_mmap"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert.EqualValues(t, 776976811, cdb_mmap.CDBHash([]byte("foo bar baz")))
	assert.EqualValues(t, 3538394712, cdb_mmap.CDBHash([]byte("The quick brown fox jumped over the lazy dog")))
}
