package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestUserAPI(t *testing.T) {
    InitializeTester[*onshape.UserAPIService](t)

    OpenAPITest{
        Call: onshape.ApiSessionRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiSessionInfoRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetUserSettingsCurrentLoggedInUserRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetUserSettingsRequest{},
        Expect: Todo(),
    }.Execute()
    
}