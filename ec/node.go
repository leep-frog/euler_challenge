package eulerchallenge

import (
	"fmt"
	"log"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p1"
	"github.com/leep-frog/euler_challenge/ec/p10"
	"github.com/leep-frog/euler_challenge/ec/p100"
	"github.com/leep-frog/euler_challenge/ec/p101"
	"github.com/leep-frog/euler_challenge/ec/p102"
	"github.com/leep-frog/euler_challenge/ec/p103"
	"github.com/leep-frog/euler_challenge/ec/p104"
	"github.com/leep-frog/euler_challenge/ec/p105"
	"github.com/leep-frog/euler_challenge/ec/p106"
	"github.com/leep-frog/euler_challenge/ec/p107"
	"github.com/leep-frog/euler_challenge/ec/p108"
	"github.com/leep-frog/euler_challenge/ec/p109"
	"github.com/leep-frog/euler_challenge/ec/p11"
	"github.com/leep-frog/euler_challenge/ec/p111"
	"github.com/leep-frog/euler_challenge/ec/p112"
	"github.com/leep-frog/euler_challenge/ec/p113"
	"github.com/leep-frog/euler_challenge/ec/p114"
	"github.com/leep-frog/euler_challenge/ec/p115"
	"github.com/leep-frog/euler_challenge/ec/p116"
	"github.com/leep-frog/euler_challenge/ec/p117"
	"github.com/leep-frog/euler_challenge/ec/p118"
	"github.com/leep-frog/euler_challenge/ec/p119"
	"github.com/leep-frog/euler_challenge/ec/p12"
	"github.com/leep-frog/euler_challenge/ec/p120"
	"github.com/leep-frog/euler_challenge/ec/p121"
	"github.com/leep-frog/euler_challenge/ec/p122"
	"github.com/leep-frog/euler_challenge/ec/p123"
	"github.com/leep-frog/euler_challenge/ec/p124"
	"github.com/leep-frog/euler_challenge/ec/p125"
	"github.com/leep-frog/euler_challenge/ec/p126"
	"github.com/leep-frog/euler_challenge/ec/p127"
	"github.com/leep-frog/euler_challenge/ec/p128"
	"github.com/leep-frog/euler_challenge/ec/p129"
	"github.com/leep-frog/euler_challenge/ec/p13"
	"github.com/leep-frog/euler_challenge/ec/p130"
	"github.com/leep-frog/euler_challenge/ec/p131"
	"github.com/leep-frog/euler_challenge/ec/p132"
	"github.com/leep-frog/euler_challenge/ec/p133"
	"github.com/leep-frog/euler_challenge/ec/p134"
	"github.com/leep-frog/euler_challenge/ec/p135"
	"github.com/leep-frog/euler_challenge/ec/p136"
	"github.com/leep-frog/euler_challenge/ec/p137"
	"github.com/leep-frog/euler_challenge/ec/p138"
	"github.com/leep-frog/euler_challenge/ec/p139"
	"github.com/leep-frog/euler_challenge/ec/p14"
	"github.com/leep-frog/euler_challenge/ec/p140"
	"github.com/leep-frog/euler_challenge/ec/p141"
	"github.com/leep-frog/euler_challenge/ec/p142"
	"github.com/leep-frog/euler_challenge/ec/p143"
	"github.com/leep-frog/euler_challenge/ec/p144"
	"github.com/leep-frog/euler_challenge/ec/p145"
	"github.com/leep-frog/euler_challenge/ec/p146"
	"github.com/leep-frog/euler_challenge/ec/p147"
	"github.com/leep-frog/euler_challenge/ec/p148"
	"github.com/leep-frog/euler_challenge/ec/p149"
	"github.com/leep-frog/euler_challenge/ec/p15"
	"github.com/leep-frog/euler_challenge/ec/p150"
	"github.com/leep-frog/euler_challenge/ec/p151"
	"github.com/leep-frog/euler_challenge/ec/p152"
	"github.com/leep-frog/euler_challenge/ec/p153"
	"github.com/leep-frog/euler_challenge/ec/p154"
	"github.com/leep-frog/euler_challenge/ec/p155"
	"github.com/leep-frog/euler_challenge/ec/p156"
	"github.com/leep-frog/euler_challenge/ec/p157"
	"github.com/leep-frog/euler_challenge/ec/p158"
	"github.com/leep-frog/euler_challenge/ec/p159"
	"github.com/leep-frog/euler_challenge/ec/p16"
	"github.com/leep-frog/euler_challenge/ec/p160"
	"github.com/leep-frog/euler_challenge/ec/p161"
	"github.com/leep-frog/euler_challenge/ec/p162"
	"github.com/leep-frog/euler_challenge/ec/p163"
	"github.com/leep-frog/euler_challenge/ec/p164"
	"github.com/leep-frog/euler_challenge/ec/p165"
	"github.com/leep-frog/euler_challenge/ec/p166"
	"github.com/leep-frog/euler_challenge/ec/p167"
	"github.com/leep-frog/euler_challenge/ec/p168"
	"github.com/leep-frog/euler_challenge/ec/p169"
	"github.com/leep-frog/euler_challenge/ec/p17"
	"github.com/leep-frog/euler_challenge/ec/p170"
	"github.com/leep-frog/euler_challenge/ec/p171"
	"github.com/leep-frog/euler_challenge/ec/p172"
	"github.com/leep-frog/euler_challenge/ec/p173"
	"github.com/leep-frog/euler_challenge/ec/p174"
	"github.com/leep-frog/euler_challenge/ec/p175"
	"github.com/leep-frog/euler_challenge/ec/p176"
	"github.com/leep-frog/euler_challenge/ec/p177"
	"github.com/leep-frog/euler_challenge/ec/p178"
	"github.com/leep-frog/euler_challenge/ec/p179"
	"github.com/leep-frog/euler_challenge/ec/p18"
	"github.com/leep-frog/euler_challenge/ec/p181"
	"github.com/leep-frog/euler_challenge/ec/p184"
	"github.com/leep-frog/euler_challenge/ec/p187"
	"github.com/leep-frog/euler_challenge/ec/p188"
	"github.com/leep-frog/euler_challenge/ec/p19"
	"github.com/leep-frog/euler_challenge/ec/p191"
	"github.com/leep-frog/euler_challenge/ec/p2"
	"github.com/leep-frog/euler_challenge/ec/p20"
	"github.com/leep-frog/euler_challenge/ec/p203"
	"github.com/leep-frog/euler_challenge/ec/p204"
	"github.com/leep-frog/euler_challenge/ec/p205"
	"github.com/leep-frog/euler_challenge/ec/p206"
	"github.com/leep-frog/euler_challenge/ec/p21"
	"github.com/leep-frog/euler_challenge/ec/p22"
	"github.com/leep-frog/euler_challenge/ec/p222"
	"github.com/leep-frog/euler_challenge/ec/p23"
	"github.com/leep-frog/euler_challenge/ec/p233"
	"github.com/leep-frog/euler_challenge/ec/p234"
	"github.com/leep-frog/euler_challenge/ec/p235"
	"github.com/leep-frog/euler_challenge/ec/p236"
	"github.com/leep-frog/euler_challenge/ec/p24"
	"github.com/leep-frog/euler_challenge/ec/p243"
	"github.com/leep-frog/euler_challenge/ec/p25"
	"github.com/leep-frog/euler_challenge/ec/p252"
	"github.com/leep-frog/euler_challenge/ec/p26"
	"github.com/leep-frog/euler_challenge/ec/p27"
	"github.com/leep-frog/euler_challenge/ec/p28"
	"github.com/leep-frog/euler_challenge/ec/p29"
	"github.com/leep-frog/euler_challenge/ec/p3"
	"github.com/leep-frog/euler_challenge/ec/p30"
	"github.com/leep-frog/euler_challenge/ec/p301"
	"github.com/leep-frog/euler_challenge/ec/p31"
	"github.com/leep-frog/euler_challenge/ec/p32"
	"github.com/leep-frog/euler_challenge/ec/p33"
	"github.com/leep-frog/euler_challenge/ec/p333"
	"github.com/leep-frog/euler_challenge/ec/p34"
	"github.com/leep-frog/euler_challenge/ec/p345"
	"github.com/leep-frog/euler_challenge/ec/p346"
	"github.com/leep-frog/euler_challenge/ec/p347"
	"github.com/leep-frog/euler_challenge/ec/p35"
	"github.com/leep-frog/euler_challenge/ec/p357"
	"github.com/leep-frog/euler_challenge/ec/p36"
	"github.com/leep-frog/euler_challenge/ec/p37"
	"github.com/leep-frog/euler_challenge/ec/p38"
	"github.com/leep-frog/euler_challenge/ec/p381"
	"github.com/leep-frog/euler_challenge/ec/p387"
	"github.com/leep-frog/euler_challenge/ec/p39"
	"github.com/leep-frog/euler_challenge/ec/p4"
	"github.com/leep-frog/euler_challenge/ec/p40"
	"github.com/leep-frog/euler_challenge/ec/p401"
	"github.com/leep-frog/euler_challenge/ec/p41"
	"github.com/leep-frog/euler_challenge/ec/p42"
	"github.com/leep-frog/euler_challenge/ec/p43"
	"github.com/leep-frog/euler_challenge/ec/p44"
	"github.com/leep-frog/euler_challenge/ec/p45"
	"github.com/leep-frog/euler_challenge/ec/p456"
	"github.com/leep-frog/euler_challenge/ec/p458"
	"github.com/leep-frog/euler_challenge/ec/p46"
	"github.com/leep-frog/euler_challenge/ec/p47"
	"github.com/leep-frog/euler_challenge/ec/p48"
	"github.com/leep-frog/euler_challenge/ec/p49"
	"github.com/leep-frog/euler_challenge/ec/p493"
	"github.com/leep-frog/euler_challenge/ec/p5"
	"github.com/leep-frog/euler_challenge/ec/p50"
	"github.com/leep-frog/euler_challenge/ec/p500"
	"github.com/leep-frog/euler_challenge/ec/p501"
	"github.com/leep-frog/euler_challenge/ec/p51"
	"github.com/leep-frog/euler_challenge/ec/p52"
	"github.com/leep-frog/euler_challenge/ec/p53"
	"github.com/leep-frog/euler_challenge/ec/p54"
	"github.com/leep-frog/euler_challenge/ec/p55"
	"github.com/leep-frog/euler_challenge/ec/p56"
	"github.com/leep-frog/euler_challenge/ec/p57"
	"github.com/leep-frog/euler_challenge/ec/p58"
	"github.com/leep-frog/euler_challenge/ec/p59"
	"github.com/leep-frog/euler_challenge/ec/p6"
	"github.com/leep-frog/euler_challenge/ec/p60"
	"github.com/leep-frog/euler_challenge/ec/p601"
	"github.com/leep-frog/euler_challenge/ec/p61"
	"github.com/leep-frog/euler_challenge/ec/p62"
	"github.com/leep-frog/euler_challenge/ec/p63"
	"github.com/leep-frog/euler_challenge/ec/p64"
	"github.com/leep-frog/euler_challenge/ec/p65"
	"github.com/leep-frog/euler_challenge/ec/p66"
	"github.com/leep-frog/euler_challenge/ec/p69"
	"github.com/leep-frog/euler_challenge/ec/p7"
	"github.com/leep-frog/euler_challenge/ec/p70"
	"github.com/leep-frog/euler_challenge/ec/p700"
	"github.com/leep-frog/euler_challenge/ec/p701"
	"github.com/leep-frog/euler_challenge/ec/p71"
	"github.com/leep-frog/euler_challenge/ec/p719"
	"github.com/leep-frog/euler_challenge/ec/p72"
	"github.com/leep-frog/euler_challenge/ec/p73"
	"github.com/leep-frog/euler_challenge/ec/p74"
	"github.com/leep-frog/euler_challenge/ec/p75"
	"github.com/leep-frog/euler_challenge/ec/p751"
	"github.com/leep-frog/euler_challenge/ec/p76"
	"github.com/leep-frog/euler_challenge/ec/p77"
	"github.com/leep-frog/euler_challenge/ec/p78"
	"github.com/leep-frog/euler_challenge/ec/p79"
	"github.com/leep-frog/euler_challenge/ec/p8"
	"github.com/leep-frog/euler_challenge/ec/p80"
	"github.com/leep-frog/euler_challenge/ec/p808"
	"github.com/leep-frog/euler_challenge/ec/p81"
	"github.com/leep-frog/euler_challenge/ec/p816"
	"github.com/leep-frog/euler_challenge/ec/p82"
	"github.com/leep-frog/euler_challenge/ec/p83"
	"github.com/leep-frog/euler_challenge/ec/p84"
	"github.com/leep-frog/euler_challenge/ec/p85"
	"github.com/leep-frog/euler_challenge/ec/p86"
	"github.com/leep-frog/euler_challenge/ec/p87"
	"github.com/leep-frog/euler_challenge/ec/p88"
	"github.com/leep-frog/euler_challenge/ec/p89"
	"github.com/leep-frog/euler_challenge/ec/p9"
	"github.com/leep-frog/euler_challenge/ec/p90"
	"github.com/leep-frog/euler_challenge/ec/p91"
	"github.com/leep-frog/euler_challenge/ec/p92"
	"github.com/leep-frog/euler_challenge/ec/p93"
	"github.com/leep-frog/euler_challenge/ec/p94"
	"github.com/leep-frog/euler_challenge/ec/p95"
	"github.com/leep-frog/euler_challenge/ec/p96"
	"github.com/leep-frog/euler_challenge/ec/p97"
	"github.com/leep-frog/euler_challenge/ec/p98"
	"github.com/leep-frog/euler_challenge/ec/p99"
	"github.com/leep-frog/euler_challenge/ec/p231"
	"github.com/leep-frog/euler_challenge/ec/p686"
	"github.com/leep-frog/euler_challenge/ec/p684"
	"github.com/leep-frog/euler_challenge/ec/p348"
	"github.com/leep-frog/euler_challenge/ec/p504"
	"github.com/leep-frog/euler_challenge/ec/p549"
	"github.com/leep-frog/euler_challenge/ec/p320"
	"github.com/leep-frog/euler_challenge/ec/p800"
	// END_IMPORT_LIST
)

