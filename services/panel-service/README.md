# Panel Service

## Event Documentation

* Events Produced:
  * **Topic:** "``panel``" | **Schema:** "``PanelEvent``" protobuf
    * Type: ``"created"`` | Data: ``Panel``
    * Type: ``"updated"`` | Data: ``Panel``
    * Type: ``"deleted"`` | Data: ``Panel`` (only with ``Panel.id`` attribute)

* Events Consumed:
  * N/A

## Configuration

### Environment Variables

**PostgreSQL:**

``POSTGRES_USER`` (Required)

* e.g. "postgres"

``POSTGRES_PASS`` (Required)

* e.g. "postgres"

``POSTGRES_HOST`` (Required)

* e.g. "localhost:5432"

``POSTGRES_DATABASE`` (Required)

* e.g. "postgres"

---

**Redis:**

``REDIS_HOST`` (Required)

* e.g. "localhost:6379"

``REDIS_PASS`` (Required)

* e.g. "redis"

---

**Kafka:**

``KAFKA_BROKERS`` (Required)

* e.g. "localhost:9092" or "localhost:9092,localhost:9093"

---

**Other:**

``LOG_LEVEL`` (Default: "info")

* i.e. "debug", "info", "warn", "error", "fatal", "panic" or "disabled"
