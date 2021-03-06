package dndlookup

import (
	// "strings"
	"fmt"
	"log"

	go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doAbilityScore(query string) string {
	searchResults, err := go5e.SearchAbilityScoreByName(query)
	if err != nil || searchResults.Count < 1 {
		log.Println(err)
		return ""
	}
	spellIndex := getBestMatch(query, searchResults)
	spell, err := go5e.GetAbilityScore(spellIndex)
	if err != nil {
		log.Println(err)
		return ""
	}
	return formatAbilityScore(spell)
}

func formatAbilityScore(res go5e.AbilityScore) string {
	formatString := "%s\n%s"

	return fmt.Sprintf(formatString,
		res.FullName,
		res.Desc[0],
	)
}
