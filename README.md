# Alexa Skill: Buyer's Guide for Apple

This repository contains a lambda function called by the Amazon Alexa skill called **Buyer's Guide for Apple**. When executed, the lambda function scrapes https://buyersguide.macrumors.com/ to extract MacRumors' recommendation for each Apple product (updated, neutral, caution, and outdated). It then allows Alexa to query this data through questions like, "Is now a good time to buy the iMac?"

## Skill Description

One-sentence: This skill helps you know whether it's a good time to buy a new Apple product by checking the MacRumors Buyer's Guide.

Detailed: The Buyer's Guide for Apple skill offers a convenient way to check whether it's a good time to buy a new Apple product. Through a real-time look at the [MacRumors Buyer's Guide](https://buyersguide.macrumors.com/), this skill allows Alexa to tell you which of four states an Apple product is in: updated, neutral, caution, and outdated. A product in the caution state, for example, has not been updated for quite some time, so it may be wise to be patient and wait for a new update. The status updated, on the other hand, means that the Apple product was just updated and you're safe to go ahead with the purchase.

### Example Phrases

* Alexa, launch Buyer's Guide for Apple.
    - Is now a good time to buy the airpods?
    - iMac Pro.
    - Should I buy the Apple TV?
* Ask Buyer's Guide for Apple is now a good time to buy the iMac?

## Running locally

```
$ go build && ./alexa-applebuyersguide

Name: imac pro Status: Neutral
Name: apple watch Status: Updated
Name: ipod touch Status: Caution
Name: iphone xr Status: Updated
Name: ipad Status: Neutral
Name: imac Status: Outdated
Name: airpods Status: Caution
Name: macbook air Status: Outdated
Name: macbook Status: Outdated
Name: mac pro Status: Outdated
Name: iphone xs Status: Updated
Name: mac mini Status: Outdated
Name: ipad mini Status: Outdated
Name: apple tv Status: Caution
Name: ipad pro Status: Outdated
Name: homepod Status: Neutral
Name: macbook pro Status: Updated
```

## Preparing the binary to upload to the Lambda Management Console

Running the `prepare-lambda-binary` script will create a binary for the Linux architecture (with `GOOS=linux`) and zip it up so it can be uploaded to the Lambda Management Console.

```
$ ./prepare-lambda-binary
```
