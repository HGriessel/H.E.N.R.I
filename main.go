package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var Project = Category{
	Type:         "Project",
	RootDir:      "/home/henri/Notes/Projects",
	noteTemplate: GetProjectTemplate,
}

var Area = Category{
	Type:         "Areas",
	RootDir:      "/home/henri/Notes/Areass",
	noteTemplate: GetAreasTemplate,
}
var Resource = Category{
	Type:         "Resources",
	RootDir:      "/home/henri/Notes/Resourcess",
	noteTemplate: GetResourcesTemplate,
}

var Archive = Category{
	Type:         "Project",
	RootDir:      "/home/henri/Notes/Projects",
	noteTemplate: GetProjectTemplate,
}

var Inbox = Category{
	Type:         "Inbox",
	RootDir:      "/home/henri/Notes/Inbox",
	noteTemplate: GetInboxTemplate,
}

var catArg string
var titleArg string
var subCat string
var helpArg string

func main() {
	flag.StringVar(
		&catArg,
		"h",
		"i",
		fmt.Sprintf("must specify title -t and sub category -s type of category will default to inbox or specify  \n\t- i = Inbox\n\t- p = Project\n\t- a = Area\n\t- r=Resource\n\t- z=archive "),
	)

	flag.StringVar(
		&catArg,
		"c",
		"i",
		fmt.Sprintf("the type of note defaults to Inbox\n\t- i = Inbox\n\t- p = Project\n\t- a = Area\n\t- r=Resource\n\t- z=archive"),
	)
	flag.StringVar(
		&titleArg,
		"t",
		"",
		"title of note",
	)
	flag.StringVar(
		&subCat,
		"s",
		"",
		"name of SubCatergory",
	)
	flag.Parse()
	if titleArg == "" || subCat == "" {
		log.Fatalf("Please use specify both title and subcat")
	}

	current_time := time.Now().Local()
	subCat := subCat
	title := titleArg

	note := Note{
		Category:    determineCategory(catArg),
		SubCategory: subCat,
		Title:       title,
		Path:        fmt.Sprintf("%v/%v/%v-%v.md", Project.RootDir, subCat, title, current_time.Format("2006-01-02")),
	}
	err := note.WriteNote()
	handleErr(err)
}

func determineCategory(t string) Category {
	switch t {
	case "i":
		return Inbox
	case "p":
		return Project
	case "a":
		return Area
	case "z":
		return Archive
	case "r":
		return Resource
	default:
		return Inbox
	}
}

func handleErr(err error) {
	if err != nil {
		log.Printf("%v\n", err)
	}
}
