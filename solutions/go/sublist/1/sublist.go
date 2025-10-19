package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	// Check if lists are equal
	if isEqual(l1, l2) {
		return RelationEqual
	}
	
	// Check if l1 is a sublist of l2
	if isSublist(l1, l2) {
		return RelationSublist
	}
	
	// Check if l2 is a sublist of l1 (making l1 a superlist)
	if isSublist(l2, l1) {
		return RelationSuperlist
	}
	
	// No relationship found
	return RelationUnequal
}

// isEqual checks if two slices are equal
func isEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

// isSublist checks if needle is a contiguous sublist of haystack
func isSublist(needle, haystack []int) bool {
	// Empty list is a sublist of any list
	if len(needle) == 0 {
		return true
	}
	
	// If needle is longer than haystack, it can't be a sublist
	if len(needle) > len(haystack) {
		return false
	}
	
	// Check each possible starting position in haystack
	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	
	return false
}