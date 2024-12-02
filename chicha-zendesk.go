package main

import (
	"bytes"
	"flag"                              // Для работы с флагами командной строки (аналог аргументов main() в Java)
	"fmt"                               // Для форматированного вывода в консоль
	"golang.org/x/crypto/acme/autocert" // Пакет для работы с Let's Encrypt
	"io"                                // Для работы с потоками (в данном случае для чтения/записи тела запроса)
	"log"                               // Для логирования событий и ошибок
	"net/http"                          // Для работы с HTTP-запросами и сервером
	"os"                                // Для взаимодействия с операционной системой (например, завершение программы)
)

// Версия программы
var version = "dev"

// Функция-обработчик запросов
// Она отвечает за прием запросов, их пересылку на целевой URL и возврат ответа клиенту.
func proxyHandler(targetURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Копируем тело входящего запроса (если есть)
		var body []byte
		if r.Body != nil {
			var err error
			body, err = io.ReadAll(r.Body)
			if err != nil {
				// В случае ошибки отправляем клиенту ответ с кодом 500 (Internal Server Error)
				http.Error(w, "Не удалось прочитать тело запроса", http.StatusInternalServerError)
				log.Printf("Ошибка чтения тела запроса: %v", err)
				return
			}
		}

		// Строим начальный URL для пересылки (целевой URL + путь запроса)
		originalURL := targetURL + r.URL.Path
		currentURL := originalURL

		// Создаем HTTP-клиент
		client := &http.Client{}

		for {
			// Создаем новый HTTP-запрос, копируя метод, заголовки и тело из оригинального запроса
			req, err := http.NewRequest(r.Method, currentURL, bytes.NewReader(body))
			if err != nil {
				http.Error(w, "Не удалось создать запрос", http.StatusInternalServerError)
				log.Printf("Ошибка создания запроса: %v", err)
				return
			}

			// Копируем заголовки из оригинального запроса
			for header, values := range r.Header {
				for _, value := range values {
					req.Header.Add(header, value)
				}
			}

			// Копируем параметры строки запроса (query string)
			req.URL.RawQuery = r.URL.RawQuery

			// Выполняем пересылку запроса
			resp, err := client.Do(req)
			if err != nil {
				http.Error(w, "Ошибка пересылки запроса", http.StatusBadGateway)
				log.Printf("Ошибка пересылки запроса: %v", err)
				return
			}
			defer resp.Body.Close()

			// Если целевой сервер вернул редирект (HTTP 3xx), обрабатываем его
			if resp.StatusCode >= 300 && resp.StatusCode < 400 {
				location, err := resp.Location()
				if err != nil {
					http.Error(w, "Не удалось обработать редирект", http.StatusInternalServerError)
					log.Printf("Ошибка обработки редиректа: %v", err)
					return
				}
				currentURL = location.String()
				log.Printf("Перенаправление на: %s", currentURL)
				continue
			}

			// Копируем заголовки из ответа целевого сервера в ответ клиенту
			for header, values := range resp.Header {
				for _, value := range values {
					w.Header().Add(header, value)
				}
			}

			// Устанавливаем HTTP-код ответа
			w.WriteHeader(resp.StatusCode)

			// Копируем тело ответа и отправляем клиенту
			responseBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Ошибка чтения тела ответа: %v", err)
				return
			}
			_, err = w.Write(responseBody)
			if err != nil {
				log.Printf("Ошибка записи тела ответа: %v", err)
			}
			return
		}
	}
}

func main() {
	// Определяем флаги командной строки
	port := flag.String("port", "8080", "Порт для запуска прокси-сервера")
	targetURL := flag.String("target-url", "https://ovmsupport.zendesk.com", "Целевой URL для пересылки запросов")
	domain := flag.String("domain", "", "Домен для автоматического получения сертификата Let's Encrypt")
	showVersion := flag.Bool("version", false, "Показать версию программы")

	// Разбираем параметры командной строки
	flag.Parse()

	// Если указан флаг --version, выводим версию программы и выходим
	if *showVersion {
		fmt.Printf("Версия программы: %s\n", version)
		os.Exit(0)
	}

	// Проверяем, что целевой URL указан
	if *targetURL == "" {
		log.Fatal("Целевой URL (--target-url) не указан")
	}

	// Если указан домен, включаем поддержку HTTPS с Let's Encrypt
	if *domain != "" {
		// Настраиваем менеджер сертификатов Let's Encrypt
		m := &autocert.Manager{
			Cache:      autocert.DirCache("certs"), // Директория для хранения сертификатов
			Prompt:     autocert.AcceptTOS,         // Автоматическое принятие условий использования
			HostPolicy: autocert.HostWhitelist(*domain),
		}

		// Запускаем HTTPS-сервер
		server := &http.Server{
			Addr:      ":443",
			TLSConfig: m.TLSConfig(),
			Handler:   proxyHandler(*targetURL),
		}

		log.Printf("Запуск HTTPS-прокси на домене %s и порту 443", *domain)
		log.Fatal(server.ListenAndServeTLS("", ""))
	} else {
		// Запускаем HTTP-сервер, если домен не указан
		http.HandleFunc("/", proxyHandler(*targetURL))
		log.Printf("Запуск HTTP-прокси на порту %s", *port)
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	}
}
