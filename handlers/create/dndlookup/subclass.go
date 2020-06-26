package dndlookup

import (
  // "strings"
  "log"
  // "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doSubclass(query string) string {
  searchResults, err := go5e.SearchSubclassByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetSubclass(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatSubclass(spell)
}

func formatSubclass(res go5e.Subclass) string {
  return res.Name
}
