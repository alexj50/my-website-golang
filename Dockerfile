FROM golang:1.10 AS builder

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/alexj50/my-website-golang
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
COPY ./views ./views

ENTRYPOINT ["./app"]
