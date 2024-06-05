package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type People struct {
	Name  string    `json:"name"`
	Pos   int       `json:"pos"`
	Child []*People `json:"child"`
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	_ = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
				//Left: &TreeNode{
				//	Val: 1,
				//	Right: &TreeNode{
				//		Val: 1,
				//	},
				//},
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	_ = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	//rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	//print(ans)

	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()

	ans := isAnagrams("anagram", "nagaram")
	fmt.Println("Counter:", ans)

}
func isAnagrams(s string, t string) bool {
	strs := strings.Split(s, "")
	strt := strings.Split(t, "")
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	sort.Slice(strt, func(i, j int) bool {
		return strt[i] < strt[j]
	})
	if strings.Join(strs, "") == strings.Join(strt, "") {
		return true
	}
	return false
}
func combine(n int, k int) [][]int {
	var backtracks = func(res *[][]int, arr []int, n, k, pos int) {}
	backtracks = func(res *[][]int, arr []int, n, k, pos int) {
		if k == len(arr) {
			p := make([]int, len(arr))
			copy(p, arr)
			*res = append(*res, p)
			return
		}

		for i := pos; i <= n; i++ {
			arr = append(arr, i)
			backtracks(res, arr, n, k, i+1)
			arr = arr[:len(arr)-1]
		}
	}
	var res = &[][]int{}

	backtracks(res, []int{}, n, k, 1)
	return *res

}

func permute(nums []int) [][]int {
	var backtracks = func(nums []int, arr []int, res *[][]int, mapCheck map[int]struct{}) {}
	backtracks = func(nums []int, arr []int, res *[][]int, mapCheck map[int]struct{}) {
		if len(arr) == len(nums) {
			p := make([]int, len(arr))
			copy(p, arr)
			*res = append(*res, p)
			return
		}
		for i := 0; i < len(nums); i++ {
			if _, ok := mapCheck[nums[i]]; !ok {
				mapCheck[nums[i]] = struct{}{}
				arr = append(arr, nums[i])
				backtracks(nums, arr, res, mapCheck)
				delete(mapCheck, nums[i])
				arr = arr[:len(arr)-1]
			}
		}
	}

	var res = &[][]int{}
	var arr = []int{}
	backtracks(nums, arr, res, map[int]struct{}{})

	return *res
}
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	arr := []int{}
	for head != nil {

		arr = append(arr, head.Val)
		head = head.Next

	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	root := &ListNode{}
	res := root
	for i := 0; i < len(arr); i++ {
		root.Val = arr[i]
		if i == len(arr)-1 {
			break
		}
		root.Next = &ListNode{}
		root = root.Next
	}
	return res
}
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var sumNums = func(root *TreeNode, res *int, arr []int) {}
	sumNums = func(root *TreeNode, res *int, arr []int) {
		arr = append(arr, root.Val)
		if root.Left != nil {
			sumNums(root.Left, res, arr)
		}
		if root.Right != nil {

			sumNums(root.Right, res, arr)

		}
		if root.Right == nil && root.Left == nil {
			sum := 0
			if len(arr) != 1 {
				for i := 0; i < len(arr); i++ {
					sum += int(math.Pow10(len(arr)-i-1) * float64(arr[i]))
				}
			}
			*res += sum
		}

	}
	res := new(int)
	var arr = []int{}
	sumNums(root, res, arr)
	return *res
}

func longestConsecutive(nums []int) int {
	mapCheck := make(map[int]bool)
	for _, val := range nums {
		mapCheck[val] = true
	}
	ans := 0
	for _, n := range nums {
		if mapCheck[n+1] {
			continue
		}
		dem := 0
		for curr := n; mapCheck[curr]; curr-- {
			dem++
		}
		if dem > ans {
			ans = dem
		}
	}
	return ans
}
func containsNearbyDuplicate(nums []int, k int) bool {
	var mapCheck = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mapCheck[nums[i]]; ok {
			if int(math.Abs(float64(mapCheck[nums[i]]-i))) <= k {
				return true
			}
			mapCheck[nums[i]] = i
		} else {
			mapCheck[nums[i]] = i
		}
	}
	return false
}

