/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	."agenda/entity"
	"github.com/spf13/cobra"
)


// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		users, err := ReadUserFromFile()
		if err != nil {
			panic(err)
		}
		username, _ := cmd.Flags().GetString("username")
		password, _:= cmd.Flags().GetString("password")
		login(users, username, password)
	},
}

func isLogin() bool{
	//check if already login
	curUser, err := ReadCurUserToFile()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return string(curUser) != "logout"
}

func login(users []User, username string, password string) {
	fmt.Println("user login")
	Login.Println("user login")
	//check from cache whether the status is login.
	if isLogin() {
		fmt.Println("Please logout first!")
		Login.Println("Login failed, Already login!")
		return
	}
	//validate username and password
	if len(username) == 0 || len(password) == 0 {
		fmt.Println("Need a username and a password")
		Login.Println("Login failed, vacant username or pasword.")
		return
	}

	for _,user := range users{
		if user.Username == username && user.Password == password{
			WriteCurUserToFile(user.Username)
			fmt.Println("Login success! Username is " + user.Username + "!")
			Login.Println(user.Username + " Login success.")
			return
		}
	}
	fmt.Println("Username or password wrong!")
	Login.Println("Username or password wrong")
}


func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username","u","","Help message for username")
	loginCmd.Flags().StringP("password","p","","Help message for password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
