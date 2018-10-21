set -x

rm -rf alexa-applebuyersguide.zip
GOOS=linux go build -o alexa-applebuyersguide -a -ldflags "-w -s -extldflags \"-static\"" -installsuffix cgo
zip -r alexa-applebuyersguide.zip alexa-applebuyersguide
