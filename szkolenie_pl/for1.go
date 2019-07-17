package main
import "fmt"
import "time"

var names = []string{"ala", "wojtek", "krzysiek"}

func main() {
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
	}
	
	time.Sleep(3 * time.Second)
}

