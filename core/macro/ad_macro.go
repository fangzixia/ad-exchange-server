package macro

import "strings"

const MAC = "__I__MAC__"

var platformReplacers = map[string]*strings.Replacer{}

func RegisterPlatformMacro() {

}

func CreateReplacer(replaceMap map[string]string) *strings.Replacer {
	replacerArgs := make([]string, 0, len(replaceMap)*2)
	for oldVal, newVal := range replaceMap {
		replacerArgs = append(replacerArgs, oldVal, newVal)
	}
	replacer := strings.NewReplacer(replacerArgs...)
	return replacer
}

func ReplacePlatformMacro(replaceMap map[string]string, url string) string {
	if url == "" || len(replaceMap) == 0 {
		return url
	}
	return CreateReplacer(replaceMap).Replace(url)
}

func ReplacePlatformMacros(replaceMap map[string]string, urls []string) []string {
	if urls == nil || len(urls) == 0 {
		return urls
	}
	newUrls := make([]string, len(urls))
	for i, url := range urls {
		newUrls[i] = ReplacePlatformMacro(replaceMap, url)
	}
	return newUrls
}

func ReplaceMediaMacro(replaceMap map[string]string, url string) string {
	if url == "" || len(replaceMap) == 0 {
		return url
	}
	return CreateReplacer(replaceMap).Replace(url)
}

func ReplaceMediaMacros(replaceMap map[string]string, urls []string) []string {
	if urls == nil || len(urls) == 0 {
		return urls
	}
	newUrls := make([]string, len(urls))
	for i, url := range urls {
		newUrls[i] = ReplaceMediaMacro(replaceMap, url)
	}
	return newUrls
}
