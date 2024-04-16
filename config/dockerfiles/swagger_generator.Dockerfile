FROM golang:1.22.1-alpine
WORKDIR /app

COPY . .

RUN go install github.com/parvez3019/go-swagger3@latest
RUN export PATH="$HOME/go/bin:$PATH"

CMD go-swagger3 --module-path . --main-file-path ./cmd/service/main.go --output /app/api/swagger.yaml --schema-without-pkg --generate-yaml true