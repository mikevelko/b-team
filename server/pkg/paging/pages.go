package paging

/*
This package's purpose is to provide helper functions for paging items in a slice
*/

// getPagesCount counts how many pages is necessary to hold items
func getPagesCount(items int, itemsPerPage int) int {
	pagesNumber := items / itemsPerPage
	if items%itemsPerPage != 0 {
		pagesNumber++
	}
	return pagesNumber
}

// findStart finds first item in required page
func findStart(currPage int, itemsPerPage int) int {
	return (currPage - 1) * itemsPerPage
}

// findEnd finds first item in required page
func findEnd(currPage int, itemsPerPage int, pagesCount int, itemsCount int) int {
	if currPage == pagesCount {
		return itemsCount
	}
	return ((currPage - 1) * itemsPerPage) + itemsPerPage
}

// GetPageItems returns start and end of range of elements that should end on picked page
func GetPageItems(currPage int, itemsPerPage int, itemsCount int) (int, int) {
	if itemsCount == 0 || itemsPerPage == 0 {
		return 0, 0
	}
	pagesCount := getPagesCount(itemsCount, itemsPerPage)
	if currPage < 1 {
		currPage = 1
	} else if currPage > pagesCount {
		currPage = pagesCount
	}
	return findStart(currPage, itemsPerPage), findEnd(currPage, itemsPerPage, pagesCount, itemsCount)
}
