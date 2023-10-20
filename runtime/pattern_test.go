package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPatternMatcher(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		matcher *patternMatcher
		err     error
	}{
		{
			name: "empty",
			matcher: &patternMatcher{
				fields: []string{},
				params: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewPatternMatcher(tt.path)
			if tt.err != nil {
				assert.Error(t, err)
				assert.Nil(t, m)

				return
			}

			assert.NotNil(t, m)
			assert.NoError(t, err)
			assert.Equal(t, tt.matcher, m)
		})
	}
}
