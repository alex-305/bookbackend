# Book Backend

## Endpoints
### POST /login
**Request Body (application/json):**

    username string

	password string

	email    string

    token    string

**Response:**

	token string

### POST /signup
**Request Body (application/json):**

    username string

	password string

	email    string

**Response:**

	token string

### GET /user/\[username\]
**Request Body:**

	N/A

**Response:**

	username    string

	email       string

	description string

	join_date   time.Time

### POST /review
**Request Body (application/json):**

	content   string

	rating    uint8

	worksid   string

    token     string

**Response:**

	reviewid string

### GET /review\[reviewid\]
**Request Body:**

	N/A

**Response:**

	content string

	username string

	rating uint8


### DELETE /review\[reviewid\]
**Request Body (application/json):**

    token    string

**Response:**

	status 


*not yet implemented*