func getProblems() []*ecmodels.Problem {
	return []*ecmodels.Problem{
		// TODO: Separate packages for each of these (for better scoping)
		p1.P1(),
		p2.P2(),
		p3.P3(),
		p4.P4(),
		p5.P5(),
		p6.P6(),
		p7.P7(),
		p8.P8(),
		p9.P9(),
		p10.P10(),
		p11.P11(),
		p12.P12(),
		p13.P13(),
		p14.P14(),
		p15.P15(),
		p16.P16(),
		p17.P17(),
		p18.P18(),
		p19.P19(),
		p20.P20(),
		p21.P21(),
		p22.P22(),
		p23.P23(),
		p24.P24(),
		p25.P25(),
		p26.P26(),
		p27.P27(),
		p28.P28(),
		p29.P29(),
		p30.P30(),
		p31.P31(),
		p32.P32(),
		p33.P33(),
		p34.P34(),
		p35.P35(),
		p36.P36(),
		p37.P37(),
		p38.P38(),
		p39.P39(),
		p40.P40(),
		p41.P41(),
		p42.P42(),
		p43.P43(),
		p44.P44(),
		p45.P45(),
		p46.P46(),
		p47.P47(),
		p48.P48(),
		p49.P49(),
		p50.P50(),
		p51.P51(),
		p52.P52(),
		p53.P53(),
		p54.P54(),
		p55.P55(),
		p56.P56(),
		p57.P57(),
		p58.P58(),
		p59.P59(),
		p60.P60(),
		p61.P61(),
		p62.P62(),
		p63.P63(),
		p64.P64(),
		p65.P65(),
		p66.P66(),
		// 67 is a bigger version of problem 18
		// 68 was solved in python (TODO: in go?)
		p69.P69(),
		p70.P70(),
		p71.P71(),
		p72.P72(),
		p73.P73(),
		p74.P74(),
		p75.P75(),
		p76.P76(),
		p77.P77(),
		p78.P78(),
		p79.P79(),
		p80.P80(),
		p81.P81(),
		p82.P82(),
		p83.P83(),
		p84.P84(),
		p85.P85(),
		p86.P86(),
		p87.P87(),
		p88.P88(),
		p89.P89(),
		p90.P90(),
		p91.P91(),
		p92.P92(),
		p93.P93(),
		p94.P94(),
		p95.P95(),
		p96.P96(),
		p97.P97(),
		p98.P98(),
		p99.P99(),
		p100.P100(),
		p101.P101(),
		p102.P102(),
		p103.P103(),
		p104.P104(),
		p105.P105(),
		p106.P106(),
		p107.P107(),
		p108.P108(),
		p109.P109(),
		// P110 is P108 with a different input
		p111.P111(),
		p112.P112(),
		p113.P113(),
		p114.P114(),
		p115.P115(),
		p116.P116(),
		p117.P117(),
		p118.P118(),
		p119.P119(),
		p120.P120(),
		p121.P121(),
		p122.P122(),
		p123.P123(),
		p124.P124(),
		p125.P125(),
		p126.P126(),
		p127.P127(),
		p128.P128(),
		p129.P129(),
		p130.P130(),
		p131.P131(),
		p132.P132(),
		p133.P133(),
		p134.P134(),
		p135.P135(),
		p136.P136(),
		p137.P137(),
		p138.P138(),
		p139.P139(),
		p140.P140(),
		p141.P141(),
		p142.P142(),
		p143.P143(),
		p144.P144(),
		p145.P145(),
		p146.P146(),
		p147.P147(),
		p148.P148(),
		p149.P149(),
		p150.P150(),
		p151.P151(),
		p152.P152(),
		p153.P153(),
		p155.P155(),
		p184.P184(),
		p222.P222(),
		p234.P234(),
		p235.P235(),
		p236.P236(),
		p252.P252(),
		p333.P333(),
		p456.P456(),
		p154.P154(),
		p156.P156(),
		p157.P157(),
		p158.P158(),
		p159.P159(),
		p160.P160(),
		p161.P161(),
		p162.P162(),
		p165.P165(),
		p164.P164(),
		p163.P163(),
		p169.P169(),
		p243.P243(),
		p233.P233(),
		p166.P166(),
		p167.P167(),
		p168.P168(),
		p170.P170(),
		p171.P171(),
		p173.P173(),
		p172.P172(),
		p174.P174(),
		p175.P175(),
		p176.P176(),
		p177.P177(),
		p178.P178(),
		p179.P179(),
		p458.P458(),
		p301.P301(),
		p181.P181(),
		p401.P401(),
		p501.P501(),
		p601.P601(),
		p701.P701(),
		p719.P719(),
		p700.P700(),
		p816.P816(),
		p808.P808(),
		p206.P206(),
		p205.P205(),
		p187.P187(),
		p204.P204(),
		p191.P191(),
		p751.P751(),
		p357.P357(),
		p345.P345(),
		p493.P493(),
		p387.P387(),
		p188.P188(),
		p347.P347(),
		p381.P381(),
		p346.P346(),
		p500.P500(),
		p203.P203(),
		p231.P231(),
		p686.P686(),
		p684.P684(),
		p348.P348(),
		p504.P504(),
		p549.P549(),
		p320.P320(),
		p800.P800(),
		// END_LIST (needed for file_generator.go)
	}
}

func Branches() map[string]command.Node {
	m := map[string]command.Node{}
	for i, p := range getProblems() {
		pStr := fmt.Sprintf("%d", p.Num)
		if _, ok := m[pStr]; ok {
			log.Fatalf("Duplicate problem entry: %d, %d", i, p.Num)
		}
		m[pStr] = p.Node()
	}
	return m
}
