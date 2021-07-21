package Utils

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"tetris/server/Global"
)

const UserInfoDir = "/UserInfo/"
var kindtype map[string]string
func InitUtils(){
	kindtype = make(map[string]string)
	kindtype["sixrandom"] = ""
	kindtype["eightrandom"] = ""
	kindtype["qimen"] = ""
	kindtype["taiyi"] = ""
	kindtype["sixcourse"] = ""
}


func SHA256Encode(s string) string {
	sha256 := crypto.SHA256.New()
	sha256.Write([]byte(s))
	return hex.EncodeToString(sha256.Sum(nil))
}

func Url(url string, params ...interface{}) string {
	queryString := ""
	for index, item := range params {
		if index%2 == 0 {
			queryString += item.(string) + "="
		} else {
			queryString += ToString(item) + "&"
		}
	}
	if url != "/" {
		url = strings.TrimRight(url, "/")
	}
	queryString = strings.TrimRight(queryString, "&")
	return url + "?" + queryString
}

func ToString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	}
	return ""
}

func TimeDiffForHumans(t time.Time) string {
	unix := t.Unix()
	now := time.Now().Unix()
	b := now - unix
	if b < 0 {
		return t.Format("2006-01-01 15:04:05")
	}
	if b < 60 {
		return fmt.Sprintf("%d秒前", b)
	}
	// 单位：分钟
	if b < 3600 {
		b = b / 60
		return fmt.Sprintf("%d分钟前", b)
	}
	// 单位：小时
	b = b / 3600
	if b < 24 {
		return fmt.Sprintf("%d个小时前", b)
	}
	// 单位：天
	b = b / 24
	if b < 30 {
		return fmt.Sprintf("%d天前", b)
	}
	// 单位：月
	b = b / 30
	if b < 12 {
		return fmt.Sprintf("%d个月前", b)
	}
	// 单位：年
	b = b / 12
	if b > 3 {
		return t.Format("2006-01-01 15:04:05")
	}
	return fmt.Sprintf("%d年钱", b)
}

func Pwd() string {
	return filepath.Dir(os.Args[0])
}

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
//mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func CreateCaptcha() string {
	for index:=0 ;index<100 ;index++{
		numrand:= fmt.Sprintf("%0*v",6, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		if 0!=strings.Index( numrand,"0") {
			return numrand;
		}
	}
	return "987654"
}

func Debug() (string, bool) {
	conn := os.Getenv("CGO_CFLAGS")
	Global.Logger.Info("Debug: ",conn)
	if conn == "-O0 -g" {
		return conn, true
	}
	return conn, false
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {

	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Getdatapath() string  {
	var datapath string
	datapath = UserInfoDir;
	_,ret := Debug()
	if true== ret {
		datapath = "datas" + UserInfoDir;
	}
	Global.Logger.Info("url: ", getCurrentDirectory())
	return datapath;
}

func Checkfilesynckind(kind string) bool {
	if _, ok := kindtype[kind]; ok {
		return true
	}
	return false
}

func CheckJsonSync(Jstr []byte)([]byte , bool) {

	var ret []byte;
	var r interface{}
	var err error;
	err = json.Unmarshal(Jstr, &r)
	b_ret := false
	if err == nil {
		gobook, ok := r.(map[string]interface{})
		if ok {
			for k, _ := range gobook {
				if k=="sync" {
					gobook[k] = true
					b_ret = true
				}
			}
			ret,err=json.Marshal(gobook)
			return ret,b_ret;
		}
	}
	return ret,b_ret
}