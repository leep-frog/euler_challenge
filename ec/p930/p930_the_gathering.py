from typing import Dict, List

import numpy as np

smartCodeAliases: Dict[str, str] = {}


class BallArrangement:

    def __init__(self, balls, n, m, solved=False):
        self.balls = balls
        self.n = n
        self.m = m
        self.solved = solved

    def __str__(self):
        return f"{self.balls}"

    def __repr__(self):
        return f"{self.balls}"

    def clone(self):
        return BallArrangement([x for x in self.balls], self.n, self.m, solved=self.solved)

    def calcInitProb(self):
        # oddsForSingleOrder = 1 / (self.n ^ self.m)

        counts = [x for x in self.balls if x != 0]

        ballPlacementOrderingCount = permCount(counts)

        arrangements = len({c for c in self.smartCodes()})

        return ballPlacementOrderingCount * arrangements # * oddsForSingleOrder

    def smartCodes(self) -> List[str]:
        codeAliases = []
        shifted = [0 for _ in range(self.n)]
        revShifted = [0 for _ in range(self.n)]
        for shift in range(self.n):
            for idx in range(self.n):
                shifted[(idx+shift)%self.n] = self.balls[idx]
                revShifted[(self.n-1-idx+shift)%self.n] = self.balls[idx]

            # code = f"{shifted}"
            code = shifted.__repr__()
            revCode = revShifted.__repr__()
            codeAliases.append(code)
            codeAliases.append(revCode)
        return codeAliases

    def smartCode(self):
        baseCode = self.balls.__repr__()
        if baseCode in smartCodeAliases:
            return smartCodeAliases[baseCode]

        allCodes = self.smartCodes()
        best = max(allCodes)

        for codeAlias in allCodes:
            smartCodeAliases[codeAlias] = best
        return best


def solve(n, m):
    print("==================== SOLVING FOR", n, m)
    uniqueArrs: List[BallArrangement] = []

    balls = [0 for _ in range(n)]
    balls[0] = m

    init = BallArrangement(balls, n, m)
    countsMapTwo: Dict[str, np.float64] = {}

    # pf.Start("calcProbs")
    # // fmt.Println("  CALC PROBS", time.Now())
    findAllArrangements(init, countsMapTwo, uniqueArrs)

    # print(uniqueArrs)
    # print(countsMapTwo)

    codeToIdx: Dict[str, int] = {ba.smartCode(): i for i, ba in enumerate(uniqueArrs)}

    l = len(uniqueArrs)
    m = np.zeros(shape=(l, l), dtype=np.float64)
    b = np.zeros(shape=(l,), dtype=np.float64)

    for i, ba in enumerate(uniqueArrs):
        row = codeToIdx[ba.smartCode()]
        m[row, row] = -1
        if ba.solved:
            # print(m)
            continue

        # print("=========================")
        # print("DOING FOR", ba)
        for j, v in enumerate(ba.balls):
            if v == 0:
                continue

            oddsOfMove = np.float64((v/ba.m)/2.)

            # print("oom", oddsOfMove)

            leftIdx = (j + ba.n - 1) % ba.n
            rightIdx = (j + 1) % ba.n

            # print("PRE", ba.balls, i, v, j)

            # Move to the left
            ba.balls[j] -= 1
            ba.balls[leftIdx] += 1
            # print("PRE LEFT", ba)
            leftCode = ba.smartCode()
            ba.balls[j] += 1
            ba.balls[leftIdx] -= 1

            # Move to the left
            ba.balls[j] -= 1
            ba.balls[rightIdx] += 1
            # print("PRE RGHT", ba)
            rightCode = ba.smartCode()
            ba.balls[j] += 1
            ba.balls[rightIdx] -= 1

            leftCol = codeToIdx[leftCode]
            # print("ADDING", oddsOfMove, "TO", row, leftCol)
            m[row, leftCol] = m[row, leftCol] + oddsOfMove
            rightCol = codeToIdx[rightCode]
            # print("ADDING", oddsOfMove, "TO", row, rightCol)
            m[row, rightCol] = m[row, rightCol] + oddsOfMove

        b.put(i, 1)
        # print(m)
        # print(b)

    soln = np.linalg.solve(m, b)
    # print("SOLVE", soln)

    # print(countsMapTwo)

    sum = np.float64(0)
    for ba in uniqueArrs:
        sc = ba.smartCode()
        # print("SOLN FOR", sc, countsMapTwo[sc], soln[codeToIdx[sc]])
        sum += countsMapTwo[sc]*soln[codeToIdx[sc]]
    # print("SUM", sum, pow(ba.n, ba.m))
    return sum / (pow(ba.n, ba.m))



    # pf.Start("CreateEqs")
    # // fmt.Println("  CREATE EQS", time.Now())
    # var eqs []*equations.Equation
    # for _, ba := range uniqueArrs {
    #     eqs = append(eqs, createEquation(ba))
    # }

    # /*fmt.Println("\n\nFOR FUN SOLVE")
    # oSoln := equations.SolveLin(eqs)
    # // var oSum float64
    # oSum := fraction.NewRational(0, 1)
    # for k, freq := range countsMapTwo {
    #     // oSum += freq.Float64() * oSoln[equations.Variable(k)]
    #     oSum = oSum.Plus(freq.Times(fraction.NewRationalFromFloat(oSoln[equations.Variable(k)])))
    # }

    # return oSum*/
    # // fmt.Println("END FOR FUN SOLVE", oSum, "\n\n")

    # pf.Start("Solve   ")
    # // fmt.Println("  Solve", time.Now())
    # solns := equations.Solve(eqs)

    # // fmt.Println("  Sum", time.Now())
    # pf.Start("Sum     ")
    # sum := fr(0, 1)
    # for k, freq := range countsMapTwo {
    #     sum = sum.Plus(freq.Times(solns[equations.Variable(k)]))
    # }

    # pf.End()

    # fmt.Println("  Done", n, m, sum, time.Now())

    # return sum


