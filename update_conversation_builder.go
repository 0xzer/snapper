package snapper

import (
	"github.com/0xzer/snapper/crypto"
	"github.com/0xzer/snapper/data/methods"
	"github.com/0xzer/snapper/protos"
	"github.com/0xzer/snapper/snaperror"
)

type UpdateConversationBuilder struct {
	client *Client

	updateRequest *protos.UpdateConversationRequest

	err *snaperror.UpdateConversationBuilderError
}

func (c *Client) NewUpdateConversationBuilder(conversation *protos.Conversation) *UpdateConversationBuilder {
	clientResolutionId, _ := crypto.GenerateBitUUID()
	return &UpdateConversationBuilder{
		client: c,
		updateRequest: &protos.UpdateConversationRequest{
			ConversationId: conversation.GetConversationId(),
			CurrentVersion: conversation.GetVersion(),
			ClientResolutionId: clientResolutionId.Uint64(),
		},
	}
}

// for reuse
func (cb *UpdateConversationBuilder) SetConversation(conversation *protos.Conversation) *UpdateConversationBuilder {
	cb.updateRequest.ConversationId = conversation.GetConversationId()
	cb.updateRequest.CurrentVersion = conversation.GetVersion()
	return cb
}

func (cb *UpdateConversationBuilder) SetRetentionPolicy(policy *protos.RetentionPolicy) *UpdateConversationBuilder {
	cb.updateRequest.UpdateConversationAction = &protos.UpdateConversationRequest_RetentionPolicy{
		RetentionPolicy: &protos.UpdateConversationRetentionPolicy{
			SelfUserId: cb.client.Session.CurrentUser.EncodedUserId,
			RetentionPolicy: policy,
		},
	}
	return cb
}

func (cb *UpdateConversationBuilder) SetRead(lastMessageId int64) *UpdateConversationBuilder {
	cb.updateRequest.UpdateConversationAction = &protos.UpdateConversationRequest_Read{
		Read: &protos.UpdateConversationRead{
			SelfUserId: cb.client.Session.CurrentUser.EncodedUserId,
			ReadConversationMessageData: &protos.ReadConversationMessageData{
				LastMessageId: lastMessageId,
			},
		},
	}
	return cb
}

func (cb *UpdateConversationBuilder) SetNotificationPreference(preference protos.ChatNotificationPreference) *UpdateConversationBuilder {
	notificationPref := &protos.UpdateConversationRequest_NotificationPreference{
		NotificationPreference: &protos.UpdateConversationNotificationPreference{
			SelfUserId: cb.client.Session.CurrentUser.EncodedUserId,
			SomeInt2: 1,
			Setting: preference,
		},
	}

	if preference == protos.ChatNotificationPreference_ALL_MESSAGES {
		notificationPref.NotificationPreference.SomeInt = 1
	}

	cb.updateRequest.UpdateConversationAction = notificationPref
	return cb
}

func (cb *UpdateConversationBuilder) SetTitle(newTitle string, oldTitle string) *UpdateConversationBuilder {
	cb.updateRequest.UpdateConversationAction = &protos.UpdateConversationRequest_ConversationTitle{
		ConversationTitle: &protos.UpdateConversationTitle{
			SelfUserId: cb.client.Session.CurrentUser.EncodedUserId,
			OldTitle: oldTitle,
			NewTitle: newTitle,
		},
	}
	return cb
}

func (cb *UpdateConversationBuilder) SetClearFromFeed() *UpdateConversationBuilder {
	cb.updateRequest.UpdateConversationAction = &protos.UpdateConversationRequest_ClearFromChatFeed{
		ClearFromChatFeed: &protos.UpdateConversationClearFromFeed{
			SelfUserId: cb.client.Session.CurrentUser.EncodedUserId,
			UpdateTimestamp: methods.GetTimestamp(),
		},
	}
	return cb
}

func (cb *UpdateConversationBuilder) SetLeaveConversation() *UpdateConversationBuilder {
	cb.updateRequest.UpdateConversationAction = &protos.UpdateConversationRequest_LeaveConversation{
		LeaveConversation: &protos.UpdateConversationLeave{
			SelfUserId: cb.client.Session.CurrentUser.EncodedUserId,
			SelfUserIdLeft: cb.client.Session.CurrentUser.EncodedUserId,
		},
	}
	return cb
}

func (cb *UpdateConversationBuilder) AddParticipants(participants []*protos.UUID) *UpdateConversationBuilder {
	cb.updateRequest.UpdateConversationAction = &protos.UpdateConversationRequest_AddParticipants{
		AddParticipants: &protos.UpdateConversationAddParticipants{
			SelfUserId: cb.client.Session.CurrentUser.EncodedUserId,
			Participants: participants,
		},
	}
	return cb
}

func (cb *UpdateConversationBuilder) build() ([]byte, error) {
	if cb.updateRequest.UpdateConversationAction == nil {
		cb.err = &snaperror.UpdateConversationBuilderError{ErrorMessage: "builder must have some action set before building"}
	}

	return protos.EncodeProtoMessage(cb.updateRequest)
}