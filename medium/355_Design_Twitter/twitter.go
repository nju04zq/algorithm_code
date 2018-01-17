package main

import "fmt"

type Tweet struct {
	tid, time int
}

type TwitterUser struct {
	uid     int
	tweets  []*Tweet
	follows map[int]bool
}

type Twitter struct {
	time  int
	users map[int]*TwitterUser
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	t := new(Twitter)
	t.users = make(map[int]*TwitterUser)
	return *t
}

func (t *Twitter) createUser(userId int) *TwitterUser {
	u := new(TwitterUser)
	u.uid = userId
	u.tweets = make([]*Tweet, 0)
	u.follows = make(map[int]bool)
	t.users[userId] = u
	return u
}

func (t *Twitter) getTime() int {
	time := t.time
	t.time++
	return time
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	user, ok := t.users[userId]
	if !ok {
		user = t.createUser(userId)
	}
	tweet := &Tweet{tweetId, t.getTime()}
	user.tweets = append(user.tweets, tweet)
}

type FeedEntry struct {
	uid, fid, time, idx int
}

type FeedHeap struct {
	buf []*FeedEntry
}

func (h *FeedHeap) init() *FeedHeap {
	h.buf = make([]*FeedEntry, 0)
	return h
}

func (h *FeedHeap) dump() {
	fmt.Println("=====Heap=====")
	for _, e := range h.buf {
		fmt.Printf("uid %d, fid %d, idx %d\n", e.uid, e.fid, e.time, e.idx)
	}
	fmt.Println("==============")
}

func (h *FeedHeap) size() int {
	return len(h.buf)
}

func (h *FeedHeap) parent(child int) int {
	return (child - 1) / 2
}

func (h *FeedHeap) lchild(parent int) int {
	return parent*2 + 1
}

func (h *FeedHeap) rchild(parent int) int {
	return parent*2 + 2
}

func (h *FeedHeap) priority(i int) int {
	return h.buf[i].time
}

func (h *FeedHeap) swap(i, j int) {
	h.buf[i], h.buf[j] = h.buf[j], h.buf[i]
}

func (h *FeedHeap) add(uid, fid, time, idx int) {
	e := &FeedEntry{uid, fid, time, idx}
	h.buf = append(h.buf, e)
	i := len(h.buf) - 1
	for i > 0 {
		parent := h.parent(i)
		if h.priority(parent) > h.priority(i) {
			break
		}
		h.swap(parent, i)
		i = parent
	}
}

func (h *FeedHeap) pop() (uid int, fid, idx int) {
	size := h.size()
	if size == 0 {
		return -1, -1, -1
	}
	uid, fid, idx = h.buf[0].uid, h.buf[0].fid, h.buf[0].idx
	h.swap(0, size-1)
	h.buf = h.buf[:size-1]
	size--
	i := 0
	for i < size {
		highest := i
		lchild, rchild := h.lchild(i), h.rchild(i)
		if lchild < size && h.priority(lchild) > h.priority(highest) {
			highest = lchild
		}
		if rchild < size && h.priority(rchild) > h.priority(highest) {
			highest = rchild
		}
		if i == highest {
			break
		}
		h.swap(i, highest)
		i = highest
	}
	return
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	u, ok := t.users[userId]
	if !ok {
		return []int{}
	}
	n := 10
	feeds := make([]int, 0, n)
	heap := new(FeedHeap).init()
	if m := len(u.tweets); m > 0 {
		tweet := u.tweets[m-1]
		heap.add(userId, tweet.tid, tweet.time, m-1)
	}
	for uid, _ := range u.follows {
		if m := len(t.users[uid].tweets); m > 0 {
			tweet := t.users[uid].tweets[m-1]
			heap.add(uid, tweet.tid, tweet.time, m-1)
		}
	}
	for i := 0; i < n; i++ {
		if heap.size() == 0 {
			break
		}
		uid, fid, idx := heap.pop()
		feeds = append(feeds, fid)
		idx--
		if idx >= 0 {
			tweet := t.users[uid].tweets[idx]
			heap.add(uid, tweet.tid, tweet.time, idx)
		}
	}
	return feeds
}

func (t *Twitter) hasUser(userId int) bool {
	_, ok := t.users[userId]
	return ok
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if !t.hasUser(followerId) {
		t.createUser(followerId)
	}
	if !t.hasUser(followeeId) {
		t.createUser(followeeId)
	}
	u := t.users[followerId]
	u.follows[followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if !t.hasUser(followerId) {
		return
	}
	if !t.hasUser(followeeId) {
		return
	}
	u := t.users[followerId]
	if _, ok := u.follows[followeeId]; ok {
		delete(u.follows, followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func testCase1() {
	fmt.Println("test cases 1")
	obj := Constructor()
	t := &obj
	fmt.Println(t.GetNewsFeed(1))
	t.PostTweet(1, 5)
	fmt.Println(t.GetNewsFeed(1))
	t.Follow(1, 2)
	t.PostTweet(2, 6)
	t.PostTweet(2, 7)
	t.PostTweet(1, 8)
	t.PostTweet(2, 9)
	t.PostTweet(2, 10)
	t.PostTweet(1, 11)
	t.PostTweet(1, 12)
	t.PostTweet(1, 13)
	t.PostTweet(1, 14)
	t.PostTweet(2, 15)
	t.PostTweet(2, 16)
	t.PostTweet(2, 17)
	t.PostTweet(1, 18)
	fmt.Println(t.GetNewsFeed(1))
	t.Unfollow(1, 2)
	fmt.Println(t.GetNewsFeed(1))
}

func testCase2() {
	//[505 94 10 13 5]
	//[22 333 505 94 2 10 13 101 3 5]
	fmt.Println("test cases 2")
	obj := Constructor()
	t := &obj
	t.PostTweet(2, 5)
	t.PostTweet(1, 3)
	t.PostTweet(1, 101)
	t.PostTweet(2, 13)
	t.PostTweet(2, 10)
	t.PostTweet(1, 2)
	t.PostTweet(2, 94)
	t.PostTweet(2, 505)
	t.PostTweet(1, 333)
	t.PostTweet(1, 22)
	fmt.Println(t.GetNewsFeed(2))
	t.Follow(2, 1)
	fmt.Println(t.GetNewsFeed(2))
}

func main() {
	testCase1()
	testCase2()
}
