<img src="https://github.com/matveynator/chicha-zendesk/blob/main/chicha-zendesk.png?raw=true" alt="chicha-zendesk" width="50%" align="right" />


## chicha-zendesk

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

Для запуска HTTP-прокси:
```bash
chicha-zendesk --port=8080 --target-url=https://ovmsupport.zendesk.com
```

Для запуска HTTPS-прокси с автоматическим сертификатом:
```bash
chicha-zendesk --domain=example.com --target-url=https://ovmsupport.zendesk.com
```

- **Сертификаты**: автоматически настраиваются и обновляются с помощью Let's Encrypt.
- Порт 443 должен быть открыт для работы через HTTPS, сертификат выдается на 90 дней и обновляется в конце срока автоматически.
