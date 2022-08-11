# Ory CTF for Def Con 30

Details will follow soon.

## Installation

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

## Running Remote

To run the set up on a remote system, use:

```bash
$ docker compose -f docker-compose.remote.yml up --build -d --force-recreate
```
