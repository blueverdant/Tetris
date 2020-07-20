package main

import (
	_ "fmt"
	"math"
	"container/list"
)
type GameInit struct {
	side 	int64;
	width 	int64
	height 	int64
	speed 	int64
	num_block int64
	type_color string
	ident int64
	direction  int64
	grade		int64
	over 		bool
	arr_bX		[]int
	arr_bY 		[]int
	arr_store_X  []int
	arr_store_Y []int
	arr_store_color []int

}
var gGame *GameInit
func createRandom(stype string){
	 temp := gGame.width/2-1

	if stype == "rBlock" {
		gGame.num_block = math.round(math.random()*4+1)
		switch(gGame.num_block){
			case 1:
				gGame.arr_bX.push(temp,temp-1,temp,temp+1)
				gGame.arr_bY.push(0,1,1,1)
				break
			case 2:
				gGame.arr_bX.push(temp,temp-1,temp-1,temp+1)
				gGame.arr_bY.push(1,0,1,1)
				break
			case 3:
				gGame.arr_bX.push(temp,temp-1,temp+1,temp+2)
				gGame.arr_bY.push(0,0,0,0)
				break
			case 4:
				gGame.arr_bX.push(temp,temp-1,temp,temp+1)
				gGame.arr_bY.push(0,0,1,1)
				break
			case 5:
				gGame.arr_bX.push(temp,temp+1,temp,temp+1)
				gGame.arr_bY.push(0,0,1,1)
				break
		}
	}
	if stype == "rColor"{
		num_color := math.round(math.random()*8+1)

		switch(num_color){
			case 1:
				gGame.type_color="#3EF72A"
				break
			case 2:
				gGame.type_color="yellow"
				break
			case 3:
				gGame.type_color="#2FE0BF"
				break
			case 4:
				gGame.type_color="red"
				break
			case 5:
				gGame.type_color="gray"
				break
			case 6:
				gGame.type_color="#C932C6"
				break
			case 7:
				gGame.type_color= "#FC751B"
				break
			case 8:
				gGame.type_color= "#6E6EDD"
				break
			case 9:
				gGame.type_color= "#F4E9E1"
				break
		}
	}
}

 func judgeCollision_down() bool{
	for  i := 0; i < len(gGame.arr_bX); i++ {
		if (gGame.arr_bY[i] + 1 == gGame.height){
			return false
		}
		if (len(gGame.arr_store_X) != 0) {
			for j := 0; j < len(gGame.arr_store_X); j++{
				if (gGame.arr_bX[i] == gGame.arr_store_X[j]) {
					if (gGame.arr_bY[i] + 1 == gGame.arr_store_Y[j]) {
						return false
					}
				}

			}
		}
	}
	return true
}


func judgeCollision_other( num int64) bool{
	for i := 0; i < len(gGame.arr_bX); i++{
		if (num == 1) {
			if (gGame.arr_bX[i] == gGame.width - 1)
				return false
		}
		if (num == -1) {
			if gGame.arr_bX[i] == 0 {
				return false
			}

		}
		if (len(gGame.arr_store_X) != 0) {
			for j := 0; j < len(gGame.arr_store_X); j++{
				if (gGame.arr_bY[i] == gGame.arr_store_Y[j]) {
					if (gGame.arr_bX[i] + num == gGame.arr_store_X[j]) {
						return false
					}
				}
			}
		}
	}
	return true;
}

func game()  {
	gGame =	&GameInit{}
}