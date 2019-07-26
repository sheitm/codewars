package main

import "fmt"

// Machine a learning machine
type Machine struct {
	confirmedCommands map[int]machineAction
	failedActions     map[string]bool
	pendingContext    executionContext
}

type executionContext struct {
	wasLearned    bool
	learnedCmd    int
	pendingAction int
	pendingCmd    int
}

type machineAction func(int) int

var allActions = map[int]machineAction{
	0: actionMult0,
	1: actionDiv2,
	2: actionMult100,
	3: actionMod2,
	4: actionPlus1,
}

// NewMachine function called to get your machine initialised
func NewMachine() Machine {
	return Machine{
		confirmedCommands: map[int]machineAction{},
		failedActions:     map[string]bool{},
	}
}

// Command executes the command
func (m *Machine) Command(cmd int, num int) int {
	if act, ok := m.confirmedCommands[cmd]; ok {
		m.pendingContext = executionContext{
			wasLearned: true,
			learnedCmd: cmd,
		}
		return act(num)
	}

	for i := 0; i < len(allActions); i++ {
		a := allActions[i]
		k := fmt.Sprintf("%d_%d", cmd, i)
		if _, ok := m.failedActions[k]; !ok {
			m.pendingContext = executionContext{
				wasLearned:    false,
				pendingCmd:    cmd,
				pendingAction: i,
			}
			return a(num)
		}
	}
	return -100
}

// Response receives the command
func (m *Machine) Response(res bool) {
	if res && m.pendingContext.wasLearned {
		return
	}

	if res && !m.pendingContext.wasLearned {
		m.confirmedCommands[m.pendingContext.pendingCmd] = allActions[m.pendingContext.pendingAction]
		return
	}

	if m.pendingContext.wasLearned {
		delete(m.confirmedCommands, m.pendingContext.pendingCmd)
	}

	k := fmt.Sprintf("%d_%d", m.pendingContext.pendingCmd, m.pendingContext.pendingAction)
	m.failedActions[k] = true
}

func actionMult0(num int) int {
	return 0
}

func actionDiv2(num int) int {
	return int(num / 2)
}

func actionMult100(num int) int {
	return num * 100
}

func actionMod2(num int) int {
	return num % 2
}

func actionPlus1(num int) int {
	return num + 1
}
