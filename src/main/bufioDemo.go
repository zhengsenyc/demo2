package main

import (
	"fmt"
	/*"os"
	"bufio"*/
)

func main(){
	/*count := map[string]int{}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		count[input.Text()]++
	}

	for k,v := range count{
		fmt.Printf("key=%s    value=%d\n",k,v)
	}*/

	myMap := make(map[string]int)
	fmt.Println(myMap)

	runBufio(myMap)
	fmt.Println(myMap)
}

func runBufio(myMap map[string]int){
	myMap["hello"] = 1
}
