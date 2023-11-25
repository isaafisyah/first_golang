package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("isa")

	if result != "Hello isa" {
		// panic("Result is not Hello isa")
		// t.FailNow()
		t.Error("Result is not Hello isa")
	}
}

func TestHelloWorldAssert(t *testing.T)  {
	result := HelloWorld("isa")
	assert.Equal(t, "Hello isa", result, "Result is not Hello isa")
}

func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("isa")
	require.Equal(t, "Hello isa", result, "Result is not Hello isa")
}

func TestMain(m *testing.M)  {
	//before
	fmt.Println("Before Unit Test")
	m.Run()
	//after
	fmt.Println("After Unit Test")
}
//subtest
func TestSubTest(t *testing.T)()  {
	//kalau ingin menjalankan salah satu sub test //go test -v -run TestSubTest/isa
	t.Run("isa", func(t *testing.T) {
		result := HelloWorld("isa")

		if result != "Hello isa" {

			t.Error("Result is not Hello isa")
		}
	})
	t.Run("afisyah", func(t *testing.T) {
		result := HelloWorld("afisyah")

		if result != "Hello afisyah" {
			t.Error("Result is not Hello afisyah")
		}
	})
}

//table test
func TestTableHelloWorld(t *testing.T)  {
	tests := []struct{
		name string
		request string
		expect string
	}{
		{
			name: "ISA",
			request: "ISA",
			expect: "Hello ISA",
		},
		{
			name: "AFISYAH",
			request: "AFISYAH",
			expect: "Hello AFISYAH",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expect, result)
		})
	}
}