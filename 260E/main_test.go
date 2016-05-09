package main

import (
	"fmt"
	"testing"
)

func TestMain(*testing.T) {
	d := NewDoor()

	if d.State != "CLOSED" {
		fmt.Errorf("State is not CLOSED: %v\n", d.State)
	}
	d.Click()
	if d.State != "OPENING" {
		fmt.Errorf("State is not OPENING: %v\n", d.State)
	}

	d.CompleteCycle()
	if d.State != "OPEN" {
		fmt.Errorf("State is not OPEN: %v\n", d.State)
	}
	d.Click()
	if d.State != "CLOSING" {
		fmt.Errorf("State is not CLOSING: %v\n", d.State)
	}
	d.Click()
	if d.State != "STOPPED_WHILE_CLOSING" {
		fmt.Errorf("State is not STOPPED_WHILE_CLOSING: %v\n", d.State)
	}
	d.Click()
	if d.State != "OPENING" {
		fmt.Errorf("State is not OPENING: %v\n", d.State)
	}
	d.Click()
	if d.State != "STOPPED_WHILE_OPENING" {
		fmt.Errorf("State is not STOPPED_WHILE_OPENING: %v\n", d.State)
	}
	d.Click()
	if d.State != "CLOSING" {
		fmt.Errorf("State is not CLOSING: %v\n", d.State)
	}
	d.CompleteCycle()
	if d.State != "CLOSED" {
		fmt.Errorf("State is not CLOSED: %v\n", d.State)
	}

	return
}
