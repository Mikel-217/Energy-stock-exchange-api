package handlers

import "net/http"

// TODO:
// implement that this handler can serve following stuff (for DateAndPrice stuff)
// - serve all prices for a date ?all(=date) -> because its optional, if there is no date we use the current date
// - serve prices for a date range ?start=date1&end=date2
func HandlePriceRequests(w http.ResponseWriter, r *http.Request) {

}
