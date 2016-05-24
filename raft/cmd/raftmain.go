package main

import (
	"flag"
	"fmt"

	"github.com/PatrickMcCormack/goplay/raft"
)

func main() {

	var verbose bool
	clusterPtr := flag.String("cluster", "", "comma seperated list of nodes")
	localPtr := flag.String("local", "127.0.0.1:2000", "This node's IP:Port")
	flag.BoolVar(&verbose, "verbose", false, "Verbose output")

	flag.Parse()

	serverStateFile := *localPtr + ".state"

	if verbose {
		fmt.Println("cluster:", *clusterPtr)
		fmt.Println("local:", *localPtr)
		fmt.Println("verbose:", verbose)
		fmt.Println("local:", serverStateFile)
	}

	// instantiate the state for this node
	var node raft.Node
	node.State = raft.FOLLOWER
	fmt.Printf("node is %v\n", node)

	node.WritePersistantState(serverStateFile)
	fmt.Printf("node is %v\n", node)

	node.ReadPersistantState(serverStateFile)
	fmt.Printf("node is %v\n", node)

}

// Functions of Go RPC have to comply with the following requirements
// otherwise the remote calls will be ignored.
//
// 1. Functions are exported (capitalize).
// 2. Functions have to have two arguments with exported types.
// 3. The first argument is for receiving from the client,
//    and the second one has to be a pointer and is for replying to the client.
// 4. Functions have to have a return value of error type.

/*

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"time"
)
func main() {

	// args for this simple example are:
	// the port this server is listening on and the port of the second server
	// ports to be expressed as :1234
	var localPort string
	var remotePort string
	if len(os.Args) == 3 {
		localPort = os.Args[1]
		remotePort = os.Args[2]
	} else {
		fmt.Println("Usage: ", os.Args[0], ":this-service-port :remote-service-port")
		os.Exit(1)
	}

	go heartbeat(remotePort, localPort)

	// The server main loop
	ae := new(AppendEntries)
	rpc.Register(ae)

	tcpAddr, err := net.ResolveTCPAddr("tcp", localPort)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	call := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		call++
		go func(c int) {
			fmt.Printf("In go routine serving call number %v\n", call)
			rpc.ServeConn(conn)
		}(call)
	}

}

// heartbeat other server
func heartbeat(remoteServicePort string, localServerPort string) {

	service := "127.0.0.1" + remoteServicePort
	localInfo := "127.0.0.1" + localServerPort

	for {
		time.Sleep(time.Millisecond * 1000)
		client, err := rpc.Dial("tcp", service)
		if err != nil {
			//log.Fatal("dialing:", err)
			fmt.Printf("remote conntection failed - error is - %v\n", err)
		} else {
			// Synchronous call
			var result string
			err = client.Call("AppendEntries.HeartBeat", localInfo, &result)
			if err != nil {
				// log.Fatal("AppendEntries error:", err)
				fmt.Printf("remote call failed - error is - %v\n", err)
			}
			fmt.Printf("AppendEntries.Echo = %v\n", result)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

*/
