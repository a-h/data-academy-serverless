package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
)

var flagAction = flag.String("action", "", "Set the action - 'print' to print JSON to the console, 'send' to send messages to EventBridge.")
var flagN = flag.Int("n", 100, "Number of records to generate.")
var flagRegion = flag.String("region", "eu-west-2", "Set the AWS region.")

func main() {
	flag.Parse()

	data := generateRecords(*flagN)

	if *flagAction == "print" {
		for _, d := range data {
			j, err := json.Marshal(d)
			if err != nil {
				fmt.Println("error marshalling", err)
				os.Exit(1)
			}
			fmt.Println(string(j))
		}
		return
	}
	if *flagAction == "send" {
		fmt.Printf("Sending %d Transactions to EventBridge...\n", *flagN)
		sess, err := session.NewSession(aws.NewConfig().WithRegion(*flagRegion))
		if err != nil {
			fmt.Println("could not create AWS session", err)
			os.Exit(1)
		}
		eb := eventbridge.New(sess)

		var i int
		for i = 0; i < len(data); i += 10 {
			max := i + 10
			if max > len(data) {
				max = len(data)
			}
			batch := data[i:max]

			entries := make([]*eventbridge.PutEventsRequestEntry, len(batch))
			for j := 0; j < len(batch); j++ {
				s, err := json.Marshal(batch[j])
				if err != nil {
					fmt.Println("error marshalling", err)
				}
				entries[j] = &eventbridge.PutEventsRequestEntry{
					Source:     aws.String("Transactions"),
					Detail:     aws.String(string(s)),
					DetailType: aws.String("transaction"),
				}
			}
			_, err = eb.PutEvents(&eventbridge.PutEventsInput{Entries: entries})
			if err != nil {
				fmt.Println("failed to put events", err)
				os.Exit(1)
			}
			fmt.Printf("Sent %d messages...\n", max)
		}
		return
	}
	flag.PrintDefaults()
}

func generateRecords(n int) []Data {
	var data []Data
	for i := 0; i < n; i++ {
		basket, total := randomBasket()
		data = append(data, Data{
			ID:            fmt.Sprintf("%d", i),
			Date:          randomTime(today, time.Now()),
			Location:      locations[rnd.Intn(len(locations))],
			CustomerName:  names[rnd.Intn(len(names))],
			PaymentMethod: paymentMethods[rnd.Intn(len(paymentMethods))],
			Basket:        basket,
			Total:         total,
		})
	}
	return data
}

func randomBasket() (b []BasketItem, total int) {
	itemCount := rnd.Intn(10) + 1
	for i := 0; i < itemCount; i++ {
		quantity := rnd.Intn(5) + 1
		item := randomItem()
		bi := BasketItem{
			Item:     item,
			Quantity: quantity,
		}
		total += item.Cost * quantity
		b = append(b, bi)
	}
	return
}

func randomItem() Item {
	itemIndex := rnd.Intn(len(products))
	return Item{
		ID:   fmt.Sprintf("item%d", itemIndex),
		Cost: rnd.Intn(6000) + 59,
		Name: products[itemIndex],
	}
}

var locations = []string{"Leeds", "London", "Manchester", "Edinburgh"}

var names = []string{
	"Cher Cluck",
	"Hattie Eagles",
	"Tamara Modisette",
	"Rosamond Alvares",
	"Catrice Sprenger",
	"Miyoko Bulger",
	"Marjorie Kole",
	"Cecille Loop",
	"Ellsworth Delpozo",
	"Venessa Retherford",
	"Pauletta Drake",
	"Jefferson Langlinais",
	"Syble Swartz",
	"Clarence Justiniano",
	"Ilana Tracy",
	"Tami Auger",
	"Tanisha Rudder",
	"Carlton Shadle",
	"Gracia Pontes",
	"Debbie Millener",
	"Marisha Mascarenas",
	"Tameka Schmoll",
	"Debora Castro",
	"Angele Betancourt",
	"Santina Ruple",
	"Brian Taubman",
	"Collene Bou",
	"Kira Reagl",
	"Leticia Weiner",
	"Tanya Fro",
}

