# myapp
```
├── cmd
│   └── main.go          # จุดเริ่มต้นของโปรแกรม
├── pkg
│   └── config.go        # การตั้งค่าคอนฟิก (เช่น การเชื่อมต่อกับฐานข้อมูล)
└── internal
    ├── core
    │   ├── domain
    │   │   └── user.go # โมเดลข้อมูลผู้ใช้
    │   ├── ports
    │   │   └── user.go # พอร์ตของแอพพลิเคชัน (interface ที่ใช้กับ Repository และ Service)
    │   └── services
    │       └── user_service.go # บริการที่ใช้ในการจัดการผู้ใช้ (Business Logic)
    ├── handlers
    │   └── user_handler.go # Handlers สำหรับจัดการ HTTP Request
    └── repositories
        └── user_repository.go # การจัดการฐานข้อมูลสำหรับผู้ใช้ (ใช้ GORM)
```