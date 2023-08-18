run command

rename .env.example -> .env

```
go run cmd/main.go
```

- POST user/register

  - from data
  - ```
    login = exampleuser
    password = 123456
    name = user 1
    age  = 18
    ```

- POST user/auth
  - ```
      {
          "login":"exampleuser",
          "password":"123456"
      }
    ```
- GET user/:name
- GET user/phone?phone_id=1
- POST user/phone
  - ```
      {
          "phone": "998999999999",
          "isFax": false,
          "description": "other desc",
          "userId": 1
      }
    ```
- PUT user/phone

  - ```
    {
        "id":1,
        "phone": "998999999999",
        "isFax": false,
        "description": "other desc",
        "userId": 1
    }
    ```

- DELETE user/phone?phone_id=1
