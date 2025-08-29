package main

type BlogPost struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Excerpt  string `json:"excerpt"`
	Date     string `json:"date"`
	ReadTime string `json:"readTime"`
	Slug     string `json:"slug"`
}

func serveJson() []BlogPost {

	posts := []BlogPost{
		{
			Id:       1,
			Title:    "Understanding REST APIs",
			Excerpt:  "A beginner-friendly guide to how REST APIs work and why they matter in modern applications.",
			Date:     "January 10, 2025",
			ReadTime: "6 min read",
			Slug:     "understanding-rest-apis",
		},
		{
			Id:       2,
			Title:    "C++ Memory Management Tips",
			Excerpt:  "Learn how to handle memory safely in C++ with smart pointers and best practices.",
			Date:     "January 18, 2025",
			ReadTime: "8 min read",
			Slug:     "cpp-memory-management-tips",
		},
		{
			Id:       3,
			Title:    "Getting Started with Spring Boot",
			Excerpt:  "A concise introduction to building microservices with Spring Boot for Java developers.",
			Date:     "January 22, 2025",
			ReadTime: "7 min read",
			Slug:     "getting-started-with-spring-boot",
		},
		{
			Id:       4,
			Title:    "Graph Theory in Machine Learning",
			Excerpt:  "How graph theory concepts are applied in ML models for brain networks and social data.",
			Date:     "January 27, 2025",
			ReadTime: "10 min read",
			Slug:     "graph-theory-in-ml",
		},
		{
			Id:       5,
			Title:    "Optimizing SQL Queries",
			Excerpt:  "Practical strategies to make your SQL queries faster and more efficient.",
			Date:     "February 2, 2025",
			ReadTime: "5 min read",
			Slug:     "optimizing-sql-queries",
		},
	}

	return posts

}
