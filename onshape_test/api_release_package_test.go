package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestReleasePackageAPI(t *testing.T) {
    InitializeTester[*onshape.ReleasePackageApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetCompanyReleaseWorkflowRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCreateObsoletionPackageRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCreateReleasePackageRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetReleasePackageRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdateReleasePackageRequest{},
        Expect: Todo(),
    }.Execute()
    
}