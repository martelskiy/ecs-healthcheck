FROM --platform=$BUILDPLATFORM golang:1.20.2@sha256:a78b213b71d9382d72aeb1e1370d339e9cd03d1f35af9c11306d690dc2215da6 as build
WORKDIR /healthcheck
ADD . .

ARG TARGETOS TARGETARCH
ENV GOOS $TARGETOS
ENV GOARCH $TARGETARCH

RUN CGO_ENABLED=0 go build -o bin/ecs-healthcheck cmd/ecs-healthcheck/main.go

FROM scratch
COPY --from=build /healthcheck/bin/ecs-healthcheck /