# Gateway Service

## Configuration

### Environment Variables

**RPC Service Addresses:**

``PANEL_SVC_ADDR`` (Required)

* e.g. "panel-service:9090"

``POST_SVC_ADDR`` (Required)

* e.g. "post-service:9090"

``USER_SVC_ADDR`` (Required)

* e.g. "user-service:9090"

``AUTH_SVC_ADDR`` (Required)

* e.g. "auth-service:9090"

``COMMENT_SVC_ADDR`` (Required)

* e.g. "comment-service:9090"

**Auth:**

``JWT_PUBLIC_KEY`` (Required)

* RSA Public Key encoded in base64 (used to verify JWT tokens created by the ``auth-service``)
