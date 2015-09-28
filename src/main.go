package main
import (
	"fmt"
	"os"
)


func main(){
	fmt.Println("Hello server!")
	stream_id:= os.Args[1]
	fmt.Println("stream id: %s",stream_id)




}