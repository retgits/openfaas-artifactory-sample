package function

import (
	"fmt"
	"net/http"
	"testing"

	handler "github.com/openfaas-incubator/go-function-sdk"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	fmt.Println("TestHandler")
	assert := assert.New(t)

	message := "Hello World!"

	req := handler.Request{
		Body:   []byte(message),
		Method: http.MethodPost,
	}

	res, err := Handle(req)
	assert.NoError(err)
	assert.Equal(res.StatusCode, 200)
	assert.Contains(string(res.Body), message)
}
