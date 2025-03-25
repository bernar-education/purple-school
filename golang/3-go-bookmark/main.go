package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	bookmarks := make(bookmarkMap, 3)
	fmt.Println("___ Bookmark manager ___")
	Menu: for {
		choice := getMenu()
		switch choice {
			case 1:
				printBookmarks(bookmarks)
			case 2:
				addBookmark(bookmarks)
			case 3:
				removeBookmark(bookmarks)
			case 4:
				break Menu
		}
	}
}

func getMenu() (choice int) {
	fmt.Println("Choose action: ")
	fmt.Println("1. Watch bookmarks")
	fmt.Println("2. Add bookmark")
	fmt.Println("3. Delete bookmark")
	fmt.Println("4. Exit")
	fmt.Scan(&choice)
	return
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Bookmarks are empty")
	}
	for key, value := range bookmarks {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkName string
	var newBookmarkUrl string
	fmt.Print("Input bookmark name: ")
	fmt.Scan(&newBookmarkName)
	fmt.Print("Add bookmark url: ")
	fmt.Scan(&newBookmarkUrl)
	bookmarks[newBookmarkName] = newBookmarkUrl
	fmt.Println("Bookmark added")
}

func removeBookmark(bookmarks bookmarkMap) {
	var bookmarkKeyToDelete string
	fmt.Print("Which bookmark you want to delete: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
	fmt.Println("Bookmark deleted")
}
