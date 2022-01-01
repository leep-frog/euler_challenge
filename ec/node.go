package eulerchallenge

import "github.com/leep-frog/command"

func Branches() map[string]*command.Node {
	return map[string]*command.Node{
		"fg": FileGenerator(),
		"1":  P1(),
		"2":  P2(),
		"3":  P3(),
		"4":  P4(),
		"5":  P5(),
		"6":  P6(),
		"7":  P7(),
		"8":  P8(),
		"9":  P9(),
		"10": P10(),
		"11": P11(),
		"12": P12(),
		"13": P13(),
		"14": P14(),
		"15": P15(),
		"16": P16(),
		"17": P17(),
		"18": P18(),
		"19": P19(),
		"20": P20(),
		"21": P21(),
		"22": P22(),
		"23": P23(),
		"24": P24(),
		"25": P25(),
		"26": P26(),
		// END_LIST (needed for file_generator.go)
	}
}
