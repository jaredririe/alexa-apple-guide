package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// BuyersGuideScraper scrapes the MacRumors' Buyers Guide.
type BuyersGuideScraper struct {
	collector *colly.Collector
}

// StatusEnum enumerates all the possible status values.
type StatusEnum string

var Status = struct {
	Updated  StatusEnum
	Neutral  StatusEnum
	Caution  StatusEnum
	Outdated StatusEnum
	Unknown  StatusEnum
}{
	"Updated",
	"Neutral",
	"Caution",
	"Outdated",
	"Unknown",
}

func NewBuyersGuideScraper() *BuyersGuideScraper {
	// Instantiate default collector
	collector := colly.NewCollector(
		colly.AllowedDomains("buyersguide.macrumors.com"),
	)

	return &BuyersGuideScraper{
		collector: collector,
	}
}

func (bgs *BuyersGuideScraper) Scrape() map[string]StatusEnum {

	idToName := make(map[string]string)
	idToStatus := make(map[string]StatusEnum)

	// Check all <a> tags for a name attribute that starts with the
	// # element selector. These tags include a title attribute that
	// will be useful later.
	//
	// Example: <a name="#iphone" title="iPhone"></a>
	// (iphone -> iPhone will be used later)
	bgs.collector.OnHTML("a", func(e *colly.HTMLElement) {
		name := e.Attr("name")
		title := e.Attr("title")

		if strings.HasPrefix(name, "#") {
			idToName[name[1:]] = title
		}
	})

	// <a> tags immediately followed by <div> tags are used for each
	// product MacRumors tracks (iPhone, iMac Pro, AirPods, etc.).
	// We extract the id attribute of the <div> tag and search for
	// child <div> tags with a class attribute matching "status".
	//
	// Example:
	// <a name="#iphone" title="iPhone"></a>
	// <div id="iphone">
	// ...
	//     <div class="status updated">...</div>
	// ...
	// </div>
	// (iphone -> "status updated" will be used later)
	bgs.collector.OnHTML("a + div", func(e *colly.HTMLElement) {
		id := e.Attr("id")

		divClasses := e.ChildAttrs("div", "class")
		for _, divClass := range divClasses {
			if strings.Contains(divClass, "status") {
				idToStatus[id] = parseStatus(divClass)
			}
		}
	})

	// Start scraping the Buyer's Guide
	bgs.collector.Visit("https://buyersguide.macrumors.com")

	fmt.Printf("%#v\n%#v\n", idToName, idToStatus)

	return produceNameToStatus(idToName, idToStatus)
}

func produceNameToStatus(idToName map[string]string, idToStatus map[string]StatusEnum) map[string]StatusEnum {
	nameToStatus := make(map[string]StatusEnum)

	for id, name := range idToName {
		stat := idToStatus[id]

		// fix special case where " was not escaped on MacRumor's Buyer's Guide,
		// causing the incorrect value to be captured
		if name == "10.5" {
			name = "iPad Pro"
		} else if name == "12.9" {
			name = "iPad Pro"
		}

		nameToStatus[strings.ToLower(name)] = stat
	}

	return nameToStatus
}

func parseStatus(stat string) StatusEnum {
	switch stat {
	case "status updated":
		return Status.Updated
	case "status":
		return Status.Neutral
	case "status caution":
		return Status.Caution
	case "status outdated":
		return Status.Outdated
	default:
		return Status.Unknown
	}
}
