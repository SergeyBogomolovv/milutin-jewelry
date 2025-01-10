# Milutin Jewellery

Данное приложение предтавляет собой сайт milutin-jewellery, админ панель и api

Сайт и админ панель разработаны на Next.js с использованием server actions.
API разработан на golang, используются redis и postgresql.

Сборка проекта:

- Скрипты для сборки и пуша контейнеров расположены в Makefile
- Необходимо выполнить `make build` и `make push`, при этом в .env файле должна быть указана переменная NEXT_PUBLIC_IMAGE_URL, так как она понадобиться при сборке фронтенда

Запуск проекта:

- Необходимо создать .env файл по примеру из .env.example
- Запустить docker compose `docker compose up -d`
