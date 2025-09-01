package parsinglogfiles

import (
    "fmt"
    "regexp"
)

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^\[(?:TRC|DBG|INF|WRN|ERR|FTL)\]\s.*`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
    re := regexp.MustCompile(`<[-=*~]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    cnt := 0
    re := regexp.MustCompile(`"[^"]*(?i)password[^"]*"`)
    for _, line := range lines {
        matches := re.FindAllString(line, -1)
        cnt += len(matches)
    }
    return cnt
}

func RemoveEndOfLineText(text string) string {
    re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`(?i)\bUser\s+([A-Za-z0-9]{6,})\b`)
    result := []string{}

    for _, line := range lines {
        matches := re.FindStringSubmatch(line)
        if len(matches) > 1 {
            username := matches[1]
            newLine := fmt.Sprintf("[USR] %s %s", username, line)
            result = append(result, newLine)
        } else {
            result = append(result, line)
        }
    }
    return result
}
