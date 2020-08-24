package main

import(
	"app/ServerSystemV2/UserManagement/Config"
	"app/ServerSystemV2/UserManagement/Model/DB"
	"app/ServerSystemV2/UserManagement/Model/Calculator"
	//"github.com/gorilla/mux"
	//jwt "github.com/dgrijalva/jwt-go"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"
	//guuid "github.com/google/uuid"
)

type Responce struct {

}
type ErrResponce struct {
	Status int
	Message string
}

var conf Config.Configuration

func init() {
	conf = Config.LoadConfig()
	DB.LoadConfData(conf.Postgres.Po_host,conf.Postgres.Po_user,conf.Postgres.Po_password,conf.Postgres.Po_dbname,conf.Postgres.Po_port)
}

func main(){

	http.HandleFunc("/users/get_user_data",           UserData)

	server := http.Server{
		Addr:         conf.Serve_port,
		ReadTimeout:  conf.Read_Timeout  * time.Second,
		WriteTimeout: conf.Write_Timeout * time.Second,
	}

	//router.Use(JwtAuthentication)

	log.Println("Starting server..." , conf.Serve_port)
	log.Panic(server.ListenAndServe())
}

func UserData(w http.ResponseWriter, r * http.Request){

	uuid := r.FormValue("uuid")
	if(uuid != ""){
		status,DataSet :=DB.GetUserData(conf.Postgres.Po_host,conf.Postgres.Po_user,conf.Postgres.Po_password,conf.Postgres.Po_dbname,uuid,conf.Postgres.Po_port)
		if status != 0{

			ResponceMethodError := ErrResponce{200,"Join Set"}
			JsonSet,_ := json.Marshal(ResponceMethodError)
			w.WriteHeader(http.StatusOK)
			w.Write(JsonSet)
		}else{
			ResponceMethodError := ErrResponce{503,"User Not Exist"}
			JsonSet,_ := json.Marshal(ResponceMethodError)
			w.WriteHeader(http.StatusSeeOther)
			w.Write(JsonSet)
		}
	}else{
		ResponceMethodError := ErrResponce{400,"Not Valid"}
		JsonSet,_ := json.Marshal(ResponceMethodError)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JsonSet)
	}
}
