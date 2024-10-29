package ui

func languageToFlag(languageCode string) string {
	langFlags := map[string]string{
		"ar":    "🇸🇦",
		"de":    "🇩🇪",
		"el":    "🇬🇷",
		"en":    "🇺🇸",
		"es-la": "🇪🇸",
		"es":    "🇪🇸",
		"fa":    "🇮🇷",
		"fr":    "🇫🇷",
		"hu":    "🇭🇺",
		"id":    "🇮🇩",
		"it":    "🇮🇹",
		"ja":    "🇯🇵",
		"pl":    "🇵🇱",
		"pt-br": "🇧🇷",
		"ro":    "🇷🇴",
		"ru":    "🇷🇺",
		"tr":    "🇹🇷",
		"uk":    "🇬🇧",
		"zh-hk": "🇭🇰",
	}

	flag, ok := langFlags[languageCode]
	if !ok {
		return languageCode
	}

	return flag
}
