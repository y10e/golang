package main

import "fmt"

type Status int

const (
	stop Status = iota
	running
	pause
)

func (s Status) String() string {
    switch s {
		case stop:
			return "stop"
		case running:
			return "running"
		case pause:
			return "pause"
		default:
			return "Unknown"
    }
}

type process struct {
	id int
	name string
	status Status

}

func main(){
	p := process{id:0 ,name:"p1", status:stop}

	// initial status
	fmt.Println(p.id)
	fmt.Println(p.name)
	fmt.Println(p.status)
	
	// stop > runnnig
	p.status = running
	fmt.Println(p.status)

	// running > pause
	p.status = pause
	fmt.Println(p.status)
}