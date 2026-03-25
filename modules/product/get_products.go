package product

import (
	"ecommerce/util"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var count int

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 50
	}

	list, err := h.productService.List(page, limit)

	if err != nil {
		util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// count, err := h.productService.Count()
	// if err != nil {
	// 	util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		defer wg.Done()

		time.Sleep(7 * time.Second)
		count1, err := h.productService.Count()
		if err != nil {
			util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		count = count1
	}()

	go func() {
		defer wg.Done()

		time.Sleep(7 * time.Second)
		count2, err := h.productService.Count()
		if err != nil {
			util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Println(count2)
	}()

	go func() {
		defer wg.Done()

		time.Sleep(7 * time.Second)
		count3, err := h.productService.Count()
		if err != nil {
			util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Println(count3)
	}()

	go func() {
		defer wg.Done()

		time.Sleep(7 * time.Second)
		count4, err := h.productService.Count()
		if err != nil {
			util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Println(count4)

	}()

	wg.Wait()

	util.SendPaginatedData(w, r, list, page, limit, count)
}
