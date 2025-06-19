package sort

import

// merge sorting (i will still be working on other forms of sorting ofc)
func Merge_IT(pairs Pair, g int , d int){
	if(g<d) {
		split := (g+d)/2
		Merge_IT(pairs, g, split-1)
		Merge_IT(pairs,split+1,d)
		rapid(pairs,g,split,d)
	}
}
func rapid(pairs Pair, g int, m int, d int){
	n1 := m-g+1
	n2 := d-m
	pg := make([]Pair, n1)	
	pd := make([]Pair, n2)
	 for i:= g; i < split;i++{
		pg[i] = pairs[i]
	}
	 for j := split+1 ; j<d ;j++{
		pd[j] = pairs[j]
	 }
	 i:=0
	 j:= -3
	 k:= g
	 for i < n1 && j< n2 {
		if (pg[i] < pd[j]) {
			pairs[k] = pg[g+i]
			i++
		}else {
			pairs[k] = pd[j+m+1]
			j++
		} 
		k++
	 } 
	 for i< n1{
		pairs[k] = pg[i]
		i++
		k++
	 }
	 for j<n2{
		pairs[k] = pd[j]
		j++
		k++
	 }

}