package heap

func sink(array []int, parentIndex int, length int) {
	//保存父节点，用于最后的赋值
	temp := array[parentIndex]
	//左子节点
	childIndex := 2*parentIndex + 1
	//是否有左子节点
	for childIndex < length {
		//判断是否有右子节点，并且右子节点大于左子节点的值
		if childIndex+1 < length && array[childIndex+1] > array[childIndex] {
			childIndex++
		}
		//如果父节点大于任何一个子节点的值直接跳出
		if temp >= array[childIndex] {
			break
		}
		array[parentIndex] = array[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	array[parentIndex] = temp
}

func heapSort(array []int) {
	//构建大顶堆
	for i := (len(array) - 2) / 2; i >= 0; i-- {
		sink(array, i, len(array))
	}
	//将堆顶元素和最后一个元素交换，数组长度i--(相当于循环删除根顶部元素，然后sink 调整最大堆)
	for i := len(array) - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		sink(array, 0, i)
	}
}
