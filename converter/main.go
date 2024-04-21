package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kokoichi206/searcher/converter/model"
	"github.com/kokoichi206/searcher/converter/os"
	"github.com/kokoichi206/searcher/converter/rdb"
)

func main() {
	db, err := rdb.New()
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	ctx := context.Background()

	members, err := db.SelectAllMembers(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client, err := os.New()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	for _, m := range members {
		fmt.Printf("m.Name: %v\n", m.Name)

		// if err := client.IndexMember(
		// 	&model.Member{
		// 		Birthday:   m.Birthday,
		// 		BlogUrl:    m.Link,
		// 		BloodType:  m.Blood,
		// 		Generation: m.Cate,
		// 		ImgUrl:     m.Img,
		// 		Name:       m.Name,
		// 	},
		// ); err != nil {
		// 	log.Fatalf("failed to index member: %v", err)
		// }

		fmt.Printf("m.Code: %v\n", m.Code)

		blogs, err := db.SelectAllBlogs(ctx, m.Code)
		if err != nil {
			log.Fatalf("failed to select all blogs: %v", err)
		}

		fmt.Printf("len(blogs): %v\n", len(blogs))

		for _, b := range blogs {
			if err := client.IndexBlog(
				&model.Blog{
					ArtiCode:  b.Code,
					Cate:      b.Cate,
					Code:      b.Code,
					Date:      b.Timestamp,
					EndTime:   b.EndTime,
					Link:      b.Link,
					Member:    m.Name,
					StartTime: b.StartTime,
					Text:      b.Content,
					Title:     b.Title,
				},
			); err != nil {
				log.Fatalf("failed to index blog: %v", err)
			}
		}
	}
}
