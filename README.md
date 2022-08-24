# Capture The Flag // Voting Village // Def Con 30

Welcome to the github.com/ory CTF at DEF CON 30! Explore a vulnerable, open-source digital election system and capture the flag!

Join the [community slack](https://slack.ory.sh) or have a look at the video summary:

[![Ory Capture The Flag Interactive Summary](https://img.youtube.com/vi/Mx8LNRndsO8/0.jpg)](https://www.youtube.com/watch?v=Mx8LNRndsO8 "Ory Capture The Flag Interactive Summary")

## Targets

This challenge runs five services. They mock a basic election system used by **authenticated** users (election workers) to submit their voting districts results. This is not a service for voters. However, everyone can sign up and see the already submitted results.

The services are all open source:

- [Ory Oathkeeper](https://github.com/ory/oathkeeper): reverse proxy for all other services
- [Ory Kratos](https://github.com/ory/kratos): authentication and session management
- [Ory Keto](https://github.com/ory/keto): authorization and access control
- Backend (this repo): the actual election system backend
- Postgres: the database

The target of this CTF is the **backend** service. Vulnerabilities found in the open source **Ory Oathkeeper**, **Ory Kratos**, and **Ory Keto** projects can be reported through our [bug bounty program](https://hackerone.com/ory_corp) and give you bounties between 100$ (low) and 3,000$ (critical). On top, we will add another 100$ for any submission done during DEF CON 30 after you talked to us personally at the Voting Machine village.

## Running Locally

Open source also means you can investigate the services locally.

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
