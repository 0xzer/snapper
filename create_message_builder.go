package snapper

import (
	"github.com/0xzer/snapper/crypto"
	"github.com/0xzer/snapper/protos"
	"github.com/0xzer/snapper/snaperror"
	"github.com/google/uuid"
)

type destination int

const (
	Conversation destination = iota
	Story
	PhoneNumber
)

type CreateMessageBuilder struct {
	client *Client

	featureAttachment []*protos.FeatureAttachment
	contentEnvelope *protos.ContentEnvelope
	destinations []*protos.DeliveryDestination

	err *snaperror.CreateMessageBuilderError
	isRetry bool
}

func (c *Client) NewCreateMessageBuilder() *CreateMessageBuilder {
	return &CreateMessageBuilder{
		client: c,
		contentEnvelope: &protos.ContentEnvelope{},
		featureAttachment: make([]*protos.FeatureAttachment, 0),
		destinations: make([]*protos.DeliveryDestination, 0),
	}
}

func (c *CreateMessageBuilder) AddConversationDestination(conversation *protos.Conversation) *CreateMessageBuilder {
	dest := c.buildDestination(Conversation, conversation).(*protos.DeliveryDestination)
	c.destinations = append(c.destinations, dest)
	return c
}

func (c *CreateMessageBuilder) SetSavePolicy(policy protos.ContentEnvelope_SavePolicy) *CreateMessageBuilder {
	c.contentEnvelope.SavePolicy = policy
	return c
}

// messageId = the message id of the message to reply to
func (c *CreateMessageBuilder) SetReplyMessage(messageId int64) *CreateMessageBuilder {
	replyFeatureAttachment := &protos.FeatureAttachment{
		Attachment: &protos.FeatureAttachment_ReplyMessageInfo{
			ReplyMessageInfo: &protos.ReplyMessageInfo{
				QuotedMessageId: messageId,
			},
		},
	}
	c.featureAttachment = append(c.featureAttachment, replyFeatureAttachment)
	return c
}

func (c *CreateMessageBuilder) SetTextMessage(text string) *CreateMessageBuilder {
	content := &protos.Contents_Text{
		Text: &protos.Text{
			Text: text,
		},
	}

	err := c.buildContents(protos.ContentType_CHAT, content)
	if err != nil {
		c.err = &snaperror.CreateMessageBuilderError{ErrorMessage: "failed to encode textmessage content bytes"}
		return c
	}

	return c
}

func (c *CreateMessageBuilder) SetRetry() *CreateMessageBuilder {
	c.isRetry = true
	return c
}

func (c *CreateMessageBuilder) build() ([]byte, error) {

	if c.err != nil {
		return nil, c.err
	}

	if len(c.destinations) <= 0 {
		c.err = &snaperror.CreateMessageBuilderError{ErrorMessage: "provide at least 1 destination to send the message to"}
		return nil, c.err
	}

	sendMessageAttemptId, _ := crypto.EncodeUUID(uuid.NewString())
	clientResolutionId, _ := crypto.GenerateBitUUID()

	c.contentEnvelope.EnvelopeEncryption = &protos.EnvelopeEncryption{
		Method: &protos.EnvelopeEncryption_None{None: &protos.Empty{}},
	}
	
	createMessage := &protos.CreateContentMessageRequest{
		SenderId: c.client.Session.CurrentUser.EncodedUserId,
		ClientResolutionId: clientResolutionId.Uint64(),
		Destinations: c.destinations,
		Content: c.contentEnvelope,
		FeatureAttachment: c.featureAttachment,
		IsRetry: c.isRetry,
		CreateContentMessageBlizzardData: &protos.CreateContentMessageBlizzardData{
			SendMessageAttemptId: &protos.UUID{
				EncodedId: sendMessageAttemptId,
			},
		},
	}

	c.client.logger.Debug().Any("data", createMessage).Msg("createmessage payload")
	createMessagesBytes, err := protos.EncodeProtoMessage(createMessage)
	if err != nil {
		return nil, err
	}

	return createMessagesBytes, nil
}

func (c *CreateMessageBuilder) buildDestination(destType destination, data interface{}) interface{} {
	dest := &protos.DeliveryDestination{
		EncryptionInfo: &protos.EncryptionInfo{
			Method: &protos.EncryptionInfo_Fidelius{Fidelius: &protos.Empty{}},
		},
	}
	switch destType {
		case Conversation:
			d := data.(*protos.Conversation)
			dest.Destination = &protos.DeliveryDestination_ConversationDestination{
				ConversationDestination: &protos.ConversationDestination{
					ConversationId: d.GetConversationId(),
					CurrentVersion: d.GetVersion(),
				},
			}
	}

	return dest
}

func (c *CreateMessageBuilder) buildContents(contentType protos.ContentType, data interface{}) error {
	content := &protos.Contents{}
	switch contentType {
		case protos.ContentType_CHAT:
			textProto := data.(*protos.Contents_Text)
			content.Content = textProto
	}

	contentBytes, err := protos.EncodeProtoMessage(content)
	if err != nil {
		return err
	}

	c.contentEnvelope.ContentType = contentType
	c.contentEnvelope.Contents = contentBytes
	return nil
}