package com.lfpc.lfpc_code.laboratories.lab2;

import com.lfpc.lfpc_code.laboratories.lab2.lib.FiniteAutomaton;
import com.lfpc.lfpc_code.laboratories.lab2.lib.State;
import guru.nidi.graphviz.attribute.*;
import guru.nidi.graphviz.engine.Format;
import guru.nidi.graphviz.engine.Graphviz;
import guru.nidi.graphviz.model.Graph;
import org.json.JSONObject;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.*;

import static guru.nidi.graphviz.attribute.Rank.RankDir.LEFT_TO_RIGHT;
import static guru.nidi.graphviz.model.Factory.graph;
import static guru.nidi.graphviz.model.Factory.node;
import static guru.nidi.graphviz.model.Link.to;

public class lab2 {
    public static FiniteAutomaton readFA(String jsonAddress) throws IOException {
        JSONObject finiteAutomataObj;
        Path fileNameFA;
        FiniteAutomaton fa = new FiniteAutomaton();

        fileNameFA = Path.of(jsonAddress);
        finiteAutomataObj = new JSONObject(Files.readString(fileNameFA));

        fa.starting_state = finiteAutomataObj.getString("q0");

        for (Object symbol : finiteAutomataObj.getJSONArray("Q")){
            fa.states.add(symbol.toString());
        }

        for (Object symbol : finiteAutomataObj.getJSONArray("E")){
            fa.alphabet.add(symbol.toString());
        }

        for (Object symbol : finiteAutomataObj.getJSONArray("F")){
            fa.finalStates.add(symbol.toString());
        }

        for (String states : fa.states){
            fa.transitions.put(states, new HashMap<>());
            for (String a : fa.alphabet){
                fa.transitions.get(states).put(a, new ArrayList<>());
            }
        }

        for (int i = 0; i < finiteAutomataObj.getJSONArray("D").length(); i++) {
            fa.transitions.get(finiteAutomataObj.getJSONArray("D").getJSONArray(i).getString(0))
                    .get(finiteAutomataObj.getJSONArray("D").getJSONArray(i).getString(1))
                    .add(finiteAutomataObj.getJSONArray("D").getJSONArray(i).getString(2));
        }

        return fa;
    }

    public static FiniteAutomaton convertTransitions(FiniteAutomaton nfa) throws IOException {
        FiniteAutomaton dfa = nfa;

        String newState = "";
        HashSet<String> tempset = new HashSet<>();
        ArrayList<String> temp = new ArrayList<>();
        ArrayList<String[]> keypair = new ArrayList<>();
        int iterations;
        for (String t : dfa.transitions.keySet()) {
            for (String alph : dfa.transitions.get(t).keySet()) {
                if (dfa.transitions.get(t).get(alph).size() > 1) {
                    keypair.add(new String[]{t, alph});

                }
            }
        }

        for(String[] pair :keypair){
            String t = pair[0];
            String alph = pair[1];

            temp = dfa.transitions.get(t).get(alph);
            temp = (ArrayList<String>) temp.clone();
            iterations = temp.size();
            for (int i = 0; i < iterations; i++) {
                String s = dfa.transitions.get(t).get(alph).get(i);
                if (s.contains("-"))
                    temp.addAll(Arrays.asList(s.split("-")));
                else
                    temp.add(s);
            }
            Collections.sort(temp);
            tempset.addAll(temp);
            temp.clear();
            temp.addAll(tempset);
            tempset.clear();

            for (String s : temp) {
                newState = newState + "-" + s;
            }

            newState = newState.substring(1, newState.length());
            dfa.transitions.get(t).get(alph).clear();
            dfa.transitions.get(t).get(alph).add(newState);

            if (!dfa.states.contains(newState)) {
                dfa.states.add(newState);
                Collections.sort(dfa.states);
                for (String s : temp) {
                    dfa.transitions.put(
                            newState,
                            dfa.transitions.get(s)
                    );
                }

                for (String s : temp) {
                    if(dfa.finalStates.contains(s))
                        dfa.finalStates.add(newState);
                }

            }
            temp.clear();
            newState = "";
            dfa = convertTransitions(dfa);
        }


        return dfa;
    }

    public static FiniteAutomaton conversionDFA(FiniteAutomaton nfa) throws IOException {
        FiniteAutomaton dfa = new FiniteAutomaton();
        dfa.alphabet = nfa.alphabet;
        dfa.states = nfa.states;
        dfa.finalStates = nfa.finalStates;
        dfa.transitions = nfa.transitions;

        dfa = convertTransitions(nfa);
        return dfa;
    }

    public static void drawGraph(FiniteAutomaton fa, String name) throws IOException {
        name = "src/main/java/com/lfpc/lfpc_code/laboratories/lab2/output/" + name + ".png";

        Graph g = graph(name).directed()
                .graphAttr().with(Rank.dir(LEFT_TO_RIGHT))
                .nodeAttr().with(Font.name("arial"))
                .linkAttr().with("class", "link-class")
                .with(node(fa.starting_state).with(Color.GREEN));

        for(String s : fa.transitions.keySet()){
            for(String s2 : fa.transitions.get(s).keySet()){
                for(String s3 :fa.transitions.get(s).get(s2))
                g = g.with(
                        node(s).link(to(node(s3)).with(Label.of(s2)))
                );
            }
        }

        for (String s : fa.finalStates){
            g = g.with(
                    node(s).with(Shape.DOUBLE_CIRCLE)
            );
        }

        Graphviz.fromGraph(g).height(100).render(Format.PNG).toFile(new File(name));
    }

    public static void main(String[] args) throws IOException, CloneNotSupportedException {
        String jsonAddress = "src/main/java/com/lfpc/lfpc_code/laboratories/lab2/input/FA.json";
        FiniteAutomaton nfa = new FiniteAutomaton(), dfa = new FiniteAutomaton();
        HashMap<String, State> states;

        nfa = readFA(jsonAddress);
        dfa = conversionDFA((FiniteAutomaton) nfa.clone());

        drawGraph(dfa, "DFA-graph");
        drawGraph(nfa, "NFA-graph");

        System.out.println("NFA");
        System.out.println(nfa.states);
        System.out.println(nfa.alphabet);
        System.out.println(nfa.finalStates);
        System.out.println(nfa.transitions);

        System.out.println("\nDFA");
        System.out.println(dfa.states);
        System.out.println(dfa.alphabet);
        System.out.println(dfa.finalStates);
        System.out.println(dfa.transitions);

    }
}