func groupAnagrams(strs []string) [][]string {
	mapCheck := make(map[string][]string)
	var res [][]string
	for i := 0; i < len(strs); i++ {
		str := sortString(strs[i])
		if _, ok := mapCheck[str]; ok {
			mapCheck[str] = append(mapCheck[str], strs[i])
		} else {
			mapCheck[str] = []string{strs[i]}
		}
	}
	for _, val := range mapCheck {
		res = append(res, val)
	}
	return res
}
func sortString(s string) string {
	characters := []rune(s)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	out := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			if out == 0 || (i-start+1) < out {
				out = i - start + 1
			}
			sum -= nums[start]
			start++
		}
	}

	return out
}
func wordPattern(pattern string, s string) bool {
	str := strings.Split(s, " ")
	mapP := map[string]string{}

	for i, v := range str {
		if val, ok := mapP[string(pattern[i])]; ok && val != v {
			return false
		} else {
			mapP[string(pattern[i])] = v
			dem := 0
			for _, value := range mapP {
				if value == v {
					dem++
					if dem > 1 {
						return false
						//delete(mapP, string(pattern[i]))
					}
				}

			}

		}
	}
	return true

}
func canConstruct(ransomNote string, magazine string) bool {

	var mapCheck = make(map[string]int, len(magazine))
	for _, v := range magazine {
		mapCheck[string(v)]++
	}
	for _, v := range ransomNote {
		if c, ok := mapCheck[string(v)]; !ok || c == 0 {
			return false
		} else {
			mapCheck[string(v)]--
		}
	}
	return true
}

func rotate(nums []int, k int) {
	nums = append(nums, nums[:k+1]...)
	nums = nums[k+1:]

}
func reverseStr(s string, k int) string {
	l, r := 0, k
	var rev = make([]string, len(s))
	for i := 0; i < len(s); i += k {
		if i+k > len(s) {
			rev = append(rev, s[i:len(s)])
			break
		}
		for l < r {
			rev = append(rev, string(s[r-1]))
			r--
		}

		l = i + k
		r = l + k
	}
	return strings.Join(rev, "")
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return false
	}
	countLeft, countRight := 0, 0
	check := false
	balance(root, countLeft, countRight, &check)
	if (countLeft > countRight || countRight > countLeft) && !check {
		return false
	}
	return true

}
func balance(head *TreeNode, countLeft int, countRight int, check *bool) {
	if head.Left != nil {
		countLeft++
		balance(head.Left, countLeft, countRight, check)
	}
	if head.Right != nil {
		countRight++
		balance(head.Right, countRight, countRight, check)
	}
	if head.Left != nil && head.Right != nil {
		if head.Left.Val == head.Right.Val {
			*check = true
		}
	}

}
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}

	}
	if head == nil {
		return nil
	}
	root := head
	nextNode := head.Next

	for nextNode != nil {
		if nextNode.Val == val {
			nextNode = nextNode.Next
			root.Next = nextNode
		} else {
			nextNode = nextNode.Next
			root = root.Next
		}
	}
	if head.Next == nil && head.Val == val {
		return nil
	}
	return head
}
func isIsomorphic(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	mapS := make(map[byte]byte, len(s))
	mapT := make(map[byte]byte, len(t))
	for i := 0; i < len(s); i++ {
		if si, ok := mapS[s[i]]; ok && si != t[i] {
			return false
		}
		mapS[s[i]] = t[i]

		if ti, ok := mapT[t[i]]; ok && ti != s[i] {
			return false
		}
		mapT[t[i]] = s[i]
	}
	return true
}
func preorderTraversal(root *TreeNode) []int {
	var preord = func(root *TreeNode, arr *[]int) {}
	preord = func(root *TreeNode, arr *[]int) {
		if root == nil {
			return
		}
		*arr = append(*arr, root.Val)
		preord(root.Left, arr)
		preord(root.Right, arr)
	}
	var arr = &[]int{}
	preord(root, arr)
	sort.Slice(*arr, func(i, j int) bool {
		return i > j
	})
	return *arr
}

