# Выполненное тестовое задание для backend-стажёра в команду Avito Advertising

Сервис для хранения и подачи объявлений. В качестве языка программирования использован Go с использованием библиотек Gin и Gorm. В качестве базы данных используется SQLite.

## Запуск
Сервис контейниризован, так что запустить его можно при помощи `docker-compose up`

Также можно скомпилировать проект `go build -o main`, после чего запусить исполняемый файл `./main`

## Структура
```
├── models
│   ├── advertisment.go // Структуры и таблицы для объявлений
|	├── setup.go // Соединение с БД
├── Controllers
│   └── advertisment.go // Контроллер для объявлений
└── main.go
```

## API
Все требования выполнены. API имеет три метода.

**Получение списка объявлений**
----

  Возвращает JSON со всеми объявлениями

* **URL**

  /api/v1/ads

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**

    None

   **Optional:**

    `page=[integer]` где `page` номер страницы(так как объявления выводятся по 10)

    `order_by=[string]`, где `order_by` - то, как сортируется список
    
    имеются четыре способа сортировки 

    1. `time_asc` (стандартное значене) - сортирует список по возрастанию даты создания

    2. `time_desc` - сортирует список по убыванию даты создания

    3. `price_asc` - сортирует список по возрастанию цены

    4. `price_desc` - сортирует список по убыванию цены

* **Data Params**

  None

* **Success Response:**
  * **Condition** : Все объявления отображаются<br />
    **Code:** 200 <br />
    **Content:** 
    ```json
    [
    {
        "id": 1,
        "name": "Велосипед stels Navigator 500 V 27.5 V020",
        "price": 11000,
        "created_at": "2021-02-18T04:12:06.329034139+03:00",
        "main_picture": "https://23.img.avito.st/image/1/5n1rObaySpRdjsiZKwbtNJOaSpLLmEg"
    },
    {
        "id": 2,
        "name": "Петух Плимутрок полосатый",
        "price": 700,
        "created_at": "2021-02-18T04:12:21.133535582+03:00",
        "main_picture": "https://44.img.avito.st/image/1/jUK6r7ayIauMBuOurP3Df0AMIa0aDiM"
    }
    ]
    ```
 
* **Error Response:**
  * **Condition** : Сортировка задана некорректно<br />
    **Code:** 422 Unprocessable Entity <br />
    **Content:** `{"error": "order_by is incorrect"}`

  OR
  * **Condition** : Страница задана некорректно<br />
    **Code:** 422 Unprocessable Entity <br />
    **Content:** `{"error": "page is incorrect"}`

* **Sample Call:**

    `curl -i http://localhost:8080/api/v1/ads?page=2&order_by=price_asc`

**Получение одного объявления**
----

Возвращает JSON с одним объявлением

* **URL**

  /api/v1/ads/:pk/

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**

    `pk=[integer]` где `pk` - это ID объявления на сервере 

   **Optional:**

    `fields=[string]` где `fields` - дополнительные поля для отображения

    имеется два дополнительных поля(поля можно указать вместе через запятую)

    1. `description` добавляет в возвращаемый json поле description с описанием

    2. `pictures` добавляет в возвращаемый json поле pictures с полным списком фотографий, однако убирает поле main_picture
    
    

* **Data Params**

  None

* **Success Response:**
  * **Condition** : Объявление с данным ID существует, дополнительные поля(при их использовании) заданы корректно<br />
    **Code:** 200 <br />
    **Content:** 
    ``` json
    {
        "id": 7,
        "name": "Органайзер или просто для декора",
        "price": 200,
        "created_at": "2021-02-18T04:13:27.86550877+03:00",
        "main_picture": "https://65.img.avito.st/image/1/wCJcZ7aybMtqzq7OUCy-N6PEbM38xm4"
    }
    ```
 
* **Error Response:**

  * **Condition** : Объявление с данным ID не существует <br />
    **Code:** 404 Not found <br />
    **Content:** `{"error": "ad not found"}`

  OR

  * **Condition** : Неверно заданы дополнительные поля fields <br />
    **Code:** 422 Unprocessable Entity <br />
    **Content:** `{"error": "fields are incorrect"}`

* **Sample Call:**

    `curl -i http://localhost:8080/api/v1/ads/7?fields=description,pictures`

**Cоздание объявления**
----
  Создает новое объявление

* **URL**

  /api/v1/ads/

* **Method:**

  `POST`
  
*  **URL Params**

    None
    

* **Data Params**

  Требуется предоставить название, описание, не более трех ссылок на фотографии, цену 
    ```json
        {
            "name": "[unicode максимум 200 символов]",
            "description": "[unicode максимум 1000 символов]",
            "pictures":"[unicode не более трех ссылок через запятую без пробелов]",
            "price":"[intger]"
        }
    ```

* **Data example** Все поля должны быть отправлены

    ```json
    {
        "name": "Nokian Nordman 7 175/65 r14 на дисках",
        "description": "Продаю колеса в сборе: штамповка + Nokian Nordman 7 175/65 r14 Состояние отличное, покупались на прошлый сезон на Ford Fiesta, но машиной практически не пользовались.",
        "pictures": "https://89.img.avito.st/image/1/CIGv3LaypGiZdWZts_xAo1B_pG4PfaY,https://44.img.avito.st/image/1/0dlI97ayfTB-Xr81ZtCZ-7dUfTboVn8,https://01.img.avito.st/image/1/mPLJibayNBv_IPYe9a_Q0DYqNB1pKDY",
        "price": 12000,
    }
    ```
* **Success Response:**

    **Condition** : Объявление создано успешно</br>
    **Code:** 200 <br />
    **Content:** 
    ```json
        {
            "id": 22
        }
    ```
 
* **Error Response:**

    **Condition** : Одно из полей не заполнено или не прошло валидцию</br>
    **Code:** 400 Bad request <br />
    **Content:** `{"error": "fields are incorrect"}`

* **Sample Call:**

    `curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"Nokian Nordman 7 175/65 r14 на дисках\", \"description\": \"Продаю колеса в сборе: штамповка + Nokian Nordman 7 175/65 r14 Состояние отличное, покупались на прошлый сезон на Ford Fiesta, но машиной практически не пользовались.\",\"pictures\": \"https://89.img.avito.st/image/1/CIGv3LaypGiZdWZts_xAo1B_pG4PfaY,https://44.img.avito.st/image/1/0dlI97ayfTB-Xr81ZtCZ-7dUfTboVn8,https://01.img.avito.st/image/1/mPLJibayNBv_IPYe9a_Q0DYqNB1pKDY\",\"price\":12000}" http://localhost:8080/api/v1/ads`