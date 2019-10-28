package entity

import (
	"log"
	"os"
)

var (
	Error * log.Logger
	Login * log.Logger
)


var logpath string = "entity/data/login.log"
var errpath string = "entity/data/err.log"
	

func init(){
	loginlog, err :=os.OpenFile(logpath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err !=nil{
		log.Fatalln(err)
	}

	errlog, err := os.OpenFile(errpath,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("file open error : %v", err)
	}

	Error = log.New(errlog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Login = log.New(loginlog, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

}
