package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		err  error
	}{
		{
			name: "empty",
			in:   ``,
		},
		{
			name: "r2",
			in: `service cloudflare.storage {
        match /b/{bucket}/o {
          match /{allPaths=**} {
            allow read, write: if request.auth != null;
          }
        }
      }`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := FromString(tt.in)
			if tt.err != nil {
				assert.Nil(t, i)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())

				return
			}

			assert.NotNil(t, i)
			assert.NoError(t, err)
		})
	}
}
