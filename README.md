# 🧠 Memo App Backend

## 🛠️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/thyms-c/be-memo-app.git
cd be-memo-app
```

### 2. Install Dependencies
```bash
go mod download
go mod tidy
```

### 3. Set up environment variables

```bash
cp .env.example .env
```

### 4. Run with Docker Compose

```bash
docker compose up -d
```
---

## 🖍️ How to import Node-Red Flow

### 1. open Node-Red on 

```bash
http://localhost:1880
```

### 2. import flow from file
```bash
node-red/auth_flow.json
```

### 3. Deploy flow 
```bash
click on 'Deploy' button
```
---


## 🔐 Roles & Permissions

- **Admin**: Full access to all memo operations
- **User**: Can only manage their own memos

---

## 🧪 Testing

For testing purposes, mock data or tools like [Postman](https://www.postman.com/) can be used to interact with the API.

