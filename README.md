# Hello API
This is an improved version of the current hello-api we use in production. It will use less memory and be cheaper to run in production, and it will scale, expand to additional words, and be more stable:

## Dependencies
- Go Version 1.18

## Setup

### Install Go
`sudo make setup`

### Upgrade Go
`sudo make install-go`

## Release Milestone

### V0 (1 day)
- [ ] Onboarding Documentation
- [ ] Simple API Response
- [ ] Unit tests
- [ ] Running somewhere other than the dev machine

### V1 (7 days)
- [ ] Create translation endpoint
- [ ] Store translations in short-term storage 
- [ ] Call existing service for translation
- [ ] Move towards long-term storage
