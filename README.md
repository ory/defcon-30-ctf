# Ory CTF for Def Con 30

Welcome to the Ory CTF at DEF CON 30! This is your chance to get some awesome Ory swag and bounty! Head over to the [remote challenge](https://defcon.ory.dev/) to get started.

## Targets

This challenge consists of five services. They mock a basic election system used by **authenticated** users (election workers) to submit their voting districts results. This is not a service for voters. However, everyone can sign up and see the already submitted results.

The services are:

- [Ory Oathkeeper](https://github.com/ory/oathkeeper): reverse proxy for all other services
- [Ory Kratos](https://github.com/ory/kratos): authentication and session management
- [Ory Keto](https://github.com/ory/keto): authorization and access control
- Backend (this repo): the actual election system backend
- Postgres: the database

The target of this CTF is the **backend** service. Vulnerabilities found in **Oathkeeper**, **Kratos**, and **Keto** can be reported through our [bug bounty program](https://hackerone.com/ory_corp) and give you bounties between 100$ (low) and 3,000$ (critical). On top, we will add **another** 100$ for any submission done during DEF CON 30 after you talked to us personally at the Voting Machine village.

## Out of Scope

We kindly ask you **not to do** denial of service or brute-force attacks against our remote services. Doing so will ruin the fun for everyone.

## Local Installation

You'll need to have Docker installed and this repository checked out to start the challenge:

```bash
$ git clone https://github.com/ory/defcon-30-ctf.git
$ cd defcon-30-ctf
$ docker compose up -d --build --force-recreate
```

Once the services are running, you are able to access them at:

```
http://localhost:5050
```
