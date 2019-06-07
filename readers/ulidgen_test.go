package readers

import (
	"testing"
	"time"

	"github.com/flyingmutant/rapid"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewReader(t *testing.T) {
	entropy := NewXidReader()
	id, err := ulid.New(ulid.Now(), entropy)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestXidReader_Monolith(t *testing.T) {
	entropy := NewXidReader()
	rapid.Check(t, func(t *rapid.T) {
		now := ulid.Timestamp(time.Now())
		id1, err := ulid.New(now, entropy)
		assert.NoError(t, err)
		id2, err := ulid.New(now, entropy)
		assert.NoError(t, err)
		assert.Equal(t, id1.Compare(id2), -1)
	})
}

func TestMonolithReader_Monolith(t *testing.T) {
	entropy := NewMonolithReader()
	rapid.Check(t, func(t *rapid.T) {
		now := ulid.Timestamp(time.Now())
		id1, err := ulid.New(now, entropy)
		assert.NoError(t, err)
		id2, err := ulid.New(now, entropy)
		assert.NoError(t, err)
		assert.Equal(t, id1.Compare(id2), -1)
	})
}

func BenchmarkXidReader_ULIDGen(b *testing.B) {
	entropy := NewXidReader()
	for i := 0; i < b.N; i++ {
		ulid.MustNew(ulid.Now(), entropy)
	}
}

func BenchmarkMonolithReader_ULIDGen(b *testing.B) {
	entropy := NewMonolithReader()
	for i := 0; i < b.N; i++ {
		ulid.MustNew(ulid.Now(), entropy)
	}
}
