# SNMP-Connection

## 📡 SNMP Connector (Go + PostgreSQL)

A simple SNMP data polling microservice written in Go using **Hexagonal Architecture**.
It retrieves SNMP data from enabled devices and stores it in a PostgreSQL database via REST APIs.

---

### 🧱 Project Structure

```plaintext
snmp-connector/
├── cmd/
│   └── main.go               # App entrypoint (wires everything)
├── internal/
│   ├── adapters/
│   │   ├── api/              # HTTP handlers (input adapter)
│   │   ├── snmp/             # SNMP communication logic (output adapter)
│   │   └── storage/          # PostgreSQL adapter (output adapter)
│   └── domain/               # Business logic + ports + models
```

---

### 🚀 Features

* Connect to SNMP-enabled devices
* Retrieve specific OIDs (like `sysName` and `sysUpTime`)
* Store data with timestamp and type in PostgreSQL
* List stored SNMP data via `/data` API
* Designed with **Hexagonal Architecture (Ports and Adapters)**

---

### 🧪 Example APIs

| Method | URL                                       | Description                              |
| ------ | ----------------------------------------- | ---------------------------------------- |
| `GET`  | `http://localhost:8080/poll?ip=127.0.0.1` | Fetch SNMP data from device and store it |
| `GET`  | `http://localhost:8080/data`              | List all stored SNMP data                |

---

### 🗂 Sample Output (JSON)

```json
[
  {
    "ID": 1,
    "DeviceIP": "127.0.0.1",
    "OID": ".1.3.6.1.2.1.1.5.0",
    "Value": "simulator-device",
    "Type": "OctetString",
    "RetrievedAt": "2025-07-12T13:40:12Z"
  },
  ...
]
```

---

### 🛠 How to Run

#### 1. Install Dependencies

```bash
go mod tidy
```

#### 2. Run PostgreSQL (Docker)

```bash
docker run --name pg-snmp -e POSTGRES_PASSWORD=123 -e POSTGRES_USER=mariam -e POSTGRES_DB=mydb -p 5432:5432 -d postgres
```

#### 3. Create Table

In `psql` or pgAdmin, run:

```sql
CREATE TABLE snmp_data (
    id SERIAL PRIMARY KEY,
    device_ip TEXT,
    oid TEXT,
    value TEXT,
    type TEXT,
    retrieved_at TIMESTAMP
);
```

#### 4. Start SNMP Simulator

Install:

```bash
pip install snmpsim
```

Run:

```bash
snmpsim-command-responder --agent-udpv4-endpoint=127.0.0.1:161
```

#### 5. Run the Go Application

```bash
go run cmd/main.go
```

---

### 🧪 Test with Postman or curl

```bash
curl "http://localhost:8080/poll?ip=127.0.0.1"
curl "http://localhost:8080/data"
```

---

### 📚 References

* SNMP Protocol: [https://en.wikipedia.org/wiki/Simple\_Network\_Management\_Protocol](https://en.wikipedia.org/wiki/Simple_Network_Management_Protocol)
* GoSNMP Library: [https://github.com/gosnmp/gosnmp](https://github.com/gosnmp/gosnmp)
* PostgreSQL Driver: [https://pkg.go.dev/github.com/lib/pq](https://pkg.go.dev/github.com/lib/pq)
* SNMP Simulator: [https://docs.lextudio.com/snmpsim/quick-start](https://docs.lextudio.com/snmpsim/quick-start)

---

### 🧠 Architecture Used: Hexagonal (Ports & Adapters)

| Layer       | Responsibility              | Files                                      |
| ----------- | --------------------------- | ------------------------------------------ |
| Domain      | Business logic & interfaces | `domain/*.go`                              |
| Application | Use cases                   | `domain/snmp_service.go`                   |
| Adapters    | SNMP logic + Storage + HTTP | `adapters/snmp`, `adapters/storage`, `api` |
| Entrypoint  | Bootstrap everything        | `cmd/main.go`                              |

This decouples the business logic from technology choices.

---



