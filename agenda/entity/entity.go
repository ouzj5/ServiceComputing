package entity

import(
	"io/ioutil"
	"encoding/json"
	"os"
)

const curUserPath string = "entity/data/session.txt"
const UserInfoPath string = "entity/data/user.txt"

type User struct{
	Username string
    Password string
    Email string
    Tel string
}
//use for file io
func ReadUserFromFile () ([]User, error){
	//read every users info
	var users []User	
	if data, err := ioutil.ReadFile(UserInfoPath); err == nil {
		str := string(data)
		json.Unmarshal([]byte(str), &users)
		return users, nil
	} else {
		Error.Println(err)
		return users, err
	}
}

func WriteUserToFile (users []User){
	//write user info to file
	if data, err:=json.Marshal(users); err == nil {
		ioutil.WriteFile(UserInfoPath,[]byte(data),os.ModeAppend)
	} else {
		Error.Println(err)
		panic(err)
	}
	
}
func WriteCurUserToFile (curUser string) {
	//set the login status
	if err := ioutil.WriteFile(curUserPath,[]byte(curUser),os.ModeAppend); err != nil {
		Error.Println(err)
		panic(err)
	}
}

func ReadCurUserToFile() (string, error){
	//check the login status
	if data, err := ioutil.ReadFile(curUserPath); err == nil {
		return string(data), nil
	} else {
		Error.Println(err)
		return string(data), err
	}
}