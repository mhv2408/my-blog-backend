package main

type BlogPost struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Post    string `json:"post"`
}

func serveJson() []BlogPost {

	posts := []BlogPost{
		{
			Title:   "My First Blog Post",
			Summary: "The first blog post I've written on my blogging software!",
			Post:    "This is my first blog post content...",
		},
		{
			Title:   "My Second Blog Post",
			Summary: "The second blog post I've written on my blogging software!",
			Post:    "This is my second blog post content...",
		},
		{
			Title:   "My Thrid Blog Post",
			Summary: "The third blog post I've written on my blogging software!",
			Post:    "This is my third blog post content...",
		},
	}

	return posts

}
