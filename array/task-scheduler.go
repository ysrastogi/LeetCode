type MaxHeap []int
func (h MaxHeap) Len() int{
    return len(h)
}
func (h MaxHeap) Less(i, j int) bool{
    return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int){
    h[i], h[j] = h[j], h[i]
}
func(h *MaxHeap) Push(x interface{}){
    *h = append(*h, x.(int))
}
func(h *MaxHeap) Pop() interface{}{
    old := *h
    n := len(old)
    val := old[n-1]
    *h = old[:n-1]
    return val
}

func leastInterval(tasks []byte, n int) int {
    freq := make(map[byte]int)
    for _, t := range tasks{
        freq[t]++
    }
    h := &MaxHeap{}
    heap.Init(h)
    for _, v:= range freq{
        heap.Push(h, v)
    }

    time :=0
    for h.Len()>0{
        k :=n+1
        temp := []int{}
        for k>0 && h.Len()>0 {
            top := heap.Pop(h).(int)
            top --
            if top >0 {
                temp = append(temp, top)
            }
            time ++
            k --
        }

        for _, v := range temp {
            heap.Push(h, v)
        }
        if h.Len()>0 {
            time += k
        }
    }
    return time
}