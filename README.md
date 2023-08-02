# Валидатор исходного кода репозитория. Консольная утилита.
## Статическая сборка
env GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o prjctvld main.go
## Параметры запуска: 
prjctvld cmd -h
## Параметры конфигурации задаются в config.json
{

    "WhiteList": "whitelist.txt", // путь до списка

    "ScipList": "sciplist.txt", // путь до списка

    "Project": "." // путь до проекта

}

## Список форматов файлов, которые можно задавать в whitelist
https://github.com/gabriel-vasile/mimetype/blob/master/supported_mimes.md
