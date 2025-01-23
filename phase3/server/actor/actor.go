package actor

//----

// mutex lock pattern
// each handler is a go routine that talks directly to the crud and synchronously updates data

// the actor pattern
// one go routine responsible for interacting with the data
// use channel messages system between the handler and the crud logic
// the handlers post to a channel and the actor reads from it
// the message contains what they want to do and any data to be changed and a response channel

// share mem by communication rather than communicate by sharing memory
