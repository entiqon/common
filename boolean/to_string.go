package boolean

// ToString returns trueStr if b is true, otherwise falseStr.
//
//	ToString(true, "enabled", "disabled") // "enabled"
//	ToString(false, "yes", "no")          // "no"
func ToString(b bool, trueStr, falseStr string) string {
	if b {
		return trueStr
	}

	return falseStr
}
