package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"sync"
	"time"
)

func main() {

	fmt.Println(5 % 5)
}

// work schedule
func a1(s string) string {
	var str []string
	l, r := len(s), len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		l, r = i, i+1

		for l >= 0 && s[l] != ' ' {
			l--
		}

		str = append(str, s[l+1:r])
		i = l

	}

	return strings.Join(str, " ")
}

func a4(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	right[len(nums)-1] = 1
	left[0] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
		right[len(nums)-1-i] = right[len(nums)-i] * nums[len(nums)-i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = left[i] * right[i]
	}
	return nums
}

func a5(nums []int) bool {

	triplet1, triplet2 := math.MaxInt32, math.MaxInt32

	for i := 0; i < len(nums); i++ {
		if nums[i] <= triplet1 {
			triplet1 = nums[i]
		} else if nums[i] <= triplet2 {
			triplet2 = nums[i]
		} else {
			return true
		}
	}
	return false

}

func sleepRandom(fromFunction string, ch chan int) {
	defer func() {
		fmt.Println(fromFunction, "sleepRandom complete")
	}()

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleeptime := randomNumber + 100
	fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms")
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "Walking up, slept for", sleeptime, "ms")
	if ch != nil {
		ch <- sleeptime
	}
}

func sleepRandomContext(ctx context.Context, ch chan bool) {
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()

	sleeptimeChan := make(chan int)
	go sleepRandom("sleepRandomContext", sleeptimeChan)

	select {
	case <-ctx.Done():
		fmt.Println("sleepRandomContext: Time to return")
	case sleeptime := <-sleeptimeChan:
		fmt.Println("Slept for ", sleeptime, "ms")
	}
}
func doWorkContext(ctx context.Context) {
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)

	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()

	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)
	select {
	case <-ctx.Done():
		fmt.Println("doWorkContext: Time to return")
	case <-ch:
		fmt.Println("sleepRandomContext returned")
	}
}
func isPalindrome(x int) bool {
	var arr []int
	for {
		du := x % 10
		x = x / 10
		arr = append(arr, du)
		if x == 0 {
			break
		}
	}
	for i := 0; i < len(arr)/2; i++ {
		if arr[len(arr)-1-i] != arr[i] {
			return false
		}
	}
	return true
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits

}

func BinarySearchs(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}

		center := (left + right) / 2
		if nums[center] == target {
			return center
		}
		if nums[center] > target {
			right = center
		}
		if nums[center] <= target {
			left = center
		}
		if left >= right-1 {
			return -1
		}
	}
	return -1
}
func isAnagram(s string) bool {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	s1 := strings.Fields(nonAlphanumericRegex.ReplaceAllString(s, ""))

	newS := strings.Split(strings.ToLower(strings.Join(s1, "")), "")
	for i := 0; i < len(newS)/2; i++ {
		if newS[i] != newS[len(newS)-i-1] {
			return false
		}
	}
	return true
}
func containsDuplicate(nums []int) bool {
	mapCheck := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := mapCheck[nums[i]]; ok {
			return true
		}
		mapCheck[nums[i]] = i
	}
	return false
}
func findDifference(nums1 []int, nums2 []int) [][]int {
	var (
		mapCheck = map[int]int{}
	)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				if _, ok := mapCheck[nums1[i]]; ok {
					break
				}
				mapCheck[nums1[i]] = nums1[i]
				break
			}
		}
	}

	return [][]int{nums1, nums2}
}
func removeDuplicates(nums []int) int {
	prev := nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if prev != nums[i] {
			nums[l] = nums[i]

			l++
		}
		prev = nums[i]
	}
	return l
}
func removeDuplicates1(nums []int) []int {
	var (
		result   []int
		mapCheck = map[int]int{}
	)

	for i := 0; i < len(nums); i++ {
		if _, ok := mapCheck[nums[i]]; ok {
			continue
		}
		mapCheck[nums[i]] = nums[i]
		result = append(result, nums[i])
	}
	for i := 0; i < len(mapCheck); i++ {
		nums[i] = mapCheck[i]
	}
	return nums[:len(mapCheck)]
}
func lastStoneWeight(stone []int) int {

	sort.Slice(stone, func(i, j int) bool {
		return stone[i] > stone[j]
	})

	for i := 0; i < len(stone); i++ {
		if len(stone) == 1 {
			return stone[0]
		}
		if stone[0] == stone[1] {
			if len(stone) == 2 {
				return 0
			}
			stone = stone[2:len(stone)]
			sort.Slice(stone, func(i, j int) bool {
				return stone[i] > stone[j]
			})
			i--
			continue
		}
		stone[1] = stone[0] - stone[1]
		stone = stone[1:len(stone)]
		i--
		sort.Slice(stone, func(i, j int) bool {
			return stone[i] > stone[j]
		})

	}
	return stone[0]
}

