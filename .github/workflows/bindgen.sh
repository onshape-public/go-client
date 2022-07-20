if [ -z "$repo" ]; then
  repo=$PWD
fi

if [ -z "$packageVersion" ]; then
  packageVersion=0.0.0
fi

java -cp ${repo}/go-oapi-codegen.jar:${repo}/openapi-generator-cli.jar org.openapitools.codegen.OpenAPIGenerator generate -i ${repo}/openapi.json.tmp -g go-oapi-codegen -o ${repo}/onshape --type-mappings DateTime=JSONTime --additional-properties=packageVersion=${packageVersion} --additional-properties=useOneOfDiscriminatorLookup=true -c ${repo}/openapi_config.json
go fmt ${repo}/onshape
