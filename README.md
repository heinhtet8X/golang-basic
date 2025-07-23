# Go REST API Project Folder Structure

Go project တွေမှာ တရားဝင် သတ်မှတ်ထားတဲ့ folder structure ကြီးရယ်လို့ မရှိပေမယ့်၊ REST API project တွေအတွက်တော့ အောက်ပါဖွဲ့စည်းပုံက အသုံးအများဆုံးနဲ့ အဆင်ပြေဆုံး ဖြစ်ပါတယ်။ ဒီ structure က project ကို maintain လုပ်ရလွယ်ကူစေပြီး scalability ကိုလည်း မြှင့်တင်ပေးပါတယ်။

## 📁 Project Structure Overview

```
my-api-project/
├── cmd/
│   └── api/
│       └── main.go         # API server ကို run ဖို့ main function
├── internal/
│   ├── app/                # Business logic (Service Layer)
│   ├── domain/             # Data models & interfaces (Domain Layer)
│   ├── repository/         # Data access logic (Repository Layer)
│   ├── transport/          # API handlers (Controller Layer)
│   │   └── http/
│   ├── config/             # Configuration loading
│   └── middleware/         # HTTP middleware
├── api/
│   └── openapi.yaml        # OpenAPI/Swagger API definition
├── configs/
│   ├── config.dev.yaml     # Development config
│   └── config.prod.yaml    # Production config
├── migrations/
│   ├── 001_create_users_table.sql
│   └── 002_create_products_table.sql
├── scripts/                # Utility scripts
├── .env.example            # Environment variables example
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
└── README.md               # Project's main README file
```

## 📂 Folder တစ်ခုချင်းစီရဲ့ ရည်ရွယ်ချက်

### `cmd/`

သင့် project ရဲ့ executable application တွေအတွက် main.go ဖိုင်တွေကို ဒီမှာထားပါတယ်။ REST API project အတွက်ကတော့ api ဆိုတဲ့ sub-folder ထဲမှာ main.go ဖိုင်ကို ထားလေ့ရှိပါတယ်။  
**`cmd/api/main.go`**: သင့် REST API server ကို စတင်ဖို့အတွက် entry point ပါ။ ဒီမှာ HTTP router ကို initialize လုပ်တာ၊ database connection ဖွင့်တာ၊ service နဲ့ repository တွေကို ချိတ်ဆက်တာတွေ လုပ်ဆောင်ပါတယ်။

---

### `internal/`

Go ရဲ့ "internal" package rule ကြောင့် ဒီထဲက code တွေကို သင့်ရဲ့ my-api-project module ထဲက package တွေကလွဲလို့ အခြား external Go module တွေကနေ import လုပ်လို့မရပါဘူး။ ဒါဟာ application ရဲ့ core logic တွေကို external dependencies တွေကနေ ကာကွယ်ပေးပြီး encapsulation ကို မြှင့်တင်ပေးပါတယ်။

- **`internal/app/`** – Business Logic (Service Layer)  
  Use case တွေနဲ့ business rules တွေကို ကိုင်တွယ်တဲ့နေရာ။  
  ဥပမာ - `user_service.go`, `product_service.go` စတာတွေ။

- **`internal/domain/`** – Domain Models & Interfaces (Domain Layer)  
  Go structs (e.g. `User`) နဲ့ interfaces (e.g. `UserRepository`) တွေကို ဒီမှာသတ်မှတ်ထားပါတယ်။  
  `errors.go` မှာ custom error type တွေ ထည့်ထားလို့ရပါတယ်။

- **`internal/repository/`** – Repository Layer  
  Database access logic (CRUD) တွေပါ။  
  ဥပမာ - `user_repo.go` မှာ `domain.UserRepository` ကို implement လုပ်ပါတယ်။  
  Sub-folder များ `postgres/`, `mysql/` စသဖြင့် DB အလိုက်ခြားနိုင်ပါတယ်။

- **`internal/transport/http/`** – API Controllers  
  Incoming HTTP requests များကို ကိုင်တွယ်တဲ့ handler logic ပါ။  
  `user_handler.go` – `/users` route  
  `router.go` – Router setup, middleware တွေထည့်ခြင်း။

- **`internal/config/`** – Configuration Loader  
  Environment variables ဒါမှမဟုတ် YAML config files များကို parse/load လုပ်ခြင်း။

- **`internal/middleware/`** – HTTP Middleware  
  Authentication, logging, rate limiting စတာတွေကို middleware အနေနဲ့ implement လုပ်ထားပါတယ်။

---

### `api/`

- **`openapi.yaml`** – OpenAPI Specification  
  API documentation လုပ်ဖို့နဲ့ client SDKs တွေ generate လုပ်ဖို့အသုံးပြု။

---

### `configs/`

- `config.dev.yaml`, `config.prod.yaml` – YAML format config files  
  (Development / Production settings)

---

### `migrations/`

- Database schema ပြောင်းလဲမှုများအတွက် SQL files  
  ဥပမာ - `001_create_users_table.sql`, `002_create_products_table.sql`

---

### `scripts/`

- Project build, deployment, setup shell scripts  
  ဥပမာ - `build.sh`, `deploy.sh`

---

### `.env.example`

- Environment variables များရဲ့ ဥပမာဖိုင်။  
  Developers တွေ setup လုပ်ရာမှာ အထောက်အကူဖြစ်စေပါတယ်။

---

### `go.mod / go.sum`

- Go module တွေအတွက် dependency metadata  
  `go mod init` နဲ့ project စတင်လိုက်တာနဲ့ auto-generate ဖြစ်ပါတယ်။

---

ဒီ Folder Structure က Project ရဲ့ maintainability, scalability, modularity တွေကို အထောက်အကူပြုနိုင်ပါတယ်။

@copy from gemini ai
