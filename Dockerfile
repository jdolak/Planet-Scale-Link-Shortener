FROM cimg/go:1.23.1
USER 0
RUN mkdir /pspbalsaas
COPY ./libs /pspbalsaas/libs
COPY ./go.* /pspbalsaas/
COPY ./src /pspbalsaas/src
WORKDIR /pspbalsaas
ENV GOPATH=/pspbalsaas/libs
RUN go build ./src/main.go
USER circleci
CMD ["./main"]
