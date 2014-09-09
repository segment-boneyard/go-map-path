package path

import "github.com/bmizerany/assert"
import "testing"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestPath(t *testing.T) {
	m := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": map[string]interface{}{
				"baz": "hello",
			},
		},
	}

	v := Path(m, "foo.bar.baz")
	assert.Equal(t, "hello", v)
}

func TestPathRootNil(t *testing.T) {
	m := map[string]interface{}{}
	v := Path(m, "foo.bar.baz.some.other.thing")
	assert.Equal(t, nil, v)
}

func TestPathNil(t *testing.T) {
	m := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": map[string]interface{}{
				"baz": "hello",
			},
		},
	}

	v := Path(m, "foo.bar.baz.some.other.thing")
	assert.Equal(t, nil, v)
}
