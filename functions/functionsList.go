package functions

func GetAutoFixingFunctions() map[string]StringModificationFunc {
	return map[string]StringModificationFunc{
		"hex":   ReplaceHexWithDecimal,
		"bin":   ReplaceBinWithDecimal,
		"up":    ToUppercase,
		"low":   ToLowercase,
		"cap":   Capitalization,
		"capn":  CapitalizationWithNumber,
		"upn":   ToUppercaseWithNumber,
		"lown":  ToLowercaseWithNumber,
		"punct": CorrectPunctuationsSpaces,
		"apos":  CorrectApostrophesSpaces,
	}
}
