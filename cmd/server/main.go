package main

import (
	r "TestTask/app/api"
	"TestTask/app/conf"
	"TestTask/app/server"
	"log"
	"net/http"
	"sync"

	_ "TestTask/docs"
)

func main() {
	// Инициализируем конфигурацию приложения.
	config := conf.New()

	// Парсим параметры сервера (например, для миграций).
	server.ParseParams(config)

	// Настраиваем маршрутизатор с помощью конфигурации.
	router := r.SetRouter(config)

	// Получаем порт для HTTP сервера из конфигурации.
	hp := config.HTTP()
	port := ":" + hp.HTTP_PORT

	// Используем WaitGroup для ожидания завершения работы сервера.
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// Запускаем HTTP сервер в отдельной горутине.
	go func() {
		// Запускаем HTTP сервер на указанном порту и передаем маршрутизатор.
		http.ListenAndServe(port, router)
		log.Printf("Start on port " + port)
		wg.Done()
	}()

	// Ожидаем завершения работы серверной горутины.
	wg.Wait()
}
