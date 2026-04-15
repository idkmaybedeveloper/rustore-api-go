# rustore-api-go

неофициальный Go клиент для rustore api

апи ручки получены путём реверс-инжиниринга официального клиента `ru.vk.store` (obviously)

## возможности

- получить информацию о приложении по package name
- получить ссылку на скачивание APK
- поиск приложений
- автодополнение поиска

## установка

```bash
go get github.com/idkmaybedeveloper/rustore-api-go
```

## использование как библиотека

```go
import "github.com/idkmaybedeveloper/rustore-api-go/rustore"

// краткая сводка + download url
summary, err := rustore.GetAppSummary("ru.ozon.app.android")
fmt.Println(summary.DownloadUrl)

// полная информация
info, err := rustore.GetAppInfo("ru.ozon.app.android")

// ссылка на скачивание по appId
dl, err := rustore.GetDownloadLink(info.AppId)
fmt.Println(dl.DownloadUrls[0].Url)

// поиск
results, err := rustore.SearchApps("telegram", 0, 20)

// автодополнение
suggestions, err := rustore.GetSearchSuggestions("yand")
```

## cli

```bash
go build -o rustore-api ./cmd/main.go
```

```bash
# инфо о приложении
./rustore-api info ru.ozon.app.android

# поиск
./rustore-api search telegram

# автодополнение
./rustore-api suggest yande
```

## api

### apps

| функция                                                    | описание                       |
| ------------------------------------------------------------| --------------------------------|
| `GetAppInfo(packageName string) (AppInfo, error)`          | полная информация о приложении |
| `GetDownloadLink(appId int) (DownloadLinkResponse, error)` | ссылка на скачивание APK       |
| `GetAppSummary(packageName string) (AppSummary, error)`    | краткая сводка + download url  |

### search

| функция                                                                      | описание         |
| ---------------------------------------------------------------------------- | ---------------- |
| `SearchApps(query string, pageNumber, pageSize int) (SearchResponse, error)` | поиск приложений |
| `GetSearchSuggestions(query string) (SearchSuggestResponse, error)`          | автодополнение   |

### utils

| функция                              | описание                     |
| ------------------------------------ | ---------------------------- |
| `FormatFileSize(bytes int64) string` | форматирование размера файла |

## эндпоинты

| метод | эндпоинт                                     | описание             |
| -------| ----------------------------------------------| ----------------------|
| get   | `/applicationData/overallInfo/{packageName}` | инфо о приложении    |
| post  | `/applicationData/v2/download-link`          | ссылка на скачивание |
| get   | `/applicationData/apps?query=...`            | поиск приложений     |
| get   | `/search/suggest?query=...`                  | автодополнение       |

## дисклеймер

это **неофициальный** клиент. апи может измениться в любой момент.

## P.S.
based on [https://code.wejust.rest/lain/rustore-api](https://code.wejust.rest/lain/rustore-api)
