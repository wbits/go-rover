package main

import (
	"reflect"
	"strconv"
	"testing"
)

type roverTestCase struct {
	rover rover
}

func TestRoverInit(t *testing.T) {
	tc := roverTestCase{newRover()}
	tc.assertEqual(0, 0, "N", "", t)
}

func TestRoverMoveForward(t *testing.T) {
	tc := roverTestCase{newRover()}
	tc.assertEqual(0, 1, "N", "f", t)
}

func TestRoverTurnRight(t *testing.T) {
	tc := roverTestCase{newRover()}
	tc.assertEqual(0, 0, "E", "r", t)
}

func TestRoverTurnLeft(t *testing.T) {
	tc := roverTestCase{newRover()}
	tc.assertEqual(0, 0, "W", "l", t)
}

func TestRoverTurnLeftAndMoveBackward(t *testing.T) {
	tc := roverTestCase{newRover()}
	tc.assertEqual(3, 0, "W", "lbbb", t)
}

func TestRoverTurnsAndMoves(t *testing.T) {
	tc := roverTestCase{newRover()}
	tc.assertEqual(-3, -3, "W", "bbblfff", t)
}

func (tc roverTestCase) assertEqual(x, y int, h string, command string, t *testing.T) {
	expected := tc.asMap(x, y, h)
	result := tc.rover.exec(command)
	if reflect.DeepEqual(expected, result) {
		t.Log("Ok")
	} else {
		t.Errorf("%v, does not match expectation %v", result, expected)
	}
}

func (tc roverTestCase) asMap(x, y int, h string) map[string]string {
	return map[string]string{
		"x": strconv.Itoa(x),
		"y": strconv.Itoa(y),
		"h": h,
	}
}
