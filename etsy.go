package etsy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Image struct {
	Url_75x75     string
	Url_170x135   string
	Url_570xN     string
	Url_fullxfull string
	Full_height   int
	Full_width    int
}

type Listing struct {
	Url         string
	Listing_id  int
	State       string
	User_id     int
	Title       string
	Description string
	Images      []Image
}

type Listings struct {
	Count   int
	Results []Listing
}

type Etsy struct {
	apiKey string
}

func New(apiKey string) *Etsy {
	return &Etsy{
		apiKey: apiKey,
	}
}

func (e *Etsy) authenticate(url string) string {
	return url + "&api_key=" + e.apiKey
}

var api_url = "https://openapi.etsy.com/v2/"

func (e *Etsy) GetStoreListings(storeID string, limit int) (Listings, error) {
	url := api_url + fmt.Sprintf("shops/%s/listings/active?fields=title,description,url&includes=Images&limit=%d", storeID, limit)
	url = e.authenticate(url)
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return Listings{}, err
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return Listings{}, err
	}

	// marshal into structs
	var listings Listings
	json.Unmarshal([]byte(result), &listings)

	return listings, nil
}
