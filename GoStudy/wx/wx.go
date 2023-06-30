package wx

type wx struct {
	Name string
	Age  int
}

func GetWx() *wx {
	return &wx{
		Name: "wx",
		Age:  10,
	}
}
