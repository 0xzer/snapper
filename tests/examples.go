package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/0xzer/snapper"
	"github.com/0xzer/snapper/debug"
	"github.com/0xzer/snapper/protos"
)


var logger = debug.NewLogger()
var cli *snapper.Client

func main() {
	sess, err := snapper.NewSessionFromFile("session.json")
	if err != nil {
		log.Fatal(err)
	}

	cli = snapper.NewClient(sess, logger, nil)

	log.Println(getExampleConversations())

	err = cli.SaveSession("session.json")
	if err != nil {
		log.Fatal(err)
	}
}

func sendTypingNotification(conversations []*protos.ConversationEntry) {
	conversation := conversations[0]
	sent, err := cli.Messaging.SendTypingNotification(conversation.GetVersionInfo().GetConversationId(), 79)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sent)
	os.Exit(1)
}

func saveMessage() {
	batchResponse := getExampleBatch()

	conversationData := batchResponse[0].GetSuccessResponse()

	conversation := conversationData.GetConversation()
	messages := conversationData.GetContentMessages()
	messageToSaveId := messages[8].GetMessageId() // random message

	updateMessageBuilder := cli.NewUpdateMessageBuilder(messageToSaveId).
	SetConversation(conversation).
	SetSave()

	updated, err := cli.Messaging.UpdateContentMessage(updateMessageBuilder)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(updated)
	os.Exit(1)
}

func removeReaction() {
	batchResponse := getExampleBatch()

	conversationData := batchResponse[0].GetSuccessResponse()

	conversation := conversationData.GetConversation()
	messages := conversationData.GetContentMessages()
	messageToReactToId := messages[8].GetMessageId() // random message

	updateMessageBuilder := cli.NewUpdateMessageBuilder(messageToReactToId).
	SetConversation(conversation).
	SetMessage(messageToReactToId).
	SetRemoveReaction()

	updated, err := cli.Messaging.UpdateContentMessage(updateMessageBuilder)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(updated)
	os.Exit(1)
}

func reactToMessage() {
	batchResponse := getExampleBatch()

	conversationData := batchResponse[0].GetSuccessResponse()

	conversation := conversationData.GetConversation()
	messages := conversationData.GetContentMessages()
	messageToReactToId := messages[8].GetMessageId() // random message

	updateMessageBuilder := cli.NewUpdateMessageBuilder(messageToReactToId).
	SetConversation(conversation).
	SetMessage(messageToReactToId).
	SetReaction(3)

	updated, err := cli.Messaging.UpdateContentMessage(updateMessageBuilder)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(updated)
	os.Exit(1)
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
	os.Exit(1)
}

func createGroup() {
	/*
	createBuilder := cli.NewCreateConversationBuilder().
	SetConversationType(protos.ConversationType_USER_CREATED_GROUP).
	SetTitle("the group title").
	AddParticipant(... some participant id).
	AddParticipants(... a slice of participant ids)

	created, err := cli.Messaging.CreateConversation(createBuilder)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(created)
	*/
}

func replyToMessage() {

	batchResponse := getExampleBatch()

	conversationData := batchResponse[0].GetSuccessResponse()

	conversation := conversationData.GetConversation()
	messages := conversationData.GetContentMessages()
	messageToReplyToId := messages[5].GetMessageId() // random message

	messageBuilder := cli.NewCreateMessageBuilder().
	AddConversationDestination(conversation).
	SetTextMessage("test reply").
	SetReplyMessage(messageToReplyToId).
	SetSavePolicy(protos.ContentEnvelope_SavePolicy_LIFETIME)
	
	sent, err := cli.Messaging.SendContentMessage(messageBuilder)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(sent)
	os.Exit(1)
}

func deltaSync(conversations []*protos.ConversationEntry) *protos.Conversation {
	conversation := conversations[0]
	conversationId := conversation.GetVersionInfo().GetConversationId()

	deltaSyncedConversation, syncErr := cli.Messaging.DeltaSync(conversationId, nil)
	if syncErr != nil {
		log.Fatal(syncErr)
	}

	deltaSyncData, err := json.Marshal(deltaSyncedConversation)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("deltaSync.json", deltaSyncData, os.ModePerm)
	return deltaSyncedConversation.GetConversation()
}

func batchDeltaSync(conversations []*protos.ConversationEntry) *protos.Conversation {
	batchRequests := cli.Messaging.BatchRequestFromEntry(conversations)
	batchDeltaSyncedConversations, syncErr := cli.Messaging.BatchDeltaSync(batchRequests)
	if syncErr != nil {
		log.Fatal(syncErr)
	}

	deltaSyncData, err := json.Marshal(batchDeltaSyncedConversations)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("batchDeltaSync.json", deltaSyncData, os.ModePerm)
	return batchDeltaSyncedConversations.GetDeltaSyncResponses()[0].GetSuccessResponse().GetConversation()
}

func queryConversations(syncResponse *protos.SyncConversationsResponse) {
	requestedPageSize := 40 // default
	conversationResponse, syncErr := cli.Messaging.QueryConversations(syncResponse.GetSyncToken(), requestedPageSize, nil)
	if syncErr != nil {
		log.Fatal(syncErr)
	}

	queryConversationsData, err := json.Marshal(conversationResponse)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("queryConversations.json", queryConversationsData, os.ModePerm)
}

func queryMessages(conversation *protos.Conversation) {
	messagesQueried, syncErr := cli.Messaging.QueryMessages(conversation.GetConversationId(), 100, conversation.GetVersion())
	if syncErr != nil {
		log.Fatal(syncErr)
	}

	messagesData, err := json.Marshal(messagesQueried)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("queryMessages.json", messagesData, os.ModePerm)
}

func getExampleBatch() []*protos.DeltaSyncResponseWrapper {
	conversationEntries, err := cli.Messaging.SyncConversations(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	conversationsBatch := cli.Messaging.BatchRequestFromEntry(conversationEntries.GetConversations())
	batchDeltaSync, err := cli.Messaging.BatchDeltaSync(conversationsBatch)
	if err != nil {
		log.Fatal(err)
	}

	return batchDeltaSync.DeltaSyncResponses
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