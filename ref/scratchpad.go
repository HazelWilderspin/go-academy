package ref


import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func scratch() {
	// --------------------------------------------------------------------------------------------
	msgChannel := make(chan string, 2)
	msgChannel <- "mmmmmm"

	msg1, ok := <-msgChannel
	fmt.Printf("-1- Reading channel message: %s\nChannel read successful: %t\n", msg1, ok)

	msgChannel <- "mmmmmm2"

	close(msgChannel)

	select {
	case msg2, ok := <-msgChannel:
		fmt.Printf("-2- Reading channel message: %s\nChannel read successful: %t\n", msg2, ok)
	default:
		fmt.Println("No messages left on the channel")
	}
	// --------------------------------------------------------------------------------------------



    type Item struct {
    }


    // / Main
    func main() {
        fmt.Println("Hello World")
        http.HandleFunc("/get", getHandler)
        requestChan = make(chan request, 100)
        defer close(requestChan)
        go actor()
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        <-c
    }


    // / Handlers
    func getHandler(w http.ResponseWriter, r *http.Request) {
        items := get()
        data, _ := json.Marshal(items)
        w.Write(data)
    }


    // / Concurrency
    type request struct {
        response chan []Item
        action   string
    }

    var requestChan chan request

    func get() []Item {
        response := make(chan []Item)
        req := request{
            response: response,
            action:   "get",
        }
        requestChan <- req
        data := <-response
        return data
    }

    func actor() {
        for req := range requestChan {
            switch req.action {
            case "get":
                data := loadData()
                req.response <- data
            case "delete":
				del()
            }
            close(req.response)
        }
    }


    // / Client/repo
    func loadData() []Item {
        // load from disk
        return []Item{}
    }

	func del(i Item){

	}
}
