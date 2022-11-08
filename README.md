## Chula SSO

A standalone chula sso based on [chula-sso](https://account.it.chula.ac.th/wiki/doku.php?id=how_does_it_work)

Being used in development by [Piyaphat Pinyo](https://www.github.com/remove158)

## Usage

Docker image is available at [Docker](https://github.com/remove158/chula-sso/pkgs/container/chula-sso)

1. Install docker
2. Download the image: `docker pull ghcr.io/remove158/chula-sso:latest`
3. Run: `docker run -p 8080:8080 ghcr.io/remove158/chula-sso:latest`

## Preview

![img](./preview-1.png)

## Configuration

The configuration is done by the following environment variables. (see [`docker run`](https://docs.docker.com/engine/reference/commandline/run/#set-environment-variables--e---env---env-file) on how to set it)

| Name         | Environment Variable | Default |
| ------------ | -------------------- | ------- |
| DeeAppId     | DEE_APP_ID           | test    |
| DeeAppSecret | DEE_APP_SECRET       | test    |
| Port         | PORT                 | 8080    |

## Endpoint

### GET /login

- Request

```sh
curl --request GET \
  --url 'http://localhost:8080/login?service=https://www.google.com'
```

- Response 302 Found

```js
// (on-success) 302 Location: https://www.google.com?ticket=86966dc5-2049-428f-88fe-2d78a5985d38
```

- Response 400 Bad Request

```json
{
	"error": "Key: 'GetLoginRequest.Service' Error:Field validation for 'Service' failed on the 'required' tag"
}
```

### GET, POST /serviceValidation

- Request

```sh
curl --request POST \
  --url 'http://localhost:8080/serviceValidation' \
  --header 'DeeAppId: test' \
  --header 'DeeAppSecret: test' \
  --header 'DeeTicket: test'
```

- Response 200 (application/json)

```json
{
	"uid": "6000000021",
	"username": "6000000021",
	"gecos": "admin",
	"password": "123456",
	"disable": false,
	"roles": ["faculty"],
	"firstname": "Faculty",
	"lastname": "จุฬา",
	"firstnameth": "คณะ",
	"lastnameth": "จุฬา",
	"ouid": "6000000021",
	"email": "admin@chula.ac.th"
}
```

- Response 401 Unauthorized

```json
{
	"error": "ticket not found"
}
```

## Diagram

### Login

```mermaid
sequenceDiagram
    Client->>SSO: GET: /login?service=https://www.google.com
    alt provide service
        SSO->>Client: 302 Found
    else didn't provide service
        SSO->>Client: 400 Bad Request
    end
```
