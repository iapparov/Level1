package main

import "fmt"


func quickSort(arr []int) []int { 
  if len(arr) <= 1{
    return arr
  }
  pivot := arr[len(arr)/2]
  var less []int
  var greater []int
  for i, elem := range arr{
    if i == len(arr)/2{
      continue 
    }
    if elem > pivot{
      greater = append(greater, elem)
    } else {
      less = append(less, elem)
    }
  }
  sorted := append(quickSort(less), pivot)
  sorted = append(sorted, quickSort(greater)...)
  return sorted
} //Минус этой реализации – создание новых слайсов, что сильно влияет на память

func quickSort2(arr []int) []int{
  copyArr := make([]int, len(arr))
  copy(copyArr, arr)
  sorter(copyArr, 0, len(copyArr)-1)
  return copyArr
} // по заданию нам нужно, чтобы функция quickSort2(arr []int) []int была такой

func sorter(arr []int, low int, high int) {
  if low < high{
    pivot := arr[high]
    i := low-1
    for j:=low;j<high;j++{
      if arr[j]<pivot{
        i++
        arr[i],arr[j] = arr[j],arr[i]
      }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1]
    p := i+1
    sorter(arr, low, p-1)
    sorter(arr, p+1, high)
  }
} // in-place сортировка без создания новых слайсов

func main() {
  arr1 := []int{10, 29, 38, 2342, 56, 1246, 6, 0, 12, 5654}
  arr2 := quickSort(arr1)
  fmt.Println(arr2)
  arr3 := quickSort2(arr1)
  fmt.Println(arr3)
}