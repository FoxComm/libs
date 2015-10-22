package seeds

import (
	"os"
	"strconv"
)

func StoreID() int {
	storeIDStr := os.Getenv("StoreID")
	if storeIDStr == "" {
		panic("Forgot to set env['StoreID']?")
	}
	storeID, _ := strconv.Atoi(storeIDStr)
	return storeID
}
