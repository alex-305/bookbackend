# Book Backend

## Endpoints
### POST /login
**Headers**

	Authorization: Bearer token

**Request Body (application/json):**

    username string

	password string

	email    string

**Response:**

	token string

### POST /signup
**Headers**

	N/A

**Request Body (application/json):**

    username string

	password string

	email    string

**Response:**

	token string

### GET /user/\[username\]
**Headers**

	N/A

**Request Body:**

	N/A

**Response:**

	username    string

	email       string

	description string

	join_date   time.Time

### POST /review
**Headers**

	Authorization: Bearer token

**Request Body (application/json):**

	content   string

	rating    uint8

	volumeid   string

    token     string

**Response:**

	reviewid string

### GET /review\[reviewid\]
**Headers**

	N/A

**Request Body:**

	N/A

**Response:**

	content string

	username string

	rating uint8


### DELETE /review\[reviewid\]
**Headers**

	Authorization: Bearer token

**Request Body**

	N/A

**Response:**

	status 


### GET /user/\[username\]/reviews
**Headers**

	N/A

**Request Body**
	
	N/A
**Params**

	?sort_order=  **OPTIONS: post_date | rating**

	?sort_dir= **OPTIONS: ascending | descending**

	?page= integer>=0

	?amount= integer>=0


**Response**

	username  string

    volumeid   string

    reviewID  string

    content   string

    rating 	  uint8

    post_date date
