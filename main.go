package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jaredririe/alexa-applebuyersguide/alexa"
	"github.com/jaredririe/alexa-applebuyersguide/scraper"
)

var (
	nameToStatus  map[string]scraper.StatusEnum
	knownProducts []string
)

func main() {
	bgs := scraper.NewBuyersGuideScraper()
	nameToStatus = bgs.Scrape()
	for name, status := range nameToStatus {
		fmt.Println("Name:", name, "Status:", status)

		knownProducts = append(knownProducts, name)
	}

	// sort list of products for consistency
	sort.Strings(knownProducts)

	lambda.Start(handler)
}

// handler handles requests from AWS Lambda.
func handler(request alexa.Request) (alexa.Response, error) {
	var response alexa.Response

	switch request.Body.Type {
	case alexa.LaunchRequestType:
		response = alexa.NewSimpleResponse(
			"Apple Buyer's Guide",
			"Welcome to Apple Buyer's Guide with data from MacRumors. You can ask me whether it's a good time to buy a particular Apple product. For example, you could ask 'is now a good time to buy the iMac?'",
			false,
		)
	case alexa.IntentRequestType:
		response = dispatchIntents(request)
	case alexa.SessionEndedRequestType:
	}

	return response, nil
}

// dispatchIntents dispatches each intent to the right handler
func dispatchIntents(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "productRecommendation":
		response = handleRecommendation(request)
	case alexa.HelpIntent:
		response = handleHelp()
	case alexa.CancelIntent, alexa.NoIntent, alexa.StopIntent:
		response = handleStop()
	case alexa.FallbackIntent:
		response = handleFallback()
	}

	return response
}

func handleRecommendation(request alexa.Request) alexa.Response {
	product := request.Body.Intent.Slots["product"].Value

	unknownProductResponse := alexa.NewSimpleResponse(
		"Unknown product",
		fmt.Sprintf("I'm not aware of the %s.", product),
		false,
	)

	status, ok := nameToStatus[strings.ToLower(product)]
	if !ok {
		return unknownProductResponse
	}

	switch status {
	case scraper.Status.Updated:
		return alexa.NewSimpleResponse(
			"Just updated!",
			fmt.Sprintf("Now is a great time to buy the %s! It was recently updated.", product),
			false,
		)

	case scraper.Status.Neutral:
		return alexa.NewSimpleResponse(
			"Neutral",
			fmt.Sprintf("I'm neutral on the %s. It is in the middle of its usual release cycle.", product),
			false,
		)

	case scraper.Status.Caution:
		return alexa.NewSimpleResponse(
			"Caution",
			fmt.Sprintf("I would use caution in purchasing the %s. It has been a long time since it was updated.", product),
			false,
		)

	case scraper.Status.Outdated:
		return alexa.NewSimpleResponse(
			"Outdated",
			fmt.Sprintf("It's probably not a good idea to buy the %s. It is outdated. A new version may be released soon.", product),
			false,
		)

	case scraper.Status.Unknown:
	}

	return unknownProductResponse
}

func handleHelp() alexa.Response {
	return alexa.NewSimpleResponse(
		"Help",
		fmt.Sprintf("You can ask me whether it's a good time to buy a particular Apple product. For example, you could ask 'is now a good time to buy the iMac?' You can also simply say the name of a product. I'm aware of the following Apple products: %s.",
			strings.Join(knownProducts, ", ")),
		false,
	)
}

func handleStop() alexa.Response {
	return alexa.NewSimpleResponse(
		"Bye!",
		"Best of luck!",
		true,
	)
}

func handleFallback() alexa.Response {
	return alexa.NewSimpleResponse(
		"I don't quite understand",
		"I can't help you with that. Try rephrasing your question or ask for help by saying 'help'",
		false,
	)
}
