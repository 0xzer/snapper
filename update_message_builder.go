package snapper

import (
	"github.com/0xzer/snapper/crypto"
	"github.com/0xzer/snapper/data/methods"
	"github.com/0xzer/snapper/protos"
	"github.com/0xzer/snapper/snaperror"
)

type UpdateMessageBuilder struct {
	client *Client
	conversationVersion *protos.ConversationVersionInfo
	updateAction *protos.UpdateAction

	err *snaperror.UpdateMessageBuilderError
}

func (c *Client) NewUpdateMessageBuilder(messageId int64) *UpdateMessageBuilder {
	return &UpdateMessageBuilder{
		client: c,
		conversationVersion: &protos.ConversationVersionInfo{},
		updateAction: &protos.UpdateAction{
			MessageId: messageId,
			SenderId: c.Session.CurrentUser.EncodedUserId,
		},
	}
}

func (m *UpdateMessageBuilder) SetConversation(conversation *protos.Conversation) *UpdateMessageBuilder {
	m.conversationVersion.ConversationVersion = conversation.GetVersion()
	m.conversationVersion.ConversationId = conversation.GetConversationId()
	return m
}

func (m *UpdateMessageBuilder) SetMessage(messageId int64) *UpdateMessageBuilder {
	m.updateAction.MessageId = messageId
	return m
}

func (m *UpdateMessageBuilder) SetReaction(intentType int64) *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_MessageReaction{
		MessageReaction: &protos.ReactionRequest{
				NewReaction: &protos.NewReaction{
					Reaction: &protos.ReactionType{
						Reaction: &protos.ReactionType_IntentType{
							IntentType: intentType,
						},
					},
			},
		},
	}
	return m
}

func (m *UpdateMessageBuilder) SetRemoveReaction() *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_RemoveReaction{
		RemoveReaction: &protos.RemoveReactionRequest{},
	}
	return m
}

func (m *UpdateMessageBuilder) SetRead() *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_Read{
		Read: &protos.Read{},
	}
	return m
}

func (m *UpdateMessageBuilder) SetRelease() *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_Release{
		Release: &protos.Release{},
	}
	return m
}

func (m *UpdateMessageBuilder) SetSave() *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_Save{
		Save: &protos.Save{
			EnvelopeEncryption: &protos.EnvelopeEncryption{
				Method: &protos.EnvelopeEncryption_None{None: &protos.Empty{}},
			},
		},
	}
	return m
}

func (m *UpdateMessageBuilder) SetUnsave() *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_Unsave{
		Unsave: &protos.Unsave{},
	}
	return m
}

// delete message
func (m *UpdateMessageBuilder) SetErase() *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_Erase{
		Erase: &protos.Erase{},
	}
	return m
}

func (m *UpdateMessageBuilder) SetSaveToCameraRoll() *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_SaveToCameraRoll{
		SaveToCameraRoll: &protos.SaveToCameraRoll{},
	}
	return m
}

func (m *UpdateMessageBuilder) SetScreenshot() *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_Screenshot{
		Screenshot: &protos.Screenshot{},
	}
	return m
}

func (m *UpdateMessageBuilder) SetScreenRecord() *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_ScreenRecord{
		ScreenRecord: &protos.ScreenRecord{},
	}
	return m
}

func (m *UpdateMessageBuilder) SetReplay(currentReplayCount int32) *UpdateMessageBuilder {
	m.updateAction.Update = &protos.UpdateAction_Replay{
		Replay: &protos.Replay{
			CurrentReplayCount: currentReplayCount,
		},
	}
	return m
}

func (m *UpdateMessageBuilder) build() ([]byte, error) {
	if m.updateAction.Update == nil {
		m.err = &snaperror.UpdateMessageBuilderError{ErrorMessage: "builder must have some action set before building"}
	}

	if m.updateAction.MessageId == 0 {
		m.err = &snaperror.UpdateMessageBuilderError{ErrorMessage: "messageId must be set before building"}
	}

	if m.err != nil {
		return nil, m.err
	}

	clientResolutionId, _ := crypto.GenerateBitUUID()
	nowTimestamp := methods.GetTimestamp()

	m.updateAction.UpdateTimestamp = nowTimestamp
	m.updateAction.ConversationId = m.conversationVersion.ConversationId
	
	updateMessage := &protos.UpdateContentMessageRequest{
		ClientResolutionId: clientResolutionId.Uint64(),
		CurrentVersion: m.conversationVersion.GetConversationVersion(),
		Update: m.updateAction,
	}

	return protos.EncodeProtoMessage(updateMessage)
}