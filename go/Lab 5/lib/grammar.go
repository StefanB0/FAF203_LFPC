package lib

const Empty = "empty"

type Grammar struct {
	Vt []string
	Vn []string
	Transitions []Transition
}

type Transition struct {
	InitS string
	FinalS []string
}

func ContainsElement(s string, slice []string) int {
	for i, ss := range slice {
		if s == ss {
			return i
		}
	}

	return -1
}

func (g *Grammar) GetTransitions(vn string) []Transition {
	result:= []Transition{}

	for _, t := range g.Transitions {
		if vn == t.InitS {
			result = append(result, t)
		}
	}

	return result
}

func (g *Grammar) LeftFactoring() {
	vn_L:= len(g.Vn)

	for i := 0; i <  vn_L; i++ {
		tt:= g.GetTransitions(g.Vn[i])
		if checkSymbolOfIndex(tt, 0) {
			m:= maxLen(tt)

			for j := 0; j < m; j++ {
				if !checkSymbolOfIndex(tt, j) {
					g.Vn = append(g.Vn, g.Vn[i] + "'")
					g.Transitions = append(g.Transitions, addLeftFactoring(tt, j)...)
					newT:= Transition{InitS: g.Vn[i], FinalS: append(tt[0].FinalS[:j], g.Vn[i] + "'")}
					g.removeLeftFactoring( g.Vn[i], j, newT)
					vn_L++
					break;
				}
			}
		}
	}
}

func maxLen(tt []Transition) int {
	max:= 0
	for _, t := range tt {
		if len(t.FinalS) > max {
			max = len(t.FinalS)
		}
	}
	return max
}

func checkSymbolOfIndex(tt []Transition, id int) bool {
	var s string
	for _, t := range tt {
		if len(t.FinalS) > id {
			s = t.FinalS[id]
			break
		}
	}

	var s1 string
	for _, t := range tt {
		if len(t.FinalS) <= id {
			return false
		} else {
			s1 = t.FinalS[id]
		}

		if s != s1 {
			return false
		}
	}

	return true
}

func addLeftFactoring(tt []Transition, id int) []Transition {
	newTT:= []Transition{}
	newT:= Transition{}
	newVn:= tt[0].InitS + "'"
	for _, t := range tt {
		if len(t.FinalS) > id {
			newT = Transition{InitS: newVn, FinalS: t.FinalS[id:]}
		} else {
			newT = Transition{InitS: newVn, FinalS: []string{ Empty }}
		}

		newTT = append(newTT, newT)
	}

	return newTT
}

func (g *Grammar) removeTransitionIndex(id int) []Transition {
	ret:= make([]Transition, 0)
	ret = append(ret, g.Transitions[:id]...)
	ret = append(ret, g.Transitions[id+1:]...)
	return ret
}

func (g *Grammar) removeLeftFactoring(Vn string, id int, t Transition) {
	l:= len(g.Transitions)
	for i := 0; i < l; i++ {
		if g.Transitions[i].InitS == Vn {
			g.Transitions = g.removeTransitionIndex(i)
			i--
			l--
		}
	}

	g.Transitions = append(g.Transitions, t)
}
