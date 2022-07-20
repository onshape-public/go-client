package onshape_test

import (
	"context"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
)

func initializeAssemblyTests(t *testing.T) TestingContext {
	client, err := onshape.NewAPIClientFromEnv()
	assert.NoError(t, err)

	ctx := context.Background()

	return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.AssemblyApi,
	}
}

func TestAssemblyAPI(t *testing.T) {
	ctx := initializeAssemblyTests(t)

	OpenAPITest{
		Call:   onshape.ApiGetNamedViewsRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiCreateAssemblyRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetOrCreateBillOfMaterialsElementRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiUpdateFeatureRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiDeleteFeatureRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiDeleteInstanceRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiCreateInstanceRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiTransformOccurrencesRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiInsertTransformedInstancesRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetAssemblyDefinitionRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetBillOfMaterialsRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetAssemblyBoundingBoxesRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetExplodedViewsRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetFeaturesRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiAddFeatureRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetFeatureSpecsRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetAssemblyMassPropertiesRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetNamedPositionsRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetAssemblyShadedViewsRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiTranslateFormatRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

}
