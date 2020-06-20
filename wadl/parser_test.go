package wadl

import (
	"github.com/grahambrooks/scheme/search"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadingWADL(t *testing.T) {
	t.Run("storage-service", func(t *testing.T) {
		f, err := os.Open("test/storage-service.wadl")
		assert.NoError(t, err)

		parser := Parser{}
		model, err := parser.Parse(f)

		assert.NoError(t, err)
		assert.Equal(t, search.WADL, model.Kind)
		assert.Equal(t, 3, len(model.Resources))
	})
}
