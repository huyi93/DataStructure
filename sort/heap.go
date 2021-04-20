package sort

//堆排序
func HeapSort(s []int) []int {
	//buildMaxHeap(s, len(s)-1)
	for i := len(s) - 1; i >= 0; i-- {
		buildMaxHeap(s, i)
		s[0], s[i] = s[i], s[0]

	}
	return s
}

func buildMaxHeap(s []int, length int) {
	for i := length / 2; i >= 0; i-- {
		heapAdjust(s, i, length)
	}
}

func heapAdjust(s []int, key int, length int) {
	temp := s[key]
	for i := 2 * key; i <= length; i *= 2 {
		if i <= length-1 && s[i+1] > s[i] {
			i++
		}
		if temp >= s[i] {
			break
		} else {
			s[key] = s[i]
			key = i
		}
	}
	s[key] = temp
}

func HeapSort2(a []int) []int {
	// 1、将无序数组a构建为一个大根堆
	for i := len(a)/2 - 1; i >= 0; i-- {
		sink(a, i, len(a))
	}
	// 2、交换a[0]和a[len(a)-1]
	// 3、然后把前面这段数组继续下沉保持堆结构，如此循环即可
	for i := len(a) - 1; i >= 1; i-- {
		// 从后往前填充值
		swap(a, 0, i)
		// 前面的长度也减一
		sink(a, 0, i)
	}
	return a
}
func sink(a []int, i int, length int) {
	for {
		// 左节点索引(从0开始，所以左节点为i*2+1)
		l := i*2 + 1
		// 有节点索引
		r := i*2 + 2
		// idx保存根、左、右三者之间较大值的索引
		idx := i
		// 存在左节点，左节点值较大，则取左节点
		if l < length && a[l] > a[idx] {
			idx = l
		}
		// 存在有节点，且值较大，取右节点
		if r < length && a[r] > a[idx] {
			idx = r
		}
		// 如果根节点较大，则不用下沉
		if idx == i {
			break
		}
		// 如果根节点较小，则交换值，并继续下沉
		swap(a, i, idx)
		// 继续下沉idx节点
		i = idx
	}
}
func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
