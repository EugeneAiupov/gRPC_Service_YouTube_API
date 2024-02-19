# test_case
# Описание

Данный инструмент командной строки, написанный на Go, который позволяет
получать URL превью (thumbnails) видеороликов на YouTube. Поддерживает как одиночные, так и множественные 
запросы для получения видеороликов. Для множественных запросов использован механизм аснхронности.

# Начало работы

### Предварительные требования

Для работы с инструментом убедитесь, что у вас установлен Go версии 1.15 или выше. Проверить текущую версию
Go можно, выполнив команду:

```bash
go version
```

### Установка

1. Клонируйте репозиторий проекта на свою машину:

```bash
git clone https://github.com/EugeneAiupov/test_case.git
```

2. Перейдите в директорию проекта:

```bash
cd test_case
```

3. Соберите исполняемый файл grpc-service (server.go):

```bash
go build
```

4. Соберите исполняемый файл клиента в каталоге `client`:

```bash
cd ./client
go build
```

## Использование

*Важно для коректной работы клиента необходимо чтобы был запущен server.exe!*

Для получения превью одного видеоролика, используйте слудующую команду находясь в директории где лежит собранный исполняемый файл client.exe:

```bash
./client <URL видео ролика с YouTube>
```

Для асинхронного получения превью нескольких видеороликов, добавьте флаг `--async`:

```bash
./client --async <URL видео ролика с YouTube> <...> <URL видео ролика с YouTube>
```

## Контакты 
Если у вас возникли вопросы или предложения, пожалуйста, свяжитесь со мной по (eugeneyounge@gmail.com)

---
