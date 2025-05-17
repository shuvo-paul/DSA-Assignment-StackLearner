package main

import "testing"

func TestRecentPosts(t *testing.T) {
	rp := NewRecentPosts(10)
	rp.AddPost("Post 1")
	rp.AddPost("Post 2")
	rp.AddPost("Post 3")
	rp.AddPost("Post 4")

	posts := rp.GetRecentPosts()
	if posts[3] != "Post 4" {
		t.Errorf("Expected Post 4, got %s", posts[0])
	}

	rp.AddPost("Post 5")
	rp.AddPost("Post 6")
	rp.AddPost("Post 7")
	rp.AddPost("Post 8")
	rp.AddPost("Post 9")
	rp.AddPost("Post 10")
	rp.AddPost("Post 11")

	posts = rp.GetRecentPosts()

	if len(posts) != 10 {
		t.Errorf("Expected 10 posts, got %d", len(posts))
	}

	if posts[0] != "Post 2" {
		t.Errorf("Expected Post 2, got %s", posts[0])
	}
}

type RecentPosts struct {
	capacity int
	posts    []string
}

func NewRecentPosts(capacity int) *RecentPosts {
	return &RecentPosts{
		capacity: capacity,
		posts:    make([]string, 0, capacity),
	}
}

func (rp *RecentPosts) AddPost(post string) {
	if len(rp.posts) == rp.capacity {
		rp.posts = rp.posts[1:]
	}
	rp.posts = append(rp.posts, post)
}

func (rp *RecentPosts) GetRecentPosts() []string {
	return rp.posts
}
