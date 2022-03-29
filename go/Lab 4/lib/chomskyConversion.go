package lib

import (
	"strconv"
)

func Convert(g_initial Grammar) Grammar {
	g_intermediate := eliminateEmtpy(g_initial)
	g_intermediate = eliminateRename(g_intermediate)
	g_intermediate = eliminateNonproductive(g_intermediate)
	g_intermediate = eliminateUnreachable(g_intermediate)

	g_final := normalForm(g_intermediate)
	return g_final
}

func DebugConvert(g_initial Grammar) Grammar {
	g_initial.Print()
	g_intermediate := eliminateEmtpy(g_initial)
	g_intermediate.Print()
	g_intermediate = eliminateRename(g_intermediate)
	g_intermediate.Print()
	g_intermediate = eliminateNonproductive(g_intermediate)
	g_intermediate.Print()
	g_intermediate = eliminateUnreachable(g_intermediate)
	g_intermediate.Print()

	g_final := normalForm(g_intermediate)
	g_final.Print()
	return g_final
}

func eliminateEmtpy(g Grammar) Grammar {

	l := len(g.Transitions)
	additionalTransitions := []Transition{}

	for i := 0; i < l; i++ {
		t := g.Transitions[i]
		if t.Final_state[0] == "e" && len(t.Final_state) == 1 {
			target := t.Initial_state
			for j := 0; j < l; j++ {
				tt := g.Transitions[j]
				if checkValue(tt.Final_state, target) {
					additionalTransitions = append(additionalTransitions, addEmptyTransitions(tt.copy(), target)...)
				}
			}

			for _, new_t := range additionalTransitions {
				if !checkDuplicate(new_t, g.Transitions) {
					g.Transitions = append(g.Transitions, new_t)
					l++
				}
			}

			additionalTransitions = []Transition{}

			g.Transitions = removeTransition(i, g.Transitions)
			i--
			l--
		}
	}

	return g
}

func eliminateRename(g Grammar) Grammar {

	l := len(g.Transitions)
	renamelist := [][]string{}

	for i := 0; i < l; i++ {
		t := g.Transitions[i]
		if len(t.Final_state) == 1 && checkValue(g.NonTerminal, t.Final_state[0]) {
			renamelist = append(renamelist, []string{t.Initial_state, t.Final_state[0]})
			g.Transitions = removeTransition(i, g.Transitions)

			l--
			i--
		}
	}

	for _, pair := range renamelist {
		g.Transitions = append(
			g.Transitions,
			addRenameTransitions(pair[1], pair[0], g.Transitions)...,
		)

	}

	return g
}

func eliminateNonproductive(g Grammar) Grammar {
	productive := []string{}

	for _, t := range g.Transitions {
		l := len(t.Final_state)
		symbol := t.Initial_state
		if l == 1 && !checkValue(productive, symbol) {
			productive = append(productive, symbol)
		}
	}

	productive = addProductive(productive, 0, g)
	nonproductive := []string{}

	l := len(g.NonTerminal)
	for i := 0; i < l; i++ {
		ch := g.NonTerminal[i]
		if !checkValue(productive, ch) {
			nonproductive = append(nonproductive, ch)
			g.NonTerminal = removeValue(i, g.NonTerminal)
			i--
			l--
		}
	}

	for _, n := range nonproductive {
		l = len(g.Transitions)
		for i := 0; i < l; i++ {
			t := g.Transitions[i]
			if t.Initial_state == n || checkValue(t.Final_state, n) {
				g.Transitions = removeTransition(i, g.Transitions)
				i--
				l--
			}
		}
	}

	return g
}

func eliminateUnreachable(g Grammar) Grammar {
	reachable := []string{"S"}
	for _, t := range g.Transitions {
		if t.Initial_state == "S" {
			for _, ch := range t.Final_state {
				if !checkValue(reachable, ch) {
					reachable = append(reachable, ch)
				}
			}
		}
	}

	reachable = addReachable(reachable, 0, g)

	l := len(g.NonTerminal)
	for i := 0; i < l; i++ {
		ch := g.NonTerminal[i]
		if !checkValue(reachable, ch) {
			g.NonTerminal = removeValue(i, g.NonTerminal)
			i--
			l--
		}
	}

	l = len(g.Terminal)
	for i := 0; i < l; i++ {
		ch := g.Terminal[i]
		if !checkValue(reachable, ch) {
			g.NonTerminal = removeValue(i, g.Terminal)
			i--
			l--
		}
	}

	l = len(g.Transitions)
	for i := 0; i < l; i++ {
		ch := g.Transitions[i].Initial_state
		if !checkValue(g.NonTerminal, ch) {
			g.Transitions = removeTransition(i, g.Transitions)
			i--
			l--
		}
	}

	return g
}

