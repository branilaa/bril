package tags

import (
	"flag"
	"fmt"
	"os"

	"github.com/branila/bril/db"
	"github.com/branila/bril/types"
)

type AddTagFlags struct {
	Priority int
}

func getAddFlags() AddTagFlags {
	var flags AddTagFlags

	fs := flag.NewFlagSet("add", flag.ExitOnError)

	fs.IntVar(&flags.Priority, "priority", -1, "Priority of the tag")
	fs.IntVar(&flags.Priority, "p", -1, "Priority of the tag")

	fs.Parse(os.Args[4:])

	return flags
}

func Add() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: bril tag add <tag>")
		return
	}

	tag := os.Args[3]

	flags := getAddFlags()

	tags := db.GetTags()

	for _, t := range tags {
		if t.Name == tag {
			fmt.Printf("Tag %s already exists\n", tag)
			return
		}
	}

	db.CreateTag(types.Tag{
		Name:     tag,
		Priority: flags.Priority,
	})

	fmt.Printf("Tag %s created\n", tag)
}

func List() {
	fmt.Println("Coming soon")
}

func Delete() {
	fmt.Println("Coming soon")
}

func Handle() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril tag <command> [args]")
		return
	}

	switch os.Args[2] {
	case "add":
		Add()
	case "ls", "list":
		List()
	case "rm", "delete":
		Delete()
	default:
		fmt.Println("Unknown command")
	}
}
