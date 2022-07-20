package onshape_test

import (
	"context"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
)

func initializeUserTests(t *testing.T) TestingContext {
	client, err := onshape.NewAPIClientFromEnv()
	assert.NoError(t, err)

	ctx := context.Background()

	return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.UserApi,
	}
}

func TestUserAPI(t *testing.T) {
	ctx := initializeUserTests(t)

	OpenAPITest{
		Call:   onshape.ApiSessionInfoRequest{},
		Expect: NoAPIError,
	}.Execute(ctx, t)
}

/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiSessionRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiGetUserSettingsCurrentLoggedInUserRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiGetUserSettingsRequest{},
    Expect: Todo,
},

***/
