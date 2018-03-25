package test

import (
	"fmt"
	"log"
	"sync"

	"github.com/lnxjedi/gopherbot/bot"
)

// Global persistent map of user name to user index
var userMap = make(map[string]int)

type testUser struct {
	Name                                        string // username / handle
	InternalID                                  string // connector internal identifier
	Email, FullName, FirstName, LastName, Phone string
}

type config struct {
	StartChannel string // the initial channel
	StartUser    string // the initial userid
	BotName      string // the short name used for addressing the robot
	BotFullName  string // the full name of the bot
	Users        []testUser
}

var lock sync.Mutex // package var lock

func init() {
	bot.RegisterConnector("test", Initialize)
}

// Initialize sets up the connector and returns a connector object
func Initialize(robot bot.Handler, l *log.Logger) bot.Connector {
	var c config

	err := robot.GetProtocolConfig(&c)
	if err != nil {
		robot.Log(bot.Fatal, fmt.Errorf("Unable to retrieve protocol configuration: %v", err))
	}
	found := false
	for i, u := range c.Users {
		userMap[u.Name] = i
		if c.StartUser == u.Name {
			found = true
		}
	}
	if !found {
		robot.Log(bot.Fatal, "Start user \"%s\" not listed in Users array")
	}

	tc := &TestConnector{
		currentChannel: c.StartChannel,
		currentUser:    c.StartUser,
		channels:       make([]string, 0),
		botName:        c.BotName,
		botFullName:    c.BotFullName,
		botID:          "deadbeef",
		users:          c.Users,
		Listener:       make(chan string),
		Speaking:       make(chan string),
	}

	tc.Handler = robot
	tc.SetFullName(tc.botFullName)
	tc.Log(bot.Debug, "Set bot full name to", tc.botFullName)
	tc.SetName(tc.botName)
	tc.Log(bot.Info, "Set bot name to", tc.botName)

	return bot.Connector(tc)
}