func maxPro(prices []int) int {
	max := 0
	profit := prices[0]
	for i := 1; i < len(prices); i++ {
		if profit > prices[i] {
			profit = prices[i]
		}
		max = int(math.Max(float64(max), float64(prices[i]-profit)))
	}
	return max
}

func generate(numRows int) [][]int {
	var dp = make([][]int, 5)
	dp[0] = []int{1}
	if numRows == 1 {
		return dp
	}
	dp[1] = []int{1, 1}
	if numRows == 2 {
		return dp
	}
	for i := 2; i < numRows; i++ {
		var arr = make([]int, i+1)
		dp[i] = arr
		for j := 0; j <= len(dp[i-1]); j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if j == len(dp[i-1]) {
				dp[i][j] = 1
			} else {

				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}

	}
	return dp
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	left := false
	hasPathSumDFS(root, targetSum, []int{}, &left)

	return left
}
func hasPathSumDFS(root *TreeNode, targetSum int, arr []int, check *bool) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	sum := 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum += arr[i]
		if sum == targetSum {
			*check = true
			return
		}
	}

	hasPathSumDFS(root.Left, targetSum, arr, check)
	hasPathSumDFS(root.Right, targetSum, arr, check)
}

func minDepth(root *TreeNode) int {
	var queue = []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		currentNodes := len(queue)
		level++
		for i := 0; i < currentNodes; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Right == nil && node.Left == nil {
				return level
			}
		}
	}
	return level
}

func countMinDepth(root *TreeNode, demLeft int, demRight int) (int, int) {
	if root.Left != nil {
		demLeft++
		countMinDepth(root.Left, demLeft, demRight)

	}
	if root.Right != nil {
		demRight++
		countMinDepth(root.Right, demLeft, demRight)
	}
	return demRight, demLeft

}

func deleteDuplicates(head *ListNode) *ListNode {
	root := head
	next := head.Next
	for next != nil {
		if root.Val == next.Val {
			next = next.Next
			root.Next = next
		} else {
			root = next
			next = next.Next
		}
	}
	return head
}

func climbStairs(n int) int {
	var dp = make([]int, n)
	if n == 0 {
		return 0
	}
	dp[1] = 1
	dp[2] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
func longestCommonPrefix(strs []string) string {
	p := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {
		}
		p = p[:i]
	}
	return p
}
func addBinary(a string, b string) string {
	m, n := len(a)-1, len(b)-1

	du := 0
	var ans = ""

	sum := 0
	for m >= 0 || n >= 0 {
		sum = du
		if m >= 0 {
			sum += int(a[m] - '0')
		}
		if n >= 0 {
			sum += int(b[n] - '0')
		}

		ans = string(byte(sum%2)+'0') + ans
		if sum <= 1 {
			du = 0
		} else {
			du = 1
		}
		m--
		n--

	}
	if du == 1 {
		ans = string(byte(du+'0')) + ans
	}

	return ans
}
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	pair := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}
	for _, r := range s {
		if _, ok := pair[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pair[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}
func das(nums [][]int) {
	n := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			println(nums[j][n-j-1])

		}
	}
}
func sumText(text1, text2 string) string {
	m, n := len(text1)-1, len(text2)-1
	var (
		du     = 0
		nguyen = 0
		str    []string
	)

	for m >= 0 || n >= 0 {
		var (
			a, b int
		)
		if m >= 0 {
			a = int(text1[m] - '0')
		}
		if n >= 0 {
			b = int(text2[n] - '0')
		}
		sum := a + b + du
		du = sum / 10
		nguyen = sum % 10
		str = append(str, strconv.Itoa(nguyen))
		m--
		n--
	}
	return strings.Join(str, "")
}
func maxOperation(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1

	count := 0
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			count++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}

	return count
}
func threeSum(nums []int) [][]int {
	var (
		result [][]int
	)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1

		for l < r {
			target := nums[i] + nums[l] + nums[r]
			if target == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				r--
				l++
				for l < r && nums[l] == nums[r] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if target > 0 {
				r--
			} else {
				l++

			}

		}
	}
	return result

}
func tribonacci(n int) int {
	var dp = make([]int, n+1)
	if n == 2 || n == 1 {
		return 1
	}

	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n-1]
}
func fibo(n int) int {
	var (
		memo = make([]int, n+1)
	)

	memo[1] = 1

	for i := 2; i < n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n-1]
}
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func subSetSumProblem(arr []int, s int) int {
	var dp = make([]int, s)
	dp[0] = 1
	for i := 0; i < len(arr); i++ {
		for j := s - 1; j >= arr[i]; j-- {
			if dp[j-arr[i]] == 1 {
				dp[j] = 1
			}
		}
	}
	if dp[s] == 1 {
		return 1
	}
	return 0
}

