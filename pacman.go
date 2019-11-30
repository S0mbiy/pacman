package main

import (
	"fmt"
	"math/rand"
  "time"
)

var pacmap = [324]int{-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,
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
var enemies [324]int

func main() {
  ch := make(chan string)

	for i := 0; i < len(enemies); i++ {
    enemies[i] = 0
  }
	for n := 0; n < 5; n++ {
    fmt.Println("Running ghost ", n)
    go ghost(n, ch)
  }
  <-ch

}

func print(pacmap, enemies [324]int){
  fmt.Print("\033[2J")
	for i := 0; i < 324; i++ {
		if(i%18==0){
			fmt.Println()
		}
    if(enemies[i]== 2) {
      fmt.Print("ᗕᗒ")
    }else{
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
	}
	fmt.Println()

}

func ghostUp(position int) (int) {
  if position < 18 {
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
  if position%18 == 17 {
    return ghostLeft(position)
  }
  return position + 1
}

func ghostLeft(position int) (int) {
  if position%18 == 0 {
    return ghostUp(position)
  }
  return position -1
}

func ghost(id int, cha chan string) {
  position := 15
  lastPosition := position
  // fmt.Println(position)

  // Change position random
  source := rand.NewSource(time.Now().UnixNano())
  generator := rand.New(source)
  for {
    // fmt.Println("antes")
    time.Sleep(time.Second/2)
    changePosition := generator.Intn(4)
    // fmt.Println(changePosition)
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
      enemies[lastPosition] = 0
      lastPosition = position
      enemies[position] = 2
      print(pacmap, enemies)
      // fmt.Println("Actual position of ghost ", id, ": ", position)
    }
  }

}
