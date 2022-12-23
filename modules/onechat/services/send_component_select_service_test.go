package services

import (
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func TestOneChatService_SendComponentSelect(t *testing.T) {
	env.Options = configs.Options{
		OneChatAPIKey: "A0db10fcc4ac55676adef031ed6567f075907c88d76544044bfd2e9fe9d7831a976490f652f2b475bb70f6d68576c7e38",
	}

	oneID := "3066360402"
	botID := "B9f96ca1b641c5827a452e182da2273a4"

	choice := []one_entities.OneChatChoice{
		{
			Url:   "https://google.co.th/",
			Sign:  "true",
			Size:  "full",
			Type:  "link",
			Label: "ทดสอบปุ่ม",
		},
		{
			Url:   "https://google.co.th/",
			Sign:  "true",
			Size:  "full",
			Type:  "link",
			Label: "ทดสอบปุ่ม2",
		},
	}

	temp := []one_entities.OneChatRequestElementSendComponent{
		{
			Title:  "ทดสอบ",
			Detail: "detail test",
			Choice: choice,
			Image:  "https://cdn.vox-cdn.com/thumbor/lFQ3pGOjzP1XQ5puQms14ad7Wv8=/0x0:1400x788/920x613/filters:focal(588x282:812x506):format(webp)/cdn.vox-cdn.com/uploads/chorus_image/image/70412073/0377c76083423a1414e4001161e0cdffb0b36e1f_760x400.0.png",
		},
	}

	oneChatService := NewOneChatService()
	body, statusCode := oneChatService.SendComponentSelect("ทดสอบ", []string{oneID}, botID, temp)

	t.Errorf("%#v\n", body)

	t.Errorf("%#v", statusCode)

}
