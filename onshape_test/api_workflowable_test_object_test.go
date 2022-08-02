package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestWorkflowableTestObjectAPI(t *testing.T) {
    InitializeTester[*onshape.WorkflowableTestObjectApiService](t)

    OpenAPITest{
        Call: onshape.ApiCreateWorkflowableTestObjectRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetWorkflowableTestObjectRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdateWorkflowableTestObjectRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiTransitionWorkflowableTestObjectRequest{},
        Expect: Todo(),
    }.Execute()
    
}