func lis(arr []int) int {
	var (
		max     = math.MinInt
		listLIS = make([]int, len(arr))
	)
	for i := 0; i < len(arr); i++ {
		listLIS[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				listLIS[i] = int(math.Max(float64(listLIS[i]), float64(listLIS[j]+1)))
				max = int(math.Max(float64(max), float64(listLIS[i])))
			}
		}
	}
	return max
}
func combinationSum3(k int, n int) [][]int {
	var (
		arr    []int
		arrSum = &[][]int{}
	)
	var backtracks = func(arr []int, res *[][]int, k int, n int, pos int) {}
	backtracks = func(arr []int, res *[][]int, k int, n int, pos int) {
		if n == 0 && len(arr) == k {
			var mapCheck = make(map[int]int)
			for i := 0; i < len(arr); i++ {
				if _, ok := mapCheck[arr[i]]; ok {
					return
				}
				mapCheck[arr[i]] = i
			}
			*res = append(*res, append([]int{}, arr...))
			return
		}
		for i := pos; i <= 9; i++ {

			arr = append(arr, i)
			if len(arr) > k {
				return
			}
			backtracks(arr, res, k, n-i, pos+1)
			pos = pos + 1
			arr = arr[:len(arr)-1]

		}
		return
	}

	backtracks(arr, arrSum, k, n, 1)
	return *arrSum
}

