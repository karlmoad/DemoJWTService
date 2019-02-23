# DemoJWTService


This is a demonstration web (REST) service showing the methods used to integrate and verify JWT tokens into the service calls.  Microsoft Azure Active Directory is used as the token idp, hence the validation logic is specific to the Microsoft implementation for now.

**NOTE:** This service is compliant with [IETF RFC6750](https://tools.ietf.org/html/rfc6750) and utilizes the Bearer authorization scheme

The service will extract the JWT authorization token from the request header and verify the signature against the appropriate microsoft certificates as well as verifying the token's **aud** value against the configured MS Azure app key

Prerequisites:
- A registered Microsoft Azure Active Directory application key.
- the app key must be stored in an environment variable, **AZURE_APP_KEY**

```
    $ export AZURE_APP_KEY=<app key here>

```
- GO (minimal version 1.11) language compiler (see https://golang.org/doc/install for your specific platform).

**NOTE:** This project utilizes the go modules dependency management system. 


Optional:
- Make (see https://www.gnu.org/software/make/ , windows users: [see here](http://gnuwin32.sourceforge.net/packages/make.htm) )


## Build for local execution

Using Make:

``` 
    $ make 
```

Using Go Build:

```
    $ go build -o ./dist/DemoJWTService service.go
```


This will build the application to run as a local executable, once built navigate to the **dist** directory and execute the **WorkspaceService** executable

```
    $ cd dist
    $ ./DemoJWTService

```

Service will be available at **http://localhost:30200/...**

