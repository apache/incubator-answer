package schema

type Paginator struct {
	Pages      []int
	Totalpages int
	Firstpage  int
	Lastpage   int
	Currpage   int
}
