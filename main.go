package main

import(
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os/exec"
	"time"
)

var ROOT_DIR = "/var/www/html/websocket"

func Temp(ws *websocket.Conn) {
	for {
		msg, _ := exec.Command("date").Output()
		fmt.Println("Sending to client: " + string(msg[:]) )
		err := websocket.Message.Send(ws, string(msg[:]) )
		if err != nil {
			fmt.Println("Can't send")
			break
		}
		time.Sleep(2 * 1000 * 1000 * 1000)
		
		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from Client: " + reply)
	} // end for
} // end Temp

func main() {
	fileServer := http.FileServer(http.Dir(ROOT_DIR) )
	http.Handle("/Temp", websocket.Handler(Temp) )
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
} // end of main

func checkError(err error){
	if err != nil {
		fmt.Println("Fatal error: ", err.Error() )
	}
} // end checkError

