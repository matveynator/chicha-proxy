<img src="https://github.com/matveynator/chicha-http-proxy/blob/main/chicha-proxy.png?raw=true" alt="chicha-proxy" width="50%" align="right" />

## **chicha-http-proxy**

**chicha-http-proxy** — прокси-сервер для пересылки запросов на Zendesk, созданный с целью помочь порядочным людям, столкнувшимся с ограничениями, продолжать использовать качественный сервис для работы и общения. Кроме того можно создавать зеркала любых других сайтов.

---

### **Скачивание**

Выберите нужную версию:

- **Linux**: [AMD64](http://files.zabiyaka.net/chicha-proxy/latest/no-gui/linux/amd64/chicha-proxy), [ARM64](http://files.zabiyaka.net/chicha-proxy/latest/no-gui/linux/arm64/chicha-proxy)
- **Windows**: [AMD64](http://files.zabiyaka.net/chicha-proxy/latest/no-gui/windows/amd64/chicha-proxy.exe), [ARM64](http://files.zabiyaka.net/chicha-proxy/latest/no-gui/windows/arm64/chicha-proxy.exe)
- **MacOS**: [Intel](http://files.zabiyaka.net/chicha-proxy/latest/no-gui/mac/amd64/chicha-proxy), [M1/M2](http://files.zabiyaka.net/chicha-proxy/latest/no-gui/mac/arm64/chicha-proxy)

Другие варианты доступны в [полном списке](http://files.zabiyaka.net/chicha-proxy/latest/no-gui).

---

### **Установка**

1. **Linux/macOS**:
   ```bash
   sudo curl http://files.zabiyaka.net/chicha-proxy/latest/no-gui/linux/amd64/chicha-proxy > /usr/local/bin/chicha-proxy; 
   sudo chmod +x /usr/local/bin/chicha-proxy; chicha-proxy --version;
   ```

2. **Windows**: 
   Скачайте файл `.exe` и добавьте его в `PATH`.

---

### **Использование**

#### Основные флаги

- `--target-url` (обязательный): URL, на который будут пересылаться запросы. Например: `--target-url=https://testsupport.zendesk.com`.
- `--http-port`: порт для запуска HTTP-сервера. По умолчанию `80`. Например: `--http-port=8080`.
- `--https-port`: порт для запуска HTTPS-сервера. По умолчанию `443`. Например: `--https-port=8443`.
- `--domain`: домен, на который будет выпущен автоматический сертификат Let's Encrypt. Например: `--domain=example.com`.
- `--version`: вывод текущей версии программы.

#### Примеры запуска

1. **Для HTTP-прокси**:
   ```bash
   chicha-proxy --http-port=8080 --target-url=https://testsupport.zendesk.com
   ```

2. **Для HTTPS-прокси с автоматическим сертификатом**:
   ```bash
   chicha-proxy --domain=example.com --target-url=https://testsupport.zendesk.com
   ```

---

### **Особенности**

- **Сертификаты**: автоматически настраиваются и обновляются с помощью Let's Encrypt. Сертификаты выдаются на 90 дней и обновляются автоматически.
- **Порты**: для работы через HTTPS должен быть открыт порт 443 (или другой, указанный в `--https-port`).

---
