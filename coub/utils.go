package coub

import "strings"

func getExtFromUrl(url string) string {
	sl := strings.Split(url, ".")
	return sl[len(sl)-1]
}

func getUrlFromMetadata(fileVersion map[string]interface{}) string {
	for _, key := range []string{"higher", "high", "med"} {
		obj, ok := fileVersion[key]

		if ok {
			file := obj.(map[string]interface{})
			return file["url"].(string)
		}
	}
	return ""
}
