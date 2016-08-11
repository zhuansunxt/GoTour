// fibonacci() function returns a function that will generate
// fibonacci number based on the order of invocation.
// The purpose of this exercise is to get familiar with golang's
// support for function closure.

package fibonacci

func fibonacci() func() int {
	invoked_time := 1
	p1 := 0
	p2 := 1

	f := func() int {
		if invoked_time == 1 {
			invoked_time = 2
			return p1
		}

		if invoked_time == 2 {
			invoked_time = 3
			return p2
		}

		
		ret := p1+p2
		p1 = p2
		p2 = ret
		return ret
	}

	return f
}
