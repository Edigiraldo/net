package chat

import "flag"

var mode *string = flag.String("mode", "client", "Use 'client' for chat client or 'server' for chat server")

// Run this function with a flag mode with the value 'client' or 'server'
func RunChatMode() {
	flag.Parse()
	switch *mode {
	case "client":
		CreateClient()
	case "server":
		RunServer()
	}
}
