# GoCoop Backend

GoCoop is a platform designed to list and categorize cooperatives around the globe, making them more accessible to users looking to learn more about cooperatives in general.

This repository contains the **backend API** for GoCoop, built using **Go** with `net/http` for handling HTTP requests, and **PostgreSQL** as the database.

## 🚀 Project Overview

The backend is responsible for:
- Handling API requests for cooperatives and categories (at the moment).
- Providing data to the frontend via a RESTful API.
- Interacting with a PostgreSQL database.

## 🛠 Tech Stack

- **Language:** Go (`net/http`)
- **Database:** PostgreSQL
- **ORM/Query Builder:** `pgx`
- **Environment Management:** `.env`

## 📡 API Endpoints

| Method | Endpoint | Description | Query Params
|--------|---------|-------------|--------------|
| GET | `/coops` | Search for cooperatives | `search` (optional), `category` (optional)
| GET | `/coops/:slug` | Get details of a cooperative | N/A
| GET | `/categories` | Get all coops categories | N/A

## 📬 Contact

For questions or suggestions, open an issue or reach out!

---
**GoCoop** - Empowering cooperatives through technology. 🌲

