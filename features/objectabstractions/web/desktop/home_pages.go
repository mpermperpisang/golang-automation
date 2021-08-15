package pages

import (
	"github.com/golang-automation/features/helper"
	desktopmessages "github.com/golang-automation/features/helper/messages/web/desktop"
	webaction "github.com/golang-automation/features/helper/web/actions"
)

type HomePage struct {
	Page webaction.Page
}

var (
	followButton    = "Follow"
	subscribeButton = "(//*[@stroke-linecap='round'])[1]"
	followersLink   = "Followers"
	aboutLink       = "About"
	loginWarning    = "Never miss a story from"
)

func (s HomePage) ClickButton(button string) *HomePage {
	switch button {
	case "Follow":
		s.Page.ClickByText(followButton)
	case "Subscribe":
		s.Page.ClickByXpath(subscribeButton)
	default:
		helper.LogPanicln(desktopmessages.NotExistButton(button))
	}

	return &HomePage{Page: s.Page}
}

func (s HomePage) ClickLink(link string) *HomePage {
	var element string

	switch link {
	case "Followers":
		element = followersLink
	case "About":
		element = aboutLink
	default:
		helper.LogPanicln(desktopmessages.NotExistLink(link))
	}

	s.Page.ClickByText(element)

	return &HomePage{Page: s.Page}
}

func (s HomePage) ValidateWarning() *HomePage {
	s.Page.IsElementDisplayedByContainsText(loginWarning)

	return &HomePage{Page: s.Page}
}
