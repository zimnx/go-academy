package main 
import "log"

type myError struct { msg string }
func (e *myError) Error() string {
	return "!!!!" + e.msg + "!!!!"
}

func foo() error {
	return &myError{"oh no!"} // HL
}

func main() {
	err := foo() 
	if err != nil { 
		log.Fatalln("got error", err)
	}	
	log.Println("hurray! no errors")
}