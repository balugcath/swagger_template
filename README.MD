# swagger template

```sh
cmd
    /main_service               main.go
src
    /main_service
        /api_handler            http ручки
        /client                 враппер для клиента, можно не делать
        /middleware             бизнес логика
        /sql                    база
        /swagger                тут вся кодогенерация
            /client_users       клиент для внешнего сервиса
            /models_server      модели main_service
            /models_users       модели для внешнего сервиса
            /restapi            сервер и операции main_service
api_swagger.yml                 api сервиса main_service
external_api_swagger.yml        api внешнего сервиса для main_service
```
