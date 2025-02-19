package cypher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNumberCypher(t *testing.T) {
	testcase := []struct {
		name  string
		input string

		expect struct {
			result  string
			isError bool
		}
	}{
		{
			name:  "Happy case1",
			input: "LLRR=",

			expect: struct {
				result  string
				isError bool
			}{
				result:  "210122",
				isError: false,
			},
		}, {
			name:  "Happy case2",
			input: "==RLL",

			expect: struct {
				result  string
				isError bool
			}{
				result:  "000210",
				isError: false,
			},
		}, {
			name:  "Happy case3",
			input: "=LLRR",

			expect: struct {
				result  string
				isError bool
			}{
				result:  "221012",
				isError: false,
			},
		}, {
			name:  "Happy case4",
			input: "RRL=R",

			expect: struct {
				result  string
				isError bool
			}{
				result:  "012001",
				isError: false,
			},
		}, {
			name:  "Happy case5",
			input: "==RLL=R=",

			expect: struct {
				result  string
				isError bool
			}{
				result:  "000210011",
				isError: false,
			},
		}, {
			name:  "Happy case6",
			input: "==RLL=R=L==",

			expect: struct {
				result  string
				isError bool
			}{
				result:  "000210011000",
				isError: false,
			},
		}, {
			name:  "input error",
			input: "==RLLA",

			expect: struct {
				result  string
				isError bool
			}{
				result:  "",
				isError: true,
			},
		},
	}

	for _, v := range testcase {
		t.Run(v.name, func(t *testing.T) {
			ret, err := CreateNumberCypher(v.input)
			if !v.expect.isError {
				assert.Nil(t, err)
				assert.Equal(t, ret, v.expect.result)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