def findAllArrangementsOld(ba: BallArrangement, countsMap: Dict[str, np.float64], bas: List[BallArrangement]):
    # if _, ok := countsMap[ba.smartCode()]; ok {
    #     return
    # }
    if ba.smartCode() in countsMap:
        return

    # print("finding for", ba)


    # var solved bool
    solved = False

    # for _, b := range ba.balls {
    #     if b == ba.m {
    #         solved = true
    #     }
    # }
    for b in ba.balls:
        if b == ba.m:
            solved = True

    bac = BallArrangement([b for b in ba.balls], ba.n, ba.m, solved=solved)
    #     balls:  bread.Copy(ba.balls),
    #     n:      ba.n,
    #     m:      ba.m,
    #     solved: solved,
    # }

    # *bas = append(*bas, bac)
    bas.append(bac)

    countsMap[ba.smartCode()] = ba.calcInitProb()

    # for i, v := range ba.balls {
    for i, v in enumerate(ba.balls):
        if v == 0:
            continue

        # leftIdx := (i + ba.n - 1) % ba.n
        # rightIdx := (i + 1) % ba.n
        leftIdx = (i + ba.n - 1) % ba.n
        rightIdx = (i + 1) % ba.n

        # Move to the left
        ba.balls[i] -= 1
        ba.balls[leftIdx] += 1
        findAllArrangements(ba, countsMap, bas)
        ba.balls[i] += 1
        ba.balls[leftIdx] -= 1

        # Move to the right
        ba.balls[i] -= 1
        ba.balls[rightIdx] += 1
        findAllArrangements(ba, countsMap, bas)
        ba.balls[i] += 1
        ba.balls[rightIdx] -= 1


