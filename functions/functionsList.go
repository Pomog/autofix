package functions

/*
ApplyAutoFixingFunctions applies a series of auto-fixing functions to the input string.
*/
func ApplyAutoFixingFunctions(str string) string {
	for _, fixingFunc := range getAutoFixingFunctions() {
		str = fixingFunc(str)
	}
	return str
}

/*
getAutoFixingFunctions returns a map of auto-fixing functions keyed by their names.
*/
func getAutoFixingFunctions() map[string]StringModificationFunc {
	return map[string]StringModificationFunc{
		"hex":     ReplaceHexWithDecimal,
		"bin":     ReplaceBinWithDecimal,
		"up":      ToUppercase,
		"low":     ToLowercase,
		"cap":     Capitalization,
		"capn":    CapitalizationWithNumber,
		"upn":     ToUppercaseWithNumber,
		"lown":    ToLowercaseWithNumber,
		"punct":   CorrectPunctuationsSpaces,
		"apos":    CorrectApostrophesSpaces,
		"article": ArticlesCorrection,
	}
}
