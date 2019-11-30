package main

import (
	"fmt"
)

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
