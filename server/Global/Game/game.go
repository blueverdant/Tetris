package Game
import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"math/rand"
	"time"

	"github.com/fv0008/AWS_Russia/server"
)
type GameInit struct {
	Side            int      `json:"side"`
	Width           int      `json:"width"`
	Height          int      `json:"height"`
	Speed           int      `json:"speed"`
	Num_block       int      `json:"num_block"`
	Type_color      string   `json:"type_color"`
	//Ident           int      `json:"ident"`
	Direction       int      `json:"direction"`
	Grade           int      `json:"grade"`
	Over            bool     `json:"Over"`
	Arr_bX          []int    `json:"arr_bX"`
	Arr_bY          []int    `json:"arr_bY"`
	Arr_store_X     []int    `json:"arr_store_X"`
	Arr_store_Y     []int    `json:"arr_store_Y"`
	Arr_store_color []string `json:"arr_store_color"`

}
var gGame *GameInit
var loop bool

func initBackground(){

}
func initgame(event server.IM_protocol)server.IM_protocol {
	gGame.Side=35
	gGame.Width=10
	gGame.Height =25
	gGame.Speed =400
	gGame.Type_color = "000000"
	//map black by client
	initBackground()
	initBlock()



	jsons,error := json.Marshal(gGame)
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Println(string(jsons))
	event.Msg = string(jsons)
	return event
}

func Down_speed_up_tick(){
 	flag_all_down := true
	flag_all_down = JudgeCollision_down()

	if (flag_all_down) {
		//gGame.initBackground()
		for i := 0; i < len(gGame.Arr_bY); i++{
			gGame.Arr_bY[i] = gGame.Arr_bY[i] + 1
		}
	}else{
		for i:=0; i < len(gGame.Arr_bX); i++ {
			gGame.Arr_store_X = append(gGame.Arr_store_X,gGame.Arr_bX[i])
			 gGame.Arr_store_Y = append(gGame.Arr_store_Y,gGame.Arr_bY[i])
			 gGame.Arr_store_color =append(gGame.Arr_store_color,gGame.Type_color)
		}
		gGame.Arr_bX = gGame.Arr_bX[0:len(gGame.Arr_bX)]
		gGame.Arr_bY = gGame.Arr_bY[0:len(gGame.Arr_bY)]
		initBlock()
	}
	 ClearUnderBlock()
	 //gGame.drawBlock(this.Type_color)
	 //gGame.drawStaticBlock()
	 gameover()
}
func initBlock(){
	createRandom("rColor")        //生成颜色字符串，
	createRandom("rBlock")
}

func gameover(){
	for i:=0; i < len(gGame.Arr_store_X); i++{
		if (gGame.Arr_store_Y[i] == 0) {
			loop = false
			gGame.Over = true
		}
	}
}

//方向键为左右的左移动函数
func move(dir_temp int){
		//initBackground()

		if (dir_temp == 1) {                    //右
			flag_all_right := true
			flag_all_right = JudgeCollision_other(1)

			if flag_all_right {
				for i := 0; i < len(gGame.Arr_bY); i++{
					gGame.Arr_bX[i] = gGame.Arr_bX[i]+1
				}
			}
		}else{
			flag_all_left := true
			flag_all_left = JudgeCollision_other(-1)

			if flag_all_left {
				for i:=0; i < len(gGame.Arr_bY); i++{
					gGame.Arr_bX[i] = gGame.Arr_bX[i]-1
				}
			}
		}
		//drawBlock(gGame.Type_color)
		//drawStaticBlock()
	}


