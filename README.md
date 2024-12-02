<img src="https://github.com/matveynator/chicha-zendesk/blob/main/chicha-zendesk.png?raw=true" alt="chicha-zendesk" width="50%" align="right" />

## **chicha-zendesk**

**chicha-zendesk** — прокси-сервер для пересылки запросов на Zendesk, созданный с целью помочь порядочным людям, столкнувшимся с ограничениями, продолжать использовать качественный сервис для работы и общения.

---

### **Скачивание**

Выберите нужную версию:

- **Linux**: [AMD64](http://files.zabiyaka.net/chicha-zendesk/latest/no-gui/linux/amd64/chicha-zendesk), [ARM64](http://files.zabiyaka.net/chicha-zendesk/latest/no-gui/linux/arm64/chicha-zendesk)
- **Windows**: [AMD64](http://files.zabiyaka.net/chicha-zendesk/latest/no-gui/windows/amd64/chicha-zendesk.exe), [ARM64](http://files.zabiyaka.net/chicha-zendesk/latest/no-gui/windows/arm64/chicha-zendesk.exe)
- **MacOS**: [Intel](http://files.zabiyaka.net/chicha-zendesk/latest/no-gui/mac/amd64/chicha-zendesk), [M1/M2](http://files.zabiyaka.net/chicha-zendesk/latest/no-gui/mac/arm64/chicha-zendesk)

Другие варианты доступны в [полном списке](http://files.zabiyaka.net/chicha-zendesk/latest).

---

### **Установка**

1. **Linux/macOS**:
   ```bash
   sudo curl http://files.zabiyaka.net/chicha-zendesk/latest/no-gui/linux/amd64/chicha-zendesk > /usr/local/bin/chicha-zendesk; 
   sudo chmod +x /usr/local/bin/chicha-zendesk; chicha-zendesk --version;
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
   chicha-zendesk --http-port=8080 --target-url=https://ovmsupport.zendesk.com
   ```

2. **Для HTTPS-прокси с автоматическим сертификатом**:
   ```bash
   chicha-zendesk --domain=example.com --target-url=https://ovmsupport.zendesk.com
   ```

---

### **Особенности**

- **Сертификаты**: автоматически настраиваются и обновляются с помощью Let's Encrypt. Сертификаты выдаются на 90 дней и обновляются автоматически.
- **Порты**: для работы через HTTPS должен быть открыт порт 443 (или другой, указанный в `--https-port`).

---
