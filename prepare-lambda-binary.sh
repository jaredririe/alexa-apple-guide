set -x

rm -rf alexa-apple-guide.zip
GOOS=linux go build -o alexa-apple-guide -a -ldflags "-w -s -extldflags \"-static\"" -installsuffix cgo
zip -r alexa-apple-guide.zip alexa-apple-guide
