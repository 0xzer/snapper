# Snapper
Snapper is a lightweight, easy-to-use Go library for interacting with Snapchat's Web API.

## Installation

Use the [package manager](https://golang.org/dl/) to install matchmaker.
```bash
go get github.com/0xzer/snapper
```

# Features
- [x] Fetch account information
- [x] Fetch information about a snapchatter
- [x] Fetch a users bitmoji
- [x] List Friends (friendmojis, streaks, display name & more!)
- [x] List Conversations/Groups
- [x] Sync/DeltaSync Conversations/Groups (Batch & Single)
- [x] Send typing notifications (Groups & Conversations)
- [x] Query messages in conversations/groups (with message history support!)
- [x] Update Conversations/Groups
    - [x] Change group title
    - [x] Mark as read
    - [x] Leave Group
    - [x] Change Retention Policy
    - [x] Clear from chat feed
    - [x] Change notification preference
    - [x] Add participants
- [x] Update Messages
    - [x] Read
    - [x] Release
    - [x] Save
    - [x] Unsave
    - [x] Erase (delete message)
    - [x] Save to camera roll
    - [x] Screenshot
    - [x] Screenrecord
    - [x] Replay
    - [x] React
    - [x] Remove reaction
- [x] Create Conversations/Groups
- [x] Send Messages
    - [x] Chat text messages
    - [x] Send to multiple conversations/groups
    - [x] Reply to messages
    - [x] Set save policy on each message you send
    - [ ] Snaps (coming soon!)
    - [ ] Sticker (coming soon!)
    - [ ] Share (coming soon!)
    - [ ] Note (coming soon!)
    - [ ] SnapReply (coming soon!)
    - [ ] Location (coming soon!)
- [x] Generate fidelius keys

# Usage
- More examples here

```go
package main

import (
	"encoding/json"
	"log"
        "github.com/rs/zerolog"
	"github.com/0xzer/snapper"
	"github.com/0xzer/snapper/protos"
)

var cli *snapper.Client

func main() {
        cookieStr := "sc-cookies-accepted=; EssentialSession=; Preferences=; Performance=; Marketing=; __Host-X-Snap-Client-Cookie=; __Host-sc-a-session=; sc-a-nonce=; __Host-sc-a-nonce=;sc-a-csrf=; blizzard_client_id="
	sess, err := snapper.NewSessionFromCookies(cookieStr) // or snapper.NewSessionFromFile("session.json")
	if err != nil {
		log.Fatal(err)
	}

	cli = snapper.NewClient(sess, zerolog.Logger{}, nil)

	sendChatMessage()

	err = cli.SaveSession("session.json")
	if err != nil {
		log.Fatal(err)
	}
}

func sendChatMessage() {
	conversations := getExampleConversations()
	firstConversation := conversations[0]

	messageBuilder := cli.NewCreateMessageBuilder().
	AddConversationDestination(firstConversation).
	SetTextMessage("hello!").
	SetSavePolicy(protos.ContentEnvelope_SavePolicy_LIFETIME)
	
	sent, err := cli.Messaging.SendContentMessage(messageBuilder)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(sent)
}

func getExampleConversations() []*protos.Conversation {
	conversationEntries, err := cli.Messaging.SyncConversations(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	conversationsBatch := cli.Messaging.BatchRequestFromEntry(conversationEntries.GetConversations())
	batchDeltaSync, err := cli.Messaging.BatchDeltaSync(conversationsBatch)
	if err != nil {
		log.Fatal(err)
	}

	return cli.Messaging.BatchResponseToConversationSlice(batchDeltaSync)
}
```