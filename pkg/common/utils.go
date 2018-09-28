package common

func ChunkSlice(strings []string, size int) [][]string {
	var divided [][]string

	for i := 0; i < len(strings); i += size {
		end := i + size

		if end > len(strings) {
			end = len(strings)
		}

		divided = append(divided, strings[i:end])
	}

	return divided
}
