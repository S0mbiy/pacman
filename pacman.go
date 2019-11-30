package main

import (
	"fmt"
	"math/rand"
  "time"
  "log"
  term "github.com/nsf/termbox-go"
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
  update := make(chan int)
  for i := 0; i < len(enemies); i++ {
    enemies[i] = 0
  }
  go print(update)
  for n := 0; n < 15; n++ {
    fmt.Println("Running ghost ", n)
    go ghost(n, ch, update)
  }
  <-ch

}

func reset() {
    term.Sync() // cosmestic purpose
}

func pacman(update chan int){
	curPos := 225
	enemies[curPos] = 1
	update <- 1
	err := term.Init()
    if err != nil {
		log.Panicln("Couldn't init keyboard listener")
		return
    }
	defer term.Close()
	for {
		switch ev := term.PollEvent(); ev.Type {
        case term.EventKey:
            switch ev.Key {
			case term.KeyArrowUp:
				if pacmap[curPos-18] > -1{
					enemies[curPos] = 0
					curPos = curPos-18
					enemies[curPos] = 1
					update <- 1
				}
			case term.KeyArrowDown:
				if pacmap[curPos+18] > -1{
					enemies[curPos] = 0
					curPos = curPos+18
					enemies[curPos] = 1
					update <- 1
				}
			case term.KeyArrowLeft:
				if pacmap[curPos-1] > -1{
					enemies[curPos] = 0
					curPos = curPos-1
					enemies[curPos] = 1
					update <- 1
				}
			case term.KeyArrowRight:
				if pacmap[curPos+1] > -1{
					enemies[curPos] = 0
					curPos = curPos+1
					enemies[curPos] = 1
					update <- 1
				}
			}
		}
	}
}

func print(update chan int){
  for {
	  select {
	    case <-update:
		  fmt.Print("\033[2J")
		  for i := 0; i < 324; i++ {
			if(i%18==0){
			  fmt.Println()
			}
		    if(enemies[i]== 2) {
		      fmt.Print("ᗕᗒ")
		    }else if(enemies[i]== 1) {
		      fmt.Print("◖◗")
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
  }
}

func ghost(id int, cha chan string, update chan int) {
  position := 153
  lastPosition := position
	options := make([]int, 0, 4)

  // Change position random
  source := rand.NewSource(time.Now().UnixNano())
  generator := rand.New(source)
  for {
		time.Sleep(time.Second/2)
		// check move options
		if pacmap[position - 18] > 0 && (position - 18) != lastPosition {
			options = append(options, position - 18)
		}
		if pacmap[position + 18] > 0 && (position + 18) != lastPosition{
			options = append(options, position + 18)
		}
		if pacmap[position - 1] > 0 && (position -1) != lastPosition{
			options = append(options, position - 1)
		}
		if pacmap[position + 1] > 0 && (position + 1) != lastPosition{
			options = append(options, position + 1)
		}

    changePosition := generator.Intn(len(options))
		position = options[changePosition]
		enemies[lastPosition] = 0
		lastPosition = position
		enemies[position] = 2
		options = options[:0]
		update <- 1
  }

}
