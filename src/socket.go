package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"math"
	"sort"
	"strings"
	"sync"
	"time"
)

func moveZeroes(nums []int) {
	var arr []*int

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[1:], 0)
		}
	}
	fmt.Println(nums)
	_ = copier.Copy(arr, nums)
}

func reverseList(head *ListNode) *ListNode {
	var revHead *ListNode
	if head == nil {
		return nil
	}
	if head != nil {
		tmp := head.Next
		head.Next = revHead
		revHead = head
		head = tmp

	}
	return revHead
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if val == nums[i] {
			nums = append(nums[:i], nums[i+1:len(nums)]...)
			i--
		}
	}
	return len(nums)
}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		} else if i == len(nums)-1 {
			nums = append(nums, target)
			sort.Ints(nums)
			break
		}
	}
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
	}
	return 0
}

func strStr(haystack string, needle string) int {
	if ok := strings.Contains(haystack, needle); !ok {
		return -1
	}
	for i := len(haystack); i >= 0; i-- {
		var checkStr string
		checkStr = haystack[:i-1]
		if ok := strings.Contains(checkStr, needle); ok {
			continue
		}
		return i - len(needle)
	}
	return -1
}

func lengthOfLastWord(s string) int {
	newStr := strings.Fields(s)
	return len(newStr[len(newStr)-1])
}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func singleNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		var temp int
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			} else if nums[i] == nums[j] {
				break
			}
			temp++
		}
		if temp == len(nums)-1 {
			return nums[i]
		}
	}

	return 0
}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 1 && nums[0] > 0 {
		return nums[0] - 1
	}
	for i := 0; i <= len(nums)-1; i++ {
		if i == len(nums)-1 {
			if nums[0] > 0 {
				return nums[0] - 1
			}
			return nums[i] + 1
		} else if nums[i]+1 != nums[i+1] {
			return nums[i] + 1
		}
	}
	return 0

}

func runningSum(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		var total int

		for j := 0; j <= i; j++ {

			total += nums[j]
		}
		result = append(result, total)
	}
	return result
}

//func mergeAlternately(word1 string, word2 string) string {
//	var str string
//
//	for i := 0; i < len(word1)+len(word2)-1; i++ {
//
//		if len(word2) <= i {
//			str += word1[i:]
//			break
//		} else if len(word1) <= i {
//			str += word2[i:]
//			break
//		} else if len(word1) > i && len(word2) > i {
//			str += word1[i:i+1] + word2[i:i+1]
//		}
//	}
//	return str
//}

type SnapshotArray struct {
	value    []int
	index    int
	snapID   int
	mapArray map[int]int
	snapShot map[int]map[int]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		value:    make([]int, length),
		mapArray: make(map[int]int, length),
		snapShot: make(map[int]map[int]int, length),
		snapID:   0,
		index:    0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.index = index
	this.value[index] = val
}

func (this *SnapshotArray) Snap() int {
	this.mapArray[this.index] = this.value[this.index]
	if _, ok := this.snapShot[this.snapID]; ok {
		this.snapID++
	}
	this.snapShot[this.snapID] = this.mapArray
	return this.snapID
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	for _, val := range this.snapShot[snap_id] {
		if val == this.mapArray[index] {
			return this.mapArray[index]
		}
	}
	return 0
}

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	temp := math.Abs(float64(arr[0] - arr[1]))
	for i := 0; i < len(arr)-1; i++ {
		if math.Abs(float64(arr[i]-arr[i+1])) != temp {
			return false
		}
	}
	return true
}

const (
	layoutDate = "2006-01-02 15:04:05"
)

func main() {
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
		//fallthrough
	case "Tuesday":
		fmt.Println("It's Tuesday.")
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	default:
		fmt.Println("It's another day.")
	}

}

//func swapPairs(head *ListNode) *ListNode {
//	if head.Next == nil {
//		return head
//	}
//
//}

