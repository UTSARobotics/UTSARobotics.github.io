# New Architecture

This document presents a proposal for the future structure of the infrastructure for the utsaras.org website.

## Old Architecture

Under the current (old) model, the website is being hosted in a work-in-progress state in which our site is located within web/pub and forwarded to the open internet via some reverse proxy program. The plan if this architecture is to stay in place is for the router for the site to eventually encompass all of the site's functionality and be bundled into a Docker image to be forwarded by some reverse proxy.

## New Architecture

Under the new architecture that this document proposes, all static pages (within reason[^1]) should be handled by the reverse proxy program. The custom backend/router will only be responsible for dynamic pages such as those that require a user to login.

This architecture requires a new folder structure:

```
```
.
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── Makefile
├── users.db
└── web
    ├─── static
    ├─── assets
    └─── templates # this folder holds golang html/templates
```
```

[^1]: When it _makes sense_ and does not hinder the development and maintenance of the project.
