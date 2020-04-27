package glist
//
// // This is a very simple implementation, see the source code to better understand the role
// func Run(n int, fn func(i int) (_stop bool) ) {
// 	for i:=0; i<n; i++ {
// 		if fn(i) {
// 			break
// 		}
// 	}
// }
//
// func Find(sliceLen int, finder func(i int) (find bool) ) (firstMatchIndex int, found bool) {
// 	firstMatchIndex = -1
// 	for i:=0; i<sliceLen; i++ {
// 		find := finder(i)
// 		if find {
// 			firstMatchIndex = i
// 			found = true
// 			return
// 		}
// 	}
// 	return
// }
// func CheckAll(sliceLen int, checker func(i int) (pass bool) ) (allPass bool) {
// 	// 如果 n 是 0 则通过
// 	allPass = true
// 	for i:=0; i<sliceLen; i++ {
// 		pass := checker(i)
// 		if pass {
// 			continue
// 		} else {
// 			allPass = false
// 			return
// 		}
// 	}
// 	return
// }
