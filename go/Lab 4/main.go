package main

import (
	"lab4/lib"
)

func main(){
	g:= lib.Grammar{
		NonTerminal: []string{"S", "A", "B", "C", "D"},
		Terminal: []string{"a", "b"},
		Transitions: []lib.Transition{
			{Initial_state: "S", Final_state: []string{"a", "B"}},
			{Initial_state: "S", Final_state: []string{"b", "A"}},
			{Initial_state: "A", Final_state: []string{"b"}},
			{Initial_state: "A", Final_state: []string{"B"}},
			{Initial_state: "A", Final_state: []string{"a", "D"}},
			{Initial_state: "A", Final_state: []string{"A", "S"}},
			{Initial_state: "A", Final_state: []string{"b", "A", "A", "B"}},
			{Initial_state: "A", Final_state: []string{"e"}},
			{Initial_state: "B", Final_state: []string{"b"}},
			{Initial_state: "B", Final_state: []string{"b", "S"}},
			{Initial_state: "C", Final_state: []string{"A", "B"}},
			{Initial_state: "D", Final_state: []string{"B", "B"}},
		}}

	g = lib.Convert(g)
	g.Print()
}