func createRandom(stype string){
	 temp := gGame.Width/2-1

	if stype == "rBlock" {
		gGame.Num_block =  rand.Intn(5)+1
		switch(gGame.Num_block){
			case 1:
				gGame.Arr_bX =append(gGame.Arr_bX,temp,temp-1,temp,temp+1)
				gGame.Arr_bY =append(gGame.Arr_bY,0,1,1,1)
				break
			case 2:
				gGame.Arr_bX =append(gGame.Arr_bX,temp,temp-1,temp-1,temp+1)
				gGame.Arr_bY =append(gGame.Arr_bY,1,0,1,1)
				break
			case 3:
				gGame.Arr_bX =append(gGame.Arr_bX,temp,temp-1,temp+1,temp+2)
				gGame.Arr_bY =append(gGame.Arr_bY,0,0,0,0)
				break
			case 4:
				gGame.Arr_bX =append(gGame.Arr_bX,temp,temp-1,temp,temp+1)
				gGame.Arr_bY =append(gGame.Arr_bY,0,0,1,1)
				break
			case 5:
				gGame.Arr_bX =append(gGame.Arr_bX,temp,temp+1,temp,temp+1)
				gGame.Arr_bY =append(gGame.Arr_bY,0,0,1,1)
				break
		}
	}
	if stype == "rColor"{
		num_color := rand.Intn(8)+1

		switch(num_color){
			case 1:
				gGame.Type_color ="#3EF72A"
				break
			case 2:
				gGame.Type_color ="yellow"
				break
			case 3:
				gGame.Type_color ="#2FE0BF"
				break
			case 4:
				gGame.Type_color ="red"
				break
			case 5:
				gGame.Type_color ="gray"
				break
			case 6:
				gGame.Type_color ="#C932C6"
				break
			case 7:
				gGame.Type_color = "#FC751B"
				break
			case 8:
				gGame.Type_color = "#6E6EDD"
				break
			case 9:
				gGame.Type_color = "#F4E9E1"
				break
		}
	}
}

func JudgeCollision_down() bool{
	for  i := 0; i < len(gGame.Arr_bX); i++ {
		if (gGame.Arr_bY[i] + 1 == gGame.Height){
			return false
		}
		if (len(gGame.Arr_store_X) != 0) {
			for j := 0; j < len(gGame.Arr_store_X); j++{
				if (gGame.Arr_bX[i] == gGame.Arr_store_X[j]) {
					if (gGame.Arr_bY[i] + 1 == gGame.Arr_store_Y[j]) {
						return false
					}
				}

			}
		}
	}
	return true
}
func  ClearUnderBlock(){
//删除低层方块
 	var arr_row []int
 	var line_num int
	if len(gGame.Arr_store_X)!= 0 {
		for j := gGame.Height -1; j >= 0; j--{
			for i := 0; i < len(gGame.Arr_store_color); i++{
				if gGame.Arr_store_Y[i] == j {
					arr_row = append(arr_row,i )
				}
			}
			if len(arr_row) == gGame.Width {
				line_num = j
				break
			}else{
				arr_row = arr_row[0:len(arr_row)]
			}
		}
	}
	if (len(arr_row) == gGame.Width) {
		//计算成绩grade
		gGame.Grade++

		for i := 0; i < len(arr_row); i++{
			gGame.Arr_store_X = gGame.Arr_store_X[arr_row[i]-i:1]
			gGame.Arr_store_Y = gGame.Arr_store_Y[arr_row[i]-i: 1]
			gGame.Arr_store_color = gGame.Arr_store_color[arr_row[i]-i: 1]
		}

		//让上面的方块往下掉一格
		for i := 0; i < len(gGame.Arr_store_color); i++{
			if (gGame.Arr_store_Y[i] < line_num) {
				gGame.Arr_store_Y[i] = gGame.Arr_store_Y[i]+1
			}
		}
	}
}

func JudgeCollision_other( num int) bool{
	for i := 0; i < len(gGame.Arr_bX); i++{
		if (num == 1) {
			if gGame.Arr_bX[i] == gGame.Width - 1{
				return false
			}

		}
		if (num == -1) {
			if gGame.Arr_bX[i] == 0 {
				return false
			}

		}
		if (len(gGame.Arr_store_X) != 0) {
			for j := 0; j < len(gGame.Arr_store_X); j++{
				if (gGame.Arr_bY[i] == gGame.Arr_store_Y[j]) {
					if gGame.Arr_bX[i] + num == gGame.Arr_store_X[j] {
						return false
					}
				}
			}
		}
	}
	return true;
}

func Start(event server.IM_protocol)server.IM_protocol{
	if false==loop && "start" == event.Msg{
		return initgame(event)
	}
	if true==loop {

	}
	//这里返回要客户端重新开始游戏
	return event
}

func GameRussia()  {
	gGame =	&GameInit{}
	for{
		time.Sleep(40 * time.Millisecond)
		if true == loop {
			Down_speed_up_tick()
		}else{
			return
		}
	}
}