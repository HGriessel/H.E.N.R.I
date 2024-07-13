package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Category struct {
	Type         string `json:"type"`
	RootDir      string `json:"root_dir"`
	noteTemplate func(*Note) string
}

type Note struct {
	Path        string   `json:"path"`
	Title       string   `json:"title"`
	Category    Category `json:"category"`
	SubCategory string   `json:"sub_category"`
}

func (n *Note) WriteNote() error {

	content := []byte(n.Category.noteTemplate(n))
	log.Printf("Checking if %v exists\n", n.GetSubCatPath())
	if _, err := os.Stat(n.GetSubCatPath()); errors.Is(err, os.ErrNotExist) {

		log.Printf("Creating %v\n", n.GetSubCatPath())
		errWr := os.Mkdir(n.GetSubCatPath(), os.ModePerm)

		if errWr != nil {
			return fmt.Errorf("Create sub dir failed with %v", errWr)
		}

	}
	log.Printf("Checking if %v exists\n", n.Path)
	if _, err := os.Stat(n.Path); errors.Is(err, os.ErrNotExist) {

		log.Printf("Creating %v\n", n.Path)
		errWr := os.WriteFile(n.Path, content, os.ModePerm)

		if errWr != nil {
			return errWr
		}
	} else {
		return fmt.Errorf("File: %v already exists", n.Path)
	}
	return nil
}

func (n *Note) GetSubCatPath() string {
	return fmt.Sprintf("%v/%v/", n.Category.RootDir, n.SubCategory)
}

func GetProjectTemplate(n *Note) string {
	current_time := time.Now().Local()
	return fmt.Sprintf("#%v\n- Created:%v\n- Category:%v\n- SubCatergory:%v\n\n\n", n.Title, current_time, n.Category.Type, n.SubCategory)
}

func GetArchivesTemplate(n *Note) string {
	current_time := time.Now().Local()
	return fmt.Sprintf("#%v\n- Created:%v\n- Category:%v\n- SubCatergory:%v\n", n.Title, current_time, n.Category.Type, n.SubCategory)
}

func GetResourcesTemplate(n *Note) string {
	current_time := time.Now().Local()
	return fmt.Sprintf("#%v\n- Created:%v\n- Category:%v\n- SubCatergory:%v\n", n.Title, current_time, n.Category.Type, n.SubCategory)
}

func GetAreasTemplate(n *Note) string {
	current_time := time.Now().Local()
	return fmt.Sprintf("#%v\n- Created:%v\n- Category:%v\n- SubCatergory:%v\n", n.Title, current_time, n.Category.Type, n.SubCategory)
}
func GetInboxTemplate(n *Note) string {
	current_time := time.Now().Local()
	return fmt.Sprintf("#%v\n- Created:%v\n- Category:%v\n- SubCatergory:%v\n", n.Title, current_time, n.Category.Type, n.SubCategory)
}
