package dndlookup

import (
  // "strings"
  "log"
  // "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doTrait(query string) string {
  searchResults, err := go5e.SearchTraitByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetTrait(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatTrait(spell)
}

func formatTrait(res go5e.Trait) string {
  return res.Name
}
