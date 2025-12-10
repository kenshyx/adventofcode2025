package day10

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aclements/go-z3/z3"
	"github.com/kenshyx/adventofcode2025/utils"
)

const (
	SwitchOn = '#'
)

func parseIndices(s string) []int {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("bad int %q: %v", p, err)
		}
		out = append(out, n)
	}
	return out
}

func GetSolution(authenticatedR *utils.UrlWithAuth) utils.Solution {
	reader, resp := authenticatedR.FetchInput()
	if resp != nil {
		defer resp.Body.Close()
	}
	p1, p2 := 0, 0
	cfg := z3.NewContextConfig()
	ctx := z3.NewContext(cfg)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}

		goal := words[0]
		if len(goal) < 2 {
			continue
		}
		goal = goal[1 : len(goal)-1]

		goalN := 0
		for i, c := range goal {
			if c == SwitchOn {
				goalN += 1 << i
			}
		}

		buttonWords := words[1 : len(words)-1]
		B := make([]int, 0, len(buttonWords))
		NS := make([][]int, 0, len(buttonWords))

		for _, button := range buttonWords {
			if len(button) < 2 {
				B = append(B, 0)
				continue
			}
			inner := button[1 : len(button)-1]
			ns := parseIndices(inner)
			if strings.TrimSpace(inner) == "" {
				B = append(B, 0)
				continue
			}

			parts := strings.Split(inner, ",")
			mask := 0
			for _, part := range parts {
				part = strings.TrimSpace(part)
				if part == "" {
					continue
				}
				x, er := strconv.Atoi(part)
				if er != nil {
					continue
				}
				mask |= 1 << x
			}
			B = append(B, mask)
			NS = append(NS, ns)
		}

		n := len(B)
		score := n

		limit := 1 << n
		for subset := 0; subset < limit; subset++ {
			an := 0
			aScore := 0

			for i := 0; i < n; i++ {
				if (subset>>i)&1 == 1 {
					an ^= B[i]
					aScore++
				}
			}

			if an == goalN && aScore < score {
				score = aScore
			}
		}

		p1 += score

		joltageTok := words[len(words)-1]
		if len(joltageTok) < 2 {
			log.Fatalf("bad joltage token: %q", joltageTok)
		}
		jInner := joltageTok[1 : len(joltageTok)-1]
		joltageNs := parseIndices(jInner)

		s := z3.NewSolver(ctx)
		nButtons := len(buttonWords)
		V := make([]z3.Int, nButtons)
		intSort := ctx.IntSort()
		for i := 0; i < nButtons; i++ {
			name := fmt.Sprintf("B%d", i)
			V[i] = ctx.IntConst(name)
			_ = intSort
		}

		zero := ctx.FromInt(0, ctx.IntSort()).(z3.Int)

		// For each joltage index i, equation sum(V[j] for i in NS[j]) == joltageNs[i]
		for i := 0; i < len(joltageNs); i++ {
			var sumExpr z3.Int
			first := true
			for j := 0; j < nButtons; j++ {
				for _, idx := range NS[j] {
					if idx == i {
						if first {
							sumExpr = V[j]
							first = false
						} else {
							sumExpr = sumExpr.Add(V[j])
						}
						break
					}
				}
			}

			rhs := ctx.FromInt(int64(joltageNs[i]), ctx.IntSort()).(z3.Int)

			if first {
				if joltageNs[i] != 0 {
					log.Fatalf("unsatisfiable constraints at index %d", i)
				}
				continue
			}
			s.Assert(sumExpr.Eq(rhs))
		}

		for _, v := range V {
			s.Assert(v.GE(zero))
		}

		var sumAll z3.Int
		if nButtons > 0 {
			sumAll = V[0]
			for i := 1; i < nButtons; i++ {
				sumAll = sumAll.Add(V[i])
			}
		} else {
			sumAll = zero
		}

		bestSum := int64(-1)

		for {
			ok, _ := s.Check()
			if ok != true {
				break
			}
			m := s.Model()

			var curSum int64
			for _, v := range V {
				valV := m.Eval(v, true).(z3.Int)
				vInt, isLit, ok2 := valV.AsInt64()
				if !ok2 || !isLit {
					log.Fatalf("can't be solved")
				}
				curSum += vInt
			}

			if bestSum == -1 || curSum < bestSum {
				bestSum = curSum
			}

			bound := ctx.FromInt(curSum-1, ctx.IntSort()).(z3.Int)
			s.Assert(sumAll.LE(bound))
		}

		if bestSum < 0 {
			log.Fatalf("no solution found by Z3 for line: %s", line)
		}
		p2 += int(bestSum)
	}

	return utils.Solution{
		Part1: p1,
		Part2: p2,
	}
}
