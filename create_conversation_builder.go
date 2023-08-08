package snapper

import (
	"github.com/0xzer/snapper/crypto"
	"github.com/0xzer/snapper/protos"
	"github.com/0xzer/snapper/snaperror"
	"github.com/google/uuid"
)


type CreateConversationBuilder struct {
	client *Client

	createRequest *protos.CreateConversationRequest

	err *snaperror.CreateConversationBuilderError
}

// THIS WILL AUTOMATICALLY ADD YOU INTO THE PARTICIPANTS SO YOU DO NOT HAVE TO DO THAT YOURSELF
func (c *Client) NewCreateConversationBuilder() *CreateConversationBuilder {
	meId := c.Session.CurrentUser.EncodedUserId
	createConversationAttemptId, _ := crypto.EncodeUUID(uuid.NewString())
	clientResolutionId, _ := crypto.GenerateBitUUID()
	return &CreateConversationBuilder{
		client: c,
		createRequest: &protos.CreateConversationRequest{
			SelfUserId: meId,
			ClientResolutionId: clientResolutionId.Uint64(),
			Participants: []*protos.UUID{
				meId,
			},
			CreateConversationAttemptId: &protos.UUID{EncodedId: createConversationAttemptId},
		},
	}
}

func (c *CreateConversationBuilder) SetTitle(title string) *CreateConversationBuilder {
	c.createRequest.ConversationTitle = title
	return c
}

func (c *CreateConversationBuilder) SetConversationType(conversationType protos.ConversationType) *CreateConversationBuilder {
	c.createRequest.Type = conversationType
	return c
}

func (c *CreateConversationBuilder) AddParticipant(participant *protos.UUID) *CreateConversationBuilder {
	hasParticipant := c.client.HasUUID(c.createRequest.Participants, participant)
	if !hasParticipant {
		c.createRequest.Participants = append(c.createRequest.Participants, participant)
	}
	return c
}

func (c *CreateConversationBuilder) AddParticipants(participants []*protos.UUID) *CreateConversationBuilder {
	c.createRequest.Participants = append(c.createRequest.Participants, participants...)
	return c
}

func (c *CreateConversationBuilder) build() ([]byte, error) {
	if len(c.createRequest.Participants) <= 1 {
		c.err = &snaperror.CreateConversationBuilderError{ErrorMessage: "participants must be at least one other person"}
	}

	return protos.EncodeProtoMessage(c.createRequest)
}