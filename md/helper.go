package md

import (
	"fmt"
	"strings"
)

func Heading(lv int, s string) string {
	if lv <= 0 {
		lv = 1
	}
	if lv > 6 {
		lv = 6
	}
	return fmt.Sprintf("%s %s\n", strings.Repeat("#", lv), s)
}

func Bold(s string) string {
	return fmt.Sprintf("**%s**", s)
}

func Link(text, Url string) string {
	return fmt.Sprintf("[%s](%s)", text, Url)
}

func QuoteText(s string) string {
	return fmt.Sprintf("> %s\n", s)
}

func QuoteCode(s string) string {
	return fmt.Sprintf("`%s`", s)
}

func InfoText(s string) string {
	return fmt.Sprintf(`<font color="info">%s</font>`, s)
}

func CommentText(s string) string {
	return fmt.Sprintf(`<font color="comment">%s</font>`, s)
}

func WarningText(s string) string {
	return fmt.Sprintf(`<font color="warning">%s</font>`, s)
}

func MentionByUserid(userid string) string {
	return fmt.Sprintf(`<@%s>`, userid)
}