func normalForm(g Grammar) Grammar {

	substitutionTable := map[string]string{}
	g = normalizeTerminal(g, substitutionTable)

	j:= 1
	l:= len(g.Transitions)
	for i := 0; i < l; i++ {
		t:= g.Transitions[i]
		if len(t.Final_state) > 2 {
			g.Transitions[i] = normalize(t, substitutionTable, &j)
		}
	}

	for ch, name := range substitutionTable {
		g.Transitions = append(g.Transitions, Transition{Initial_state: name, Final_state: []string{ch}})
		g.NonTerminal = append(g.NonTerminal, name)
	}
	return g
}

func normalizeTerminal(g Grammar, st map[string]string) Grammar {
	for i, ch := range g.Terminal {
		name := "X" + strconv.Itoa(i+1)
		st[ch] = name
	}

	l:= len(g.Transitions) - len(st)
	for i := 0; i < l; i++ {
		t:= g.Transitions[i]
		for j, ch := range t.Final_state {
			if checkValue(g.Terminal, ch) && len(t.Final_state) > 1{
				g.Transitions[i].Final_state[j] = st[ch]
			}
		}
	}
	return g
}

func normalize(t Transition, st map[string]string, j *int) Transition {

	l:= len(t.Final_state)-1
	for i := 0; i < l; i++ {
		pair:= t.Final_state[i] + t.Final_state[i+1]
		if st[pair] == "" {
			name := "Y" + strconv.Itoa(*j)
			st[pair] = name
			*j++
		}
		t.Final_state[i] = st[pair]
		t.Final_state = removeValue(i+1, t.Final_state)
		l--
	}

	return t
}

func addProductive(productive []string, lastIteration int, g Grammar) []string {
	currentIteration := lastIteration
	nonproductive := false

	for _, t := range g.Transitions {
		if !checkValue(productive, t.Initial_state) {
			for _, ch := range t.Final_state {
				if !checkValue(g.Terminal, ch) && !checkValue(productive, ch) {
					nonproductive = true
				}
			}
			if !nonproductive {
				productive = append(productive, t.Initial_state)
				currentIteration++
			}

		}
		nonproductive = false
	}

	if currentIteration > lastIteration {
		return addProductive(productive, currentIteration, g)
	} else {
		return productive
	}
}

func addReachable(reachable []string, lastIteration int, g Grammar) []string {
	l := len(reachable)

	for i := 1; i < l; i++ {
		ch := reachable[i]
		if checkValue(g.NonTerminal, ch) {
			for _, t := range g.Transitions {
				if t.Initial_state == ch {
					for _, ch2 := range t.Final_state {
						if !checkValue(reachable, ch2) {
							reachable = append(reachable, ch2)
							l++
						}
					}
				}
			}
		}
	}

	return reachable
}

func checkValue(array []string, target string) bool {
	for _, ch := range array {
		if ch == target {
			return true
		}
	}
	return false
}

func checkDuplicate(t Transition, ttable []Transition) bool {
	for _, tt := range ttable {
		if t.Initial_state == tt.Initial_state && compareStringSlices(t.Final_state, tt.Final_state) {
			return true
		}
	}
	return false
}

func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func addRenameTransitions(initial, final string, tt []Transition) []Transition {
	l := len(tt)
	newTransitions := []Transition{}

	for i := 0; i < l; i++ {
		if tt[i].Initial_state == initial {
			t := tt[i].copy()
			t.Initial_state = final
			if !checkDuplicate(t, tt) {
				newTransitions = append(newTransitions, t)
			}
		}

	}

	return newTransitions
}

func addEmptyTransitions(t Transition, target string) []Transition {
	newTransitions := []Transition{}

	for i, ch := range t.Final_state {
		if ch == target {
			t_new := t.copy()
			t_new.Final_state = removeValue(i, t_new.Final_state)
			newTransitions = append(newTransitions, t_new)
		}
	}

	return newTransitions
}

func removeValue(id int, array []string) []string {
	l := len(array)
	for i := id; i < l-1; i++ {
		array[i] = array[i+1]
	}
	array[l-1] = ""
	array = array[:l-1]
	return array
}

func removeTransition(id int, t []Transition) []Transition {
	l := len(t)
	for i := id; i < l-1; i++ {
		t[i] = t[i+1]
	}

	t = t[:l-1]
	return t
}
