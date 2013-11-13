package aloha

import (
	"fmt"
	"math/rand"
	"time"
)

// Host structure 
type Host struct {
	sendIn     int
	frames     int
	framesSent int
	wantToSend bool
}

func Simulate(hostnum int, probability float64, slots int, delay int) {
	rand.Seed(time.Now().UTC().UnixNano())
	var hosts = []Host{}
	var collision bool
	var firstToSend int
	for h := 0; h < hostnum; h++ {
		hosts = append(hosts, Host{
			-1,
			0,
			0,
			false})
	}
	for s := 0; s < slots; s++ {
		for num := range hosts {
			hosts[num].wantToSend = false
		}
		for num := range hosts {
			if hosts[num].sendIn >= 0 {
				hosts[num].sendIn = hosts[num].sendIn - 1
				if hosts[num].sendIn == -1 {
					hosts[num].wantToSend = true
				}
			}
			if hosts[num].wantToSend == false {
				if rand.Float64() < probability {
					hosts[num].frames = hosts[num].frames + 1
					hosts[num].wantToSend = true
				}
			}

		}
		collision = false
		firstToSend = -1
		for num := range hosts {
			if hosts[num].wantToSend == true {
				if firstToSend == -1 {
					firstToSend = num
				} else {
					collision = true
				}
			}
		}
		if (!collision) && (firstToSend != -1) {
			hosts[firstToSend].framesSent += 1
		} else if collision {
			for num := range hosts {
				if hosts[num].wantToSend {
					hosts[num].sendIn = 1 + rand.Intn(delay-1)
				}
			}
		}
	}
	totalsent := 0.0
	totalframes := 0.0
	for num := range hosts {
		fmt.Println("host ", num, ": ", hosts[num].frames, " attempted and ", hosts[num].framesSent, " sent")
		totalframes += float64(hosts[num].frames)
		totalsent += float64(hosts[num].framesSent)
	}
	fmt.Println("bottom line: ", 100*totalsent/totalframes, "% frames got through and ", 100*totalsent/float64(slots), "% slots utilized")
}
