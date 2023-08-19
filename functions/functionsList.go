package functions

func ApplyAutoFixingFunctions(str string) string {
	for _, fixingFunc := range getAutoFixingFunctions() {
		str = fixingFunc(str)
	}
	return str
}

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
