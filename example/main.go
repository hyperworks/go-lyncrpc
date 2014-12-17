package main

import (
	"fmt"
	"github.com/hyperworks/go-lyncrpc"
	"time"
)

const (
	ExampleServerUrl = "sip.reuasmb.net:443"
	ExampleUsername  = "rm-qa_7.thomsonreuters.com@reuasmb.net"
	ExamplePassword  = "Welcome1"
	ExampleRecipient = "rm-qa_9.thomsonreuters.com@reuasmb.net"
)

func main() {
	fmt.Println("dialing...")
	client, e := lyncrpc.Dial("tcp", "0.0.0.0:7331")
	noError(e)

	fmt.Println("testing connection...")
	noError(client.Ping())
	noError(client.Ping())
	noError(client.Ping())

	sleep := func() {
		fmt.Println("sleeping...")
		time.Sleep(3 * time.Second)
	}

	fmt.Println("signing in...")
	e = client.SignIn(ExampleServerUrl, ExampleUsername, ExampleUsername, ExamplePassword)
	noError(e)

	sleep()

	fmt.Println("getting contacts...")
	contacts, e := client.Contacts()
	noError(e)

	for _, contact := range contacts {
		fmt.Printf("contact: %#v\n", contact.Name)
	}

	fmt.Println("chatting...")
	noError(client.BeginConversation(ExampleRecipient))
	noError(client.SendMessage("HELLO"))
	noError(client.SendMessage("WORLD"))
	noError(client.SendMessage("THIS"))
	noError(client.SendMessage("IS"))
	noError(client.SendMessage("TEST"))
	noError(client.SendMessage("MESSAGE"))
	noError(client.EndConversation())

	// fmt.Println("getting availability...")
	// availability, e := client.Availability()
	// noError(e)
	// fmt.Println("availability: " + availability)

	// fmt.Println("changing availability...")
	// e = client.SetAvailability("Busy")
	// noError(e)

	// sleep()

	// fmt.Println("getting availability...")
	// availability, e = client.Availability()
	// noError(e)
	// fmt.Println("availability: " + availability)

	sleep()

	fmt.Println("signing out...")
	e = client.SignOut()
	noError(e)

	fmt.Println("done.")
}

func noError(e error) {
	if e != nil {
		panic(e)
	}
}
