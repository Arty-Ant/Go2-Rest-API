## Инструкция для установки  и настройки MySQL на Ubuntu 

### Установка  

1. Обновить кеш пакетов Linux
```bash
sudo apt update
```

2. Установить `mysql-server`
```bash
sudo apt install mysql-server
```

3. Запустить сервис MySQL
```bash
sudo systemctl start mysql.service
```


### Настройка

1. Запуск `mysql cli` клиента
```bash
sudo mysql
```

2. Установка пароля для пользователя `root`
```mysql
mysql> ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
```

3. Выход из `cli` клиента
```mysql
mysql> exit
```

4. Опционально. Запуск security script
```bash
sudo mysql_secure_installation
```

5. Если нужен вход через `sudo mysql`
```mysql
mysql> ALTER USER 'root'@'localhost' IDENTIFIED WITH auth_socket;
```

6. Вход с помощью `cli` клиента по паролю
```bash
mysql -u root -p
```

7. Создание пользователя для работы с БД
```mysql
mysql> CREATE USER 'username'@'localhost' IDENTIFIED WITH mysql_native_password BY 'user_password';
```

8. Назначение прав доступа. Общий синтаксис
```mysql
mysql> GRANT PRIVILEGE ON database.table TO 'username'@'host';
```

9. Пример 1
```mysql
mysql> GRANT CREATE, ALTER, DROP, INSERT, UPDATE, INDEX, DELETE, SELECT, REFERENCES, RELOAD on *.* TO 'username'@'localhost' WITH GRANT OPTION;
```

10. Пример 2
```mysql
mysql> GRANT ALL PRIVILEGES ON *.* TO 'sammy'@'localhost' WITH GRANT OPTION;
```

### Создание БД

1. Запуск от пользователя
```bash
mysql -u username -p
```

2. Создание БД
```mysql
mysql> CREATE DATABASE mydb;
```

3. Выход
```mysql
mysql> exit
```
