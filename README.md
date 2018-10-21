# Alexa Skill: Apple Buyer's Guide

This repository contains a lambda function called by the Amazon Alexa skill called **Apple Buyer's Guide**. When executed the lambda function scrapes https://buyersguide.macrumors.com/ to extract MacRumor's recommendation for each Apple product (updated, neutral, caution, and outdated). It then allows Alexa to query this data through questions like, "Is now a good time to buy the iMac?"

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
