package main

import (
	"lab5/lib"
)

func main(){
	g:= lib.Grammar{}
	s:= "dbaacbaaa"
	g.Vn = []string{"S", "A", "B", "D"}
	g.Vt = []string{"a", "b", "c", "d"}
	g.Transitions = []lib.Transition {
		{InitS: "S", FinalS: []string{"d", "A"}},
		{InitS: "A", FinalS: []string{"D"}},
		{InitS: "A", FinalS: []string{"D", "c"}},
		{InitS: "D", FinalS: []string{"b", "B"}},
		{InitS: "B", FinalS: []string{"a"}},
		{InitS: "B", FinalS: []string{"a"}},
	}
	g.LeftFactoring()

	LL1:= lib.LL1Matrix{}
	LL1.Run(g)
	LL1.Print()
	LL1.AnalyzeWord(s)

}
