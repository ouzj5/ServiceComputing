package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"os"
	"os/exec"
)
var start = flag.IntP("start", "s", -1, "start page")
var end = flag.IntP("end", "e", -1, "end page")
var length = flag.IntP("length", "l", -1, "page length")
var f = flag.IntP("fend", "f", 0, "use f as flag of an end page")
var dest = flag.StringP("dest", "d", "", "output destination")


func runf(name string) string{
//use '\f' as end of a page
	file,_:= os.Open(name)
	str := make([]byte, 0) //output slice
	for i := 1; i <= *end; i ++{
		data := make([]byte, 1)  
		count,_ := file.Read(data)
		if i >= *start {
			//among the target pages
			for count == 1 && data[0] != '\f' {
				str = append(str, data[0])  //add to the output
				count,_ = file.Read(data)
			}
			if i != *end {
				str = append(str, '\f')  //flag of an end print page
			}
		} else { //before the start page
			for count == 1 && data[0] != '\f' {
				count,_ = file.Read(data) //skip the pre page
			}			
		}
	}
	return string(str)
}
func runl(name string) string{
//static row number as one page
	file,_:= os.Open(name)
	str := make([]byte, 0) //output slice
	for i := 1; i <= *end; i ++{
		data := make([]byte, 1)  
		if i >= *start {
			//among the target pages
			for j := 1; j <= *length; j ++ {
				//read n line
				count,_ := file.Read(data)
				for count == 1 && data[0] != '\n' {
					str = append(str, data[0])  //add to the output
					count,_ = file.Read(data)
				}
				str = append(str, '\n')
			}
			if i != *end {
				str = append(str, '\f')  //flag of an end print page
			}
		} else { //before the start page, discard		
			for j := 1; j <= *length; j ++ {
				//read n line
				count,_ := file.Read(data)
				for count == 1 && data[0] != '\n' {
					count,_ = file.Read(data) //skip the pre page
				}
			}			
		}
	}
	return string(str)
}


func main(){
	var fname string  //the name of input file
	var ret string    //the output string
	flag.Lookup("fend").NoOptDefVal = "1"  //set flag if input a '-f' without arg
	flag.Parse()  //parse

	dslice := []byte(*dest)  //destination string
	/*
		check whether the flag argument is legal
	*/
	if *start > *end || *start < 0  || *end < 0 { //check the essential flag
		fmt.Printf("start or end argument wrong!")
		return
	}
	if *f == 1 && *length != -1 { //check using one method to input
		fmt.Printf("flag -f and -l only can be used one!")
		return
	}
	if *length == -1 { //set default to 72 length per page
		(*length) = 72	
	}
	if flag.NArg() == 0 {
	//if no file arg, ask user to input
		fmt.Scanln(&fname)
	} else {
		fname = flag.Arg(0)
	}

	/*
		start to read the data
	*/
	if *f == 1{
		// '\f' as end signal
		ret = runf(fname)
	} else {
		//static page length
		ret = runl(fname)
	}

	/*
		choose output destination
	*/
	if len(dslice) == 0 {
	//no -d flag, output to stdin
		fmt.Println(ret)
	} else {
	//set -d flag, use lp command with -d arg flag
		tem := "-d"
		arg := []byte(tem)
		arg = append(arg, dslice...)
		cmd := exec.Command("lp", string(arg)) 
		in,_  := cmd.StdinPipe()
		go func() { //send the message just read to the subprocess bu pipe
			defer in.Close()	
			fmt.Fprint(in, ret)
		}()
		cmd.Run()
	}
}
