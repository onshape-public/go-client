package onshape_test

import (
	"net/http"
	"reflect"
	"testing"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

type TestingContext map[string]interface{}

func (o TestingContext) SetDefault(name string, value interface{}) TestingContext {
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
			if v, ok := tc[name]; ok {
				rf.Set(reflect.ValueOf(v))
			}
		}
	}

	return vp.Interface()
}

type TemplatableTest interface {
	Execute(ctx TestingContext, t *testing.T) bool
}

func Ptr[T any](val T) *T {
	return &val
}

type OpenAPITest struct {
	Name   string
	Call   interface{}
	Expect interface{}
	Pass   interface{}
	Fail   interface{}
}

func (o OpenAPITest) Execute(ctx TestingContext, t *testing.T) bool {
	call := ApplyDefaultsToObject(ctx, o.Call)

	name := o.Name

	if name == "" {
		name = "Test " + reflect.TypeOf(call).Elem().Name()
	}

	if o.Pass == nil {
		o.Pass = Nothing
	}
	if o.Fail == nil {
		o.Fail = Spew
	}
	if o.Expect == nil {
		o.Expect = NoAPIError
	}

	var rs []reflect.Value

	ret := t.Run(name, func(t *testing.T) {
		validator := o.Expect

		if reflect.ValueOf(validator).Pointer() == reflect.ValueOf(Todo).Pointer() {
			t.SkipNow()
		}

		rs = reflect.ValueOf(call).MethodByName("Execute").Call([]reflect.Value{})
		rs = append([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(t)}, rs...)

		reflect.ValueOf(o.Expect).Call(rs)
	})

	if reflect.ValueOf(o.Expect).Pointer() != reflect.ValueOf(Todo).Pointer() {
		if ret {
			reflect.ValueOf(o.Pass).Call(rs)
		} else {
			reflect.ValueOf(o.Fail).Call(rs)
		}
	}

	return ret
}

func NoAPIError(ctx TestingContext, t *testing.T, _ interface{}, _ *http.Response, err error) {
	assert.NoError(t, err)
}

func APIError(ctx TestingContext, t *testing.T, _ interface{}, _ *http.Response, err error) {
	assert.Error(t, err)
}

func Value[T any](val T) func(TestingContext, *testing.T, T, *http.Response, error) {
	return func(ctx TestingContext, t *testing.T, cmp T, _ *http.Response, err error) {
		assert.NoError(t, err)
		assert.EqualValues(t, val, cmp)
	}
}

func Spew(_ TestingContext, t *testing.T, _ interface{}, v *http.Response, _ error) {
	t.Log(spew.Sdump(v))
}

func Todo(ctx TestingContext, t *testing.T, _ interface{}, _ *http.Response, err error) {
	t.SkipNow()
}

func Nothing(_, _, _, _, _ interface{}) {}
