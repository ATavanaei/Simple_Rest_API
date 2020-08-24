package DB

import (
	_ "github.com/lib/pq"
	"database/sql"
	"strconv"
	"fmt"
)

var XpMap map[string]int
var PlatMap map[string]int

type UserClass struct {
	Id       int
	Xp       int
	Platform string
	Class    string
}

type Xperiance struct {
	Id int
	Level_set string
}

type Platform struct {
	Id int
	level_set string
}

func GetUserData(host,user,pass,db,uuid string, port int) (status int,respData[]string){

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, db)
	connection,err := sql.Open("postgres", psqlInfo)
	if err!=nil{
		//panic(err)
		return 0,respData
	}
	defer connection.Close()

	sqlStatement := "SELECT id,xp,platform,class FROM users WHERE uuid=$1"
	row1 := connection.QueryRow(sqlStatement,uuid)

	var userClass UserClass
	err1 := row1.Scan(&userClass.Id,&userClass.Xp,&userClass.Platform,&userClass.Class)
	IdString := strconv.Itoa(userClass.Id)
	XpString := strconv.Itoa(userClass.Xp)
	respData = []string{IdString,XpString,userClass.Platform,userClass.Class}
	if err1!=nil{
		//panic(err1)
		return 0,respData
	}else{
		return 1,respData
	}
}

func LoadConfData(host,user,pass,db string, port int) {
	XpMap = make(map[string]int)
	PlatMap = make(map[string]int)

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, db)
	connection,err := sql.Open("postgres", psqlInfo)
	if err!=nil{
		panic(err)
	}
	defer connection.Close()

	sqlStatement_platform := "SELECT id,level_set FROM platform"
	row_platform,err_platform := connection.Query(sqlStatement_platform)
	if err_platform !=nil{
		panic(err)
	}else{
		for row_platform.Next(){
			var platform Platform
			err = row_platform.Scan(&platform.Id,&platform.level_set)
			if err !=nil{
				panic(err)
			}
			PlatMap[platform.level_set] = platform.Id
		}
	}

	sqlStatement_xp := "SELECT id,level_set FROM xperiance"
	row_xp,err_xp := connection.Query(sqlStatement_xp)
	if err_xp !=nil{
		panic(err)
	}else{
		for row_xp.Next(){
			var xperiance Xperiance
			err = row_xp.Scan(&xperiance.Id,&xperiance.Level_set)
			if err !=nil{
				panic(err)
			}
			XpMap[xperiance.Level_set] = xperiance.Id
		}
	}
}