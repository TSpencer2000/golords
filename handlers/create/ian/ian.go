package ian

import (
  "strings"

  "golords/handlers/create/handler"

  "github.com/bwmarrin/discordgo"
)

/*
  TODO only enable this in my server.
*/

func New() handler.CreateHandler {
  return IanHandler{}
}

type IanHandler struct {
  handler.DefaultHandler
}

func (h IanHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  content := strings.ToLower(m.Content)

  buyWords := []string{"buy", "bought", "purchase", "get a new", "house", "pay"}

  shouldTrigger := false

  for _, word := range buyWords {
    shouldTrigger = shouldTrigger || strings.Contains(content, word)
  }

  shouldTrigger = shouldTrigger && strings.Contains(content, "ian")

  if !shouldTrigger {
    return
  }

  s.ChannelMessageSend(m.ChannelID, "<@208773246009475072>, don't buy that thing!!!!!")
}

func (h IanHandler) GetPrompts() []string {
  return []string{}
}

func (h IanHandler) Help() string {
  return ""
}

func (h IanHandler) Should(hint string) bool {
  // Always call Do() from the handler
  return true
}
