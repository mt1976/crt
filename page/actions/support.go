package actions

func IsInActions(value *Action, list []*Action) bool {
	// loop through each action in the list
	for i := range list {
		// if the given action matches an action in the list, return true
		if list[i].Is(value) {
			return true
		}
	}
	// if no match was found, return false
	return false
}
