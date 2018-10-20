package dedup

// Leading returns a slice with leading duplicates removed in.
// TODO: Could do this by deleting values from "in", though this could cause a
// lot of slice creation
// Note: LEADING EMPTY STRINGS WILL BE REMOVED
func Leading(in []string) []string {
	// create array of input size
	ordered := make([]string, len(in), len(in))

	// grab elements from the end first, i.e. len(in)-1-i
	var check checker
	last := len(in) - 1
	orderedIndex := last
	for i := range in {
		line := in[last-i]
		if check.IsDup(line) {
			continue
		}
		// when doing dedupe, insert from bottom (last index len-1) to top (first index 0)
		ordered[orderedIndex] = line
		orderedIndex--
	}

	// iterate over array, slice it at first non empty string basically removes
	// the deduped lines that were not inserted into ordered
	for i, v := range ordered {
		if v != "" {
			ordered = ordered[i:]
			break
		}
	}
	// we would only get here if ordered is all "" though that is fine
	return ordered
}

type checker struct {
	dups map[string]int
}

func (c *checker) IsDup(s string) bool {
	if c.dups == nil {
		c.dups = make(map[string]int)
	}

	if c.dups[s] == 0 {
		// set this so that it will be a duplicate in the future
		c.dups[s] = 1
		return false
	}

	// counting how often a line happens might be interesting to keep around later
	c.dups[s] = c.dups[s] + 1
	return true
}
