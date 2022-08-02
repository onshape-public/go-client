# onshape_test_gen

This is the Onshape test generator. It is a simple program that automatically generates test stubs in the `onshape_test` folder. It will generate one file per Onshape API and one test per REST endpoint. If test files already exist, they will be scanned via a simple text search for endpoint names. In this case, stubs for any nonexistant endpoint tests will be appended as a comment to the bottom of the test files.

## How to use

The Onshape test generator requires as input a valid OpenAPI specification located at `openapi.json` in the root of the repository. This file should exist automatically and be up-to-date for the latest master. Then, in this folder, run the command:

`go run .`

This will scan the repository test folder, generate new tests, and output any files that were added/modified.