//func pairSum(head *ListNode) int {
//	node := head
//	var count int
//	for node != nil {
//		node = node.Next
//		count++
//	}
//	return count
//}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var longest = func(node *TreeNode, leftRight bool, count *int) {}
	longest = func(node *TreeNode, leftRight bool, count *int) {
		if node == nil {
			return
		}

		if node.Left != nil && !leftRight {
			*count++
			longest(node.Right, true, count)
		}
		if node.Right != nil && leftRight {
			*count++
			longest(node.Left, false, count)
		}
		return

	}
	queue := []*TreeNode{root}
	var max = math.MinInt
	for len(queue) > 0 {
		currentNodeLevel := len(queue)
		for i := 0; i < currentNodeLevel; i++ {
			var (
				count2 = new(int)
				count1 = new(int)
			)
			node := queue[0]
			queue = queue[1:]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			longest(node, true, count1)
			longest(node, false, count2)
			*count1 += *count2
			max = int(math.Max(float64(max), float64(*count1)))
		}

	}
	return max
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var countGoodNodes = func(*TreeNode, int, *int) {}
	countGoodNodes = func(node *TreeNode, prev int, count *int) {
		if node == nil {
			return
		}
		for node != nil {
			if prev <= node.Val {
				prev = node.Val
				*count++

			}
			if node.Left != nil {
				countGoodNodes(node.Left, prev, count)
			}
			if node.Right != nil {
				countGoodNodes(node.Right, prev, count)
			}

			return
		}
	}
	var count = new(int)
	*count = 0
	countGoodNodes(root, math.MinInt, count)
	return *count
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var pathSumDPS = func(*TreeNode, int, []int, *int) {}
	pathSumDPS = func(node *TreeNode, targetSum int, arr []int, count *int) {
		if node == nil {
			return
		}
		arr = append(arr, node.Val)
		sum := 0
		for i := len(arr) - 1; i >= 0; i-- {
			sum += arr[i]
			if sum == targetSum {
				*count++
			}
		}
		if node.Left != nil {
			pathSumDPS(node.Left, targetSum, arr, count)
		}
		if node.Right != nil {
			pathSumDPS(node.Right, targetSum, arr, count)
		}
		return

	}

	arr := []int{}
	var count = new(int)
	pathSumDPS(root, targetSum, arr, count)
	return *count

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1, arr2 := &[]int{}, &[]int{}
	arrLeaf(root1, arr1)
	arrLeaf(root2, arr2)
	if arr1 == arr2 {
		return true
	} else {
		return false
	}
	//return compareArrays(*arr1, *arr2)
}
func compareArrays(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func arrLeaf(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	for root != nil {
		if root.Left != nil {
			arrLeaf(root.Left, arr)
		}
		if root.Right != nil {
			arrLeaf(root.Right, arr)
		}
		if root.Left == nil && root.Right == nil {
			*arr = append(*arr, root.Val)
		}
		return
	}
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil || root.Val == key {
		return nil
	}

	deletes(root, &TreeNode{}, true, key)
	deletes(root, &TreeNode{}, false, key)
	return root
}

func deletes(root *TreeNode, preNode *TreeNode, leftRight bool, key int) {
	for root != nil {
		if root.Val == key {
			newNode := &TreeNode{}
			if leftRight == true {
				newNode = root.Right
				newNode.Left = root.Left
				preNode.Left = newNode
			} else {
				newNode = root.Left
				newNode.Right = root.Right
				preNode.Right = newNode
			}
			return
		}
		if root.Left != nil {
			deletes(root.Left, root, true, key)
		}
		if root.Right != nil {
			deletes(root.Right, root, false, key)
		}
		return
	}
}
func maxBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		level    int
		resLevel int
		res      = math.MinInt
	)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nodesCurrentLevel := len(queue)
		sum := 0
		level++
		for i := 0; i < nodesCurrentLevel; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if res < sum {
			res = sum
			resLevel = level
		}
	}
	return resLevel
}
func rightSideBFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nodesCurrentLevel := len(queue)
		prev := 0
		for i := 0; i < nodesCurrentLevel; i++ {
			node := queue[0]
			queue = queue[1:]
			prev = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		ans = append(ans, prev)
	}
	return ans
}

func PrintName(people []*People, arr *[]string) {
	if people == nil {
		return
	}
	for _, val := range people {
		PrintName(val.Child, arr)
		*arr = append(*arr, val.Name)
	}
}
func a(s string) []string {
	var (
		res []string
	)
	mapStruct := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var result []string
	for i := 0; i < len(s); i++ {
		if _, ok := mapStruct[string(s[i])]; ok {
			result = append(result, mapStruct[string(s[i])])
		}
	}
	solutions(result, 0, "", res)
	return res
}

func solutions(arr []string, index int, s string, res []string) {

	newStr := arr[index]
	for i := 0; i < len(newStr); i++ {
		if len(s) == len(arr) {
			res = append(res, s)
			solutions(arr, index+1, s, res)
			return
		}
		aa := fmt.Sprintf("%s%s", s, newStr[i])
		solutions(arr, index+1, aa, res)
	}
}
func removeStars(s string) string {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = append(stack[:len(stack)-1])
		} else {
			stack = append(stack, s[i])
		}

	}

	return string(stack)
}

func equalPairs(grid [][]int) int {
	mapCols := make(map[[200]int]int)
	mapRows := make(map[[200]int]int)
	res := 0
	for i := 0; i < len(grid); i++ {
		var arr [200]int
		copy(arr[:], grid[i])
		mapRows[arr]++
	}

	for i := 0; i < len(grid); i++ {
		var (
			a   []int
			arr [200]int
		)
		for j := 0; j < len(grid); j++ {
			a = append(a, grid[j][i])
		}
		copy(arr[:], a)
		mapCols[arr]++
		if _, ok := mapRows[arr]; ok {
			res += mapRows[arr]
		}
	}
	return res
}

