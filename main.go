package main

import (
	"fmt"
	"sort"
	"strings"

	"os"

	"github.com/fate-lovely/gofred"
)

func main() {
	keyword := strings.TrimSpace(os.Args[1])

	if keyword == "" {
		keyword = "joy"
	}

	result := search(keyword)

	sort.SliceStable(result, func(i, j int) bool {
		return getScore(keyword, result[i]) > getScore(keyword, result[j])
	})

	for _, kaomoji := range result {
		gofred.AddItem(&gofred.Item{
			Title:    kaomoji.kaomoji,
			Subtitle: kaomoji.name,
			Arg:      kaomoji.kaomoji,
			// hide icon
			Icon: &gofred.Icon{
				Path: " ",
			},
			Text: &gofred.Text{
				Copy:      kaomoji.kaomoji,
				Largetype: kaomoji.kaomoji,
			},
		})
	}

	json, _ := gofred.JSON()
	fmt.Print(json)
}

func search(keyword string) []*Kaomoji {
	result := make([]*Kaomoji, 0)

	for _, kaomoji := range kaomojis {
		if strings.Index(kaomoji.keywords, keyword) != -1 {
			result = append(result, kaomoji)
		}
	}

	return result
}

func getScore(keyword string, kaomoji *Kaomoji) int {
	if keyword == kaomoji.name {
		return 3
	}

	i := strings.Index(kaomoji.name, keyword)

	if i == 0 {
		return 2
	}

	if i > 0 {
		return 1
	}

	return 0
}
