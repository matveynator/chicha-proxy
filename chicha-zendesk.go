package main

import (
	"bytes"    // Для работы с буфером байтов, аналог ByteArrayOutputStream в Java
	"flag"     // Для обработки флагов командной строки, аналог args в Java
	"io"       // Для чтения и записи данных, аналог InputStream/OutputStream в Java
	"log"      // Для логирования, аналог java.util.logging
	"net/http" // Для работы с HTTP-запросами и ответами
)

// Главная функция программы, аналог main метода в Java
func main() {
	// Определяем флаги для параметров командной строки
	port := flag.String("port", "8080", "Порт, на котором будет работать сервер")
	// Флаг для выбора целевого URL (по умолчанию URL Zendesk)
	targetURL := flag.String("url", "https://ovmsupport.zendesk.com", "Целевой URL для проксирования")
	flag.Parse() // Анализируем переданные параметры (аналог аргументов в Java)

	// Устанавливаем обработчик запросов (аналог сервлетов или контроллеров в Java)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Вызов функции обработки запросов с переданным URL
		proxyHandler(w, r, *targetURL)
	})

	// Запускаем HTTP-сервер, аналогично использованию API серверов в Java (например, Jetty)
	log.Printf("Запуск прокси-сервера на порту %s, целевой URL: %s", *port, *targetURL)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err) // Завершаем программу при ошибке
	}
}

// Функция-обработчик для проксирования запросов
func proxyHandler(w http.ResponseWriter, r *http.Request, targetURL string) {
	// Создаем HTTP-клиент, аналог HttpClient из Apache или стандартного Java
	client := &http.Client{}

	// Формируем базовый URL из целевого URL и пути текущего запроса
	originalURL := targetURL + r.URL.Path // r.URL.Path содержит путь запроса
	currentURL := originalURL             // Переменная для отслеживания текущего URL при редиректах

	// Читаем тело запроса, если оно присутствует (для POST, PUT и других методов с телом)
	var body []byte
	if r.Body != nil { // r.Body - это аналог InputStream
		var err error
		body, err = io.ReadAll(r.Body) // Читаем все данные из тела запроса
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
			log.Printf("Ошибка чтения тела запроса: %v", err)
			return
		}
	}

	// Цикл для обработки возможных редиректов (коды 3xx)
	for {
		// Создаем новый HTTP-запрос (аналог HttpRequest в Java)
		req, err := http.NewRequest(r.Method, currentURL, bytes.NewReader(body))
		if err != nil {
			http.Error(w, "Ошибка создания запроса", http.StatusInternalServerError)
			log.Printf("Ошибка создания запроса: %v", err)
			return
		}

		// Копируем заголовки из оригинального запроса в новый
		for header, values := range r.Header {
			for _, value := range values {
				req.Header.Add(header, value)
			}
		}

		// Копируем параметры строки запроса (query parameters)
		req.URL.RawQuery = r.URL.RawQuery // RawQuery содержит строку запроса без декодирования

		// Отправляем запрос через HTTP-клиент
		resp, err := client.Do(req) // Аналог client.execute(request) в Java
		if err != nil {
			http.Error(w, "Ошибка отправки запроса", http.StatusBadGateway)
			log.Printf("Ошибка отправки запроса: %v", err)
			return
		}
		defer resp.Body.Close() // Гарантируем освобождение ресурсов (аналог try-with-resources)

		// Проверяем статус ответа для обработки редиректа (3xx)
		if resp.StatusCode >= 300 && resp.StatusCode < 400 {
			location, err := resp.Location() // Получаем новый URL из заголовка Location
			if err != nil {
				http.Error(w, "Ошибка обработки редиректа", http.StatusInternalServerError)
				log.Printf("Ошибка получения нового URL при редиректе: %v", err)
				return
			}
			currentURL = location.String() // Переключаемся на новый URL
			log.Printf("Перенаправление на: %s", currentURL)
			continue // Продолжаем цикл для отправки нового запроса
		}

		// Если редиректа нет, возвращаем ответ клиенту
		// Копируем заголовки ответа
		for header, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(header, value)
			}
		}

		// Устанавливаем код статуса ответа (аналог setStatus в Java сервлетах)
		w.WriteHeader(resp.StatusCode)

		// Копируем тело ответа
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Ошибка чтения тела ответа: %v", err)
			return
		}
		_, err = w.Write(body) // Отправляем тело ответа клиенту
		if err != nil {
			log.Printf("Ошибка отправки тела ответа: %v", err)
		}
		return // Завершаем обработку запроса
	}
}
