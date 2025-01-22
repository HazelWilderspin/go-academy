package actor

//----

// mutex lock pattern
// each handler is a go routing that talks directly to the crud and synchronously updates data

// the actor pattern
// one go routine responsible for interacting with the data
// use some kind of message system between the handler and the crud logic
// the handlers post to a channel and the crud logic reads from it
// the message contains what they want to do and any data to be changed

// share mem by communication rather than communicate by sharing memory

// func actor() {
// 	queue := make(chan string, 2)
// 	queue <- "mmmmmm"

// 	msg1, ok := <-queue
// 	fmt.Printf("-1- Reading channel message: %s\nChannel read successful: %t\n", msg1, ok)

// 	queue <- "mmmmmm2"

// 	close(queue)

// 	select {
// 	case msg2, ok := <-queue:
// 		fmt.Printf("-2- Reading channel message: %s\nChannel read successful: %t\n", msg2, ok)
// 	default:
// 		fmt.Println("No messages left on the channel")
// 	}
// }
