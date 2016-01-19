package main

// you can also use imports, for example:
import "sort"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

  // const
  const SIZE = 50001

  // ints
  var rootID int = 0
  var maxID int = 0
  var diameters = make([]int, SIZE)
  var sortedDiameters = make([]int, SIZE)

  // bools
  var seen = make([]bool, SIZE)


  // Structs
  type node struct {
      id int
      edgeList []int
  }

  // other
  var nodeMap = make([]*node, SIZE)


  func (n *node) addEdge(e int) {
    n.edgeList = append(n.edgeList, e)
  }


func Solution(A []int, B []int, K int) int {
  var n int = len(A)

 for i := 0; i < n; i++ {
    if( nodeMap[A[i]] == nil ) {
      nodeMap[A[i]] = &node{A[i],[]int{}}
    }

    if( nodeMap[B[i]] == nil ) {
      nodeMap[B[i]] =  &node{B[i],[]int{}}
    }

    var na *node = nodeMap[A[i]]
    var nb *node = nodeMap[B[i]]

    na.addEdge(B[i])
    nb.addEdge(A[i])

    maxID = int(math.Max(float64(maxID), float64(A[i])))
    maxID = int(math.Max(float64(maxID), float64(B[i])))
 }

  rootID = A[0]
  var res int = math.MaxInt64
  var low int = 0
  var high = int(math.Min(float64(900), float64(maxID)))

  for low <= high {
    var mid int = (low + high) / 2

    if(isAvailable(K, mid)) {
      res = int(math.Min(float64(res), float64(mid)))
      high = mid - 1
    } else {
      low = mid + 1
    }
  }

  return res
}


func isAvailable( camLimit int, notCovered int) bool {

  for i := 0; i < len(seen); i++{
    seen[i] = false
  }

  for i := 0; i < len(diameters); i++{
    diameters[i] = -1
  }

  if(dfs(nodeMap[rootID], notCovered) > camLimit){
    return false
  }

  return true

}

func dfs(nd *node, limit int) int {
  if( nd == nil || seen[nd.id]) {
      return 0
  }

  seen[nd.id] = true

  var counter int = 0

  for i := 0; i < len(nd.edgeList); i++ {
    counter += dfs(nodeMap[nd.edgeList[i]], limit)
  }

  var n int = 0

  for i := 0; i < len(nd.edgeList); i++ {
    sortedDiameters[n] = diameters[nd.edgeList[i]]
    n++
 }

  sort.Ints(sortedDiameters[:n])

  for i := 0; i < n / 2; i++ {
    var temp int = sortedDiameters[i]
    sortedDiameters[i] = sortedDiameters[n - 1 - i]
    sortedDiameters[n - 1 - i] = temp
  }

  var maxDiamRem int = -1

  for i := 0; i < n - 1; i++ {
    if (sortedDiameters[i] + sortedDiameters[i + 1] + 2 > limit) {
      counter++
    } else {
      maxDiamRem =  int(math.Max(float64(maxDiamRem), float64(sortedDiameters[i])))
      break
    }
  }

  if (n >= 1) {
    var i int = n - 1

    if(sortedDiameters[i] + 1 > limit) {
      counter++
    } else {
      maxDiamRem = int(math.Max(float64(maxDiamRem), float64(sortedDiameters[i])))
    }
  }

  if (maxDiamRem == -1) {
    diameters[nd.id] = 0
  } else {
    diameters[nd.id] = maxDiamRem + 1
  }

  return counter

}
