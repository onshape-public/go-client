package onshape_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/require"
)

var tester TestingInstance
var mapper *ResultMapper = &ResultMapper{}

func init() {
	populateDefaultResultMapper()
}

func populateDefaultResultMapper() {
	DefaultResultMapper().Add(MapFields(func(o *onshape.BTDocumentInfo) TestingContext {
		return TestingContext{
			"did":     o.GetId(),
			"wid":     *o.GetDefaultWorkspace().Id,
			"wmid":    *o.GetDefaultWorkspace().Id,
			"wvid":    *o.GetDefaultWorkspace().Id,
			"wvmid":   *o.GetDefaultWorkspace().Id,
			"mv":      *o.GetDefaultWorkspace().Microversion,
			"owner":   o.GetOwner().Id,
			"creator": Ptr(Ptr(o.GetCreatedBy()).GetId()),
		}
	})).Add(MapFields(func(o *onshape.BTDocumentSummaryInfo) TestingContext {
		return TestingContext{
			"did":     o.GetId(),
			"wid":     *o.GetDefaultWorkspace().Id,
			"wmid":    *o.GetDefaultWorkspace().Id,
			"wvid":    *o.GetDefaultWorkspace().Id,
			"wvmid":   *o.GetDefaultWorkspace().Id,
			"mv":      *o.GetDefaultWorkspace().Microversion,
			"owner":   o.GetOwner().Id,
			"creator": Ptr(Ptr(o.GetCreatedBy()).GetId()),
		}
	})).Add(MapFields(func(o *onshape.BTWorkspaceInfo) TestingContext {
		return TestingContext{
			"wid":   o.GetId(),
			"wmid":  o.GetId(),
			"wvid":  o.GetId(),
			"wvmid": o.GetId(),
			"mv":    o.GetMicroversion(),
		}
	})).Add(MapFields(func(o *onshape.BTRestoreFromHistoryInfo) TestingContext {
		return TestingContext{
			"mv": o.GetNewMicroversion(),
		}
	})).Add(MapFields(func(o *onshape.BTMicroversionInfo) TestingContext {
		return TestingContext{
			"mv": o.GetMicroversion(),
		}
	})).Add(MapFields(func(o *onshape.BTAppElementModifyInfo) TestingContext {
		return TestingContext{
			"eid": o.GetElementId(),
			"tid": o.GetTransactionId(),
		}
	}))
}

type TestingInstance struct {
	context TestingContext
	tests   []*testing.T
}

func InitializeTester[T any](t *testing.T) {
	client, err := onshape.NewAPIClientFromEnv()
	require.NoError(t, err)

	ctx := context.Background()

	tester = TestingInstance{
		context: TestingContext{
			"client":     client,
			"ctx":        ctx,
			"ApiService": getAPIForOnshape[T](client),
		},
		tests: []*testing.T{t},
	}
}

func getAPIForOnshape[T any](client *onshape.APIClient) T {
	rv := reflect.ValueOf(client).Elem()

	for i := 0; i < rv.NumField(); i++ {
		fi := rv.Field(i)
		rf := reflect.NewAt(fi.Type(), unsafe.Pointer(fi.UnsafeAddr())).Elem()
		if val, ok := rf.Interface().(T); ok {
			return val
		}
	}

	panic("Couldn't find API")
}

func DefaultResultMapper() *ResultMapper {
	return mapper
}

func Context() TestingContext {
	return tester.context
}

func SetContext(ctx TestingContext) {
	tester.context = ctx
}

func Tester() *testing.T {
	l := len(tester.tests)
	return tester.tests[l-1]
}

func (t *TestingInstance) pushTest(q *testing.T) {
	t.tests = append(t.tests, q)
}

func (t *TestingInstance) popTest() *testing.T {
	l := len(t.tests)
	ret := Tester()
	t.tests = t.tests[:l-1]
	return ret
}

type TestingContext map[string]interface{}

func (o TestingContext) WithDefault(name string, value interface{}) TestingContext {
	return TestingContext{name: value}.InheritDefaults(o)
}

func (o TestingContext) InheritDefaults(context TestingContext) TestingContext {
	if o == nil {
		o = make(map[string]interface{})
	}

	vals := make(map[string]interface{})

	for k, v := range context {
		vals[k] = v
	}

	for k, v := range o {
		vals[k] = v
	}

	return vals
}

