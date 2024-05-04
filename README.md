# Book Backend

## Endpoints
### POST /login
**Params:**

    Username string

	Password string

	Email    string

    token    string

### POST /signup
**Params:**

    Username string

	Password string

	Email    string

### POST /review
**Params:**

	Content   string

	Rating    uint8

	Username  string

	WorksID   string

    token     string

### GET /review\[reviewid\]
**Params:**

	N/A

### DELETE /review\[reviewid\]
**Params:**

    token    string

### PUT /review


*not yet implemented*
