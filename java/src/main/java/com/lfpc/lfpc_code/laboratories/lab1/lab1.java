package com.lfpc.lfpc_code.laboratories.lab1;

import org.json.JSONArray;
import org.json.JSONObject;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.HashMap;



public class lab1 {
    public static boolean checkString(String testString, JSONObject grammarObj, char Vn){
        boolean failsafe;
        for (int i = 0; i < testString.length(); i++) {
            failsafe = true;
            for (int j = 0; j < grammarObj.getJSONObject("P").getJSONArray(String.valueOf(Vn)).length(); j++) {
                if (testString.charAt(i) == grammarObj.getJSONObject("P").getJSONArray(String.valueOf(Vn)).get(j).toString().charAt(0)) {
                    failsafe = false;
                    if (i == testString.length() - 1 && grammarObj.getJSONObject("P").getJSONArray(String.valueOf(Vn)).get(j).toString().length() == 2)
                        return false;
                    if (i < testString.length() - 1 && grammarObj.getJSONObject("P").getJSONArray(String.valueOf(Vn)).get(j).toString().length() < 2)
                        return false;
                    if(i < testString.length() - 1)
                        Vn = grammarObj.getJSONObject("P").getJSONArray(String.valueOf(Vn)).get(j).toString().charAt(1);
                    break;
                }

            }
            if(failsafe)
                return false;
        }
        return true;
    }

    public static void derivationTree(
            JSONObject stringsObj,
            JSONObject grammarObj,
            char Vn
    ){
        String word, tempword = "";

        for (int x = 0; x < stringsObj.getJSONArray("String").length(); x++) {
            if (checkString(stringsObj.getJSONArray("String").get(x).toString(), grammarObj, Vn)) {
                System.out.println("\nValid String " + String.valueOf(x+1));
                System.out.print(Vn);
                word = stringsObj.getJSONArray("String").get(x).toString();

                for (int i = 0; i < word.length(); i++) {
                    JSONArray Vnd = grammarObj.getJSONObject("P").getJSONArray(String.valueOf(Vn));
                    for (int j = 0; j < Vnd.length(); j++) {
                        if (Vnd.get(j).toString().charAt(0) == word.charAt(i)) {
                            if (Vnd.get(j).toString().length() == 2)
                                Vn = Vnd.get(j).toString().charAt(1);
                            else Vn = ' ';
                            tempword = tempword.concat(String.valueOf(word.charAt(i)));
                        }
                    }
                    System.out.print(" -> ");
                    System.out.print(tempword);
                    System.out.print(Vn);
                }
                Vn = 'S';
                tempword = "";
            } else
                System.out.println("\nRejected String " + String.valueOf(x+1));
        }
    }

