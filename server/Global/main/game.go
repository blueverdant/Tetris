package main

import (
	_ "fmt"
	"time"

	//"container/list"
	"math/rand"
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
	arr_store_color []string

}
var gGame *GameInit
var loop bool

 func down_speed_up_tick(){
 	flag_all_down := true
	flag_all_down = judgeCollision_down()

	if (flag_all_down) {
		//gGame.initBackground()
		for i := 0; i < len(gGame.arr_bY); i++{
			gGame.arr_bY[i] = gGame.arr_bY[i] + 1
		}
	}else{
		for i:=0; i < len(gGame.arr_bX); i++ {
			gGame.arr_store_X = append(gGame.arr_store_X,gGame.arr_bX[i])
			 gGame.arr_store_Y = append(gGame.arr_store_Y,gGame.arr_bY[i])
			 gGame.arr_store_color=append(gGame.arr_store_color,gGame.type_color)
		}
		gGame.arr_bX = gGame.arr_bX[0:len(gGame.arr_bX)]
		gGame.arr_bY = gGame.arr_bY[0:len(gGame.arr_bY)]
		 //gGame.initBlock()
	}
	 clearUnderBlock()
	 //gGame.drawBlock(this.type_color)
	 //gGame.drawStaticBlock()
	 gameover()
}

func gameover(){
	for i:=0; i < len(gGame.arr_store_X); i++{
		if (gGame.arr_store_Y[i] == 0) {
			loop = false
			gGame.over = true
		}
	}
}

//方向键为左右的左移动函数
	func  move(dir_temp int){
		//initBackground()

		if (dir_temp == 1) {                    //右
			flag_all_right := true
			flag_all_right = judgeCollision_other(1)

			if flag_all_right {
				for i := 0; i < len(gGame.arr_bY); i++{
					gGame.arr_bX[i] = gGame.arr_bX[i]+1
				}
			}
		}else{
			flag_all_left := true
			flag_all_left = judgeCollision_other(-1)

			if flag_all_left {
				for i:=0; i < len(gGame.arr_bY); i++{
					gGame.arr_bX[i] = gGame.arr_bX[i]-1
				}
			}
		}
		//drawBlock(gGame.type_color)
		//drawStaticBlock()
	}


func createRandom(stype string){
	 temp := gGame.width/2-1

	if stype == "rBlock" {
		gGame.num_block =  rand.Intn(5)+1
		switch(gGame.num_block){
			case 1:
				gGame.arr_bX=append(gGame.arr_bX,temp,temp-1,temp,temp+1)
				gGame.arr_bY=append((gGame.arr_bY,0,1,1,1)
				break
			case 2:
				gGame.arr_bX=append(gGame.arr_bX,temp,temp-1,temp-1,temp+1)
				gGame.arr_bY=append(gGame.arr_bY,1,0,1,1)
				break
			case 3:
				gGame.arr_bX=append(gGame.arr_bX,temp,temp-1,temp+1,temp+2)
				gGame.arr_bY=append(gGame.arr_bY,0,0,0,0)
				break
			case 4:
				gGame.arr_bX=append(gGame.arr_bX,temp,temp-1,temp,temp+1)
				gGame.arr_bY=append(gGame.arr_bY,0,0,1,1)
				break
			case 5:
				gGame.arr_bX=append(gGame.arr_bX,temp,temp+1,temp,temp+1)
				gGame.arr_bY=append(gGame.arr_bY,0,0,1,1)
				break
		}
	}
	if stype == "rColor"{
		num_color := rand.Intn(8)+1

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
	if len(gGame.arr_store_X)!= 0 {
		for j := gGame.height-1; j >= 0; j--{
			for i := 0; i < len(gGame.arr_store_color); i++{
				if gGame.arr_store_Y[i] == j {
					arr_row = append(arr_row,i )
				}
			}
			if len(arr_row) == gGame.width {
				line_num = j
				break
			}else{
				arr_row = arr_row[0:len(arr_row)]
			}
		}
	}
	if (len(arr_row) == gGame.width) {
		//计算成绩grade
		gGame.grade++

		for i := 0; i < len(arr_row); i++{
			gGame.arr_store_X = gGame.arr_store_X[arr_row[i]-i:1]
			gGame.arr_store_Y = gGame.arr_store_Y[arr_row[i]-i: 1]
			gGame.arr_store_color = gGame.arr_store_color[arr_row[i]-i: 1]
		}

		//让上面的方块往下掉一格
		for i := 0; i < len(gGame.arr_store_color); i++{
			if (gGame.arr_store_Y[i] < line_num) {
				gGame.arr_store_Y[i] = gGame.arr_store_Y[i]+1
			}
		}
	}
}

func judgeCollision_other( num int) bool{
	for i := 0; i < len(gGame.arr_bX); i++{
		if (num == 1) {
			if gGame.arr_bX[i] == gGame.width - 1{
				return false
			}

		}
		if (num == -1) {
			if gGame.arr_bX[i] == 0 {
				return false
			}

		}
		if (len(gGame.arr_store_X) != 0) {
			for j := 0; j < len(gGame.arr_store_X); j++{
				if (gGame.arr_bY[i] == gGame.arr_store_Y[j]) {
					if gGame.arr_bX[i] + num == gGame.arr_store_X[j] {
						return false
					}
				}
			}
		}
	}
	return true;
}

func GameRussia()  {
	loop = true
	gGame =	&GameInit{}
	for;true == loop;{
		time.Sleep(40 * time.Millisecond)
		down_speed_up_tick()
	}
}