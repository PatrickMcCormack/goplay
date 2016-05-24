package raft

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type NodeType int
type Term uint64
type IndexPtr uint64

const (
	FOLLOWER NodeType = 1 << iota
	CANDIDATE
	LEADER
)

// PersistantState ... Updated on stable storage before responding to RPCs
type PersistantState struct {
	State       NodeType
	CurrentTerm Term
	VotedFor    IndexPtr
	Log         IndexPtr
}

type VolatileState struct {
	CommitIndex IndexPtr
	LastApplied IndexPtr
}

type LeaderVolatileState struct {
	NextIndex  IndexPtr
	MatchIndex IndexPtr
}

type Node struct {
	PersistantState
	VolatileState
	LeaderVolatileState
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (node *Node) extractPersistant(pstate *PersistantState) {
	pstate.State = node.State
	pstate.CurrentTerm = node.CurrentTerm
	pstate.VotedFor = node.VotedFor
	pstate.Log = node.Log
}

// WritePersistantState saves the node state to a named file
// To protect against corruption on write the previous
// state file is renamed as a backup and a new state
// file is written. To protect the backup from becoming
// corrupt by overwrite this function checks if the current
// file can be read before moving it into the backup state.
// TODO - this is just a basic implemetnation to get things
// working, need to come back and make the function comply
// with the documentation.
func (node *Node) WritePersistantState(filename string) {

	pstate := PersistantState{}
	node.extractPersistant(&pstate)
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(pstate)
	check(err)
	f, err := os.Create(filename)
	check(err)
	defer f.Close()
	n, err := f.Write(buffer.Bytes())
	check(err)
	fmt.Printf("Wrote %d bytes to %v\n", n, filename)
}

// ReadPersistantState retrieves the node state from a named file
// if the file cannot be read because it does not exist
// a default state is returned.
// If the file cannot be read because it is corrupt then
// the previous state is read (from the )
// TODO - this is just a basic implemetnation to get things
// working, need to come back and make the function comply
// with the documentation.
func (node *Node) ReadPersistantState(filename string) {
	f, err := os.Open(filename)
	check(err)
	buffer := bufio.NewReader(f)
	decoder := gob.NewDecoder(buffer)
	pstate := PersistantState{}
	err = decoder.Decode(&pstate)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("Read this from %v - %v\n", filename, pstate)
}
