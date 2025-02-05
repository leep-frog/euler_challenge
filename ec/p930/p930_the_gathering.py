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
7 7 ==> 112348.31113028523
                241124.38773524063
"""
