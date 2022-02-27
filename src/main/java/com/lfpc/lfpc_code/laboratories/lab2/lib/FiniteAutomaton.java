package com.lfpc.lfpc_code.laboratories.lab2.lib;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;

public class FiniteAutomaton implements Cloneable{
    public String starting_state;
    public ArrayList<String> states, alphabet, finalStates;
    public HashMap<String, HashMap<String, ArrayList<String>>> transitions;

    public FiniteAutomaton(){
        this.states = new ArrayList<>();
        this.alphabet = new ArrayList<>();
        this.finalStates = new ArrayList<>();
        this.transitions = new HashMap<>();
    }

    public Object clone() throws CloneNotSupportedException {
        FiniteAutomaton f2 = (FiniteAutomaton)super.clone();

        f2.starting_state = this.starting_state;

        f2.states = new ArrayList<>();
        f2.states.addAll(this.states);

        f2.alphabet = new ArrayList<>();
        f2.alphabet.addAll(this.alphabet);

        f2.finalStates = new ArrayList<>();
        f2.finalStates.addAll(this.finalStates);

        f2.transitions = new HashMap<>();
        for(String s : this.transitions.keySet()){
            f2.transitions.put(s, new HashMap<>());
            for(String s2 : this.transitions.get(s).keySet())
                f2.transitions.get(s).put(s2, (ArrayList<String>) this.transitions.get(s).get(s2).clone());
        }

        return f2;
    }

}
