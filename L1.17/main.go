package main

import "fmt"

func binarySearchRec(arr []int, target int) int{
    return search(arr, target, 0, len(arr)-1)
} //Рекурсивный бинарный поиск

func search(arr []int, target int, low int, high int) int{
    if low > high{
        return -1
    }
    mid := (low+high)/2
    if arr[mid] == target{
        return mid
    } else if arr[mid] < target{
        return search(arr, target, mid+1, high)
    } else {
        return search(arr, target, low, mid-1)
    }
} //вспомогательная функция для рекурсивного поиска

func binarySearchIter(arr []int, target int) int{
    low, high := 0, len(arr)-1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}//Итеративный бинарный поиск

func main() {
    array := []int{1,2,3,4,5,6,7,8,9}
    fmt.Println(binarySearchIter(array, 5))
    fmt.Println(binarySearchRec(array, 5))
    fmt.Println(binarySearchIter(array, 10))
    fmt.Println(binarySearchRec(array, 10))
}

