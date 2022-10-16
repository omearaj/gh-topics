package main

import (
	"fmt"
        "os"

	//"github.com/cli/go-gh"
)

func main() {

    // The first argument
    // is always program name
    myProgramName := os.Args[0]
  
    // this will take 4
    // command line arguments
    cmdArgs := os.Args[4]
  
    // getting the arguments
    // with normal indexing
    gettingArgs := os.Args[2]
  
    toGetAllArgs := os.Args[1:]
  
    // it will display
    // the program name
    fmt.Println(myProgramName)
      
    fmt.Println(cmdArgs)
      
    fmt.Println(gettingArgs)
      
    fmt.Println(toGetAllArgs)



	//fmt.Println("hi world, this is the gh-jamie extension!")
	//client, err := gh.RESTClient(nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//response := struct {Login string}{}
	//err = client.Get("user", &response)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("running as %s\n", response.Login)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
