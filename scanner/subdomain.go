package subdomain

import (
	"bufio"
	"fmt"
	"github.com/cweill/Permute-Golang"
	"net"
	"os"
)

func Bruteforce(domain string, wordList string) {

	//subdomains := genPermutatons(2)
	wordListFile, err := os.Open(wordList)

	if err != nil {
		panic(err)
	}

	defer wordListFile.Close()

	fileScanner := bufio.NewScanner(wordListFile)

	for fileScanner.Scan() {
		name := fileScanner.Text()
		host := name + "." + domain

		ips, err := net.LookupIP(host)

		if err != nil {
			fmt.Println(host)
			fmt.Println(err)
		} else {
			fmt.Println(host)
			fmt.Println(ips)
		}
	}

	if err := fileScanner.Err(); err != nil {
		panic(err)
	}
}

//leave this for now
//generates too many permutations
//we'll just use lists for brute forcing
func genPermutatons(number int) []string {
	//this needs to return permutations of all characters
	//not sure what would be a good way to do it
	//so lets do it one by one
	//validChars := "abcdefghijklmnopqrstuvwxyz0123456789-"

	fmt.Println(permute.LexicographicPermutations("abc"))
	return []string{"derp", "another"}
}
