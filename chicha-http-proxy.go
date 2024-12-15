package main

import (
	"bytes"
	"flag"                              // Для работы с флагами командной строки
	"fmt"                               // Для форматированного вывода в консоль
	"golang.org/x/crypto/acme/autocert" // Пакет для работы с Let's Encrypt
	"io"                                // Для работы с потоками данных
	"log"                               // Для логирования событий и ошибок
	"net/http"                          // Для работы с HTTP-запросами и сервером
	"os"                                // Для взаимодействия с операционной системой
)

// Версия программы
var version = "dev"

// proxyHandler возвращает функцию-обработчик для прокси-сервера
// targetURL - целевой URL для пересылки запросов
func proxyHandler(targetURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Чтение тела входящего запроса, если оно существует
		var body []byte
		if r.Body != nil {
			var err error
			body, err = io.ReadAll(r.Body)
			if err != nil {
				// Отправка ответа с ошибкой 500 и логирование ошибки
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				log.Printf("Error reading request body: %v", err)
				return
			}
		}

		// Формирование полного URL для пересылки (целевой URL + путь запроса)
		originalURL := targetURL + r.URL.Path
		currentURL := originalURL

		// Создание HTTP-клиента для пересылки запроса
		client := &http.Client{}

		for {
			// Создание нового HTTP-запроса с методом, заголовками и телом из оригинального запроса
			req, err := http.NewRequest(r.Method, currentURL, bytes.NewReader(body))
			if err != nil {
				http.Error(w, "Failed to create request", http.StatusInternalServerError)
				log.Printf("Error creating request: %v", err)
				return
			}

			// Копирование заголовков из оригинального запроса
			for header, values := range r.Header {
				for _, value := range values {
					req.Header.Add(header, value)
				}
			}

			// Копирование строки запроса (query string)
			req.URL.RawQuery = r.URL.RawQuery

			// Выполнение пересылки запроса
			resp, err := client.Do(req)
			if err != nil {
				http.Error(w, "Error forwarding request", http.StatusBadGateway)
				log.Printf("Error forwarding request: %v", err)
				return
			}
			defer resp.Body.Close()

			// Обработка редиректов (HTTP 3xx)
			if resp.StatusCode >= 300 && resp.StatusCode < 400 {
				location, err := resp.Location()
				if err != nil {
					http.Error(w, "Failed to handle redirect", http.StatusInternalServerError)
					log.Printf("Error handling redirect: %v", err)
					return
				}
				currentURL = location.String()
				log.Printf("Redirecting to: %s", currentURL)
				continue
			}

			// Копирование заголовков из ответа целевого сервера в ответ клиенту
			for header, values := range resp.Header {
				for _, value := range values {
					w.Header().Add(header, value)
				}
			}

			// Установка HTTP-кода ответа
			w.WriteHeader(resp.StatusCode)

			// Копирование тела ответа и отправка клиенту
			responseBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading response body: %v", err)
				return
			}
			_, err = w.Write(responseBody)
			if err != nil {
				log.Printf("Error writing response body: %v", err)
			}
			return
		}
	}
}

func main() {
	// Определение флагов командной строки
	httpPort := flag.String("http-port", "80", "Port for the HTTP server")
	httpsPort := flag.String("https-port", "443", "Port for the HTTPS server")
	targetURL := flag.String("target-url", "https://zendesk.com", "Target URL for forwarding requests")
	domain := flag.String("domain", "", "Domain for automatic Let's Encrypt certificate")
	showVersion := flag.Bool("version", false, "Show program version")

	// Разбор флагов
	flag.Parse()

	// Если указан флаг --version, выводим версию и выходим
	if *showVersion {
		fmt.Printf("Program version: %s\n", version)
		os.Exit(0)
	}

	// Проверка наличия целевого URL
	if *targetURL == "" {
		log.Fatal("Target URL (--target-url) is not specified")
	}

	// Создание обработчика прокси
	handler := proxyHandler(*targetURL)

	// Канал для сигналов завершения работы
	done := make(chan bool)

	// Запуск HTTP-сервера в отдельной горутине, если HTTP-порт указан
	if *httpPort != "" {
		go func() {
			httpServer := &http.Server{
				Addr:    ":" + *httpPort,
				Handler: handler,
			}
			log.Printf("Starting HTTP proxy on port %s tartgeting %s", *httpPort, *targetURL)
			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("HTTP server error: %v", err)
			}
		}()
	}

	// Запуск HTTPS-сервера в отдельной горутине, если домен указан
	if *domain != "" {
		go func() {
			// Настройка менеджера сертификатов Let's Encrypt
			m := &autocert.Manager{
				Cache:      autocert.DirCache("certs"),          // Директория для хранения сертификатов
				Prompt:     autocert.AcceptTOS,                  // Автоматическое принятие условий использования
				HostPolicy: autocert.HostWhitelist(*domain),     // Разрешенные домены
			}

			httpsServer := &http.Server{
				Addr:      ":" + *httpsPort,
				TLSConfig: m.TLSConfig(),
				Handler:   handler,
			}

			log.Printf("Starting HTTPS proxy on domain %s and port %s targeting %s", *domain, *httpsPort, *targetURL)
			if err := httpsServer.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
				log.Fatalf("HTTPS server error: %v", err)
			}
		}()
	} else {
		fmt.Println("Domain for HTTPS not specified.");
	}

	// Ожидание завершения работы серверов
	<-done
}

