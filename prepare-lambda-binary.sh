set -x

GOOS=linux go build -o linux-binary -a -ldflags "-w -s -extldflags \"-static\"" -installsuffix cgo
zip -r alexa-applebuyersguide.zip linux-binary
