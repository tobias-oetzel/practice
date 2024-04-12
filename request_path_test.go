package proxy

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestPath(t *testing.T) {

	t.Run("get path", func(t *testing.T) {
		// Arrange
		url, err := url.Parse("https://my-proxy.example.com/path%20to/resource")

		// Act
		requestPath := RequestPath{Url: url}

		result := requestPath.getPath()

		// Assert
		assert.Equal(t, "/path to/resource", result)
		assert.NoError(t, err)
	})
}
