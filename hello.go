package main

import (
	"flag"
)

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "youremail.com", "the email to scan")
	flag.Parse();

	// fmt.Println("Folder: ", folder)
	// fmt.Println("Email: ", email)

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}