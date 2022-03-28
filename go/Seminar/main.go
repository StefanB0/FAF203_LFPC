package main

func F(a, b []int, n int) []int {
	solution := make([]int, n)
	foo:= 1

	for i:= 0; i < n; i++{
		b[i] = foo
		foo *= a[i]
	}

	
	return solution
}

func main(){
	a:= []int{1, 2, 3, 4}
	b:= []int{4, 3, 2, 1}

	final:= []int{24, 12, 8, 6}



}
