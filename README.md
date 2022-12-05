# static-web

A deployment system for static webpages

## Features

- [x] upload service to place files into a storage
  - [x] upload files
  - [x] list deployments
  - [x] remove deployment
  - [x] use random subdomain
  - [x] use custom subdomain
  - [x] store meta data (user, date of last upload)
  - [ ] use custom FQDN
  - [ ] rolling update (upload new files, then switch to new version)
  - [ ] register domain in caddy
  - [ ] shutdown deployment and remove files after some time (no re-upload)
- [ ] webserver
  - [ ] 200.html for SPA
  - [ ] custom 404 page
  - [ ] automatic ssl support (should this even be done by caddy?)
  - [ ] automatic https redirect (should this even be done by caddy?)
- [x] cli to easily deploy static webpages
  - [x] upload files
  - [x] list deployments
  - [x] remove deployment
  - [x] use random subdomain
  - [x] use custom subdomain
  - [x] use custom FQDN
  - [ ] proper authentication
  - [ ] re-upload only changed / sync files
