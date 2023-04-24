package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// initialzie variables
var firstname string
var lastname string
var adress string
var house_number int
var postcode int
var email_adress string
var path string

func main() {
	getArgData()
	JsonData := create_json()
	fmt.Printf("%s", JsonData)

}

func getArgData() {
	//set commandline arguments connected to the values
	flag.StringVar(&firstname, "f", "false", "Enter you firstname(string)")
	flag.StringVar(&lastname, "l", "false", "Enter you lastname(string)")
	flag.StringVar(&adress, "a", "false", "Enter you adress(string)")
	flag.IntVar(&house_number, "hn", 0, "Enter you house_number(int)")
	flag.IntVar(&postcode, "pc", 0, "Enter you postcode(int)")
	flag.StringVar(&email_adress, "e", "false", "Enter you email_adress(string)")
	flag.StringVar(&path, "path", "", "Enter the path where the .json file should be created. If not used it just returns a commandline string(string)")
	//parse the flag values
	flag.Parse()
}

// create map with all user data
func create_json() []byte { //[]byte debug line
	//error checking for default values
	if firstname == "flase" || lastname == "false" || adress == "false" || email_adress == "false" || house_number == 0 || postcode == 0 {
		log.Fatal("missing data, use all required flags; -f -l -a -hn -pc -e. Type -help for explanation.")
	}
	userdata := map[string]interface{}{
		"firstname":    firstname,
		"lastname":     lastname,
		"adress":       adress + fmt.Sprint(house_number),
		"postcode":     postcode,
		"email_adress": email_adress}

	//create json without prefix and 1 space indent
	jsonData, err := json.MarshalIndent(userdata, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	if path != "" {
		write_json_file(jsonData)
	}

	//debug line
	return jsonData
}

// creates a new .json file with the users first- and lastname as title and data from jsonData []byte object.
func write_json_file(jsonData []byte) {
	//checks if the dir path is available, if not fatal error stop programm
	_, error := os.ReadDir(path)
	if error != nil {
		log.Fatal(error)
	}

	err := ioutil.WriteFile(fmt.Sprintf("%v"+"%v%v.json", path, firstname, lastname), jsonData, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
