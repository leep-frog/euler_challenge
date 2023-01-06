package commandths

import (
	"fmt"
	"strings"
	"testing"

	"github.com/leep-frog/command"
)

func TestExecute(t *testing.T) {
	errFunc := func(s string) string {
		return fmt.Sprintf("failed to parse int: strconv.Atoi: parsing %q: invalid syntax", s)
	}
	for _, test := range []struct {
		name string
		etc  *command.ExecuteTestCase
	}{
		// Math tests
		{
			name: "Handles single number",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17"},
				WantStdout: "17",
			},
		},
		{
			name: "Adds two numbers",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17 + 19"},
				WantStdout: "36",
			},
		},
		{
			name: "Adds two numbers",
			etc: &command.ExecuteTestCase{
				Args:       []string{"13+23"},
				WantStdout: "36",
			},
		},
		{
			name: "Adds two numbers with letter code",
			etc: &command.ExecuteTestCase{
				Args:       []string{"78 p 35"},
				WantStdout: "113",
			},
		},
		{
			name: "Subtracts two numbers",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17 - 19"},
				WantStdout: "-2",
			},
		},
		/*{
			name: "Subtracts two numbers with no spaces",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17 - 19"},
				WantStdout: "-2",
			},
		},*/
		{
			name: "Subtracts two numbers with letter code",
			etc: &command.ExecuteTestCase{
				Args:       []string{"486 - 168"},
				WantStdout: "318",
			},
		},
		{
			name: "Multiplies two numbers",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17 * -19"},
				WantStdout: "-323",
			},
		},
		{
			name: "Multiplies two numbers with no spaces",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-17*-19"},
				WantStdout: "323",
			},
		},
		{
			name: "Multiplies two numbers with letter code",
			etc: &command.ExecuteTestCase{
				Args:       []string{"13 t 10"},
				WantStdout: "130",
			},
		},
		{
			name: "Divides two numbers",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17 / 5"},
				WantStdout: "3",
			},
		},
		{
			name: "Divides two numbers with no spaces",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-17/3"},
				WantStdout: "-5",
			},
		},
		{
			name: "Divides two numbers with letter code",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-17 d -3"},
				WantStdout: "5",
			},
		},
		{
			name: "Modulos two numbers",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17 % 3"},
				WantStdout: "2",
			},
		},
		{
			name: "Modulos two numbers with no spaces",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17%3"},
				WantStdout: "2",
			},
		},
		{
			name: "Modulos two numbers with letter code",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17 o 3"},
				WantStdout: "2",
			},
		},
		{
			name: "Exponentiates two numbers",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17 ^ 3"},
				WantStdout: "4913",
			},
		},
		{
			name: "Exponentiates two numbers with no spaces",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3^17"},
				WantStdout: "129140163",
			},
		},
		{
			name: "Exponentiates negative even times",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-17 ^ 4"},
				WantStdout: "83521",
			},
		},
		{
			name: "Exponentiates negative odd times",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-17 ^ 3"},
				WantStdout: "-4913",
			},
		},
		{
			name: "Handles parens with single number",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-17 + (24)"},
				WantStdout: "7",
			},
		},
		{
			name: "Fails if close paren after operation",
			etc: &command.ExecuteTestCase{
				// Args: []string{"-17 + (24 -) 3"},
				Args: []string{"(24 -)"},
				// WantStdout: "7",
			},
		},
		{
			name: "Handles parens with expression number",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-17 + (-12 + 30)"},
				WantStdout: "1",
			},
		},
		{
			name: "Handles nested parens",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-17 + (-12 + ((3) * ((7))))"},
				WantStdout: "-8",
			},
		},
		{
			name: "Handles underscores parens",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-1_7 + (-1_2 + ((3) * ((7))))"},
				WantStdout: "-8",
			},
		},
		{
			name: "Fails for leading underscore",
			etc: &command.ExecuteTestCase{
				Args:       []string{"_17"},
				WantStderr: errFunc("_17"),
			},
		},
		{
			name: "Fails for trailing underscore",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17_"},
				WantStderr: errFunc("17_"),
			},
		},
		{
			name: "Handles starting parentheses",
			etc: &command.ExecuteTestCase{
				Args:       []string{"(-17)"},
				WantStdout: "-17",
			},
		},
		{
			name: "Unexpected close paren",
			etc: &command.ExecuteTestCase{
				Args:       []string{"-17)"},
				WantStderr: "unexpected close parentheses",
			},
		},
		{
			name: "Fails if adjacent operators",
			etc: &command.ExecuteTestCase{
				Args:       []string{"17 + + 12"},
				WantStderr: `consecutive operations`,
			},
		},
		{
			name: "Fails if adjacent numbers",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 + (17 + 12 7)"},
				WantStderr: `unexpected number 7`,
			},
		},
		// PEMDAS tests
		{
			name: "add then subtract in order",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 + 7 - 4"},
				WantStdout: "6",
			},
		},
		{
			name: "subtract then add in order",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 - 7 + 4"},
				WantStdout: "0",
			},
		},
		{
			name: "multiply before add",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 + 7 * 4"},
				WantStdout: "31",
			},
		},
		{
			name: "multiply before add",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 * 7 + 4"},
				WantStdout: "25",
			},
		},
		{
			name: "divide before add",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 + 17 / 2"},
				WantStdout: "11",
			},
		},
		{
			name: "divide before add",
			etc: &command.ExecuteTestCase{
				Args:       []string{"12 / 3 + 4"},
				WantStdout: "8",
			},
		},
		{
			name: "multiply before subtract",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 - 7 * 4"},
				WantStdout: "-25",
			},
		},
		{
			name: "multiply before subtract",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 * 7 - 4"},
				WantStdout: "17",
			},
		},
		{
			name: "divide before subtract",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 - 17 / 2"},
				WantStdout: "-5",
			},
		},
		{
			name: "divide before subtract",
			etc: &command.ExecuteTestCase{
				Args:       []string{"12 / 3 - 4"},
				WantStdout: "0",
			},
		},
		{
			name: "exp before multiply",
			etc: &command.ExecuteTestCase{
				Args:       []string{"3 * 11 ^ 2"},
				WantStdout: "363",
			},
		},
		{
			name: "exp before divide",
			etc: &command.ExecuteTestCase{
				Args:       []string{"352 / 2 ^ 5"},
				WantStdout: "11",
			},
		},
		// TODO: Separate arg tests.
		// Prime factor tests
		{
			name: "Factors a number",
			etc: &command.ExecuteTestCase{
				Args:       []string{"prime", "factor", "352"},
				WantStdout: "352: 2^5 * 11^1",
			},
		},
		{
			name: "Factors multiple numbers",
			etc: &command.ExecuteTestCase{
				Args: []string{"prime", "factor", "2048", "3125", "17"},
				WantStdout: strings.Join([]string{
					"2048: 2^11",
					"3125: 5^5",
					"17: 17^1",
				}, "\n"),
			},
		},
		// Nth prime tests
		{
			name: "Gets the 1st prime",
			etc: &command.ExecuteTestCase{
				Args:       []string{"prime", "nth", "1"},
				WantStdout: "2",
			},
		},
		{
			name: "Gets the 1001th prime",
			etc: &command.ExecuteTestCase{
				Args:       []string{"prime", "nth", "1001"},
				WantStdout: "7927",
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			test.etc.Node = CLI().Node()
			test.etc.SkipDataCheck = true
			if test.etc.WantStdout != "" {
				test.etc.WantStdout += "\n"
			}
			if test.etc.WantStderr != "" {
				test.etc.WantErr = fmt.Errorf(test.etc.WantStderr)
				test.etc.WantStderr += "\n"
			}
			command.ExecuteTest(t, test.etc)
		})
	}
}
