package main

import (
	"container/heap"
	_ "net/http/pprof"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortStr := string(s)
		mp[sortStr] = append(mp[sortStr], str)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}
func longestConsecutive(nums []int) int {
	result := 0
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if result < currentStreak {
				result = currentStreak
			}
		}
	}
	return result
}

func moveZeroes(nums []int) {
	left, right, max := 0, 0, len(nums)
	for right < max {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	if nums[0] > 0 {
		return nil
	}
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func trap(height []int) int {
	ans := 0
	n := len(height)
	if n == 0 {
		return ans
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}

	return ans
}

func lengthOfLongestSubstring(s string) int {
	left, ans := 0, 0
	cnt := [128]int{}
	for right, c := range s {
		cnt[c]++
		for cnt[c] > 1 {
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}

func findAnagrams(s string, p string) []int {
	targetLen := len(p)

	var res []int
	tmpP := []byte(p)
	sort.Slice(tmpP, func(i, j int) bool {
		return tmpP[i] < tmpP[j]
	})
	sortP := string(tmpP)

	for i, _ := range s {
		if i+targetLen-1 < len(s) {
			tmp := []byte(s[i : i+targetLen])
			sort.Slice(tmp, func(i, j int) bool {
				return tmp[i] < tmp[j]
			})
			tmpStr := string(tmp)
			if sortP == tmpStr {
				res = append(res, i)
			}
		}
	}

	return res
}

func findAnagrams2(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

func subarraySum(nums []int, k int) int {
	ans := 0
	for left, n := range nums {
		sum := n
		if sum == k {
			ans++
		}
		right := left + 1
		for sum <= k && right < len(nums) {
			sum = sum + nums[right]
			if sum == k {
				ans++
			}
		}

	}
	return ans
}

func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow2(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

func inorderTraversal2(root *TreeNode) []int {
	var recur func(node *TreeNode)
	var res []int
	recur = func(node *TreeNode) {
		if node == nil {
			return
		}

		recur(node.Left)
		res = append(res, node.Val)
		recur(node.Right)
	}

	recur(root)
	return res
}

func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	paris := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if paris[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != paris[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}

	l.head.prev = l.tail
	l.tail.next = l.head

	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++

		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		} else {
			node := this.cache[key]
			node.value = value
			this.moveToHead(node)
		}
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func findDuplicate(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i, v := range nums {
		if nums[i] == nums[i+1] {
			return v
		}
	}
	return 0
}

func findDuplicate2(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

func searchInsert(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid
		}
	}

	return 0
}

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] > prev[1] {
			res = append(res, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res

}

func maxSubArray(nums []int) int {
	ans := nums[0]
	var dp = make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

func partition2(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	ans := [][]string{}
	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
		}

		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func partition(s string) [][]string {
	n := len(s)
	f := make([][]int, 8)
	for i := range f {
		f[i] = make([]int, n)
	}

	var isPartition func(i, j int) int
	isPartition = func(i, j int) int {
		if i >= j {
			return 1
		}

		if f[i][j] != 0 {
			return f[i][j]
		}

		f[i][j] = -1
		if s[i] == s[j] {
			f[i][j] = isPartition(i+1, j-1)
		}
		return f[i][j]
	}

	ans := [][]string{}
	var splits []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
		}

		for j := i; j < n; j++ {
			if isPartition(i, j) > 0 {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}

	dfs(0)
	return ans

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := head
	for i := 0; i < length-n; i++ {
		if i == length-n-1 {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dummy.Next
}

func getLength(head *ListNode) int {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}

	return length
}

func rotate3(matrix [][]int) {
	n := len(matrix)
	temp := make([][]int, n)
	for i := range temp {
		temp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			temp[j][n-1-i] = v
		}
	}
	copy(matrix, temp)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {

	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	left, right, result := 0, len(height)-1, 0
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > result {
			result = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func main() {
	//moveZeroes([]int{0, 1, 0, 3, 12})
	rotate3([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