func maxSubArray(nums []int) int {
	var (
		sum = 0
		max = math.MinInt
	)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		max = int(math.Max(float64(sum), float64(max)))

		if sum < 0 {
			sum = 0
		}
	}
	return max
}
func reverse(x int) int {
	var (
		result  int
		arr     []int
		checkAm bool
	)
	if x < 0 {
		x = x * (-1)
		checkAm = true
	}
	for {
		du := x % 10
		nguyen := x / 10
		arr = append(arr, du)
		if nguyen <= 0 {
			break
		}
		x = nguyen
	}

	for i := 0; i < len(arr); i++ {
		result += arr[i] * int(math.Pow10(len(arr)-i-1))
	}
	if checkAm {
		return result * -1
	}

	return result
}

var (
	elasticClient *elastic.Client
)

func ReceiveAndSend(c chan int) {

	fmt.Printf("Receive: %d\n", <-c)
	fmt.Printf("Send ...2\n")
	c <- 2
}

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	value := []int64{12312, 3312, 124}
	args := make([]interface{}, 2+len(value))
	args[0] = "SET"
	args[1] = "keyHuan"

	for i, key := range value {
		args[2+i] = key
	}

	cmd := redis.NewBoolCmd(args...)
	fmt.Println(cmd)
	_ = rdb.Process(cmd)
	fmt.Print(cmd.Result())
}

func sender(c chan<- int, name string) {
	for i := 1; i <= 100; i++ {
		c <- 1
		fmt.Printf("%s has sent 1 to channel\n", name)
		runtime.Gosched()
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func receiveOnly(c <-chan int) {
	fmt.Printf("Received: %v", <-c)
}

func sendOnly(c chan<- int) {
	c <- 2
}

func receiveAndSend(c chan int) {
	fmt.Printf("Receive: %d\n", <-c)
	fmt.Printf("Send 2...\n")
	c <- 2
}

type SendNotificationV2Request struct {
	ApiKey     string `json:"-"`
	AppID      string `json:"app_id"`
	TemplateID string `json:"template_id"`
}

func timeConversion(s string) string {

	if ok := strings.Contains(s, "PM"); ok {
		str := strings.Replace(s, "PM", "", 1)
		str1 := strings.Split(str, ":")
		num, err := strconv.Atoi(str1[0])
		if err != nil {
			panic(err)
		}
		str1[0] = strconv.Itoa(num + 12)

		newStr := strings.Join(str1, ":")
		fmt.Println(newStr)
		return newStr
	} else {
		str1 := strings.Split(s, ":")
		num, err := strconv.Atoi(str1[0])
		if err != nil {
			panic(err)
		}
		if str1[0] == "12" {
			str1[0] = strconv.Itoa(num-12) + "0"
		}
		newStr := strings.Join(str1, ":")
		return strings.Replace(newStr, "AM", "", 1)
	}

}

func noonTask() {
	fmt.Println(time.Now())
	fmt.Println("do some job.")
}

func initNoon() {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), 9, 36, 0, 5, t.Location())
	d := n.Sub(t)
	if d < 0 {
		n = n.Add(10 * time.Second)
		d = n.Sub(t)
	}
	for {
		time.Sleep(d)
		d = 5 * time.Second
		noonTask()
	}
}

type limitRewardCampaign struct {
	UserID        int64
	RewardID      int64
	CampaignID    int64
	RuleID        int64
	Limit         int64
	ExcludeReward []int64
	StartDate     time.Time
	EndDate       time.Time
}

func checkURLs(urls []string) {
	var wg sync.WaitGroup
	c := make(chan string)
	for _, url := range urls {
		wg.Add(1)
		go checkURL(url, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}

func checkURL(url string, c chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	_, err := http.Get(url)
	if err != nil {
		c <- url + " can not be reached"
	} else {
		c <- url + " can be reached"
	}
}
func useRateLimit(rateLimit int64, second int64, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		key := "RATE_LIMIT_COUNT_" + clientIp
		err := incRequestCount(key, rateLimit, second, rdb)
		if err != nil {
			c.AbortWithStatus(403)
			return
		}
		c.Next()
	}
}
func incRequestCount(key string, rateLimit int64, second int64, rdb *redis.Client) error {
	err := rdb.Watch(func(tx *redis.Tx) error {
		_ = tx.SetNX(key, 0, time.Duration(second)*time.Second)
		count, err := tx.Incr(key).Result()
		fmt.Println(count)
		if count > rateLimit {
			err = errors.New("rate limited")
		}
		if err != nil {
			return err
		}
		return nil
	}, key)
	return err
}
func doSomething(ctx context.Context) {
	//Context WithValue
	//fmt.Printf("Do Something: myKey's value is %s\n", ctx.Value("myKey"))
	//
	//anotherCtx := context.WithValue(ctx, "myKey", "myAnother")
	//doAnother(anotherCtx)
	//fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

	//Context WithCancel
	ctx, cancelCtx := context.WithCancel(ctx)
	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelCtx()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Do Something: finish \n")
}
func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("Do Another err :%s\n", err)
			}
			fmt.Printf("Do Another: finished\n")
			return
		case num := <-printCh:
			fmt.Printf("Do Another: %d\n", num)
		}

	}
}
func In(arr []int) {
	for _, val := range arr {
		fmt.Println(val)
		time.Sleep(1 * time.Second)
	}
}
