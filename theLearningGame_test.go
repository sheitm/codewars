package main

import (
	"time"
	"testing"
	"math/rand"
)

func TestMachine_Response_mult100(t *testing.T) {
	commandMult100 := 1001
	commandMult0 := 0

	mch := NewMachine()

	// Act
	expect := 300
	num := 3

	for i := 0; i < 20; i++ {
		r := mch.Command(commandMult100, num)
		if r != expect {
			mch.Response(false)
			continue
		}
		mch.Response(true)
		break
	}

	// Assert
	rr := mch.Command(commandMult100, num)
	if rr != expect {
		t.Errorf("Expected %d. got %d", expect, rr)
	}

	// Act
	expect = 0

	for i := 0; i < 20; i++ {
		r := mch.Command(commandMult0, num)
		if r != expect {
			mch.Response(false)
			continue
		}
		mch.Response(true)
		break
	}

	// Assert
	rr = mch.Command(commandMult0, num)
	if rr != expect {
		t.Errorf("Expected %d. got %d", expect, rr)
	}
}

func TestMachine_comprehensive(t *testing.T) {
	_actions := []func(int) int{
		func(x int) int { return x + 1 },
		func(x int) int { return 0 },
		func(x int) int { return int(x / 2) },
		func(x int) int { return x * 100 },
		func(x int) int { return x % 2 }}

	mch := NewMachine()
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	for i := 0; i < 200; i++ {
		number := rnd.Intn(100)
		num := mch.Command(i%5, number)
		mch.Response(_actions[i%5](number) == num)
	}

	type test struct {
		action int
		num    int
		result int
		msg string
	}

	tests := []test{{0, 100, 101, "Should apply the num + 1 action to the command 0"}, 
    {1, 100, 0,"Should apply the num * 0 action to the command 1"}, 
    {2, 100, 50,"Should apply the num / 2 action to the command 2"}, 
    {3, 1, 100,"Should apply the num * 100 action to the command 3"}, 
	{4, 100, 0,"Should apply the num % 2 action to the command 4"}}
	
	// machine := NewMachine()
	for index, tst := range tests {
		r := mch.Command(tst.action, tst.num)
		if r != tst.result {
			t.Errorf("Test is %d. Expected %d, got %d", index, tst.result, r)
		}
	}
}

func TestMachine_allSameCommand(t *testing.T) {
	_actions := []func(int) int{
		func(x int) int { return x + 1 },
		func(x int) int { return x + 1 },
		func(x int) int { return x + 1 },
		func(x int) int { return x + 1 },
		func(x int) int { return x + 1 }}

	mch := NewMachine()
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	for i := 0; i < 200; i++ {
		number := rnd.Intn(100)
		num := mch.Command(i%5, number)
		mch.Response(_actions[i%5](number) == num)
	}

	type test struct {
		action int
		num    int
		result int
		msg string
	}

	tests := []test{{0, 100, 101, "Should apply the num + 1 action to the command 0"}, 
    {1, 99, 100,"Should apply the num + 1 action to the command 1"}, 
    {2, -3, -2,"Should apply the num + 1 action to the command 2"}, 
    {3, 1, 2,"Should apply the num + 1 action to the command 3"}, 
	{4, 100, 101,"Should apply the num + 1 action to the command 4"}}
	
	// machine := NewMachine()
	for index, tst := range tests {
		r := mch.Command(tst.action, tst.num)
		if r != tst.result {
			t.Errorf("Test is %d. Expected %d, got %d", index, tst.result, r)
		}
	}
}
