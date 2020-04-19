package glist



// This is a very simple implementation, see the source code to better understand the role
func Run(n int, fn func(i int) (_break bool) ) {
	for i:=0; i<n; i++ {
		if fn(i) {
			break
		}
	}
}
