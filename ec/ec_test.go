package eulerchallenge

import (
	"testing"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

var (
	runLongTests = false
)

func TestAll(t *testing.T) {
	for _, test := range []struct {
		name string
		args []string
		want []string
		long bool
	}{
		// TEST_START (needed for file_generator.go)
		{
			name: "p43",
			args: []string{"43"},
			want: []string{"16695334890"},
		},
		{
			name: "p42",
			args: []string{"42", "p42.txt"},
			want: []string{"162"},
		},
		{
			name: "p41",
			args: []string{"41"},
			want: []string{"7652413"},
		},
		{
			name: "p40",
			args: []string{"40"},
			want: []string{"210"},
		},
		{
			name: "p39",
			args: []string{"39"},
			want: []string{"840"},
		},
		{
			name: "p38",
			args: []string{"38"},
			want: []string{"932718654 9327"},
		},
		{
			name: "p37",
			args: []string{"37"},
			want: []string{"748317"},
		},
		{
			name: "p36",
			args: []string{"36", "1000000"},
			want: []string{"872187"},
		},
		{
			name: "p36 example",
			args: []string{"36", "10"},
			want: []string{"25"},
		},
		{
			name: "p35",
			args: []string{"35", "1000000"},
			want: []string{"55"},
		},
		{
			name: "p35 example",
			args: []string{"35", "100"},
			want: []string{"13"},
		},
		{
			name: "p34",
			args: []string{"34"},
			want: []string{"40730"},
		},
		{
			name: "p33",
			args: []string{"33"},
			// Answer is actually 100
			want: []string{"387296 38729600"},
		},
		{
			name: "p32",
			args: []string{"32"},
			want: []string{"45228"},
		},
		{
			name: "p31",
			args: []string{"31"},
			want: []string{"73682"},
		},
		{
			name: "p30",
			args: []string{"30", "5"},
			want: []string{"443839"},
		},
		{
			name: "p30 example",
			args: []string{"30", "4"},
			want: []string{"19316"},
		},
		{
			name: "p29",
			args: []string{"29", "100"},
			want: []string{"9183"},
		},
		{
			name: "p29 example",
			args: []string{"29", "5"},
			want: []string{"15"},
		},
		{
			name: "p28",
			args: []string{"28", "1001"},
			want: []string{"669171001"},
		},
		{
			name: "p28 example",
			args: []string{"28", "5"},
			want: []string{"101"},
		},
		{
			name: "p27",
			args: []string{"27", "1000"},
			want: []string{"-59231"},
		},
		{
			name: "p26",
			args: []string{"26", "1000"},
			want: []string{"983"},
		},
		{
			name: "p26 example",
			args: []string{"26", "10"},
			want: []string{"7"},
		},
		{
			name: "p25",
			args: []string{"25", "1000"},
			want: []string{"4782"},
		},
		{
			name: "p25 example 2",
			args: []string{"25", "2"},
			want: []string{"7"},
		},
		{
			name: "p25 example 1",
			args: []string{"25", "1"},
			want: []string{"2"},
		},
		{
			name: "p24",
			args: []string{"24", "1000000"},
			want: []string{"2783915460"},
		},
		{
			name: "p24 example 2",
			args: []string{"24", maths.Facotiral(9).Plus(maths.One()).String()},
			want: []string{"1023456789"},
		},
		{
			name: "p24 example 1",
			args: []string{"24", maths.Facotiral(9).String()},
			want: []string{"0987654321"},
		},
		{
			name: "p23",
			args: []string{"23", "28123"},
			want: []string{"4179871"},
		},
		{
			name: "p22",
			args: []string{"22", "p22.txt"},
			want: []string{"871198282"},
		},
		{
			name: "p21",
			args: []string{"21", "10000"},
			want: []string{"31626"},
		},
		{
			name: "p20",
			args: []string{"20", "100"},
			want: []string{"648"},
		},
		{
			name: "p20 example",
			args: []string{"20", "10"},
			want: []string{"27"},
		},
		{
			name: "p19",
			args: []string{"19"},
			want: []string{"171"},
		},
		{
			name: "p18 example",
			args: []string{"18", "p18_example.txt"},
			want: []string{"23"},
		},
		{
			name: "p18",
			args: []string{"18", "p18.txt"},
			want: []string{"1074"},
		},
		{
			name: "p67",
			args: []string{"18", "p67.txt"},
			want: []string{"7273"},
		},
		{
			name: "p17 example",
			args: []string{"17", "5"},
			want: []string{"19"},
		},
		{
			name: "p17",
			args: []string{"17", "1000"},
			want: []string{"21124"},
		},
		{
			name: "p16 example",
			args: []string{"16", "10"},
			want: []string{"7"},
		},
		{
			name: "p16",
			args: []string{"16", "1000"},
			want: []string{"1366"},
		},
		{
			name: "p15 example",
			args: []string{"15", "2"},
			want: []string{"6"},
		},
		{
			name: "p15",
			args: []string{"15", "20"},
			want: []string{"137846528820"},
		},
		{
			name: "p14",
			args: []string{"14", "1000000"},
			want: []string{"5537376230"},
			long: true,
		},
		{
			name: "p13",
			args: []string{"13"},
			want: []string{"5537376230"},
		},
		{
			name: "p12 example",
			args: []string{"12", "5"},
			want: []string{"28"},
		},
		{
			name: "p12",
			args: []string{"12", "500"},
			want: []string{"76576500"},
		},
		{
			name: "p11",
			args: []string{"11", "4"},
			want: []string{"70600674"},
			long: true,
		},
		{
			name: "p10 example",
			args: []string{"10", "10"},
			want: []string{"17"},
		},
		{
			name: "p10",
			args: []string{"10", "2000000"},
			want: []string{"142913828922"},
			long: true,
		},
		{
			name: "p9",
			args: []string{"9", "1000"},
			want: []string{"31875000"},
		},
		{
			name: "p8 example",
			args: []string{"8", "4"},
			want: []string{"5832"},
		},
		{
			name: "p8",
			args: []string{"8", "13"},
			want: []string{"23514624000"},
		},
		{
			name: "p7 example",
			args: []string{"7", "6"},
			want: []string{"13"},
		},
		{
			name: "p7",
			args: []string{"7", "10001"},
			want: []string{"104743"},
		},
		{
			name: "p6 example",
			args: []string{"6", "10"},
			want: []string{"2640"},
		},
		{
			name: "p6",
			args: []string{"6", "100"},
			want: []string{"25164150"},
		},
		{
			name: "p5 example",
			args: []string{"5", "10"},
			want: []string{"2520"},
		},
		{
			name: "p5",
			args: []string{"5", "20"},
			want: []string{"232792560"},
		},
		{
			name: "p4 example",
			args: []string{"4", "2"},
			want: []string{"9009"},
		},
		{
			name: "p4",
			args: []string{"4", "3"},
			want: []string{"906609"},
		},
		{
			name: "p3 example",
			args: []string{"3", "13195"},
			want: []string{"29"},
		},
		{
			name: "p3",
			args: []string{"3", "600851475143"},
			want: []string{"6857"},
		},
		{
			name: "p2",
			args: []string{"2", "4000000"},
			want: []string{"4613732"},
		},
		{
			name: "p1 example",
			args: []string{"1", "10"},
			want: []string{"23"},
		},
		{
			name: "p1",
			args: []string{"1", "1000"},
			want: []string{"233168"},
		},
		/* Useful for commenting out tests. */
	} {
		if test.long == runLongTests {
			t.Run(test.name, func(t *testing.T) {
				etc := &command.ExecuteTestCase{
					Node:          command.BranchNode(Branches(), nil, true),
					Args:          test.args,
					WantStdout:    test.want,
					SkipDataCheck: true,
				}
				command.ExecuteTest(t, etc)
			})
		}
	}
}