var products = []string{
	"Consomme printaniere royal",
	"Chicken gumbo",
	"Tomato aux croutons",
	"Onion au gratin",
	"St. Emilion",
	"Radishes",
	"Chicken soup with rice",
	"Clam broth (cup)",
	"Cream of new asparagus, croutons",
	"Clear green turtle",
	"Striped bass saute, meuniere",
	"Anchovies",
	"Fresh lobsters in every style",
	"Celery",
	"Pim-olas",
	"Caviar",
	"Sardines",
	"India chutney",
	"Pickles",
	"English walnuts",
	"Pate de foies-gras",
	"Pomard",
	"Brook trout, mountain style",
	"Whitebait, sauce tartare",
	"Clams",
	"Oysters",
	"Claremont planked shad",
	"G. H. Mumm & Co's Extra Dry",
	"Cerealine with Milk",
	"Sliced Bananas",
	"Wheat Vitos",
	"Sliced Tomatoes",
	"Russian Caviare on Toast",
	"Smoked beef in cream",
	"Apple Sauce",
	"Potage a la Victoria",
	"Breakfast",
	"Strawberries",
	"Preserved figs",
	"BLUE POINTS",
	"CONSOMME ANGLAISE",
	"CREAM OF CAULIFLOWER",
	"BROILED SHAD, A LA MAITRE D'HOTEL",
	"SLICED CUCUMBERS",
	"SALTED ALMONDS",
	"POTATOES, JULIEN",
	"Cracked Wheat",
	"Malt Breakfast Food",
	"BOILED BEEF TONGUE, ITALIAN SAUCE",
	"Young Onions",
	"Pears",
	"ROAST SIRLOIN OF BEEF, YORKSHIRE PUDDING",
	"Huhnerbruhe",
	"ROAST EASTER LAMB, MINT SAUCE",
	"Rockaways",
	"Hafergrutze",
	"BROWNED POTATOES",
	"Pampelmuse",
	"Apfelsinen",
	"Ananas",
	"Milchreis",
	"Grape fruit",
	"Oranges",
	"Clam Fritters",
	"Filet v. Schildkrote m. Truffeln",
	"Bouillon, en Tasse",
	"Spargel Suppe",
	"Kraftsuppe, konigliche Art",
	"Rissoles a la Merrill",
	"S. Julien",
	"Chambertin",
	"St. Julien",
	"Vegetable",
	"Puree of split peas aux croutons",
	"Consomme in cup",
	"Broiled shad, Maitre d'hotel",
	"Mashed potatoes",
	"Breaded veal cutlet with peas",
	"Hind-quarter of spring lamb with stuffed tomatoes",
	"Hot or cold ribs of beef",
	"Doucette salad",
	"New beets",
	"Salisbury steak au cresson",
	"Boiled rice",
	"Stewed oyster plant",
	"Boiled onions, cream sauce",
	"Old fashioned rice pudding",
	"Ice cream",
	"Coffee",
	"Tea",
	"Milk",
	"Mush",
	"Rolled Oats",
	"Small Hominy",
	"Broiled Mackerel",
	"Kippered Herring",
	"Strawberries with cream",
	"Compote of fruits",
	"Orange marmalade",
}

var paymentMethods = []string{"card", "cash"}

var now = time.Now()
var today = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
var tomorrow = today.Add(time.Hour * 24)

var rnd = rand.New(rand.NewSource(time.Now().Unix()))

func randomTime(from, to time.Time) time.Time {
	max := to.Sub(from)
	return from.Add(time.Duration(rnd.Int63n(int64(max))))
}

type Data struct {
	ID            string       `json:"id"`
	Date          time.Time    `json:"date"`
	Location      string       `json:"location"`
	CustomerName  string       `json:"customerName"`
	Basket        []BasketItem `json:"basket"`
	Total         int          `json:"total"`
	PaymentMethod string       `json:"paymentMethod"`
}

type BasketItem struct {
	Item     Item `json:"item"`
	Quantity int  `json:"quantity"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
}
