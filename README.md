# Структура проекта<br/>
__src__ - директория с файлом main.go<br/><br/>
__storage__ - директория, где находится файл db_helper.go для работы с базой данных. Также в директории есть поддиректория db, в которую будет создана база данных<br/><br/>
Для работы с базой данных используется __sqlite__<br/>

├── go.mod<br/>
├── go.sum<br/>
├── src<br/>
│   └── main.go<br/>
└── storage<br/>
|    ├── db<br/>
|    └── db_helper.go<br/>

# Пример работы<br/>
![image](https://github.com/user-attachments/assets/9b8e9806-0a6e-41c1-bec1-eb09778989d3)

## Добавление записи
![image](https://github.com/user-attachments/assets/9bbe54a4-e548-4535-b0ad-0abb67542be1)
![image](https://github.com/user-attachments/assets/e420a150-8638-4055-abd2-3a0d724db710)

## Удаление записи
![image](https://github.com/user-attachments/assets/02014af1-fa05-469d-b579-c4faf97693d5)
![image](https://github.com/user-attachments/assets/c6839ba5-42f0-4b9b-9cf7-c934c4e6a205)

## Отметка о завершённости записи
![image](https://github.com/user-attachments/assets/58de325f-afdc-4d61-8322-e9d2f38fdb7b)

## Отметка о незавершённости записи
![image](https://github.com/user-attachments/assets/f80dd398-0153-4c6e-9125-5e8429712d27)

## Изменение записи
![image](https://github.com/user-attachments/assets/5be587a4-b208-489a-9005-21d3909e07fc)<br/>
![image](https://github.com/user-attachments/assets/336c4951-0ae5-4231-888e-21c52f4a3d77)