    public static void finiteAutomata(JSONObject grammarObj) throws IOException {
        ArrayList<String> E = new ArrayList<>();
        ArrayList<String> Q = new ArrayList<>();
        HashMap<String, String> dictionary = new HashMap<>();
        String word = "";

        for (int i = 0; i < grammarObj.getJSONArray("Vn").length(); i++) {
            dictionary.put(grammarObj.getJSONArray("Vn").get(i).toString(), "q" + String.valueOf(i));
        }
        dictionary.put("X", "q" + String.valueOf(grammarObj.getJSONArray("Vn").length()));

        System.out.println("FA= (Q, Σ, δ,q0, F)");
        word = "q0";
        for (int i = 1; i <= grammarObj.getJSONArray("Vn").length(); i++) {
            word = word + ", q" + String.valueOf(i);
        }
        System.out.println("Q = {" + word + "}");

        for (int i = 0; i < grammarObj.getJSONArray("Vt").length(); i++) {
            E.add(grammarObj.getJSONArray("Vt").get(i).toString());
        }

        word = grammarObj.getJSONArray("Vt").toString();
        word = word.substring(1, word.length()-1);
        word = "Σ = {" + word + "}";
        System.out.println(word);
        System.out.println("q0 = {q0}");
        System.out.println("F = {q" + String.valueOf(grammarObj.getJSONArray("Vn").length()) + "}");

        word = "";
        for (int i = 0; i < grammarObj.getJSONArray("Vn").length(); i++) {
            word = word + "q" + String.valueOf(i) + " = {" +grammarObj.getJSONArray("Vn").get(i).toString() + "}, ";
        }

        word = word + "q" + String.valueOf(grammarObj.getJSONArray("Vn").length()) + " = {X}";
        System.out.println(word);

        for (int i = 0; i < grammarObj.getJSONArray("Vn").length(); i++) {
            Q.add(grammarObj.getJSONArray("Vn").get(i).toString());
        }
        Q.add("X");

        int length = 0;
        for (int i = 0; i < Q.toArray().length - 1; i++) {
            length = length + grammarObj.getJSONObject("P").getJSONArray(Q.get(i)).length();
        }

        String[][] stateTransitions = new String[length][3];
        int j = 0;

        for (int i = 0; i < Q.toArray().length - 1; i++) {
            for (int k = 0; k < grammarObj.getJSONObject("P").getJSONArray(Q.get(i)).length(); k++) {
                stateTransitions[j][0] = Q.get(i);
                stateTransitions[j][1] = grammarObj.getJSONObject("P").getJSONArray(Q.get(i)).get(k).toString().substring(0,1);
                if (grammarObj.getJSONObject("P").getJSONArray(Q.get(i)).get(k).toString().length() < 2)
                    stateTransitions[j][2] = "X";
                else
                    stateTransitions[j][2] = grammarObj.getJSONObject("P").getJSONArray(Q.get(i)).get(k).toString().substring(1,2);
                j++;
            }
        }

        word = "";
        for (int i = 0; i < stateTransitions.length; i++) {
            word = word + "δ(" + dictionary.get(stateTransitions[i][0]) + ","
                    + stateTransitions[i][1] + ") = {" + dictionary.get(stateTransitions[i][2]) + "}; ";
        }
        word = word.substring(0, word.length()-1) + ".";
        System.out.println(word);

        drawGraph(stateTransitions, dictionary);
    }

    public static void drawGraph(String[][] stateTransitions, HashMap<String, String> dictionary) throws IOException {
        File file = new File("src/main/java/com/lfpc/lfpc_code/laboratories/lab1/output/graph.dot");
        file.createNewFile();
        FileWriter writer = new FileWriter(file);
        String graphContent = "";
        graphContent = graphContent.concat("""
                digraph G {
                  graph [fontsize=20]
                  edge [fontsize=15]
                  node [fontsize=20]
                  q3 [shape = doublecircle]
                  ranksep = .5
                  nodesep = .6
                  edge [style="setlinewidth(2)"]
                  """);
        for (int i = 0; i < stateTransitions.length; i++) {
            graphContent = graphContent + dictionary.get(stateTransitions[i][0]) + " -> "
                    + dictionary.get(stateTransitions[i][2]) + " [label = " + stateTransitions[i][1] + "]\n";
        }
        graphContent = graphContent + "}";
        writer.write(graphContent);
        writer.flush();
        writer.close();

    }

    public static void main(String[] args) throws IOException {
        char Vn = 'S';
        HashMap<String, String[]> Grammar = new HashMap<>();
        String grammar, testStrings, tempword;
        JSONObject grammarObj, stringsObj;
        Path fileNameGrammar, fileNameTestStrings;

        String answer = "The gramar is a type 3 regular, right linear grammar";

        fileNameGrammar = Path.of("src/main/java/com/lfpc/lfpc_code/laboratories/lab1/input/grammar.json");
        grammar = Files.readString(fileNameGrammar);
        fileNameTestStrings = Path.of("src/main/java/com/lfpc/lfpc_code/laboratories/lab1/input/strings.json");
        testStrings = Files.readString(fileNameTestStrings);

        grammarObj = new JSONObject(grammar);
        stringsObj = new JSONObject(testStrings);

        System.out.println(answer);
        System.out.println("Finite automata conversion\n");
        finiteAutomata(grammarObj);
        System.out.println("\nStep by step derivation of the test strings");
        derivationTree(stringsObj, grammarObj, Vn);
    }
}
