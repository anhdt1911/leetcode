package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergeArr := append(nums1, nums2...)
	arrLen := len(mergeArr)
	if arrLen <= 1 {
		return float64(mergeArr[0])
	}

	mergeArr = quickSort(mergeArr)

	if arrLen%2 != 0 {
		return float64(mergeArr[(arrLen-1)/2])
	} else {
		num1 := float64(mergeArr[arrLen/2])
		num2 := float64(mergeArr[arrLen/2-1])
		return (num1 + num2) / 2
	}

}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}
