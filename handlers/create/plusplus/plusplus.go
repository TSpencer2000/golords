package plusplus

import (
  "strings"
  "log"

  pp "golords/plusplus"
  "golords/handlers/create/handler"

  "github.com/bwmarrin/discordgo"
)

func New() handler.CreateHandler {
  return PlusHandler{}
}

type PlusHandler struct {
  handler.DefaultHandler
}

func (h PlusHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  // Do() here is responsible for determining what needs to be done
  // It will be run for every message.
  if !strings.Contains(m.Content, "!db"){
    return
  }

  log.Println("Testing")

  pp.PlusPlus("test")
}

func (h PlusHandler) GetPrompts() []string {
  return []string{"<none>"}
}

func (h PlusHandler) Help() string {
  return "Karma"
}

func (h PlusHandler) Should(hint string) bool {
  return true
}
