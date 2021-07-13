# Gitlab Fundamentals

This is an exemplary project created for the needs of SoftServe "The Talk".

### Project contains:
- multi-stage Dockerfile which is used to test and build simple http server written in Go,
- Makefile with set of commands for the needs of Gitlab runner,
- .gitlab-ci.yaml configuration file with defined stages.

### Defined stages in Gitlab CI/CD
- test: runs app tests
- build: build the app
- tag: creates and save to artifactory `$TAG` variable with git tag value
- publish: prepares release file
- release: create Gitlab release via API and adds binaries prepared in publish stage as release package

## Author
Hubert Siwik: siwik.hubert@gmail.com
