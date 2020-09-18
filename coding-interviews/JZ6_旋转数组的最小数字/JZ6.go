package main

/**
 * 
 * @param rotateArray int整型一维数组 
 * @return int整型
*/
func minNumberInRotateArray( rotateArray []int ) int {
    //双指针，二分法
    left,right := 0,len(rotateArray)-1
    for ;left < right; {
        mid := (left + right) / 2
        if (rotateArray[mid] > rotateArray[right]) {
            left = mid + 1
        } else if (rotateArray[mid] < rotateArray[right]) {
            right = mid
        } else {
            right = right - 1
        }
    }
    return rotateArray[left]
}