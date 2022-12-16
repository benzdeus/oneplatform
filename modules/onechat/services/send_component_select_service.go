package services

import (
	one_entities "github.com/benzdeus/oneplatform/entities"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
)

func (service OneChatService) SendComponentSelect(title, oneID, botID string, templates []one_entities.OneChatRequestElementSendComponent) int {

	_, statusCode := helper.RunOneChatSendComponentHTTP("POST", title, env.BaseUrlOneChat+"/bc_msg/api/v1/broadcast_group_template", oneID, botID, templates)

	return statusCode
}