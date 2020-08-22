
# 3. password security

Date: 21-08-2020 21:36:48

## Status

Accepted

## Context

The application will need to accept a username and password combination
for authentication with the external senturion service. We need to securely 
use the credentials and not leak them.

## Decision

Use an .env file for testing and eventually migrate to the keyring. 

Use the keyring for enhanced security.

## Consequences

We have time to analyze the complexity of the keyring solution.