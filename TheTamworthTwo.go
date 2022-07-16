package main

import (
	"fmt"
)

type status struct {
	direction int
	x         int
	y         int
}

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

var _map [10][10]int

var flag [160005]bool

func main() {
	TheTamwothTwo()
}

func Move(character *status) {
	switch character.direction {
	case NORTH:
		character.y += 1
	case EAST:
		character.x += 1
	case SOUTH:
		character.y -= 1
	case WEST:
		character.x -= 1
	}

	if !(character.x < 0 || character.x > 9 || character.y < 0 || character.y > 9 || _map[character.x][character.y] == 1) {
		return
	}

	if character.x < 0 {
		character.x = 0
	}
	if character.x > 9 {
		character.x = 9
	}
	if character.y < 0 {
		character.y = 0
	}
	if character.y > 9 {
		character.y = 9
	}

	if _map[character.x][character.y] == 1 {
		switch character.direction {
		case NORTH:
			character.y -= 1
		case EAST:
			character.x -= 1
		case SOUTH:
			character.y += 1
		case WEST:
			character.x += 1
		}
	}

	character.direction = (character.direction + 1) % 4

}

func TheTamwothTwo() {
	var cow, John, d_cow, d_John status
	var time int
	var y, x int
	for y = 9; y >= 0; y-- {
		var t string
		_, err := fmt.Scan(&t)
		if err != nil {
			panic("input something wrong!")
			return
		}
		for x = 0; x < 10; x++ {

			if t[x] == '.' {
				_map[x][y] = 0
				continue
			}
			if t[x] == '*' {
				_map[x][y] = 1
				continue
			}
			if t[x] == 'C' {
				cow.direction = NORTH
				cow.x = x
				cow.y = y
				d_cow = cow
			}
			if t[x] == 'F' {
				John.direction = NORTH
				John.x = x
				John.y = y
				d_cow = cow
			}
			//执行到此处的都是空地
			_map[x][y] = 0
		}
	}

	for {
		if cow.x == John.x && cow.y == John.y {
			fmt.Println(time)
			break
		}

		Move(&cow)
		Move(&John)

		f := John.x + John.y*10 + cow.x*100 + cow.y*1000 + cow.direction*10000 + John.direction*40000

		if d_cow == cow && d_John == John || flag[f] == true {
			fmt.Println(0)
			break
		}

		flag[f] = true
		time++
	}

}