func ApplyDefaultsToObject(tc TestingContext, o interface{}) interface{} {
	vp := reflect.New(reflect.TypeOf(o))
	val := vp.Elem()
	val.Set(reflect.ValueOf(o))

	for i := 0; i < val.NumField(); i++ {
		fi := val.Field(i)
		name := val.Type().Field(i).Name
		rf := reflect.NewAt(fi.Type(), unsafe.Pointer(fi.UnsafeAddr())).Elem()
		if rf.IsZero() {
			if v, ok := tc[name]; ok && rf.CanSet() {
				rf.Set(reflect.ValueOf(v))
			}
		}
	}

	return vp.Interface()
}

type OpenAPIResultHandler func(value interface{}, response *http.Response, err error)

func APIError() OpenAPIResultHandler {
	return func(value interface{}, response *http.Response, err error) {
		require.Error(Tester(), err)
	}
}

func NoAPIError() OpenAPIResultHandler {
	return func(value interface{}, response *http.Response, err error) {
		require.NoError(Tester(), err)
	}
}

func NoAPIErrorAnd[T any](check func(value T)) OpenAPIResultHandler {
	return func(value interface{}, response *http.Response, err error) {
		require.NoError(Tester(), err)
		val, ok := value.(T)
		require.True(Tester(), ok)
		if ok && err == nil {
			check(val)
		}
	}
}

func Nothing() OpenAPIResultHandler {
	return func(value interface{}, response *http.Response, err error) {
	}
}

func MapResult() OpenAPIResultHandler {
	return func(value interface{}, response *http.Response, err error) {
		DefaultResultMapper().MapResult(value)
	}
}

func Spew() OpenAPIResultHandler {
	return func(value interface{}, response *http.Response, err error) {
		spew.Dump("Spew(): "+Tester().Name(), value, response, err)
	}
}

var todo = func(value interface{}, response *http.Response, err error) {
	spew.Dump("Spew(): "+Tester().Name(), value, response, err)
}

func Todo() OpenAPIResultHandler {
	return todo
}

type OpenAPITest struct {
	Name    string
	Call    interface{}
	Context TestingContext
	Expect  OpenAPIResultHandler
	Pass    OpenAPIResultHandler
	Fail    OpenAPIResultHandler
}

func (o OpenAPITest) Execute() bool {
	call := ApplyDefaultsToObject(o.Context.InheritDefaults(Context()), o.Call)

	name := o.Name

	if name == "" {
		name = "Test " + reflect.TypeOf(call).Elem().Name()
	}

	if o.Pass == nil {
		o.Pass = MapResult()
	}
	if o.Fail == nil {
		o.Fail = Spew()
	}
	if o.Expect == nil {
		o.Expect = NoAPIError()
	}

	var value interface{}
	var response *http.Response
	var err error

	ret := Tester().Run(name, func(t *testing.T) {
		tester.pushTest(t)
		validator := o.Expect

		if reflect.ValueOf(validator).Pointer() == reflect.ValueOf(Todo()).Pointer() {
			t.SkipNow()
		}

		rs := reflect.ValueOf(call).MethodByName("Execute").Call([]reflect.Value{})
		value = rs[0].Interface()
		response = rs[1].Interface().(*http.Response)
		err, _ = rs[2].Interface().(error)
		o.Expect(value, response, err)
	})
	defer tester.popTest()

	if ret {
		o.Pass(value, response, err)
	} else {
		o.Fail(value, response, err)
	}

	return ret
}

type ResultMapping func(value interface{})

func MapFields[T any](fm func(T) TestingContext) ResultMapping {
	return func(value interface{}) {
		if t, ok := value.(T); ok {
			SetContext(fm(t).InheritDefaults(Context()))
		}
	}
}

type ResultMapper struct {
	maps []ResultMapping
}

func (m *ResultMapper) Add(mapping ResultMapping) *ResultMapper {
	m.maps = append(m.maps, mapping)
	return m
}

func (m *ResultMapper) MapResult(value interface{}) {
	for _, mapper := range m.maps {
		mapper(value)
	}
}

func Ptr[T any](val T) *T {
	return &val
}
