ARG base_image
ARG builder_image=concourse/golang-builder

FROM ${builder_image} as builder
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ENV CGO_ENABLED 0

RUN go build -o /assets/in ./cmd/in
RUN go build -o /assets/out ./cmd/out
RUN go build -o /assets/check ./cmd/check

FROM ${base_image}

COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*
