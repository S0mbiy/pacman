package main

import (
	"fmt"
)

var pacmap [324]int
var enemies [324]int

func main() {
	pacmap := [324]int{-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,
	-1,1,1,1,1,1,1,1,-1,-1,1,1,1,1,1,1,1,-1,
	-1,1,-1,-1,1,-1,-1,1,-1,-1,1,-1,-1,1,-1,-1,1,-1,
	-1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,-1,
	-1,1,-1,-1,1,-1,1,-1,-1,-1,-1,1,-1,1,-1,-1,1,-1,
	-1,1,1,1,1,-1,1,1,-1,-1,1,1,-1,1,1,1,1,-1,
	-1,-1,-1,-1,1,-1,-1,1,-1,-1,1,-1,-1,1,-1,-1,-1,-1,
	-1,-1,-1,-1,1,-1,1,1,1,1,1,1,-1,1,-1,-1,-1,-1,
	-1,-1,-1,-1,1,-1,1,-1,-3,-3,-1,1,-1,1,-1,-1,-1,-1,
	-1,-1,-1,-1,1,1,1,-1,-2,-2,-1,1,1,1,-1,-1,-1,-1,
	-1,1,1,1,1,-1,1,1,1,1,1,1,-1,1,1,1,1,-1,
	-1,1,-1,-1,1,-1,-1,-1,1,1,-1,-1,-1,1,-1,-1,1,-1,
	-1,1,1,-1,1,1,1,1,1,1,1,1,1,1,-1,1,1,-1,
	-1,-1,1,-1,1,-1,1,-1,-1,-1,-1,1,-1,1,-1,1,-1,-1,
	-1,1,1,1,1,-1,1,1,-1,-1,1,1,-1,1,1,1,1,-1,
	-1,1,-1,-1,-1,-1,-1,1,-1,-1,1,-1,-1,-1,-1,-1,1,-1,
	-1,1,1,1,1,1,1,1,-1,-1,1,1,1,1,1,1,1,-1,
	-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1}
	print(pacmap)
	for i := 0; i < len(enemies); i++ {
    enemies[i] = 0
  }
}

func print(pacmap [324]int){
	for i := 0; i < 324; i++ {
		if(i%18==0){
			fmt.Println()
		}
		if(pacmap[i]==-1){
			fmt.Print("██")
		}else if(pacmap[i]==1){
			fmt.Print("░░")
		}else if(pacmap[i]==-2){
			fmt.Print("▄▄")
		}else if(pacmap[i]==-3){
			fmt.Print("▀▀")
		}
	}
	fmt.Println()
}

func ghostUp(position int) (int) {
  if position <= 17 {
    return ghostDown(position)
  }
  return position - 18
}
func ghostDown(position int) (int) {
  if position >= 306 {
    return ghostRight(position)
  }
  return position + 18
}

func ghostRight(position int) (int) {
  if position%17 == 0 {
    return ghostLeft(position)
  }
  return position + 1
}

func ghostLeft(position int) (int) {
  if position%17 == 1 || position == 0{
    return ghostUp(position)
  }
  return position -1
}

func ghost(id int) {
  position := 0
  lastPosition := position
  fmt.Println(position)

  // Change position random
  source := rand.NewSource(time.Now().UnixNano())
  generator := rand.New(source)
  for {
    changePosition := generator.Intn(4)
    fmt.Println(changePosition)
    switch changePosition {
    case 0:
      position = ghostUp(position)
    case 1:
      position = ghostDown(position)
    case 2:
      position = ghostLeft(position)
    case 3:
      position = ghostRight(position)
    }
    if pacmap[position] == -1 {
      position = lastPosition
    } else {
      lastPosition = position
      fmt.Println("Actual position of ghost ", id, ": ", position)
    }
  }

}
