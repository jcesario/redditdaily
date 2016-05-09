package main

type Door struct {
	State string
}

/*
If the door is OPEN or CLOSED, clicking the button will cause the door to move, until it completes the cycle of opening or closing.
Door: Closed -> Button clicked -> Door: Opening -> Cycle complete -> Door: Open.
If the door is currently opening or closing, clicking the button will make the door stop where it is. When clicked again, the door will go the opposite direction, until complete or the button is clicked again.
*/
func (d *Door) Click() {
	if d.State == "OPENING" || d.State == "CLOSING" {
		d.State = "STOPPED_WHILE_" + d.State
		return
	}

	switch {
	case d.State == "CLOSED":
		d.State = "OPENING"
		return
	case d.State == "OPEN":
		d.State = "CLOSING"
		return
	case d.State == "STOPPED_WHILE_CLOSING":
		d.State = "OPENING"
		return
	case d.State == "STOPPED_WHILE_OPENING":
		d.State = "CLOSING"
		return
	}
}

func (d *Door) CompleteCycle() {
	switch {
	case d.State == "CLOSING":
		d.State = "CLOSED"
		return
	case d.State == "OPENING":
		d.State = "OPENED"
		return
	}

	return
}

// We will assume the initial state is CLOSED.
func NewDoor() (d Door) {
	d = Door{State: "CLOSED"}

	return
}

func main() {

}
