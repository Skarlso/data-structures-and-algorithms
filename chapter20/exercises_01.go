package chapter20

import "fmt"

var basketballPlayers = []map[string]string{
	{
		"firstName": "Jill",
		"lastName":  "Huang",
		"team":      "Gators",
	},
	{
		"firstName": "Janko",
		"lastName":  "Barton",
		"team":      "Sharks",
	},
	{
		"firstName": "Wanda",
		"lastName":  "Vakulsak",
		"team":      "Sharks",
	},
	{
		"firstName": "Jill",
		"lastName":  "Moloney",
		"team":      "Sharks",
	},
	{
		"firstName": "Luuk",
		"lastName":  "Watkins",
		"team":      "Sharks",
	},
}

var footballPlayers = []map[string]string{
	{
		"firstName": "Hanzla",
		"lastName":  "Radosti",
		"team":      "32ers",
	},
	{
		"firstName": "Tine",
		"lastName":  "Watkins",
		"team":      "Barleycorns",
	},
	{
		"firstName": "Alex",
		"lastName":  "Patel",
		"team":      "32ers",
	},
	{
		"firstName": "Jill",
		"lastName":  "Huang",
		"team":      "Barleycorns",
	},
	{
		"firstName": "Wanda",
		"lastName":  "Vakulsak",
		"team":      "Barleycorns",
	},
}

func FindMultiSportsPlayers(sportOne, sportTwo []map[string]string) []string {
	// we just need the name
	namesHash := make(map[string]struct{})
	keyF := func(first, last string) string {
		return fmt.Sprintf("%s %s", first, last)
	}
	for _, n := range sportOne {
		namesHash[keyF(n["firstName"], n["lastName"])] = struct{}{}
	}

	var result []string
	for _, v := range sportTwo {
		key := keyF(v["firstName"], v["lastName"])
		if _, ok := namesHash[key]; ok {
			result = append(result, key)
		}
	}
	return result
}
