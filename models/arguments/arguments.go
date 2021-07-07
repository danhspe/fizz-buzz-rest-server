package arguments

type Arguments struct {
	int1  int
	int2  int
	limit int
	str1  string
	str2  string
}

func New(int1 int, int2 int, limit int, str1 string, str2 string) Arguments {
	return Arguments{int1, int2, limit, str1, str2}
}
