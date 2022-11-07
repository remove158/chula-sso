## Chula SSO
A standalone chula sso based on [chula-sso](https://account.it.chula.ac.th/wiki/doku.php)

Being used in development by [Piyaphat Pinyo](https://www.github.com/remove158)

A chula sso [wiki](https://account.it.chula.ac.th/wiki/doku.php?id=how_does_it_work) is available.

You can clone this project
## Preview
![img](./preview-1.png)
## Setup
- create `.env` follow this template
```
DEE_APP_SECRET=
DEE_APP_ID=
PORT=
```
> DeeAppSecret (default="test") </br>
> DeeAppId (default="test") </br>
> PORT (default="8080") 

## Endpoint
- GET: /login?service={your-service-here}
```
GET: /login?service=https://www.google.com

Response: 
(on-success) 301 Redirect to https://www.google.com?ticket=86966dc5-2049-428f-88fe-2d78a5985d38
```
- GET,POST: /serviceValidation

```
GET,POST : /serviceValidation

Header : {
	DeeAppId : string, 
	DeeAppSecret: string, 
	DeeTicket: string
}

Response: 
(on-success) 200
{ 
	"uid": "string",
	"username": "string",
	"gecos": "string",
	"disable":  false,
	"roles": ["student"],
	"firstname": "string",
	"firstnameth": "string",
	"lastname": "string",
	"lastnameth": "string",
	"ouid": "string",
	"email": "string"
}

(on-fail) 401

```
## Diagram 
![chula-sso](https://account.it.chula.ac.th/wiki/lib/plugins/plantuml/img.php?width=0&height=0&title=PlantUML%20Graph&align=&version=2011-07-16&md5=f8f62ed0420593df3f158216f286b820)
