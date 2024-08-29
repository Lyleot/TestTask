# Используем официальный образ Golang для сборки приложения
FROM golang:1.22.2 AS build-env

# Копируем исходный код в рабочую директорию контейнера
ADD . /src

# Переходим в директорию с исходным кодом и выполняем команду сборки
RUN cd /src && make build

# Используем базовый образ Alpine для финального контейнера
FROM alpine:3.13

# Устанавливаем необходимые пакеты
RUN apk --no-cache add gettext tzdata
RUN apk --no-cache add libc6-compat

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранное приложение и миграции из предыдущего этапа сборки
COPY --from=build-env /src/TestTask /app/
COPY --from=build-env /src/migrations/ /app/migrations

# Делаем бинарный файл исполняемым
RUN chmod +x /app/TestTask

# Устанавливаем команду по умолчанию для запуска приложения
CMD ["/app/TestTask"]