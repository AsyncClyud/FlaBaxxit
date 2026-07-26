package moderation

import (
	"strings"
)

var WordList_RU []string = []string{
	"сука",
	"суки",
	"сучка",
	"сучки",
	"сучара",
	"сучары",
	"бля",
	"баля",
	"быля",
	"балядоты",
	"блядь",
	"блять",
	"блядота",
	"блядота",
	"пидор",
	"пидоры",
	"пидорас",
	"пидорасы",
	"выебать",
	"выебали",
	"выебал",
	"выебала",
	"секс",
	"сексуально",
	"сексуальный",
	"сексуальная",
	"секси",
	"хохол",
	"хохлы",
	"заебал",
	"заебали",
	"чурка",
	"чурки",
	"хач",
	"хачи",
	"хуй",
	"хуйня",
	"хуеглот",
	"хуеглоты",
	"хуесос",
	"хуесосы",
}

var WordList_EN []string = []string{
	"2 girls 1 cup",
	"ass",
	"ass hole",
	"ass fuck",
	"ass-fucker",
	"anal",
	"anus",
	"bicth",
	"fuck",
	"fucking",
	"fucker",
	"nigger",
	"nigga",
	"niga",
	"niger",
	"cock",
	"cum",
	"cumming",
	"cumslut",
	"cunt",
	"cunts",
}

var leetReplacer = strings.NewReplacer(
	"0", "o",
	"3", "e",
	"4", "a",
	"1", "i",
	"@", "a",
	"$", "s",
	"!", "i",
)

func FormatText(text string) (formatted_text string) {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	text = leetReplacer.Replace(text)

	return text
}

func GetLanguage(text string) (tag string) {
	if contains := strings.ContainsAny(text, "йцукенгшщзхъфывапролджэячсмитьбю"); contains {
		return "RU"
	} else {
		return "EN"
	}
}

func ModerateText(text string) (contains bool) {
	text = FormatText(text)
	tag := GetLanguage(text)
	if tag == "RU" {
		for _, word := range WordList_RU {
			if ok := strings.Contains(text, word); ok {
				return true
			}
		}
	}
	if tag == "EN" {
		for _, word := range WordList_EN {
			if ok := strings.Contains(text, word); ok {
				return true
			}
		}
	}
	return false
}
