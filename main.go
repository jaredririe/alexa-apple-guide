package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jaredririe/alexa-apple-guide/alexa"
	"github.com/jaredririe/alexa-apple-guide/scraper"
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
		response = alexa.NewResponse(
			"Apple Guide (Unofficial)",
			"Welcome to the Unofficial Apple Guide. You can ask me whether it's a good time to buy a particular Apple product. For example, you could ask 'is now a good time to buy the iMac?'",
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

	unknownProductResponse := alexa.NewResponse(
		"Unknown product",
		fmt.Sprintf("I'm not aware of the %s.", product),
		false,
	)

	lowercaseProduct := strings.ToLower(product)
	status, ok := nameToStatus[lowercaseProduct]
	if !ok {
		noSpaces := strings.Replace(lowercaseProduct, " ", "", -1)
		status, ok = nameToStatus[noSpaces]
		if !ok {
			return unknownProductResponse
		}
	}

	switch status {
	case scraper.Status.Updated:
		return alexa.NewResponse(
			"Just updated!",
			fmt.Sprintf("Now is a great time to buy the %s! It was recently updated.", product),
			true,
		)

	case scraper.Status.Neutral:
		return alexa.NewResponse(
			"Neutral",
			fmt.Sprintf("I'm neutral on the %s. It is in the middle of its usual release cycle.", product),
			true,
		)

	case scraper.Status.Caution:
		return alexa.NewResponse(
			"Caution",
			fmt.Sprintf("I would use caution in purchasing the %s. It has been a long time since it was updated.", product),
			true,
		)

	case scraper.Status.Outdated:
		return alexa.NewResponse(
			"Outdated",
			fmt.Sprintf("It's probably not a good idea to buy the %s. It is outdated. A new version may be released soon.", product),
			true,
		)

	case scraper.Status.Unknown:
	}

	return unknownProductResponse
}

func handleHelp() alexa.Response {
	return alexa.NewResponse(
		"Help",
		fmt.Sprintf("You can ask me whether it's a good time to buy a particular Apple product. For example, you could ask 'is now a good time to buy the iMac?' You can also simply say the name of a product. Here are some of the Apple products I know about: %s. What product are you interested in?",
			strings.Join(knownProducts[:8], ", ")),
		false,
	)
}

func handleStop() alexa.Response {
	return alexa.NewResponse(
		"Bye!",
		"Best of luck!",
		true,
	)
}

func handleFallback() alexa.Response {
	return alexa.NewResponse(
		"I don't quite understand",
		"I can't help you with that. Try rephrasing your question or ask for help by saying 'help'",
		false,
	)
}
