# syntax=docker/dockerfile:1

# Указываем версию Go, которую будем использовать
ARG GO_VERSION=1.21.6
FROM golang:${GO_VERSION} AS build

# Устанавливаем рабочую директорию
WORKDIR /src

# Копируем все файлы проекта в контейнер
COPY . .

# Загружаем зависимости
RUN go mod download

# Собираем приложение (можно пропустить, если используем `go run .`)
RUN go build -o /web/app /src/web/app

################################################################################
# Create a new stage for running the application that contains the minimal
# runtime dependencies for the application. This often uses a different base
# image from the build stage where the necessary files are copied from the build
# stage.
FROM golang:${GO_VERSION} AS final

# Устанавливаем рабочую директорию на web/app
WORKDIR /src/web/app

# Копируем все файлы из этапа сборки
COPY --from=build /src /src

# Указываем команду для выполнения
ENTRYPOINT ["go", "run", "."]
