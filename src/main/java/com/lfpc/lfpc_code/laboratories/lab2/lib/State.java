package com.lfpc.lfpc_code.laboratories.lab2.lib;

import java.util.HashMap;

public class State {
    public String name;
    private boolean finalState;
    public HashMap<String, State> transitions;

    public State(String name){
        this.name = name;
        this.transitions = new HashMap<>();
        this.finalState = false;
    }

    public boolean isFinalState() {
        return finalState;
    }

    public void setFinalState(boolean finalState) {
        this.finalState = finalState;
    }

    public HashMap<String, State> getTransitions() {
        return transitions;
    }
}
