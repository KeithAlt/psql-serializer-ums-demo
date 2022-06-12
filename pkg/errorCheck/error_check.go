package errorCheck

// Lazy check for error checking
func Fatal(err error) {
	if err != nil {
		panic(err)
	}
}
