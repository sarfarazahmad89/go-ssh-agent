//go:build windows
// +build windows
package main
import (
	"log"

	"github.com/Microsoft/go-winio"
	"golang.org/x/crypto/ssh/agent"
)

const (
	sockPath = `\\.\pipe\openssh-ssh-agent`
)


func main() {
    // Start listening on the windows named pipe
	listener, err := winio.ListenPipe(sockPath, nil)
	if err != nil {
		log.Fatalf("Failed to listen on specified sockpath. Error: ", err)
	}

	log.Printf("started ssh agent on `%s`", sockPath)
	defer listener.Close()

	// Start a new keyring
	agentServer := agent.NewKeyring()

	// Serve the agent with the keyring on the socket
	for {
		conn, err := listener.Accept()
		if err != nil {
			conn.Close()
			log.Fatal("connection failed")
		}
		go agent.ServeAgent(agentServer, conn)
		defer conn.Close()
	}

}
