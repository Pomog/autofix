package functions

import (
	"regexp"
	"strings"
)

const (
	articleA  = "a"
	articleAn = "an"
)

//Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h

/*
ArticlesCorrection corrects indefinite articles in the input string.
It replaces "a" or "A" with "an" or "An" in the input string respectively before certain words
if the next word begins with a vowel (a, e, i, o, u) or a h.
returns the corrected string.
*/
func ArticlesCorrection(input string) string {
	listOfLetters := []string{"a", "e", "i", "o", "u", "h"} // TODO: add configurable punctuation, by config file or env var
	lettersPattern := strings.Join(listOfLetters, "")
	strPattern := `\b(` + articleA + `|` + strings.ToUpper(articleA) + `)(\s+)([` + lettersPattern + `])` // shoud be three groups, where the first element is the article, the second is the space, and the third is the  frist letter of the next word

	re := regexp.MustCompile(strPattern)

	return re.ReplaceAllStringFunc(input, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		return replaceArticles(match, submatches)
	})
}

/*
processes the submatches and prefix to determine the corrected output.
submatches []string: the submatches from the regular expression
shoud be of length 4, where the first element is the article,
the second is the space, and the third is the  frist letter of the next word
replaces the articles in the based on the prefix and submatches.
*/
func replaceArticles(match string, submatches []string) string {
	output := submatches[1] + submatches[2] + submatches[3]
	if strings.HasPrefix(match, articleA+" ") {
		output = "an" + submatches[2] + submatches[3]
	}
	if strings.HasPrefix(match, strings.ToUpper(articleA)+" ") {
		output = "An" + submatches[2] + submatches[3]
	}
	return output
}
