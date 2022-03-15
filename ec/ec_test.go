package eulerchallenge

import (
	"testing"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

var (
	runLongTests = false
	// to keep import
	one = maths.One()
)

type codingChallengeTest struct {
	name     string
	args     []string
	want     []string
	estimate float64
	long     bool
}

func TestAll(t *testing.T) {
	tests := []*codingChallengeTest{
		// TEST_START (needed for file_generator.go)
		{
			name:     "p118",
			args:     []string{"118", "1"},
			want:     []string{"44680"},
			estimate: 7,
		},
		/*{
			name: "p117",
			args: []string{"117", "50"},
			want: []string{"100808458960497"},
		},
		{
			name: "p117 example",
			args: []string{"117", "5"},
			want: []string{"15"},
		},
		{
			name: "p116",
			args: []string{"116", "50"},
			want: []string{"20492570929"},
		},
		{
			name: "p116 example",
			args: []string{"116", "5"},
			want: []string{"12"},
		},
		{
			name: "p115",
			args: []string{"115", "50"},
			want: []string{"168"},
		},
		{
			name: "p115 example",
			args: []string{"115", "10"},
			want: []string{"57"},
		},
		{
			name: "p115 example",
			args: []string{"115", "3"},
			want: []string{"30"},
		},
		{
			name: "p114",
			args: []string{"114", "50"},
			want: []string{"16475640049"},
		},
		{
			name: "p114 example",
			args: []string{"114", "7"},
			want: []string{"17"},
		},
		{
			name: "p113",
			args: []string{"113", "100"},
			want: []string{"51161058134250"},
		},
		{
			name: "p113 example",
			args: []string{"113", "10"},
			want: []string{"277032"},
		},
		{
			name: "p113 example",
			args: []string{"113", "6"},
			want: []string{"12951"},
		},
		{
			name: "p113 example",
			args: []string{"113", "2"},
			want: []string{"99"},
		},
		{
			name: "p113 example",
			args: []string{"113", "1"},
			want: []string{"9"},
		},
		{
			name: "p112",
			args: []string{"112", "99"},
			want: []string{"1587000"},
		},
		{
			name: "p112 example",
			args: []string{"112", "90"},
			want: []string{"21780"},
		},
		{
			name: "p112 example",
			args: []string{"112", "50"},
			want: []string{"538"},
		},
		{
			name: "p111",
			args: []string{"111", "10"},
			want: []string{"612407567715"},
		},
		{
			name:     "p111 example",
			args:     []string{"111", "4"},
			want:     []string{"273700"},
			estimate: 0.25,
		},
		{
			name: "p110",
			args: []string{"108", "4000000"},
			want: []string{"9350130049860600"},
		},
		{
			name:     "p110 example",
			args:     []string{"108", "100"},
			want:     []string{"1260"},
			estimate: 2,
		},
		{
			name:     "p109",
			args:     []string{"109", "100"},
			want:     []string{"38182"},
			estimate: 1,
		},
		{
			name:     "p109 example",
			args:     []string{"109", "6"},
			want:     []string{"11"},
			estimate: 1,
		},
		{
			name:     "p108",
			args:     []string{"108", "1000"},
			want:     []string{"180180"},
			estimate: 2,
		},
		{
			name:     "p107",
			args:     []string{"107", "p107_network.txt"},
			want:     []string{"259679"},
			estimate: 0.1,
		},
		{
			name:     "p107 example",
			args:     []string{"107", "p107_example.txt"},
			want:     []string{"150"},
			estimate: 0.1,
		},
		{
			name:     "p106",
			args:     []string{"106", "12"},
			want:     []string{"21384"},
			estimate: 0.25,
		},
		{
			name:     "p106",
			args:     []string{"106", "7"},
			want:     []string{"70"},
			estimate: 0.25,
		},
		{
			name:     "p106 example",
			args:     []string{"106", "4"},
			want:     []string{"1"},
			estimate: 0.25,
		},
		{
			name: "p105",
			args: []string{"105", "p105_sets.txt"},
			want: []string{"73702"},
		},
		{
			name:     "p104",
			args:     []string{"104"},
			want:     []string{"329468"},
			estimate: 0.5,
		},
		{
			name: "p103",
			args: []string{"103", "7"},
			want: []string{"20313839404245"},
		},
		{
			name: "p103 example",
			args: []string{"103", "6"},
			want: []string{"111819202225"},
		},
		{
			name: "p103 example",
			args: []string{"103", "5"},
			want: []string{"69111213"},
		},
		{
			name: "p103 example",
			args: []string{"103", "4"},
			want: []string{"3567"},
		},
		{
			name: "p103 example",
			args: []string{"103", "3"},
			want: []string{"234"},
		},
		{
			name: "p103 example",
			args: []string{"103", "2"},
			want: []string{"12"},
		},
		{
			name: "p103 example",
			args: []string{"103", "1"},
			want: []string{"1"},
		},
		{
			name: "p102",
			args: []string{"102", "p102_triangles.txt"},
			want: []string{"228"},
		},
		{
			name:     "p101",
			args:     []string{"101", "2"},
			want:     []string{"37076114526"},
			estimate: 0.5,
		},
		{
			name:     "p101",
			args:     []string{"101", "1"},
			want:     []string{"74"},
			estimate: 0.25,
		},
		{
			name: "p100",
			args: []string{"100"},
			want: []string{"756872327473"},
		},
		{
			name: "p99",
			args: []string{"99", "p099.txt"},
			want: []string{"709"},
		},
		{
			name: "p98",
			args: []string{"98", "words.txt"},
			want: []string{"18769"},
		},
		{
			name: "p97",
			args: []string{"97"},
			want: []string{"8739992577"},
			long: true, // 16 seconds
		},
		{
			name: "p96",
			args: []string{"96", "p96.txt"},
			want: []string{"24702"},
			long: true, // sometimes infinite loops or at least long loops
		},
		{
			name: "p96 example",
			args: []string{"96", "p96_example.txt"},
			want: []string{"483"},
		},
		{
			name: "p95",
			args: []string{"95"},
			want: []string{"28 14316"},
		},
		{
			name: "p94",
			args: []string{"94"},
			want: []string{"518408346"},
		},
		{
			name: "p93",
			args: []string{"93", "9"},
			want: []string{"1258 51"},
		},
		{
			name: "p93 example",
			args: []string{"93", "4"},
			want: []string{"1234 28"},
		},
		{
			name: "p92",
			args: []string{"92"},
			want: []string{"8581146"},
			long: true, // 6 seconds
		},
		{
			name: "p91",
			args: []string{"91", "50"},
			want: []string{"14234"},
		},
		{
			name: "p91 example",
			args: []string{"91", "2"},
			want: []string{"14"},
		},
		{
			name: "p90",
			args: []string{"90"},
			want: []string{"1217"},
		},
		{
			name: "p89",
			args: []string{"89", "p89.txt"},
			want: []string{"743"},
		},
		{
			name: "p88",
			args: []string{"88", "12000"},
			want: []string{"7587457"},
			long: true, // 20 minutes
		},
		{
			name: "p88 example 2",
			args: []string{"88", "12"},
			want: []string{"61"},
		},
		{
			name: "p88 example 1",
			args: []string{"88", "6"},
			want: []string{"30"},
		},
		{
			name: "p87",
			args: []string{"87", "50000000"},
			want: []string{"1097343"},
		},
		{
			name: "p87 example",
			args: []string{"87", "50"},
			want: []string{"4"},
		},
		{
			name: "p86",
			args: []string{"86", "1000000"},
			want: []string{"1818"},
		},
		{
			name: "p86 example",
			args: []string{"86", "2000"},
			want: []string{"100"},
		},
		{
			name: "p85",
			args: []string{"85"},
			want: []string{"2772 7999992"},
		},
		{
			// ended up getting correct answer (101524) with 6 dice lol
			// so I guess bug + wrong dice = success!
			name: "p84",
			args: []string{"84", "6"},
			want: []string{"0"},
			long: true, // doesn't actually work?
		},
		{
			name: "p84 example",
			args: []string{"84", "6"},
			want: []string{"0"},
			long: true, // doesn't actually work?
		},
		{
			name: "p83",
			args: []string{"83", "p83.txt"},
			want: []string{"425185"},
		},
		{
			name: "p82",
			args: []string{"82", "p82.txt"},
			want: []string{"260324"},
		},
		{
			name: "p81 example",
			args: []string{"81", "p81_example.txt"},
			want: []string{"2427"},
		},
		{
			name: "p81",
			args: []string{"81", "p81.txt"},
			want: []string{"427337"},
		},
		{
			name: "p80",
			args: []string{"80"},
			want: []string{"40886"},
		},
		{
			name: "p79",
			args: []string{"79", "p79.txt"},
			want: []string{"73162890 true"},
		},
		{
			name: "p78",
			args: []string{"78", "1000000"},
			want: []string{"55374 36325300925435785930832331577396761646715836173633893227071086460709268608053489541731404543537668438991170680745272159154493740615385823202158167635276250554555342115855424598920159035413044811245082197335097953570911884252410730174907784762924663654000000"},
			long: true,
		},
		{
			name: "p77",
			args: []string{"77"},
			want: []string{"71 5006"},
		},
		{
			name: "p76",
			args: []string{"76", "100"},
			want: []string{"190569291"},
		},
		{
			name: "p76 example",
			args: []string{"76", "5"},
			want: []string{"6"},
		},
		{
			name: "p75",
			args: []string{"75", "1500000"},
			want: []string{"161667"},
			long: true,
		},
		{
			name: "p75 example",
			args: []string{"75", "48"},
			want: []string{"6"},
		},
		{
			name: "p74",
			args: []string{"74", "1000000"},
			want: []string{"402"},
		},
		{
			name: "p73",
			args: []string{"73", "12000"},
			want: []string{"7295372"},
		},
		{
			name: "p73 example",
			args: []string{"73", "8"},
			want: []string{"3"},
		},
		{
			name: "p72",
			args: []string{"72", "1000000"},
			want: []string{"303963552391"},
			long: true,
		},
		{
			name: "p72 example",
			args: []string{"72", "8"},
			want: []string{"21"},
		},
		{
			name: "p71",
			args: []string{"71", "1000000"},
			want: []string{"428570/999997"},
		},
		{
			name: "p71 example",
			args: []string{"71", "8"},
			want: []string{"2/5"},
		},
		{
			name: "p70",
			args: []string{"70", "10000000"},
			want: []string{"8319823"},
		},
		{
			name: "p69",
			args: []string{"69", "1000000"},
			want: []string{"510510"},
		},
		{
			name: "p69 example",
			args: []string{"69", "10"},
			want: []string{"6"},
		},
		{
			name: "p66",
			args: []string{"66", "1000"},
			want: []string{"661"},
		},
		{
			name: "p66 example",
			args: []string{"66", "7"},
			want: []string{"5"},
		},
		{
			name: "p65",
			args: []string{"65", "100"},
			want: []string{"272"},
		},
		{
			name: "p65 example",
			args: []string{"65", "10"},
			want: []string{"17"},
		},
		{
			name: "p64",
			args: []string{"64", "10000"},
			want: []string{"1322"},
		},
		{
			name: "p64 example",
			args: []string{"64", "13"},
			want: []string{"4"},
		},
		{
			name: "p63",
			args: []string{"63"},
			want: []string{"49"},
		},
		{
			name: "p62",
			args: []string{"62", "5"},
			want: []string{"127035954683"},
		},
		{
			name: "p62 example",
			args: []string{"62", "3"},
			want: []string{"41063625"},
		},
		{
			name: "p61",
			args: []string{"61", "6"},
			want: []string{"28684 [8256 5625 2512 1281 8128 2882]"},
		},
		{
			name: "p61 example",
			args: []string{"61", "3"},
			want: []string{"19291 [8128 2882 8281]"},
		},
		{
			name: "p60",
			args: []string{"60", "3"},
			want: []string{
				"792 [673 109 7 3]",
			},
		},
		{
			name: "p60",
			args: []string{"60", "4"},
			want: []string{
				"26033 [8389 6733 5701 5197 13]",
			},
		},
		{
			name: "p59",
			args: []string{"59", "p59.txt"},
			want: []string{"129448"},
		},
		{
			name: "p58",
			args: []string{"58"},
			want: []string{"26241"},
		},
		{
			name: "p57",
			args: []string{"57"},
			want: []string{"153"},
		},
		{
			name: "p56",
			args: []string{"56"},
			want: []string{"972"},
		},
		{
			name: "p55",
			args: []string{"55"},
			want: []string{"249"},
		},
		{
			name: "p54",
			args: []string{"54", "p54.txt"},
			want: []string{"376"},
		},
		{
			name: "p54",
			args: []string{"54", "p54_example.txt"},
			want: []string{"3"},
		},
		{
			name: "p53",
			args: []string{"53"},
			want: []string{"4075"},
		},
		{
			name: "p52",
			args: []string{"52", "6"},
			want: []string{"142857"},
		},
		{
			name: "p52 example",
			args: []string{"52", "2"},
			want: []string{"125874"},
		},
		{
			name: "p51",
			args: []string{"51", "8"},
			want: []string{"_2_3_3 121313"},
		},
		{
			name: "p51 example 2",
			args: []string{"51", "7"},
			want: []string{"56__3 56003"},
		},
		{
			name: "p51 example 1",
			args: []string{"51", "6"},
			want: []string{"_3 13"},
		},
		{
			name: "p50",
			args: []string{"50", "1000000"},
			want: []string{"997651 543"},
		},
		{
			name: "p50 example 2",
			args: []string{"50", "1000"},
			want: []string{"953 21"},
		},
		{
			name: "p50 example 1",
			args: []string{"50", "100"},
			want: []string{"41 6"},
		},
		{
			name: "p49",
			args: []string{"49"},
			want: []string{"148748178147", "296962999629"},
		},
		{
			name: "p48",
			args: []string{"48", "1000"},
			want: []string{"9110846700"},
		},
		{
			name: "p48 example",
			args: []string{"48", "10"},
			want: []string{"405071317"},
		},
		{
			name: "p47",
			args: []string{"47", "4"},
			want: []string{"134043"},
		},
		{
			name: "p47 example 2",
			args: []string{"47", "3"},
			want: []string{"644"},
		},
		{
			name: "p47 example 1",
			args: []string{"47", "2"},
			want: []string{"14"},
		},
		{
			name: "p46",
			args: []string{"46"},
			want: []string{"5777"},
		},
		{
			name: "p45",
			args: []string{"45"},
			want: []string{"1533776805"},
		},
		{
			name: "p44",
			args: []string{"44"},
			want: []string{"5482660"},
		},
		{
			name: "p43",
			args: []string{"43"},
			want: []string{"16695334890"},
		},
		{
			name: "p42",
			args: []string{"42", "words.txt"},
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
			args: []string{"24", maths.Factorial(9).Plus(maths.One()).String()},
			want: []string{"1023456789"},
		},
		{
			name: "p24 example 1",
			args: []string{"24", maths.Factorial(9).String()},
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
			want: []string{
				"3_2",
				"2_1",
				"1_0",
				"0_0",
				"23",
			},
		},
		{
			name: "p18",
			args: []string{"18", "p18.txt"},
			want: []string{
				"14_9",
				"13_8",
				"12_8",
				"11_7",
				"10_6",
				"9_5",
				"8_4",
				"7_3",
				"6_3",
				"5_3",
				"4_2",
				"3_2",
				"2_2",
				"1_1",
				"0_0",
				"1074",
			},
		},
		{
			name: "p67",
			args: []string{"18", "p67.txt"},
			want: []string{
				"99_53",
				"98_52",
				"97_52",
				"96_51",
				"95_50",
				"94_49",
				"93_49",
				"92_48",
				"91_47",
				"90_47",
				"89_46",
				"88_46",
				"87_46",
				"86_46",
				"85_45",
				"84_45",
				"83_45",
				"82_45",
				"81_45",
				"80_45",
				"79_44",
				"78_43",
				"77_43",
				"76_43",
				"75_42",
				"74_42",
				"73_42",
				"72_42",
				"71_42",
				"70_42",
				"69_42",
				"68_41",
				"67_41",
				"66_40",
				"65_39",
				"64_38",
				"63_37",
				"62_36",
				"61_36",
				"60_36",
				"59_36",
				"58_36",
				"57_36",
				"56_36",
				"55_35",
				"54_34",
				"53_33",
				"52_33",
				"51_32",
				"50_32",
				"49_32",
				"48_32",
				"47_31",
				"46_30",
				"45_29",
				"44_28",
				"43_27",
				"42_27",
				"41_26",
				"40_25",
				"39_25",
				"38_25",
				"37_24",
				"36_23",
				"35_22",
				"34_21",
				"33_20",
				"32_19",
				"31_18",
				"30_17",
				"29_17",
				"28_17",
				"27_16",
				"26_15",
				"25_15",
				"24_14",
				"23_14",
				"22_13",
				"21_13",
				"20_13",
				"19_12",
				"18_12",
				"17_12",
				"16_11",
				"15_10",
				"14_9",
				"13_8",
				"12_7",
				"11_6",
				"10_6",
				"9_5",
				"8_5",
				"7_4",
				"6_4",
				"5_3",
				"4_2",
				"3_1",
				"2_0",
				"1_0",
				"0_0",
				"7273",
			},
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
			args: []string{"13", "p13.txt"},
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
	}

	for _, test := range tests {
		test.test(t)
	}
}

func (ct *codingChallengeTest) test(t *testing.T) {
	// TODO: have time cutoff instead
	if ct.long {
		return
	}
	t.Run(ct.name, func(t *testing.T) {
		if ct.estimate == 0 {
			t.Logf("Estimate is not set")
		}
		etc := &command.ExecuteTestCase{
			Node:          command.BranchNode(Branches(), nil),
			Args:          ct.args,
			WantStdout:    ct.want,
			SkipDataCheck: true,
		}
		command.ExecuteTest(t, etc)
	})
}
