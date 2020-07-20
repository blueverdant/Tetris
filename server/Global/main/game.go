package main

import (
	_ "fmt"
	"math"
	"container/list"
	"math/rand"
	"time"
)
type GameInit struct {
	side 	int;
	width 	int
	height 	int
	speed 	int
	num_block int
	type_color string
	ident int
	direction  int
	grade		int
	over 		bool
	arr_bX		[]int
	arr_bY 		[]int
	arr_store_X  []int
	arr_store_Y []int
	arr_store_color []int

}
var gGame *GameInit
var loop bool

 func down_speed_up_tick(){
 	flag_all_down := true
	flag_all_down = judgeCollision_down()

	if (flag_all_down) {
		gGame.initBackground()
		for i := 0; i < gGame.arr_bY.length; i++{
			gGame.arr_bY[i] = gGame.arr_bY[i] + 1
		}
	}else{
		for i:=0; i < gGame.arr_bX.length; i++ {
			 gGame.arr_store_X.push(this.arr_bX[i])
			 gGame.arr_store_Y.push(this.arr_bY[i])
			 gGame.arr_store_color.push(this.type_color)
		}

		 gGame.arr_bX.splice(0,this.arr_bX.length)
		 gGame.arr_bY.splice(0,this.arr_bY.length)
		 gGame.initBlock()
	}
	 clearUnderBlock()
	 //gGame.drawBlock(this.type_color)
	 //gGame.drawStaticBlock()
	 gameover()
}

func gameover(){
	for i:=0; i < gGame.arr_store_X.length; i++{
		if (gGame.arr_store_Y[i] == 0) {
			loop = false
			gGame.over = true
		}
	}
}

func createRandom(stype string){
	 temp := gGame.width/2-1

	if stype == "rBlock" {
		rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
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
func  clearUnderBlock(){
//删除低层方块
 	var arr_row []int
 	var line_num int
	if gGame.arr_store_X.length != 0 {
		for j := gGame.height-1; j >= 0; j--{
			for i := 0; i < len(gGame.arr_store_color); i++){
				if gGame.arr_store_Y[i] == j {
					arr_row.push(i)
				}
			}
			if len(arr_row) == gGame.width {
				line_num = j
				break
			}else{
				arr_row.splice(0, arr_row.length)
			}
		}
	}
	if (arr_row.length == this.width) {
		//计算成绩grade
		gGame.grade++

		for i := 0; i < arr_row.length; i++{
			gGame.arr_store_X.splice(arr_row[i]-i, 1)
			gGame.arr_store_Y.splice(arr_row[i]-i, 1)
			gGame.arr_store_color.splice(arr_row[i]-i, 1)
		}

		//让上面的方块往下掉一格
		for i := 0; i < gGame.arr_store_color.length; i++){
			if (gGame.arr_store_Y[i] < line_num) {
				gGame.arr_store_Y[i] = gGame.arr_store_Y[i]+1
			}
		}
	}
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
	loop = true
	gGame =	&GameInit{}
	for;true == loop;{
		down_speed_up_tick()
	}
}