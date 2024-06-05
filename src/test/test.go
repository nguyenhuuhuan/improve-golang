package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type Node struct {
	value int64
	edge  []*Node
}

func (n *Node) addEdge(node *Node) {
	n.edge = append(n.edge, node)
}

func DFS(current *Node, mapDfs map[*Node]bool) {
	if current == nil {
		return
	}
	mapDfs[current] = true
	fmt.Printf("%d ", current.value)
	for _, item := range current.edge {
		if !mapDfs[item] {
			DFS(item, mapDfs)
		}
	}
}

type Queue struct {
	node []*Node
}

func (q *Queue) pushBack(node *Node) {
	q.node = append(q.node, node)
}

func (q *Queue) Head() *Node {
	if len(q.node) == 0 {
		return nil
	}
	return q.node[0]
}

func (q *Queue) Pop() *Node {
	if len(q.node) == 0 {
		return nil
	}
	q.node = append(q.node[1:len(q.node)])
	return nil
}

func BFS(node *Node) {
	queue := Queue{}
	mapCheck := map[*Node]bool{}
	queue.pushBack(node)
	mapCheck[node] = true
	for len(queue.node) > 0 {
		current := queue.Head()
		fmt.Println(queue.Head().value)

		for _, item := range current.edge {
			if !mapCheck[item] {
				queue.pushBack(item)
				mapCheck[item] = true
			}
		}
		queue.Pop()
	}
}

var mapStr = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

func maxVowels(s string, k int) int {
	str := strings.Split(s, "")
	max := 0
	res := 0
	for i := 0; i < k; i++ {
		if ok := mapStr[str[i]]; ok {
			max++
		}
	}

	for i := 1; i < len(str)-k+1; i++ {
		l := i
		r := k + i - 1

		if ok := mapStr[str[r]]; ok {
			max++
		}
		if ok := mapStr[str[l-1]]; ok {
			max--
		}
		res = int(math.Max(float64(res), float64(max)))
	}
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	solution(candidates, []int{}, target, &result, 0)

	return result
}

func solution(candidates, current []int, target int, result *[][]int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*result = append(*result, append([]int{}, current...))
		return
	}

	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		solution(candidates, current, target-candidates[i], result, i)
		current = current[:len(current)-1]
	}

}

func main() {
	head := &ListNode{Val: 1}
	head.addNote(&ListNode{Val: 2})
	head.addNote(&ListNode{Val: 3})
	head.addNote(&ListNode{Val: 4})
	head.addNote(&ListNode{Val: 5})

	//fmt.Println(oddEvenList(head))
	//head.Next = &ListNode{Val: 1}
	//head.Next.Next = &ListNode{Val: 2}
	//rotateRight(head, 4)
	//i := 1
	//Try(i)

	fmt.Println(combine(3, 7))

}

func combine(k int, n int) [][]int {
	var (
		result *[][]int
	)
	backtracking(k-1, n, result, make([]int, k), 1)

	return *result
}
func backtracking(k int, n int, result *[][]int, res []int, start int) {
	if k < 0 {
		if n == 0 {
			*result = append(*result, append([]int{}, res...))
		}
	}
	for i := start; i < 10 && n >= i; i++ {
		res[k] = i
		backtracking(k-1, n-i, result, res, i+1)
	}
}
func compress(chars []byte) int {
	var mapCheck = make(map[byte]int, len(chars))
	for _, val := range chars {
		if _, ok := mapCheck[val]; ok {
			mapCheck[val]++
		} else {
			mapCheck[val] = 0
		}
	}
	return len(mapCheck) * 2
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	l := root
	for l.Right != nil && l.Left != nil {
		if l.Val == val {
			return l
		}
		if l.Val > val {
			l = l.Right
		} else {
			l = l.Left
		}
	}
	return nil
}
func a6(nums []int) int {
	var (
		max int
		sum int
	)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		max = int(math.Max(float64(max), float64(sum)))
	}
	return max
}

func a7(nums []int) int {
	var sum, sum1, sum2 int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum-nums[0] == 0 {
		return 0
	}

	for i := 0; i < len(nums)-1; i++ {
		sum1 += nums[i]
		sum2 = sum - sum1 - nums[i+1]
		if sum1 == sum2 {
			return i + 1
		}
	}
	return -1
}

func (l *ListNode) insert(node *ListNode, k int) {
	temp := l
	for i := 0; i < k-2; i++ {
		temp = temp.Next
	}
	node.Next = temp.Next
	temp.Next = node
}

func findMaxAverage1(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := 0; i < len(nums)-k; i++ {
		sum = sum - nums[i] + nums[i+k]
		maxSum = int(math.Max(float64(maxSum), float64(sum)))
	}

	return float64(maxSum) / float64(k)
}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	max := float64(math.MinInt)
	var l = 0
	for i := 0; i < len(nums); i++ {
		l = i
		if len(nums)-i < k {
			continue
		}
		var total float64 = 0

		for l < k+i {
			total += float64(nums[l])
			l++
		}
		max = math.Max(max, total/float64(k))

	}
	return max
}

func printMapKeys(key string, value interface{}, dataMap map[string]interface{}) {
	if value != nil {

	}
	if reflect.TypeOf(value).Kind() == reflect.Map {
		for k, v := range value.(map[string]interface{}) {
			newKey := fmt.Sprintf("%s.%s", key, k)
			dataMap[newKey] = v
			printMapKeys(newKey, v, dataMap)
		}
	} else if reflect.TypeOf(value).Kind() == reflect.Slice {
		for i, v := range value.([]interface{}) {
			newKey := fmt.Sprintf("%s[%d]", key, i)
			dataMap[newKey] = v
			printMapKeys(newKey, v, dataMap)
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var X = &[4]int{}

func inkq() {

	for i := 0; i < 3; i++ {
		fmt.Print(X[i])
	}
	fmt.Println()
}
func Try(index int) {
	for i := 0; i <= 1; i++ {
		X[index] = i
		if index == 3 {
			inkq()
		} else {
			Try(index + 1)
		}
	}
}

func (l *ListNode) addNote(node *ListNode) {
	tail := l
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
}

func oddEvenList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head.Next}

	even := head.Next
	odd := head
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = dummy.Next

	return head
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	count, tail, sp := 0, head, head

	for tail.Next != nil {
		tail = tail.Next
		count++
	}
	for i := 1; i < count-k%count; i++ {
		sp = sp.Next
	}
	tail.Next = head
	res := sp.Next
	sp.Next = nil

	return res
}
