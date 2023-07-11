package manager

func filter[T any](items []T, filterFn func(t *T) bool) []T {
	end := 0
	for _, item := range items {
		ok := filterFn(&item)
		if ok {
			items[end] = item
			end += 1
		}
	}

	return items[:end-1]
}
