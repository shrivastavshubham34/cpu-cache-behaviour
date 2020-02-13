package main
import "testing";

const var = 7           // for an element size of 64 Bytes
const size = 8 + (8 * N) // the struct size in bytes
type node struct {
    i [var]int64  // let us control the size of each element
    n *node      // the next element
}

func loop(el *node, b *testing.B) {
    b.ResetTimer()
	for it := 0; it < b.N; it++ {
	   // for N=0 : movq (CX), CX
	   // for N=7 : movq 56(CX), CX
	   el = el.n
	}
 }

 // compute the number of elements needed to reach the WSS
func computeLen(workingSetSize uint) int {
    return (2 << (workingSetSize - 1)) / S
}
// here we make a simple continuous array, and link all elements
// the resulting elements are layed out sequentially, densely packed
func makeContinuousArray(workingSetSize uint) *node {
    l := computeLen(workingSetSize)
    a := make([]node, l)
    for i := 0; i < l; i++ {
      a[i].n = &a[(i+1)%l]
    }
    return &a[0]
}

const PAGE_SIZE = 4096  // for an element size of 64 Bytes
func dispatchOnePerPage(workingSetSize uint) *node {
    l := computeLen(workingSetSize)
    // compute how many items fit in one page
    d := PAGE_SIZE / size
    // compute the number of items to allocate pages
    ls := d * l
    // allocate pages
    a := make([]node, ls)
    // link to the next element on the next page
    for i := 0; i < l; i++ {
        a[i*d].n = &a[((i+1)%l)*d]
    }
    return &a[0]
}

func BenchmarkCache10(b *testing.B)  { 
    loop(dispatchOnePerPage(10),b)
}

func BenchmarkCache12(b *testing.B)  { 
    loop(dispatchOnePerPage(12),b)
}

func BenchmarkCache14(b *testing.B)  { 
    loop(dispatchOnePerPage(14),b)
}

func BenchmarkCache15(b *testing.B)  { 
    loop(dispatchOnePerPage(15),b)
}

func BenchmarkCache16(b *testing.B)  { 
    loop(dispatchOnePerPage(16),b)
}

func BenchmarkCache17(b *testing.B)  { 
    loop(dispatchOnePerPage(17),b)
}

func BenchmarkCache18(b *testing.B)  { 
    loop(dispatchOnePerPage(18),b)
}

func BenchmarkCache19(b *testing.B)  { 
    loop(dispatchOnePerPage(19),b)
}

func BenchmarkCache20(b *testing.B)  { 
    loop(dispatchOnePerPage(20),b)
}

func BenchmarkCache22(b *testing.B)  { 
    loop(dispatchOnePerPage(22),b)
}

func BenchmarkCache24(b *testing.B)  { 
    loop(dispatchOnePerPage(24),b)
}

func BenchmarkCache26(b *testing.B)  { 
    loop(dispatchOnePerPage(26),b)
}
