package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P59() *problem {
	return fileInputNode(59, func(lines []string, o command.Output) {
		charStrs := strings.Split(lines[0], ",")
		var chars []int
		for _, s := range charStrs {
			chars = append(chars, parse.Atoi(s))
		}
		lower := strings.ToLower(letters)
		for _, a := range lower {
			for _, b := range lower {
				for _, c := range lower {
					/*code := maths.ToBinary(int(a)).Concat(maths.ToBinary(int(b))).Concat(maths.ToBinary(int(c)))
					for code.Len() < len()*/
					/*codes := []*maths.Binary{
						maths.ToBinary(int(a)),
						maths.ToBinary(int(b)),
						maths.ToBinary(int(c)),
					}*/
					codes := []int{
						int(a),
						int(b),
						int(c),
					}
					var codeIdx int
					var decoded []rune
					for _, char := range chars {
						decoded = append(decoded, rune(maths.XOR(codes[codeIdx], char)))
						codeIdx = (codeIdx + 1) % len(codes)
					}
					sd := string(decoded)
					if strings.Contains(strings.ToLower(sd), " and ") {
						var sum int
						for _, s := range sd {
							sum += int(s)
						}
						o.Stdoutln(sum)
						return
					}
				}
			}
		}
		o.Stderrln("nope", len(chars))
	}, []*execution{
		{
			args:     []string{"p59.txt"},
			want:     "129448",
			estimate: 1,
		},
	})
}
