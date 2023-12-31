# Effective Mobile-TASK

# Инструкция к запуску приложения 
Изначально приложение запускается на http://localhost:3005/ и представляет из себя вебсайт с возможность просмотра пользователей

Документация Swagger UI открывается по адресу http://localhost:3005/swagger/index.html

Измените конфигурацию приложения для запуска сервера и подключения к бд в следующих файлах:

[.env]
```sh
ENV=dev
USE_HTTP=true
DATABASE_DSN= your postgres patch
```
[config.yml]
```sh
http:
    port: 3005
    host: 0.0.0.0
```
Таблица с которой работает приложение находится в папке [migrations]. Перед запуском приложения ее необходимо создать заранее
```sh
CREATE TABLE IF NOT EXISTS users (
 id             bigserial   primary key,
 name           text        not null,
 surname        text        not null,
 patronymic     text,
 age            int         not null,
 gender         text        not null,
 nation         text        not null
);
```
Приложение готово к запуску.

Онсновные команды для работы с приложением прописаны в Makefile:
```sh
run:
	go run main.go

docker-build:
	docker build -t effective_mobile_task:local .

docker-compose-up:
	docker compose -f docker-compose.yml up
```
# Инструкция к запуску приложения через Docker
В выше упомянутом [.env] файле необходимо изменить host на host.docker.internal
```sh
ENV=dev
USE_HTTP=true
DATABASE_DSN=postgres://youruser:yourpassword@host.docker.internal:5432/yourdatabase
```
Если была изменена конфигурация приложения в файле [config.yml], то необходимо прокинуть новые порты в файле [docker-compose.yml]:
```sh
services:
  app:
    image: effective_mobile_task:local
    container_name: em
    ports:
    - "3005:3005" (your ports)
```
Приложение готово к запуску через Docker.
Осталось выполнить следующие команды:
```sh
docker build -t sber_task:local .
docker compose -f docker-compose.yml up
 ```
# Тестирование
Протестировать все запрос можно через Swagger UI по адресу http://localhost:3005/swagger/index.html

Или воспользоваться POSTMAN с моей [коллекцией запросов](https://drive.google.com/file/d/1Z_5r1t4ItExc552QEtyaD8XEuGlTP_Ig/view?usp=sharing)
