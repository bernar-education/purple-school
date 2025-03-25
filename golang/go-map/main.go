package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	bookmarks := make(bookmarkMap, 3)
	fmt.Println("___ Книга закладок ___")
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
	fmt.Println("Выберите действие: ")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&choice)
	return
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Закладок нет")
	}
	for key, value := range bookmarks {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkName string
	var newBookmarkUrl string
	fmt.Print("Введите название: ")
	fmt.Scan(&newBookmarkName)
	fmt.Print("Введите ссылку: ")
	fmt.Scan(&newBookmarkUrl)
	bookmarks[newBookmarkName] = newBookmarkUrl
	fmt.Println("Закладка добавлена")
}

func removeBookmark(bookmarks bookmarkMap) {
	var bookmarkKeyToDelete string
	fmt.Print("Какую закладку удалить: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
	fmt.Println("Закладка удалена")
}
