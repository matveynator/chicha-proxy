## chicha-zendesk

**chicha-zendesk** — прокси-сервер для пересылки запросов на Zendesk, созданный с целью помочь порядочным людям, столкнувшимся с ограничениями, продолжать использовать качественный сервис для работы и общения.

---

### Запуск

#### По умолчанию:
```bash
./chicha-zendesk
```

- **Порт**: `8080`
- **Целевой URL**: `https://ovmsupport.zendesk.com`

---

### Флаги и опции

- **Изменение порта**:
  ```bash
  ./chicha-zendesk -port=9090
  ```

- **Изменение целевого URL**:
  ```bash
  ./chicha-zendesk -url=https://example.zendesk.com
  ```

- **Одновременное изменение порта и URL**:
  ```bash
  ./chicha-zendesk -port=9090 -url=https://example.zendesk.com
  ```

---

### Примеры использования

- Отправка GET-запроса:
  ```bash
  curl -X GET http://localhost:8080/api/v2/users
  ```

- Отправка POST-запроса:
  ```bash
  curl -X POST http://localhost:8080/api/v2/tickets -d '{"subject":"Test Ticket"}' -H "Content-Type: application/json"
  ```

---

### Примечание

Программа создана исключительно для использования порядочными людьми, которые соблюдают законы и правила, с целью продолжать пользоваться сервисами, недоступными в их регионе. Пожалуйста, используйте её ответственно.