type TreeNodes struct {
	Val   int
	Left  *TreeNodes
	Right *TreeNodes
}

//	func max(root *TreeNodes) int {
//		for root != nil {
//			if root.Right != nil {
//				max(root.Right)
//			} else {
//				max(root.Left)
//			}
//		}
//	}
func u(arr []int) bool {
	var (
		map1   = make(map[int]int)
		map2   = make(map[int]int)
		result = true
	)

	for i := 0; i < len(arr); i++ {
		if _, ok := map1[arr[i]]; ok {
			map1[arr[i]]++
		} else {
			map1[arr[i]] = 1
		}
	}
	for _, v := range map1 {
		if _, ok := map2[v]; ok {
			result = false
		} else {
			map2[v] = v
		}
	}
	return result
}
func diff(nums1 []int, nums2 []int) [][]int {
	var (
		map1   = make(map[int]int)
		map2   = make(map[int]int)
		res1   []int
		res2   []int
		result [][]int
	)
	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]] = nums1[i]
	}
	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]] = nums2[i]
	}
	for _, v := range map1 {
		if _, ok := map2[v]; ok {
			delete(map1, v)
			delete(map2, v)
		}
	}
	for _, v := range map1 {
		res1 = append(res1, v)
	}
	for _, v := range map2 {
		res2 = append(res2, v)
	}
	result = append(result, res1, res2)
	return result
}
func maxProfit(arr []int) int {
	l := 0
	var max = 0
	for i := 1; i < len(arr); i++ {
		if arr[l] > arr[i] {
			l = i
		}
		max = int(math.Max(float64(max), float64(arr[i]-arr[l])))
	}
	return max
}

func SumTwo(arr []int, k int64) []int64 {
	var (
		hashMap = make(map[int64]int64)
		result  []int64
	)

	for i := 0; i < len(arr); i++ {
		s := k - int64(arr[i])
		if _, ok := hashMap[s]; ok {
			result = append(result, hashMap[s], int64(i))
			return result
		} else {
			hashMap[int64(arr[i])] = int64(i)
		}
	}
	return result
}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	l, r := 0, 1
	res := 0
	for l < r {
		if nums[r]+nums[l] == k {
			res++
			l++
			r--
		} else if nums[r]+nums[l] < k {
			l++
		} else {
			r--
		}
	}

	return res
}

func maxA(height []int) int {
	l, r := 0, len(height)-1
	max, res := 0, 0
	for width := len(height) - 1; width > 0; width-- {
		if height[l] < height[r] {
			max = width * height[l]
			l++
		} else {
			max = width * height[r]
			r--
		}

		res = int(math.Max(float64(max), float64(res)))
	}
	return res

}

func isS(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == len(s) {
		return true
	}
	return false
}

func mergeAlternately(word1 string, word2 string) string {
	var result []string
	if len(word2) > len(word1) {
		for i := 0; i < len(word2); i++ {
			result = append(result, string(word1[i]), string(word2[i]))
		}
	} else {
		for i := 0; i < len(word1); i++ {
			result = append(result, string(word1[i]), string(word2[i]))
		}
	}
	return strings.Join(result, "")
}

func mergeSort2Array(arr1 []int, arr2 []int) []int {
	var result []int
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[i] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}
	return result
}
func compare(a, b []int) int64 {
	var result int64
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
			continue
		} else if a[i] > b[j] {
			j++
			continue
		} else {
			dem0, dem1 := 0, 0
			for j < len(a) && a[i] == b[j] {
				dem0++
				i++
			}
			tmp := a[i-1]
			for j < len(b) && b[j] == tmp {
				dem1++
				j++
			}
			result += int64(dem0 * dem1)
		}
	}
	return result
}