def findAllArrangements(ba: BallArrangement, countsMap: Dict[str, np.float64], bas: List[BallArrangement]):
    # if _, ok := countsMap[ba.smartCode()]; ok {
    #     return
    # }

    rem = [ba]

    while rem:
        ba = rem.pop()
        if ba.smartCode() in countsMap:
            continue

        # print("finding for", ba)


        # var solved bool
        solved = False

        # for _, b := range ba.balls {
        #     if b == ba.m {
        #         solved = true
        #     }
        # }
        for b in ba.balls:
            if b == ba.m:
                solved = True

        bac = BallArrangement([b for b in ba.balls], ba.n, ba.m, solved=solved)
        #     balls:  bread.Copy(ba.balls),
        #     n:      ba.n,
        #     m:      ba.m,
        #     solved: solved,
        # }

        # *bas = append(*bas, bac)
        bas.append(bac)

        countsMap[ba.smartCode()] = ba.calcInitProb()

        # for i, v := range ba.balls {
        for i, v in enumerate(ba.balls):
            if v == 0:
                continue

            # leftIdx := (i + ba.n - 1) % ba.n
            # rightIdx := (i + 1) % ba.n
            leftIdx = (i + ba.n - 1) % ba.n
            rightIdx = (i + 1) % ba.n

            # Move to the left
            ba.balls[i] -= 1
            ba.balls[leftIdx] += 1
            rem.append(ba.clone())
            ba.balls[i] += 1
            ba.balls[leftIdx] -= 1

            # Move to the right
            ba.balls[i] -= 1
            ba.balls[rightIdx] += 1
            rem.append(ba.clone())
            ba.balls[i] += 1
            ba.balls[rightIdx] -= 1

import math
def permCount(counts: List[int])-> int:
    nonZeroes = [x for x in counts if x != 0]
    totalOps = math.factorial(sum(nonZeroes))
    for c in nonZeroes:
        totalOps //= math.factorial(c)
    return totalOps

# solve(2, 3)

def solveG(n, m):
    sum = np.float64(0)
    for i in range(2, n+1):
        for j in range(2, m+1):
            sum += solve(i, j)
    print(sum)

solveG(12, 12)

"""

        4.127581862769911
        4.127549196026351e+06
        4.127549196026351e+06
8 8 ==> 4.127581862693018e+06
        4.127581862538576e+06
                4.1275818626223775e+06
8 7 ==> 570256.9432509352
        570256.943244465
NOOOo 7 7 ==> 112348.31113028523
7 7 ==> 241124.38773652646
        241124.38773524063
"""


