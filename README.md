# Book Backend

## Endpoints
### POST /login
**Params:**

    Username string

	Password string

	Email    string

    token    string

**Response:**

	token string

### POST /signup
**Params:**

    Username string

	Password string

	Email    string

**Response:**

	token string

### POST /review
**Params:**

	Content   string

	Rating    uint8

	WorksID   string

    token     string

**Response:**

	reviewid string

### GET /review\[reviewid\]
**Params:**

	N/A

**Response:**

	Content string

	Username string

	Rating uint8


### DELETE /review\[reviewid\]
**Params:**

    token    string

**Response:**

	status

### PUT /review


*not yet implemented*
