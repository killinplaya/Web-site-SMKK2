# Сайт-визитка для ООО «СМК К2»

## Что реализовано
- Одностраничный адаптивный сайт-визитка.
- Go-бэкенд (без внешних фреймворков):
  - `GET /api/company` — реквизиты компании для фронтенда.
  - `GET /health` — проверка статуса сервера.
- Использованы материалы компании:
  - `logo-white.png`
  - `hero-photo.png`
  - `hero-slide-port.jpg`
  - `hero-slide-npp.jpg`
  - `director.jpg`

## Безопасность
На сервере включены заголовки:
- `Content-Security-Policy`
- `X-Content-Type-Options`
- `X-Frame-Options`
- `Referrer-Policy`
- `Permissions-Policy`
- `Cross-Origin-Opener-Policy`
- `Cross-Origin-Resource-Policy`
- `X-Permitted-Cross-Domain-Policies`

## Технологический стек
- Backend: Go (`net/http`, `embed`, JSON API)
- Frontend: HTML5, CSS3, Vanilla JS

## Структура
- `main.go` — сервер и API
- `main_test.go` — тесты API/безопасности
- `web/index.html` — основная страница
- `web/static/css/style.css` — стили
- `web/static/js/main.js` — клиентская логика
- `web/static/img/*` — изображения

## Проверка
```bash
go test ./...
go build ./...
```
