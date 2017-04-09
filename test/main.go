package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("go.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	fmt.Println("Started Print log")
	a := 123
	log.Println("Test Success")
	log.Printf("Integer a: %v:", a)

}

// package main
//
// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func init() {
// 	nf, err := os.Create("log.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	log.SetOutput(nf)
// }
//
// func main() {
// 	_, err := os.Open("no-file.txt")
// 	if err != nil {
// 		//		fmt.Println("err happened", err)
// 		log.Println("err happened", err)
// 		//		log.Fatalln(err)
// 		//		panic(err)
// 	}
// }
