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
	"errors"
	."agenda/entity"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
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
		email, _ := cmd.Flags().GetString("email")
		telphone,_ := cmd.Flags().GetString("telphone")
		createUser(users, username, password, email, telphone)
		fmt.Println("register called")
		Login.Println("register called")
	},
}


func createUser(users []User, username string, password string, email string, telphone string) {
	if err := validate(users,username,password,email,telphone); err != nil{
		//check user info 
		fmt.Println(err)
		return
	} else {
		//add new user to list
		users = append(users,User{username,password,email,telphone})
		//sync the user info to file
		WriteUserToFile(users)
		fmt.Println("Register success")
		Login.Println("Register " + username + " success")
	}
}

func validate(users []User,username string, password string,email string ,telphone string) error{
	//check user info
	for _, user := range users {
		//check if username repeat
		if user.Username == username{
			Login.Println("repeat username")
			return errors.New("Username has been used!")
		}
	}
	//check the flag value
	if len(password) == 0{
		Login.Println("No input password")
		return errors.New("Need a password")
	} else if len(email) == 0 {
		Login.Println("No input email")
		return errors.New("Need an email")
	} else if len(telphone) == 0 {
		Login.Println("No input telphone")
		return errors.New("Need a telphone")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username","u","","Help message for username")
	registerCmd.Flags().StringP("password","p","","Help message for password")
	registerCmd.Flags().StringP("email","e","","Help message for email")
	registerCmd.Flags().StringP("telphone","t","","Help message for telphone")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
