# Alexa Skill: Picking Apples (Unofficial)

This repository contains a lambda function called by the Amazon Alexa skill called **Picking Apples (Unofficial)**. When executed, the lambda function scrapes https://buyersguide.macrumors.com/ to extract MacRumors' recommendation for each Apple product (updated, neutral, caution, and outdated). It then allows Alexa to query this data through questions like, "Is now a good time to buy the iMac?"

## Skill Description

### One-sentence

This skill helps you know whether it's a good time to buy a new Apple product by checking the MacRumors Buyer's Guide.

### Detailed

Note: this is an unofficial skill and not affiliated with or sponsored by Apple or MacRumors.

Picking Apples (Unofficial) offers a convenient way to check whether it's a good time to buy a new Apple product. Through a real-time look at the MacRumors Buyer's Guide, this skill allows Alexa to tell you whether a product is recently updated, in the middle of its release cycle, somewhat out of date, or clearly outdated.

The MacRumors Buyer's Guide is located at https://buyersguide.macrumors.com/.

### Example Phrases

* Alexa, launch Picking Apples.
    - Is now a good time to buy the airpods?
    - iMac Pro.
    - Should I buy the Apple TV?
* Alexa, ask Picking Apples is now a good time to buy the iPhone XR?
* ask picking apples to tell me if it's a good time to buy the iphone xr

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
