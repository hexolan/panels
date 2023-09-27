# User Service

## Event Documentation

* Events Produced:
  * **Topic:** "``user``" | **Schema:** "``UserEvent``" protobuf
    * Type: ``"created"`` | Data: ``User``
    * Type: ``"updated"`` | Data: ``User``
    * Type: ``"deleted"`` | Data: ``User``

* Events Consumed:
  * N/A

## Configuration

### Environment Variables

**MongoDB:**

``MONGODB_URI`` (Required)

* e.g. "mongodb://mongo:mongo@localhost:27017/"

---

**Kafka:**

``KAFKA_BROKERS`` (Required)

* e.g. "localhost:9092" or "localhost:9092,localhost:9093"
