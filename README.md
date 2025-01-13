# Milutin Jewellery

Данное приложение предтавляет собой сайт milutin-jewellery, админ панель и api

Сайт и админ панель разработаны на Next.js с использованием server actions.
API разработан на golang, используются redis и postgresql.

## Сборка

- Скрипты для сборки и пуша контейнеров расположены в Makefile
- Необходимо выполнить `make build` и `make push`, при этом в .env файле должна быть указана переменная NEXT_PUBLIC_IMAGE_URL, так как она понадобиться при сборке фронтенда

## Деплой

1. Поставить docker, выполнить `docker login`
2. Создать `docker-compose.yml`, вставить туда содержимое из репозитория, ПОМЕНЯТЬ СЕКРЕТЫ!
3. Создать `.env` с необходимыми данными, запустить `docker compose up -d`
4. Установить nginx: `sudo apt update && sudo apt install nginx`
5. Добавить конфиг из репозитория `vim /etc/nginx/sites-available/milutin-jewellery.conf`
6. Добавить ссылку на конфиг `sudo ln -s /etc/nginx/sites-available/milutin-jewellery.conf /etc/nginx/sites-enabled/`
7. Добавить поле `client_max_body_size 100M;` в http в `/etc/nginx/nginx.conf`
8. Перезапустить nginx `sudo systemctl restart nginx`
9. Установить certbot `sudo apt update && sudo apt install certbot python3-certbot-nginx`
10. Создать сертификаты `sudo certbot --nginx -d milutin-jewellery.com -d api.milutin-jewellery.com -d admin.milutin-jewellery.com`
