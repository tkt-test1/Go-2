# 🧠 Go Task Management API

Go + Gin + GORM + JWT を使った本格派のタスク管理REST APIです。  
ユーザー登録・ログイン機能に加え、  
**タスクのCRUD・ステータス履歴管理・検索機能**を備えた中〜上級者向け構成です🚀

---

## 🌈 特徴

- 🧍‍♀️ ユーザー認証（JWT）
- ✅ タスクの作成 / 更新 / 削除 / 一覧
- 🔄 ステータス変更時の履歴記録
- 🔍 タイトル・ステータスでの検索フィルタ
- 🗂 クリーンなディレクトリ構成
- 💾 GORMでDB操作（デフォルトはSQLite）

---

## 🧩 技術スタック

| カテゴリ | 技術 |
|-----------|------|
| 言語 | Go 1.22+ |
| Webフレームワーク | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io) |
| 認証 | [JWT (golang-jwt)](https://github.com/golang-jwt/jwt) |
| DB | SQLite（簡易動作確認用） |

---

## 📁 ディレクトリ構成

```
go-task-api/
├── main.go
├── controllers/
│   ├── user_controller.go
│   └── task_controller.go
├── models/
│   ├── user.go
│   ├── task.go
│   └── history.go
├── routes/
│   └── routes.go
├── utils/
│   └── jwt.go
├── go.mod
└── go.sum
```

---

## ⚙️ セットアップ手順

### 1️⃣ クローン
```bash
git clone https://github.com/<yourname>/go-task-api.git
cd go-task-api
```

### 2️⃣ 依存関係をインストール
```bash
go mod tidy
```

### 3️⃣ 起動
```bash
go run main.go
```

サーバーが `http://localhost:8080` で起動します 🎉

---

## 🔐 API エンドポイント

### 👤 ユーザー登録
`POST /api/register`

```json
{
  "name": "Taro",
  "email": "taro@example.com",
  "password": "pass123"
}
```

✅ Response
```json
{ "message": "User registered successfully" }
```

---

### 🔑 ログイン
`POST /api/login`

```json
{
  "email": "taro@example.com",
  "password": "pass123"
}
```

✅ Response
```json
{ "token": "eyJhbGciOiJIUzI1NiIs..." }
```

---

### 🧾 タスク作成
`POST /api/tasks`

Header に JWT を渡す：
```
Authorization: <token>
```

Body:
```json
{
  "title": "Design API",
  "status": "todo"
}
```

✅ Response
```json
{
  "id": 1,
  "title": "Design API",
  "status": "todo",
  "user_id": 1
}
```

---

### 📋 タスク一覧取得（検索付き）
`GET /api/tasks?title=API&status=todo`

✅ Response
```json
[
  {
    "id": 1,
    "title": "Design API",
    "status": "todo",
    "histories": []
  }
]
```

---

### 🔄 タスク更新（履歴保存つき）
`PUT /api/tasks/:id`

```json
{
  "title": "Design API v2",
  "status": "in_progress"
}
```

✅ Response
```json
{
  "id": 1,
  "title": "Design API v2",
  "status": "in_progress"
}
```

履歴テーブルにステータス変更が保存されます👇  
```json
{
  "task_id": 1,
  "old_status": "todo",
  "new_status": "in_progress"
}
```

---

### ❌ タスク削除
`DELETE /api/tasks/:id`

✅ Response
```json
{ "message": "Task deleted" }
```

---

## 🧮 DB仕様

SQLiteを使用（`task.db` が自動生成）  
モデル関係：

```
User 1 ── * Task 1 ── * TaskHistory
```

- User: ユーザー
- Task: ユーザーのタスク
- TaskHistory: ステータス変更履歴  

---

## 🧠 学べるポイント

- Ginを使ったREST API設計  
- GORMによるモデルリレーションと検索条件指定  
- JWTによる認証とミドルウェア  
- ステータス履歴のビジネスロジック設計  
- Goらしいディレクトリ構成（controllers / models / utils 分割）

---

## 🚀 今後の拡張アイデア

- 🐳 Docker + PostgreSQL 対応  
- 🧾 Swagger（`swaggo/swag`）でAPIドキュメント自動生成  
- 🔁 Refresh Token実装  
- 🧪 CI/CDパイプライン（GitHub Actions）  
- 🌐 React or Next.jsでフロント接続  


