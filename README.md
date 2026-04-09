# 🌐 Сайт-визитка для ООО «СМК К2»

Одностраничный адаптивный сайт-визитка с Go-бэкендом и REST API.
Проект демонстрирует навыки разработки полного цикла: от фронтенда до серверной логики и базовой безопасности.

---

## 📸 Скриншот

![Скриншот сайта](./web/static/img/screenshot.png)

---

## 🚀 Что реализовано

* Одностраничный адаптивный сайт
* API на Go (без фреймворков):

  * `GET /api/company` — получение данных компании
  * `GET /health` — проверка состояния сервера
* Динамическая подгрузка данных на фронтенде
* Чистый Vanilla JS без библиотек

---

## 🔒 Безопасность

Настроены HTTP-заголовки:

* Content-Security-Policy
* X-Content-Type-Options
* X-Frame-Options
* Referrer-Policy
* Permissions-Policy
* Cross-Origin-Opener-Policy
* Cross-Origin-Resource-Policy

---

## 🛠 Технологии

* **Backend:** Go (net/http, embed, JSON API)
* **Frontend:** HTML5, CSS3, Vanilla JavaScript

---

## 📁 Структура проекта

* `main.go` — сервер и API
* `main_test.go` — тесты
* `web/index.html` — основная страница
* `web/static/css/style.css` — стили
* `web/static/js/main.js` — логика
* `web/static/img/` — изображения

---

## ⚙️ Запуск проекта

```bash
go test ./...
go build ./...
./project-name
```