func lengthOfLongestSubstring(s string) int {
	var result int
	if len(s) == 1 {
		return 1
	}
	for i := 0; i < len(s)-1; i++ {
		min := 1
		var mapCheck = make(map[byte]bool)
		mapCheck[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if _, ok := mapCheck[s[j]]; ok {
				break
			}
			mapCheck[s[j]] = true
			min++
		}
		if min > result {
			result = min
		}
	}
	return result
}

//func longestPalindrome(s string) string {
//	var result string
//	var min int
//	for i := 0; i < len(s)-1; i++ {
//		for j := i + 1; j < len(s); j++ {
//			if isPalindromic(s[i : j+1]) {
//				if min < len(s[i:j+1]) {
//					min = len(s[i : j+1])
//					result = s[i : j+1]
//				}
//			}
//		}
//	}
//	if result == "" {
//		return strings.Split(s, "")[0]
//	}
//	return result
//}

func isPalindromic(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	if len(s) == 2 {
		return true
	}
	return true
}

func convert(s string, numRows int) string {
	rows := make([]strings.Builder, numRows+1)

	idx := 0
	for idx < len(s) {
		for i := 0; i < numRows && idx < len(s); i++ {
			rows[i].WriteByte(s[idx])
			idx++
			fmt.Println(rows[i].String() + " ")

		}
		for j := numRows - 2; 0 < j && idx < len(s); j-- {
			rows[j].WriteByte(s[idx])
			idx++
			fmt.Println(rows[j].String() + " ")
		}
	}
	var zigzag strings.Builder
	for _, r := range rows {
		zigzag.WriteString(r.String())
	}
	return zigzag.String()
}
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int
	backtrack(&result, current, candidates, target, 0)
	return result
}
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func backtrack(result *[][]int, current []int, candidates []int, remain int, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		sort.Ints(current)
		for _, item := range *result {
			if Equal(item, current) {
				return
			}
		}
		*result = append(*result, append([]int{}, current...))
		return
	} else {
		for i := start; i < len(candidates); i++ {
			current = append(current, candidates[i])
			backtrack(result, current, candidates, remain-candidates[i], i+1)
			current = current[:len(current)-1]
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) []int {
	tmp := &ListNode{}
	var result []int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{}
		}
		result = append(result, tmp.Val)
		tmp = tmp.Next
	}
	return result
}

type user struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type cachedUser struct {
	user
	expireAtTimestamp int64
}

type localCache struct {
	stop  chan struct{}
	wg    sync.WaitGroup
	mu    sync.RWMutex
	users map[int64]cachedUser
}

func newLocalCache(cleanupInterval time.Duration) *localCache {
	lc := &localCache{
		users: make(map[int64]cachedUser),
		stop:  make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()
		lc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)
	return lc
}

func (lc *localCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for uid, cu := range lc.users {
				if cu.expireAtTimestamp <= time.Now().Unix() {
					delete(lc.users, uid)
				}
			}
			lc.mu.Unlock()
		}
	}
}

func (lc *localCache) stopCleanup() {
	close(lc.stop)
	lc.wg.Wait()
}

func (lc *localCache) update(u user, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.users[u.Id] = cachedUser{
		user:              u,
		expireAtTimestamp: expireAtTimestamp,
	}
}

var (
	errUserNotInCache = errors.New("the user isn't in cache")
)

func (lc *localCache) read(id int64) (user, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	cu, ok := lc.users[id]
	if !ok {
		return user{}, errUserNotInCache
	}

	return cu.user, nil
}
func (lc *localCache) delete(id int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.users, id)
}

func solution(arr []int32) map[int32]int32 {

	mapResult := map[int32]int32{}

	for i := 0; i < len(arr); i++ {
		if _, ok := mapResult[arr[i]]; ok {
			mapResult[arr[i]]++
		} else {
			mapResult[arr[i]] = 1
		}
	}
	return mapResult
}