"""
(12,2)

Done 2 2 1/2
Done 2 3 9/4
Done 2 4 13/2
Done 2 5 125/8
Done 2 6 137/4
Done 2 7 1715/24
Done 2 8 871/6
Done 2 9 4653/16
Done 2 10 4629/8
Done 2 11 91839/80
Done 2 12 45517/20

Done 3 2 4/3
Done 3 3 22/3
Done 3 4 244/9
Done 3 5 785/9
Done 3 6 266/1
Done 3 7 2387/3
Done 3 8 7100/3
Done 3 9 42215/6
Done 3 10 565351/27
Done 3 11 16847501/270
Done 3 12 8377562/45

Done 4 2 5/2
Done 4 3 63/4
Done 4 4 141/2
Done 4 5 6875/24
Done 4 6 4537/4
Done 4 7 107359/24
Done 4 8 530179/30
Done 4 9 5599737/80
Done 4 10 2222645/8
Done 4 11 123707133/112
Done 4 12 123095199/28

Done 5 2 4/1
Done 5 3 306/11
Done 5 4 151564/1045
Done 5 5 133888609/187891
Done 5 6 134080469434/38517655
Done 5 7 7105329795548341/415875121035
Done 5 8 196568838886165412596/2332643553885315
Done 5 9 5830590212560331447146443/13978755270583397690
Done 5 10 17342773089169849523503257437/8380263784714746915155
Done 5 11 2918189558702319388353590634255382517/283740696449077498114025895050
Done 5 12 14294451019888215947578516573321808846998/279342715654116796893258493676725

Done 6 2 35/6
Done 6 3 2627/60
Done 6 4 51295/198
Done 6 5 166286135/111384
Done 6 6 11232998871/1293292
Done 6 7 39440108197/772616
Done 6 8 8126262587/26910
Done 6 9 14387881259058613633/8022419605200
Done 6 10 907012852414366651/84943266408
Done 6 11 126814593092948354086471/1991456270359920
Done 6 12 2713175583127485750675923/7136051635456380

Done 7 2 8/1
Done 7 3 18324/287
Done 7 4 17214435848/40895491
Done 7 5 71989353488983754525250/25867168754170863847
Done 7 6 777223914433180501800168926827348/41341697870492695532426686571
Done 7 7 2273490635757873912539703465070105148591170701809465918/17656781873668440797191189624675935391640659851357
Done 7 8 160524864318475361551899694491773673905527957044308618580912893768548223947810232/180682458208638147251348788951645747701935721608447366349794912855129177781
Done 7 9 167496041028314165918596438326708639207971235877172602988143166466758059427000862407642527294335319363713719409/27210489033337013312580881151288952418508924141297858210936161047873130951067835312337932224548185155371

Done 8 2 21/2
Done 8 3 2469/28
Done 8 4 17489673/27370
Done 8 5 14717071765/3087336
Done 8 6 13267135559022325/361662373828
         13267135559022325/361662373828
Done 8 7 3401341427521218147/11853668823832
Done 8 8 3568221991297287266156509165081/1577145284633864467291770

Done 9 2 40/3
Done 9 3 638068/5457
Done 9 4 7424253462296/8055492111
Done 9 5 711273215875770308706514219777440205726570/92894522744297162621365277325820838703
Done 9 6 241066666266884271897420818340438752569750924390213018699835280038348/3644883227487585910238784922278452651189764854224452310016733087
Done 9 7 114188212887081919843346589635485464618494103203852528974205925234922496067895521524154137646860701473670311971904746/196278973958448793894354391408681974032882309928011697127559026731257928813883502398274212458381734511290368211
Done 9 8 3351976158664919217072492203085271118368594649583070819306924983147799705289103721134976646059497001336397710113980146921968914262873757322619441181208379511113892505896/649603038148556583622965819141304284407014520569408330619887819584469515001056067887620575214389427026123364120410372281930131575767815098311080579677637104905693

Done 10 2 33/2
Done 10 3 3641949/24244
Done 10 4 44924233825145027/35174399581310
Done 10 5 40258284392226910265788067428537/3442883530634441840343914088
Done 10 6 42913157360518001379992844735145909196702965887886919/383009825851415412691967817729415623266583341980
Done 10 7 4387704959644541966120234061930074051209676766168038267968771586953209723073/4007972589101943606940108769434392077583825609033851185668506831665480

Done 11 2 20/1
Done 11 3 101754414/540737
Done 11 4 6740973219346984757759895292/3932870988925217692187957
Done 11 5 171282080749940038689267220285166330802763109867526078444646904785638490845464964243154575/9990218087379097968292663408110252745442285366702255002461661279010727844841152243553
Done 11 6 147379365194135222742477231235273478241893888034552823007567147104331586931804456517311539456857015999752277571241933576813546547981414848406031231992174042632100664744219576258971142675410/816580499824705040838069657523220474471186812349147366243518820330053108764034454419291450131844162607652260273229055909363427843004749414017173396658026127451286537900761928097473691
Done 11 7 11733322636615788921931168800010233465131774971236200947508936936160791598279305073346083583039465570134149895099895609079293251001827568767336776578680299653472890669167665069185227807350977712793675939382698702057035323136864534213515304686257930841910063316610300851049297519465267038973971913666706322722064938793293021732579316507377629146006931562861287736254676219049647/6049784142473345879269779096712628884167251989609417569390043818117255567466518857397066841729280926692511419717194304963857061217920765194914593489666436391634891770900017394181453715196926283216568948991051229479405462232616638355928578834883767288305614928558083602027000683846077741383048857317054660927077357984058218113847053111971098402210972338187401513175736929


"""
