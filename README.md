# 📦 Open Inventory Manager
Open Inventory Manager is a lightweight, open-source platform for managing stock, assets, and supplies. It provides tools for tracking inventory levels, recording transactions, generating reports, and integrating with other systems. Designed to be modular and extensible, it can be adapted for retail, warehouse, office, or personal use.

A lightweight, high-performance inventory management backend built in **Go**, using:

* ⚡ Gin (HTTP framework)
* ⚡ BadgerDB (embedded key/value database)
* ⚡ Async persistence
* ⚡ JSON API
* ⚡ Fully thread-safe in-memory inventory model

---

## 🚀 Features

* Add, update, delete, and fetch with filters inventory items
* In-memory collection with mutex protection
* Automatic async persistence to BadgerDB
* Unique ID generation based on timestamp
* Supplier metadata support
* Fast API built with Gin
* Minimal CPU & memory usage
* Runs in Docker (multi-stage build)
* Simple to deploy & scale


# 🛠 Installation

### Install via `go install` direct method (Recomended)

```bash
go install github.com/SumukhaS291299/Open-Inventory-Manager@latest
```

Make sure `~/go/bin` is in your PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

# ▶️ Running the server

```bash
go run main.go
```

Server runs by default on:

```
http://localhost:8080
```

---

# 🔧 API Endpoints

> POST /additem
``` json
{
  "attributes": {
    "name": "Widget",
    "description": "High-quality widget",
    "color": "Red",
    "category": "Gadgets",
    "location": "A1",
    "unit_price": 19.99,
    "stock_level": 100
  },
  "supplier": {
    "name": "Supplies",
    "supplier_type": "wholesale"
  }
}
```

> DELETE /deleteitem

>GET /filteritem

>PUT /modifyitem


Use Httpie Collections.json to import the basic collection methods

---

# 💾 Persistence (BadgerDB)

All items are stored asynchronously in **BadgerDB**:

* Fast writes
* Durable
* Embedded (no external DB needed)
* Data survives server restarts

If you are using the image

By default it will persist a VOLUME in ["/data/db"] inside the container

---

Rules:

1. **All in-memory modifications happen inside the mutex.**
2. **DB writes happen asynchronously in goroutines.**
3. **Only local copies of items are used in goroutines (no shared memory).**

This ensures:

* Zero race conditions
* Maximum throughput
* Safe parallel API access

---

# 🐳 Docker Support

Build:

```bash
docker build -t inventory-manager .
```

Run:

```bash
docker run -p 8080:8080 inventory-manager
```

---

# 🧪 Python API Tester (Mock Data)

```python
pip install requests faker
```

Run:

```python
python python_mocker.py
```

Generates random inventory items and POSTs them to your server.

---

> ###  🧱 Multi-stage Dockerfile (included)

# 🙋 Contributing

PRs welcome!
Feel free to open issues for:

* Bug reports
* Feature suggestions
* Performance improvements
* Documentation updates

---

# 📜 License

MIT License — feel free to use for personal or commercial projects.
