package main

import (
	"fmt"
)

const MAX_DOOR = 100

func main() {
	Doors := [4]int32{0, 0, 0, 0}
	for step := 1; step <= MAX_DOOR; step++ {
		for door := 0; door <= MAX_DOOR; door += step {
			if Doors[door>>5]&(1<<(uint(door&31))) != 0 {
				Doors[door>>5] &= ^(1 << uint(door&31))
			} else {
				Doors[door>>5] |= (1 << uint(door&31))
			}

		}
	}
	for door := 1; door <= MAX_DOOR; door++ {
		if Doors[door>>5]&(1<<(uint(door&31))) != 0 {
			fmt.Printf("\nDoors #%d: is open\n", door)
		} else {
			// fmt.Printf("\nDoors #%d: is closed\n", door)
		}
	}
	fmt.Printf("\nAll other doors are closed\n")
}
