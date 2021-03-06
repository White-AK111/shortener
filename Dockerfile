FROM golang:1.17 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:1.17 as builder

COPY --from=modules /go/pkg /go/pkg

RUN mkdir -p /src
ADD . /src
WORKDIR /src

  # Добавляем непривилегированного пользователя
RUN useradd -u 10001 shortener

  # Собираем бинарный файл
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
go build -o /shortener ./cmd/main.go

FROM scratch

  # Не забываем скопировать /etc/passwd с предыдущего стейджа
COPY --from=builder /etc/passwd /etc/passwd
USER shortener

COPY --from=builder /shortener /shortener
COPY configs/config.yml /etc/config.yml

#CMD ["/shortener -path=/etc/config.yml"]
ENTRYPOINT ["/shortener", "--path=/etc/config.yml"]
