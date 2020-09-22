package querystacks

import (
  "strings"
  "fmt"
  "log"

  "golords/state"
  "golords/handlers"

  "github.com/bwmarrin/discordgo"
)

func init(){
  handlers.RegisterActiveModule(
    Stack{},
  )
}

type Stack struct{}

func (h Stack) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)

  var out string
  var err error
  switch data[0] {
  case "topstacks":
    // TODO specify how many with command
    results, err = state.GetTopNStacks(5)
    for user, amount := range results {
      out += fmt.Sprintf("%v: %d stacks\n", user, amount)
    }
  // TODO !mystacks; a health workaround with !stacks @me works at the moment
  case "stacks":
    // TODO this only really needs one query
    // SELECT * FROM _ WHERE --- OR --- OR --- OR --- OR
    for _, user := range m.Mentions {
      amount, err := state.GetStacksForUser(m.GuildID, user.String())
      out += fmt.Sprintf("%v: %d stacks\n", user.String(), amount)
    }
    results, err = state.GetStacksForUser(m.GuildID, m.Mentions[0].String())
    out = m.Mentions[0].String()
  default:
    err = fmt.Errorf("Bad command: %v", data[0])
  }

  // Mongo machine broke
  if err != nil {
    log.Printf("Error in query: %v", err)
    return
  }

  if out == "" {
    log.Println("No output for stack query")
    return
  }

  s.ChannelMessageSend(m.ChannelID, out)
}

func (h Stack) Help() string {
  return "Find out how many stacks someone has.\n!topstacks: top 5\n!stacks @someone: someone's stacks (can be you)"
}

func (h Stack) Prefixes() []string {
  return []string{"topstacks", "stacks"}
}
