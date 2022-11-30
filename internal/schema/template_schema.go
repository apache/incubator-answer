package schema

type Paginator struct {
	Pages      []int
	Totalpages int
	Prevpage   int
	Nextpage   int
	Currpage   int
}
