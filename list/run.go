package glist


type OP struct {
	isBreak bool
}
func Break() OP{return OP{isBreak:true}}
// This is a very simple implementation, see the source code to better understand the role
func Run(n int, fn func(i int) (_break OP) ) {
	for i:=0; i<n; i++ {
		if fn(i).isBreak {
			break
		}
	}
}
