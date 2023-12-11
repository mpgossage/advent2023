package main

import "fmt"

func reduce_sequence(seq []int) [][]int  {
	var result[][]int
	result = append(result, seq)

	all_zero := false
	for all_zero != true {
		all_zero = true
		var delta []int
		for i:=1;i<len(seq);i++{
			val := seq[i]-seq[i-1]
			delta = append(delta, val)
			if val != 0 {
				all_zero = false
			}
		}
		result = append(result, delta)
		seq =delta
	}
	return result
}

func extrapolate_sequence(seqs [][]int) [][]int  {
	for i:=len(seqs) -2; i>=0 ;i-- {
		seqs[i] = append(seqs[i], LastInt(seqs[i])+LastInt(seqs[i+1]))
		//fmt.Printf("expand %d %v\n", i, reduced)
	}
	return seqs
}

func extrapolate_sequence_prepend(seqs [][]int) [][]int  {
	for i:=len(seqs) -2; i>=0 ;i-- {
		seqs[i] = append([]int{seqs[i][0]-seqs[i+1][0]}, seqs[i]...)
		//fmt.Printf("expand %d %v\n", i, reduced)
	}
	return seqs
}


func day09a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	total := 0
	for _, line := range lines {
		reduced := reduce_sequence(StringToIntArray(line, " "))
		fmt.Printf("reduced %+v\n", reduced)
		extrapolated := extrapolate_sequence(reduced)
		fmt.Printf("extrapolated %+v\n", extrapolated)

		total += LastInt(extrapolated[0])
	}
	fmt.Printf("total %d\n", total)
}

func day09b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	total := 0
	for _, line := range lines {
		reduced := reduce_sequence(StringToIntArray(line, " "))
		fmt.Printf("reduced %+v\n", reduced)
		extrapolated := extrapolate_sequence_prepend(reduced)
		fmt.Printf("extrapolated %+v\n", extrapolated)

		total += extrapolated[0][0]
	}
	fmt.Printf("total %d\n", total)
}