package main

func main() {
}

func removeDuplicates(S string) string {
	xs := []byte(S)
	i := 0
	for {
		if i+1 == len(xs) || i >= len(xs) {
			break
		}
		if xs[i] == xs[i+1] {
			xs = append(xs[:i], xs[i+2:]...)
			i = 0
			continue
		}
		i++
	}

	return string(xs)
}
