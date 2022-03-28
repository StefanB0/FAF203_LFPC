package lib

import "fmt"

type Grammar struct {
	NonTerminal []string
	Terminal    []string
	Transitions []Transition
}

type Transition struct {
	Initial_state string
	Final_state   []string
}

func (t *Transition) copy() Transition {
	t_new := Transition{}
	t_new.Initial_state = t.Initial_state
	t_new.Final_state = append(t_new.Final_state, t.Final_state...)
	copy(t_new.Final_state, t.Final_state)

	return t_new
}

func (g *Grammar) Print() {
	fmt.Printf("Nonterminal symbols: %s\n", g.NonTerminal)
	fmt.Printf("Terminal symbols: %s\n", g.Terminal)
	fmt.Printf("Transitions(%d):\n", len(g.Transitions))

	for _, t := range g.Transitions {
		fmt.Printf("    %s -> %s\n", t.Initial_state, t.Final_state)
	}


}
