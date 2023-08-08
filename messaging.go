package snapper

import (
	"bytes"
	"github.com/0xzer/snapper/data/headers"
	"github.com/0xzer/snapper/data/paths"
	"github.com/0xzer/snapper/protos"
	"github.com/0xzer/snapper/types"
)

type Messaging struct {
	client *Client
}

// if it's the first SyncConversations request, pass nil as both syncToken & conversationVersionInfos.
func (m *Messaging) SyncConversations(syncToken []byte, conversationVersionInfos []*protos.ConversationVersionInfo) (*protos.SyncConversationsResponse, error) {
	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	payload := &protos.SyncConversationsRequest{
		SelfUserId: m.client.Session.CurrentUser.EncodedUserId,
		VersionInfos: conversationVersionInfos,
		SyncToken: []byte("useV3"),
	}

	if syncToken != nil {
		payload.SyncToken = syncToken
	}

	payloadBytes, err := protos.EncodeProtoMessage(payload)
	if err != nil {
		return nil, err
	}
	
	_, responseBody, err := m.client.Request.MakeRequest(paths.SYNC_CONVERSATIONS, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.SyncConversationsResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

// only thing needed is conversationId for now, so just pass that as nil.
// if you are having trouble, try passing in otherParticipantId aswell.
func (m *Messaging) DeltaSync(conversationId *protos.UUID, otherParticipantUserId *protos.UUID) (*protos.DeltaSyncResponse, error) {
	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	payload := &protos.DeltaSyncRequest{
		SelfUserId: m.client.Session.CurrentUser.EncodedUserId,
		OtherParticipantUserId: otherParticipantUserId,
		ConversationId: conversationId,
	}

	payloadBytes, err := protos.EncodeProtoMessage(payload)
	if err != nil {
		return nil, err
	}
	
	_, responseBody, err := m.client.Request.MakeRequest(paths.DELTA_SYNC_CONVERSATIONS, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.DeltaSyncResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

func (m *Messaging) BatchDeltaSync(deltaSyncRequests []*protos.DeltaSyncRequest) (*protos.BatchDeltaSyncResponse, error) {
	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	payload := &protos.BatchDeltaSyncRequest{
		DeltaSyncRequests: deltaSyncRequests,
	}

	payloadBytes, err := protos.EncodeProtoMessage(payload)
	if err != nil {
		return nil, err
	}
	
	_, responseBody, err := m.client.Request.MakeRequest(paths.BATCH_DELTA_SYNC_CONVERSATIONS, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.BatchDeltaSyncResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

// the sync token you got earlier already contains all the information needed to query the conversation
// im sure this endpoint will eventually support batch requests aswell since the response does
// paginationInfo is not required, just pass nil if you don't want to send it
func (m *Messaging) QueryConversations(syncToken []byte, requestedPageSize int, paginationInfo *protos.QueryConversationsRequest_PaginationInfo) (*protos.QueryConversationsResponse, error) {
	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	payload := &protos.QueryConversationsRequest{
		SelfUserId: m.client.Session.CurrentUser.EncodedUserId,
		PaginationInfo: paginationInfo,
		RequestedPageSize: int32(requestedPageSize),
		SyncToken: syncToken,
	}

	payloadBytes, err := protos.EncodeProtoMessage(payload)
	if err != nil {
		return nil, err
	}
	
	_, responseBody, err := m.client.Request.MakeRequest(paths.QUERY_CONVERSATIONS, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.QueryConversationsResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

// currentVersion is the cursor from where to query messages from
func (m *Messaging) QueryMessages(conversationId *protos.UUID, requestedCountSize int, currentVersion int64) (*protos.QueryMessagesResponse, error) {
	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	payload := &protos.QueryMessagesRequest{
		SelfUserId: m.client.Session.CurrentUser.EncodedUserId,
		ConversationId: conversationId,
		RequestedCountSize: int32(requestedCountSize),
		CurrentVersion: currentVersion,
	}

	payloadBytes, err := protos.EncodeProtoMessage(payload)
	if err != nil {
		return nil, err
	}
	
	_, responseBody, err := m.client.Request.MakeRequest(paths.QUERY_MESSAGES, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.QueryMessagesResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

func (m *Messaging) GetGroups() (*protos.GetGroupsResponseWrapper, error) {
	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	payload := &protos.GetGroupsRequest{
		SelfUserId: m.client.Session.CurrentUser.EncodedUserId,
	}

	payloadBytes, err := protos.EncodeProtoMessage(payload)
	if err != nil {
		return nil, err
	}
	
	_, responseBody, err := m.client.Request.MakeRequest(paths.GET_GROUPS_URL, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.GetGroupsResponseWrapper
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

func (m *Messaging) SendContentMessage(messageBuilder *CreateMessageBuilder) (*protos.CreateContentMessageResponse, error) {
	payloadBytes, err := messageBuilder.build()
	if err != nil {
		return nil, err
	}

	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	_, responseBody, err := m.client.Request.MakeRequest(paths.CREATE_CONTENT_MESSAGE, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.CreateContentMessageResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

func (m *Messaging) UpdateContentMessage(updateMessageBuilder *UpdateMessageBuilder) (*protos.UpdateContentMessageResponse, error) {
	payloadBytes, err := updateMessageBuilder.build()
	if err != nil {
		return nil, err
	}

	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	_, responseBody, err := m.client.Request.MakeRequest(paths.UPDATE_CONTENT_MESSAGE, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.UpdateContentMessageResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

func (m *Messaging) CreateConversation(createConversationBuilder *CreateConversationBuilder) (*protos.CreateConversationResponse, error) {
	payloadBytes, err := createConversationBuilder.build()
	if err != nil {
		return nil, err
	}

	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	_, responseBody, err := m.client.Request.MakeRequest(paths.CREATE_CONVERSATION, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.CreateConversationResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

func (m *Messaging) UpdateConversation(updateConversationBuilder *UpdateConversationBuilder) (*protos.UpdateConversationResponse, error) {
	payloadBytes, err := updateConversationBuilder.build()
	if err != nil {
		return nil, err
	}

	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	_, responseBody, err := m.client.Request.MakeRequest(paths.UPDATE_CONVERSATION, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.UpdateConversationResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

func (m *Messaging) SendTypingNotification(conversationId *protos.UUID, lastCreatedMessageId int64) (*protos.SendTypingNotificationResponse, error) {
	headers := headers.NewCoreHeaders(m.client.Session.SnapCookies, m.client.Session.SnapTokens, m.client.device)
	payload := &protos.SendTypingNotificationRequest{
		SelfUserId: m.client.Session.CurrentUser.EncodedUserId,
		ConversationId: conversationId,
		LastCreatedMessageId: lastCreatedMessageId,
	}

	payloadBytes, err := protos.EncodeProtoMessage(payload)
	if err != nil {
		return nil, err
	}
	
	_, responseBody, err := m.client.Request.MakeRequest(paths.SEND_TYPING_NOTIFICATION, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}

	var resData protos.SendTypingNotificationResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &resData, nil
}

func (m *Messaging) BatchRequestFromEntry(conversations []*protos.ConversationEntry) []*protos.DeltaSyncRequest {
	deltaRequests := make([]*protos.DeltaSyncRequest, 0)
	for _, conv := range conversations {
		deltaRequest := &protos.DeltaSyncRequest{
			ConversationId: conv.GetVersionInfo().GetConversationId(),
			SelfUserId: m.client.Session.CurrentUser.EncodedUserId,
			OtherParticipantUserId: m.GetOtherParticipantId(conv.GetParticipants()),
		}
		deltaRequests = append(deltaRequests, deltaRequest)
	}
	return deltaRequests
}

func (m *Messaging) GetOtherParticipantId(participants []*protos.UUID) *protos.UUID {
	selfUserId := m.client.Session.CurrentUser.EncodedUserId.EncodedId
	for _, id := range participants {
		if !bytes.Equal(selfUserId, id.EncodedId) {
			return id
		}
	}

	return nil
}

func (m *Messaging) BatchResponseToConversationSlice(responses *protos.BatchDeltaSyncResponse) []*protos.Conversation {
	deltaResponses := responses.GetDeltaSyncResponses()
	conversations := make([]*protos.Conversation, 0)

	if deltaResponses == nil || len(deltaResponses) <= 0 {
		return conversations
	}

	for _, res := range deltaResponses {
		successResponse := res.GetSuccessResponse()
		if successResponse == nil {
			continue
		}
		conversations = append(conversations, successResponse.GetConversation())
	}

	return conversations
}
