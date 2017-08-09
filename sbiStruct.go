
package main

import (
	"fmt"   
	)

// transaction will implement the processes
type SBITransaction struct {
}

type user struct {
	uname   		    string  `json:"uname"`
	password     		string  `json:"password"`
	firstname    		float32 `json:"firstname"`
	lastname	 		string  `json:"lastname"`	
}

func dummy() {
		fmt.Printf("dummy")
	}
	func main(){
		
	}
