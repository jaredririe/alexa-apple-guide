set -x

rm -rf alexa-picking-apples.zip
GOOS=linux go build -o alexa-picking-apples -a -ldflags "-w -s -extldflags \"-static\"" -installsuffix cgo
zip -r alexa-picking-apples.zip alexa-picking-apples
