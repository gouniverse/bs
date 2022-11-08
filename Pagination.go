package bs

import (
	"log"
	"math"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

type PaginationOptions struct {
	NumberItems       int
	CurrentPageNumber int
	PerPage           int
	PagesToShow       int
	URL               string
	FirstPageStartsAt int
}

func Pagination(options PaginationOptions) string {
	min, max, previousPageNumber, nextPageNumber := paginationMinMaxPrevNext(options.NumberItems, options.CurrentPageNumber, options.PerPage, options.PagesToShow, options.FirstPageStartsAt)

	if max < 2 {
		return ""
	}

	liStart := hb.NewLI().Class("page-item").Children([]*hb.Tag{
		hb.NewHyperlink().Class("page-link").Style("cursor:pointer;").Attr("href", options.URL+utils.ToString(previousPageNumber)).HTML("&laquo;"),
	})

	liEnd := hb.NewLI().Class("page-item").Children([]*hb.Tag{
		hb.NewHyperlink().Class("page-link").Style("cursor:pointer;").Attr("href", options.URL+utils.ToString(nextPageNumber)).HTML("&raquo;"),
	})

	pages := []*hb.Tag{}

	if previousPageNumber >= options.FirstPageStartsAt {
		pages = append(pages, liStart)
	}

	for i := min; i < max; i++ {
		active := ""
		if i == options.CurrentPageNumber {
			active = " active"
		}

		li := hb.NewLI().Class("page-item" + active).Children([]*hb.Tag{
			hb.NewHyperlink().Class("page-link").Style("cursor:pointer;").Attr("href", options.URL+utils.ToString(i)).HTML(utils.ToString(i + 1)),
		})

		pages = append(pages, li)
	}

	if nextPageNumber < max {
		pages = append(pages, liEnd)
	}

	nav := hb.NewNav().Children([]*hb.Tag{
		hb.NewUL().Class("pagination").Children(pages),
	})

	return nav.ToHTML()
}

func paginationMinMaxPrevNext(numberItems int, currentPageNumber int, perPage int, pagesToShow int, firstPageStartsAt int) (min int, max int, previousPageNumber int, nextPageNumber int) {
	previousPageNumber = currentPageNumber - 1
	nextPageNumber = currentPageNumber + 1

	numberOfPages := int(math.Ceil(float64(numberItems) / float64(perPage)))
	// adjacent := int(math.Floor(((float64(currentPageNumber) - 3) / 2)))

	min = currentPageNumber - pagesToShow/2
	if pagesToShow%2 == 0 {
		min++ // if even number
	}

	if min < firstPageStartsAt {
		min = firstPageStartsAt
	}

	max = min + pagesToShow - 1

	if max > numberOfPages {
		max = numberOfPages

		// too little pages, pad on the left
		min = max - pagesToShow
		if min < firstPageStartsAt {
			min = firstPageStartsAt
		}
	}

	return min, max, previousPageNumber, nextPageNumber
}
