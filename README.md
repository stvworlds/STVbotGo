# STVbotGo

A Discord bot written in Go — a rewrite of the original STVbot project (previously in JavaScript).

## Tech Stack

- **Go** 1.26+
- **[discordgo](https://github.com/bwmarrin/discordgo)** — Discord API client
- **[ffmpeg-go](https://github.com/u2takey/ffmpeg-go)** — media conversion
- **[godotenv](https://github.com/joho/godotenv)** + **yaml.v3** — configuration

## Features

| Module | Description |
|---|---|
| `antiscam` | Detects spam (2+ attachments), bans users, logs to designated channels |
| `tickets` | Full ticket system with modals, buttons, and role-based access |
| `newscreate` | Modal-driven news/blog post creation via external API |
| `convert` | Converts video to GIF using FFmpeg |
| `autothreadopen` | Auto-creates threads in configured channels |
| `cats` / `dog` | Responds to keywords with animal images from external API |
| `joinmessage` | Sends a message when a user joins the server |
| `whitelist` | Whitelist management commands |
| `ping` | Health-check ping/pong command |

## Configuration

Config is loaded from `./configs/` (local) or `/configs/` (Docker) as YAML, with optional `.env` overrides for:

- `TOKEN` — Discord bot token
- `APP_ENV` — environment (`dev` / `main`)
- `API_KEY` — external API key
- `LOG_LEVEL` — logger min level (`debug` / `info` / `warn` / `error`; default `debug` in `dev`, else `info`)
- `LOG_WEBHOOK_URL` — optional Discord webhook for remote application logs

## Project Structure

```
cmd/STVbotGo/          # entry point
internal/
  app/                 # bot startup, session, signal handling
  config/              # config loading
  logger/              # console + optional Discord webhook logger
  handlers/            # one package per feature module
```
