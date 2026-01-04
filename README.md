# Workshop: Go Fiber + GORM

โปรเจคนี้ทำให้:
- รันแบบปกติบนเครื่องได้
- รันด้วย Docker ได้
- ใช้ SQLite แบบ **PURE GO** (ไม่ต้อง CGO) เป็นค่าเริ่มต้น ✅
- (มีตัวเลือก) รันกับ MySQL + phpMyAdmin ใน Docker ได้ ✅

## Run บนเครื่อง (ค่าเริ่มต้น: SQLite pure-go)

```bash
go mod tidy
go run .
```

เปิด:
- http://localhost:3000
- http://localhost:3000/health
- http://localhost:3000/api/users

> ไฟล์ DB จะเป็น `workshop.db` ในโฟลเดอร์โปรเจค

## Run ด้วย Docker (ค่าเริ่มต้น: SQLite pure-go)

```bash
docker compose up -d --build
```

เปิด:
- http://localhost:3000

SQLite DB จะถูกเก็บใน volume `app_data` (path ใน container: `/app/data/workshop.db`)

## Run ด้วย Docker + MySQL + phpMyAdmin (ตัวเลือก)

```bash
docker compose -f docker-compose.mysql.yml up -d --build
```

เปิด:
- API: http://localhost:3000
- phpMyAdmin: http://localhost:8081

phpMyAdmin login:
- Username: `workshop`
- Password: `workshoppass`
- Server: เว้นว่าง (หรือเลือก `mysql` ถ้ามีให้เลือก)

## Environment

- `DB_DIALECT`: `sqlite` (default) | `mysql`
- `SQLITE_PATH`: path ไปไฟล์ sqlite (default: `workshop.db`)
- `DB_DSN`: MySQL DSN (ใช้เมื่อ `DB_DIALECT=mysql`)
"# GO-Workshop" 
