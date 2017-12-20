package lircdremotes

import (
	"fmt"
	"log"
	"strings"

	"github.com/chbmuc/lirc"
)

// Remote represents a single lircd remote and all its available commands.
type Remote struct {
	Name     string
	Commands []string
}

func parseKeyNames(reply []string) []string {

	keyNames := []string{}
	for i := 0; i < len(reply); i++ {
		keyString := strings.Split(reply[i], " ")
		keyNames = append(keyNames, keyString[1])
	}
	return keyNames
}

// RemoteCommands Gets all the available remotes and their commands from lircd,
// given the lircd connection.
func RemoteCommands(ir *lirc.Router) []Remote {
	remotesReply := ir.Command(`LIST`)
	// the ir object only keeps one Data object across replies, it seems
	// so, copy the list of remotes out to a new slice
	remotes := make([]string, len(remotesReply.Data))
	remoteCommands := make([]Remote, 0)
	copy(remotes, remotesReply.Data)

	fmt.Printf("%+v\n", remotes)

	for j := 0; j < len(remotes); j++ {
		currentRemote := remotes[j]
		log.Printf("Getting commands for %v\n", currentRemote)
		reply := ir.Command(fmt.Sprintf("LIST %v", currentRemote))
		keyNames := parseKeyNames(reply.Data)
		newRemote := Remote{Name: currentRemote, Commands: keyNames}
		remoteCommands = append(remoteCommands, newRemote)
	}

	return remoteCommands
}
