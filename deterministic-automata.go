/*
	Go script that implementes a Deterministic Finite Automata (DFA) following
	this Python implementation: https://drive.google.com/drive/folders/1lxu9hWiPRBjhgjUWyN8RLwfKnKS8AQkL

	Example: Set of Alphabet (denoted with a Sigma symbol) Sigma = { a, b }
	Language of the machine,  L = { Words where a's come before b's}

*/

package main

import (
	"fmt"
)

// Defining a struct for the DFA
type DFA struct {
	Q     map[int]bool            // Set of states
	Sigma map[string]bool         // Set of Symbols
	Delta map[stateSymbolPair]int // Transition function
	q0    int                     //Initial state
	F     map[int]bool            // Set of final states
}

// Custom type to use as a key for the transition function map Delta
type stateSymbolPair struct {
	state  int
	symbol string
}

func main() {

	// This automaton is defined s.t. all the "a's" are before the "b's"
	D_0 := DFA{
		Q:     map[int]bool{0: true, 1: true, 2: true},
		Sigma: map[string]bool{"a": true, "b": true},

		// This mapping of the transition funciton is obtained by following its transition graph
		Delta: map[stateSymbolPair]int{
			{0, "a"}: 0,
			{0, "b"}: 1,
			{1, "a"}: 2, // Example: {1, "a"}: 2 --> means in state 1 if I encounter an "a" next, then transition to state 2.
			{1, "b"}: 1,
			{2, "a"}: 2,
			{2, "b"}: 2,
		},
		q0: 0,
		F:  map[int]bool{0: false, 1: true},
	}

	// Example of accepted languages: aa, aabbb
	// Example of un-accepted languages: ba, aba

	fmt.Println(D_0.run("aabbb")) // Returns true
	fmt.Println(D_0.run("aba"))   // Returns false
	fmt.Println(D_0.run(""))      // Empty string which returns true
}

// Method to run the DFA instance 'D_0' with different input strings
func (dfa *DFA) run(w string) bool {
	q := dfa.q0                // Initialize q to the initial state q0
	for _, symbol := range w { // Iterate through the string w
		q = dfa.Delta[stateSymbolPair{q, string(symbol)}] // Update q to its new state
	}
	_, isFinal := dfa.F[q] // After processing the entire string, check if the current state q is the final state by looking up the final state F
	return isFinal
}
