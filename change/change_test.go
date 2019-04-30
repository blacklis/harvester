package change

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taxibeat/harvester/config"
)

func TestChange(t *testing.T) {
	c := New(config.SourceConsul, "key", "value", 1)
	assert.Equal(t, config.SourceConsul, c.Source())
	assert.Equal(t, "key", c.Key())
	assert.Equal(t, "value", c.Value())
	assert.Equal(t, uint64(1), c.Version())
}
