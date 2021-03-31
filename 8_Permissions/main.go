package main

import (
	"fmt"
)

const (
	Read     uint8 = 1 << iota // 0
	Write                      // 1
	Delete                     // 2
	Exectute                   // 3
	Super                      // 4
)

func name(i uint8) string {
	switch i {
	case Read:
		return "Read"
	case Write:
		return "Write"
	case Delete:
		return "Delete"
	case Exectute:
		return "Exectute"
	case Super:
		return "Super"
	default:
		return "UNKNOWN"
	}
}

func allowed(i, c uint8) bool {
	return i&c != 0
}

func main() {
	me := uint8(23)
	for _, flag := range []uint8{Read, Write, Delete, Exectute, Super} {
		var icon string
		if allowed(me, flag) {
			icon = "✅"
		} else {
			icon = "❌"
		}
		fmt.Printf("% 08b %s\n% 08b %s %s\n-----\n", me, "me", flag, name(flag), icon)
	}
}
