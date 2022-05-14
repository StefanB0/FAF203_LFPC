package lib

import (
	"fmt"
	"strings"
)

type LL1Matrix struct {
	matrix []LL1Vn
	precedenceTable [][][]string

}

type LL1Vn struct {
	name  string
	first map[string]bool
	follow map[string]bool
}

func (ll1 *LL1Matrix) Run(g Grammar) {

	for _, vn := range g.Vn {
		ll1.matrix = append(ll1.matrix, LL1Vn{
				vn,
				first(vn, g),
				map[string]bool{},
			})
	}

	ll1.precedenceTable = makeTable(g.Vn, g.Vt)

	followArray:= follow(g, *ll1)
	for i := 0; i < len(g.Vn); i++ {
		ll1.matrix[i].follow = followArray[i]
	}

	ll1.fillTable(g)

}

func (ll1m *LL1Matrix) getFirst(vn string) map[string]bool {
	for _, ll1vn := range ll1m.matrix {
		if ll1vn.name == vn {
			return ll1vn.first
		}
	}

	return map[string]bool{}
}

func (ll1m *LL1Matrix) Print() {
	for _, ll1 := range ll1m.matrix {
		first:= ""
		follow:= ""
		for s := range ll1.first {
			first += s + ", "
		}

		for s := range ll1.follow {
			follow += s + ", "
		}

		if len(first) > 0{
			first = first[:len(first) - 2]
		}
		if len(follow) > 0 {
			follow = follow[:len(follow) - 2]
		}

		fmt.Printf("%s: First(%s), Follow(%s)\n", ll1.name, first, follow)
	}
}

func joinMap(m1 map[string]bool, m2 map[string]bool) map[string]bool {
	mm := make(map[string]bool, 0)

	for k, v := range m1 {
		mm[k] = v
	}

	for k, v := range m2 {
		mm[k] = v
	}

	return mm
}

func first(vn string, g Grammar) map[string]bool {
	result := make(map[string]bool, 0)
	count := 1
	for _, t := range g.GetTransitions(vn) {
		for i := 0; i < count; i++ {
			if ContainsElement(t.FinalS[i], g.Vt) != -1 || t.FinalS[i] == Empty {
				result[t.FinalS[i]] = true
			} else if ContainsElement(t.FinalS[i], g.Vn) != -1 {
				child := first(t.FinalS[i], g)

				if child[Empty] && count < len(t.FinalS) {
					count++
					delete(child, Empty)
				}

				result = joinMap(result, child)
			}
		}
	}

	return result
}

func getLastVn(vn string, g Grammar, matrix LL1Matrix) []string {
	result:= []string{}
	for _, t := range g.GetTransitions(vn) {
		result = append(result, followEOF(t, g, matrix)...)
	}

	return result
}

func followEOF(t Transition, g Grammar, matrix LL1Matrix) []string {
	result:= []string{}

	l:= len(t.FinalS)
	stop:= l - 1
	for i := l - 1; i >= stop; i-- {
		current:= t.FinalS[i]
		result = append(result, current)
		if t.InitS != current {
			result = append(result, getLastVn(current, g, matrix)...)
		}

		first:= matrix.getFirst(current)
		if first[Empty] && stop > 0{
			stop--
		}
	}


	return result
}

func follow(g Grammar, matrix LL1Matrix) []map[string]bool {
	followArray:= make([]map[string]bool, len(g.Vn))
	for i := 0; i < len(followArray); i++ {
		followArray[i] = map[string]bool{}
	}

	followA:= make(map[string]map[string]bool)
	for _, vn := range g.Vn {
		followA[vn] = make(map[string]bool)
	}

	for _, t := range g.Transitions {
		for i := 0; i <= len(t.FinalS) - 2; i++ {
			stop:= i+2
			for j := i+1; j < stop; j++ {
				if ContainsElement(t.FinalS[i], g.Vn) != -1 {
					followA[t.FinalS[i]][t.FinalS[j]] = true
				}

				index:= 0
				for ii, v := range matrix.matrix {
					if v.name == t.FinalS[j] {
						index = ii
						break
					}
				}

				if matrix.matrix[index].first[Empty] && stop < len(t.FinalS) {
					stop++
				}
			}
		}
	}



	for i, vn := range g.Vn {
		for _, vnn := range g.Vn {
			if followA[vn][vnn] {
				index:= 0
				for ii, v := range matrix.matrix {
					if v.name == vnn {
						index = ii
						break
					}
				}

				followArray[i] = joinMap(followArray[i], matrix.matrix[index].first)
				if followArray[i][Empty] {
					delete(followArray[i], Empty)
				}
			}
		}

		for _, vt := range g.Vt {
			if followA[vn][vt] {
				followArray[i][vt] = true
			}
		}
		//!
		lastOfVn:= getLastVn(vn, g, matrix)
		for _, lvn := range lastOfVn {
			index:= ContainsElement(lvn, g.Vn)
			if index != -1 {
				followArray[index] = joinMap(followArray[index], followArray[i])
			}
		}
	}

	followArray[0]["$"] = true
	followEOFA:= getLastVn("S", g, matrix)
	for _, vn := range followEOFA {
		index:= ContainsElement(vn, g.Vn)
		if index != -1 {
			followArray[index]["$"] = true
		}
	}

	return followArray
}

func makeTable(vn, vt []string) [][][]string {
	table:= make([][][]string, len(vn))

	for i := range vn {
		table[i] = make([][]string, len(vt))
		for j := range vt {
			table[i][j] = []string{}
		}
	}
	return table
}

func (ll1m *LL1Matrix) AnalyzeWord(s string, g Grammar) {
	// pointer:= 0
	derivation:= "S"

	// for pointer < len(s) {

	// 	if ContainsElement(derivation[pointer:pointer+1], g.Vt) != -1 {
	// 		pointer++
	// 	} else {
	// 		vnn := ContainsElement(derivation[pointer:pointer+1], g.Vn)
	// 		vtt := ContainsElement(s[pointer:pointer+1], g.Vt)
	// 		ll1m.precedenceTable[vnn][vtt]
	// 	}
	// }

	fmt.Println(derivation)
}

func (ll1m *LL1Matrix) fillTable(g Grammar) {
	var i, j int
	for _, ll1vn := range ll1m.matrix {
		for k, v := range ll1vn.first {
			if k == Empty && v {
				for kk, vv := range ll1vn.follow {
					if (vv) {
						i, j = getCoordinates(ll1vn.name, kk, g)
						ll1m.precedenceTable[i][j] = append(ll1m.precedenceTable[i][j], getFinalStates(g, ll1vn.name)...)
					}
				}
			} else if v {
				i, j = getCoordinates(ll1vn.name, k, g)
				ll1m.precedenceTable[i][j] = append(ll1m.precedenceTable[i][j], getFinalStates(g, ll1vn.name)...)
			}

		}
	}
}

func getCoordinates(vn, vt string, g Grammar) (int, int) {
	var i, j int
	for i1, gvn := range g.Vn {
		if gvn == vn {
			i = i1
		}
	}

	for j1, gvt := range g.Vt {
		if gvt == vt {
			j = j1
		}
	}

	return i, j
}

func getFinalStates(g Grammar, vn string) []string {
	result := []string{}
	tt:= g.GetTransitions(vn)

	for _, t := range tt {
		s:= ""
		for _, s1 := range t.FinalS {
			s += s1
		}
		result = append(result, s)
	}

	return result
}
