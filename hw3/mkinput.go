package main

import (
	"os"
	"math/rand"
      "time"

)
func mkf() {
	var str [] byte
	for k := 0; k < 10; k ++ {
		for j := 0; j < 20; j ++ {
			rand.Seed(time.Now().UnixNano())
			x := rand.Intn(25)
			for i := x ; i < 10 + x; i ++ {
				str = append(str, byte('a' + (i % 25)))
			}
			str = append(str, '\n')
		}
		if k != 9 {
			str = append(str, '\f')
		}
	}
	file, _ := os.Create("testf.txt")
	file.Write(str)
}
func mkl() {
	var str [] byte
	for j := 0; j < 200; j ++ {
		rand.Seed(time.Now().UnixNano())
        	x := rand.Intn(25)
		for i := x ; i < 10 + x; i ++ {
			str = append(str, byte('a' + (i % 25)))
		}
		str = append(str, '\n')
	}
	file, _ := os.Create("testl.txt")
	file.Write(str)
}
func main() {
	mkf()
	mkl()
}
