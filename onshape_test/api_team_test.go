package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestTeamAPI(t *testing.T) {
    InitializeTester[*onshape.TeamAPIService](t)

    OpenAPITest{
        Call: onshape.ApiFindRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetTeamRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetMembersRequest{},
        Expect: Todo(),
    }.Execute()
    
}