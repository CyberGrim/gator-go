# Gator

Gator is a command-line RSS feed aggregator written in Go. It allows you to track your favorite blogs and news sites directly from your terminal, storing posts in a PostgreSQL database for easy browsing.

## Tech Stack

- **Language:** [Go](https://go.dev/) (1.23+)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **SQL Tooling:** [sqlc](https://sqlc.dev/) (type-safe SQL generation)
- **Migrations:** [Goose](https://github.com/pressly/goose)

## Features

- **User Management:** Register and switch between different user profiles.
- **Feed Management:** Add, follow, and unfollow RSS feeds.
- **Automated Aggregation:** Periodically scrape followed feeds and store new posts.
- **Post Browsing:** View the latest posts from all the feeds you follow.

## Prerequisites

To run Gator, you will need the following installed on your system:

- **Go:** Version 1.23 or higher.
- **PostgreSQL:** A running instance of Postgres.
- **Goose:** (Optional) For running database migrations.

## Installation

You can install the Gator CLI tool using the `go install` command:

```bash
go install github.com/cybergrim/gator@latest
```

Ensure your `$GOPATH/bin` is in your system's `PATH` to run the `gator` command from anywhere.

## Setup and Configuration

Gator requires a configuration file named `.gatorconfig.json` in your home directory. This file stores your database connection string and the currently logged-in user.

### 1. Create the Config File

Create a file at `~/.gatorconfig.json` with the following structure:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

### 2. Database Migrations

This project uses [Goose](https://github.com/pressly/goose) for migrations. To set up your schema, navigate to the schema directory and run the following:

```bash
cd sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up
```

## Usage

Once installed and configured, you can start using Gator. Here are some of the essential commands:

### User Commands

- **Register a user:**
  ```bash
  gator register <username>
  ```
- **Login as a user:**
  ```bash
  gator login <username>
  ```
- **List all users:**
  ```bash
  gator users
  ```

### Feed Commands

- **Add a new feed:**
  ```bash
  gator addfeed <name> <url>
  ```
- **Follow an existing feed:**
  ```bash
  gator follow <url>
  ```
- **Unfollow a feed:**
  ```bash
  gator unfollow <url>
  ```
- **List all feeds:**
  ```bash
  gator feeds
  ```

### Aggregation and Browsing

- **Start the aggregator:**
  ```bash
  gator agg <duration>
  ```
  *Example: `gator agg 1m` (collects feeds every 1 minute).*

- **Browse posts:**
  ```bash
  gator browse <limit>
  ```
  *Example: `gator browse 5` (shows the 5 most recent posts).*

### Maintenance

- **Reset the database:**
  ```bash
  gator reset
  ```
  *Warning: This will delete all users, feeds, and posts from the database.*
