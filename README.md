# Alexa Skill: Apple Guide (Unofficial)

This repository contains a lambda function called by the Amazon Alexa skill called **Apple Guide (Unofficial)**. When executed, the lambda function scrapes https://buyersguide.macrumors.com/ to extract MacRumors' recommendation for each Apple product (updated, neutral, caution, and outdated). It then allows Alexa to query this data through questions like, "Is now a good time to buy the iMac?"

## Skill Description

### One-sentence

This (unofficial) skill helps you know whether it's a good time to buy a new Apple product by checking the MacRumors Buyer's Guide.

### Detailed

Note: this is an unofficial skill and not affiliated with or sponsored by Apple or MacRumors.

Apple Guide (Unofficial) offers a convenient way to check whether it's a good time to buy a new Apple product. Through a real-time look at the MacRumors Buyer's Guide, this skill allows Alexa to tell you whether a product is recently updated, in the middle of its release cycle, somewhat out of date, or clearly outdated.

The MacRumors Buyer's Guide is located at https://buyersguide.macrumors.com/.

### Example Phrases

* Alexa, launch Unofficial Apple Guide.
    - Is now a good time to buy the Airpods?
    - iMac Pro.
    - Should I buy the Apple TV?
* Alexa, ask Unofficial Apple Guide is now a good time to buy the iPhone XR?
* Alexa, ask Unofficial Apple Guide to tell me if it's a good time to buy the Mac Mini

## Running locally

```
$ go build && ./alexa-apple-guide

Name: imac Status: Outdated
Name: ipad mini Status: Outdated
Name: ipad pro Status: Updated
Name: iphone xr Status: Updated
Name: homepod Status: Neutral
Name: macbook Status: Outdated
Name: ipod touch Status: Caution
Name: ipad Status: Neutral
Name: macbook air Status: Updated
Name: iphone xs Status: Updated
Name: apple tv Status: Caution
Name: imac pro Status: Neutral
Name: mac pro Status: Outdated
Name: mac mini Status: Updated
Name: airpods Status: Caution
Name: apple watch Status: Updated
Name: macbook pro Status: Neutral
```

## Preparing the binary to upload to the Lambda Management Console

Running the `prepare-lambda-binary` script will create a binary for the Linux architecture (with `GOOS=linux`) and zip it up so it can be uploaded to the Lambda Management Console.

```
$ ./prepare-lambda-binary
```
