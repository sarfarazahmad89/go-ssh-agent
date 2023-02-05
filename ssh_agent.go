//go:build linux
// +build linux
package main
import (
	"log"
	"net"
	"os"
    "flag"
    "path/filepath"

	"golang.org/x/crypto/ssh/agent"
)

func main() {
    // Get user's homedir
    homedir, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    // Default the sockpath to homedir/.ssh/ssh-auth-sock
    sockPathDefault := filepath.Join(homedir, ".ssh", "ssh-auth-sock")
    sockPath := flag.String("sshpipe", sockPathDefault, "UNIX socket for the OpenSSH agent")
    flag.Parse()

    // Start the server
    listener, err := net.Listen("unix", *sockPath)
    if err != nil {
        log.Fatalf("Failed to listen on specified sockpath. Error: %s", err)
    }
    log.Printf("started ssh agent on `%s`", *sockPath)
	defer listener.Close()

	// Start a new keyring
	agentServer := agent.NewKeyring()

	// Serve the agent with the keyring on the socket.
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
