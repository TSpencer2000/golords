package plusplus

import (
  "strings"
  "fmt"

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
  // Must mention someone to invoke
  if len(m.Mentions) == 0 {
    return
  }

  inc := strings.Contains(m.Content, "++")
  dec := strings.Contains(m.Content, "--")

  // Do not allow increment and decrement in same operation
  // for simplicity - we will need to do a lot of parsing to
  // add syntax like @someone ++ @someonelse --

  // NOT XOR lol
  if inc == dec {
    return
  }

  outStr := ""

  if inc {
    for _, user := range m.Mentions {
      if user.ID == m.Author.ID {
        // Nice try, buckwheat
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        score, _ := pp.MinusMinus(user.String())
        outStr = outStr + fmt.Sprintf("%v has lost 10 stacks for being a sleazy d-bag. They now have %d.\n", user.String(), score)
        continue
      }
      score, err := pp.PlusPlus(user.String())
      if err != nil {
        // Mongo machine broke
        return
      }
      outStr = outStr + fmt.Sprintf("%v now has %d stacks\n", user.String(), score)
    }
  }

  if dec {
    for _, user := range m.Mentions {
      if user.ID == m.Author.ID {
        // Nice try, buckwheat
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        pp.MinusMinus(user.String())
        score, _ := pp.MinusMinus(user.String())
        outStr = outStr + fmt.Sprintf("%v has lost 10 stacks for being a sleazy d-bag. They now have %d.\n", user.String(), score)
        continue
      }
      score, err := pp.MinusMinus(user.String())
      if err != nil {
        // Mongo machine broke
        return
      }
      outStr = outStr + fmt.Sprintf("%v now has %d stacks\n", user.String(), score)
    }
  }

  // Send what is hopefully not an enormous message
  s.ChannelMessageSend(m.ChannelID, outStr)
}

func (h PlusHandler) GetPrompts() []string {
  return []string{"<none>"}
}

func (h PlusHandler) Help() string {
  return "Karma"
}

func (h PlusHandler) Should(hint string) bool {
  // Always call Do() from the handler
  return true
}
