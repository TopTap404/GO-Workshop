```bash
go mod tidy
go run .
```

เปิด:
- http://localhost:3000

```bash
docker compose -f docker-compose.mysql.yml up -d --build
```

- API: http://localhost:3000
- phpMyAdmin: http://localhost:8081

phpMyAdmin login:
- Username: `workshop`
- Password: `workshoppass`
