package morsecode

func GetCode(r rune) []int {
	switch r {
	case 'a':
		return []int{1, 3}
	case 'b':
		return []int{3, 1, 1, 1}
	case 'c':
		return []int{3, 1, 3, 1}
	case 'd':
		return []int{3, 1, 1}
	case 'e':
		return []int{1}
	case 'f':
		return []int{1, 1, 3, 1}
	case 'g':
		return []int{3, 3, 1}
	case 'h':
		return []int{1, 1, 1, 1}
	case 'i':
		return []int{1, 1}
	case 'j':
		return []int{1, 3, 3, 3}
	case 'k':
		return []int{3, 1, 3}
	case 'l':
		return []int{1, 3, 1, 1}
	case 'm':
		return []int{3, 3}
	case 'n':
		return []int{3, 1}
	case 'o':
		return []int{3, 3, 3}
	case 'p':
		return []int{1, 3, 3, 1}
	case 'q':
		return []int{3, 3, 1, 3}
	case 'r':
		return []int{1, 3, 1}
	case 's':
		return []int{1, 1, 1}
	case 't':
		return []int{3}
	case 'u':
		return []int{1, 1, 3}
	case 'v':
		return []int{1, 1, 1, 3}
	case 'w':
		return []int{1, 3, 3}
	case 'x':
		return []int{1, 3, 3, 1}
	case 'y':
		return []int{3, 1, 3, 3}
	case 'z':
		return []int{3, 3, 1, 1}
	case '0':
		return []int{3, 3, 3, 3, 3}
	case '1':
		return []int{1, 3, 3, 3, 3}
	case '2':
		return []int{1, 1, 3, 3, 3}
	case '3':
		return []int{1, 1, 1, 3, 3}
	case '4':
		return []int{1, 1, 1, 1, 3}
	case '5':
		return []int{1, 1, 1, 1, 1}
	case '6':
		return []int{3, 1, 1, 1, 1}
	case '7':
		return []int{3, 3, 1, 1, 1}
	case '8':
		return []int{3, 3, 3, 1, 1}
	case '9':
		return []int{3, 3, 3, 3, 1}
	case '/':
		return []int{3, 1, 1, 3, 1}
	case '.':
		return []int{1, 3, 1, 3, 1, 3}
	case '?':
		return []int{1, 1, 3, 3, 1, 1}
	case ' ':
		return []int{-4}
	}
	return []int{}
}
