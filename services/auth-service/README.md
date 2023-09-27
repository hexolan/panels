# Auth Service

## Event Documentation

* Events Produced:
  * N/A

* Events Consumed:
  * **Topic:** "``user``" | **Schema:** "``UserEvent``" protobuf
    * Type: ``"deleted"`` | Data: ``User``

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

**Kafka:**

``KAFKA_BROKERS`` (Required)

* e.g. "localhost:9092" or "localhost:9092,localhost:9093"

---

**Other:**

``JWT_PRIVATE_KEY`` (Required)

* RSA Private Key (used to encode JWT tokens) represented in base64

* Generating:
  * Using OpenSSL: ``openssl genrsa -out private.pem 2048``
  * Getting Base64 Representation: ``python -c "import base64;private_key=open('private.pem', 'r').read();print(base64.standard_b64encode(private_key.encode('utf-8')).decode('utf-8'))"``

``JWT_PUBLIC_KEY`` (Required)

* RSA Public Key (used to verify JWT tokens) represented in base64

* Generating:
  * Using OpenSSL: ``openssl rsa -in private.pem -pubout -out public.pem``
  * Getting Base64 Representation: ``python -c "import base64;private_key=open('public.pem', 'r').read();print(base64.standard_b64encode(private_key.encode('utf-8')).decode('utf-8'))"``

``PASSWORD_PEPPER`` (Required)

* A secret salt applied globally along with generated password salts.

* Generate with ``python -c "import secrets;print(secrets.token_hex(nbytes=16))"`` or similar method.
