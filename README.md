# Milutin Jewellery

Данное приложение представляет собой сайт ювелирной мастерской milutin-jewellery, админ панель для сайта и API.

Сайт и админ панель разработаны на Next.js с использованием server actions.
API разработан на golang с использованием redis и postgresql.

## Сборка

- В проекте настроены пайплайны для автоматической проверки, сборки и отправки образов в docker-hub.
- Если необходимо выполнить сборку и пуш локально, скрипты лежат в Makefile.

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
