# Airbnb/Booking.com Clone — Microservices Backend

A production-oriented, microservices-based backend system inspired by Airbnb. Built with a polyglot architecture using **Go** and **Node.js (TypeScript)**, backed by **MySQL**, **Redis**, and **BullMQ**, designed around real-world engineering concerns like idempotency, distributed locking, RBAC, async messaging, and clean layered architecture.

---

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [Services](#services)
  - [Authentication Service (Go)](#authentication-service-go)
  - [Hotel Service (Node.js / TypeScript)](#hotel-service-nodejs--typescript)
  - [Booking Service (Node.js / TypeScript)](#booking-service-nodejs--typescript)
  - [Notification Service (Node.js / TypeScript)](#notification-service-nodejs--typescript)
  - [Review Service (Go)](#review-service-go)
- [Key Engineering Implementations](#key-engineering-implementations)
- [Infrastructure & DevOps](#infrastructure--devops)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)

---

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                        Client / API Gateway                      │
└───────────┬─────────────┬──────────────┬───────────┬────────────┘
            │             │              │           │
            ▼             ▼              ▼           ▼
  ┌──────────────┐ ┌──────────────┐ ┌──────────┐ ┌──────────────┐
  │   Auth       │ │   Hotel      │ │ Booking  │ │   Review     │
  │   Service    │ │   Service    │ │ Service  │ │   Service    │
  │   (Go)       │ │   (TS)       │ │ (TS)     │ │   (Go)       │
  └──────┬───────┘ └──────┬───────┘ └────┬─────┘ └──────┬───────┘
         │                │              │               │
         ▼                ▼              ▼               │
      MySQL            MySQL          MySQL +            │
    (Auth DB)        (Hotel DB)    Redis (Redlock)       ▼
                                        │             MySQL
                                        │           (Review DB)
                                        ▼
                              ┌──────────────────┐
                              │  Notification    │
                              │  Service (TS)    │
                              │  BullMQ + Redis  │
                              └──────────────────┘
```

Each service is fully **independent** — its own database, its own Docker container, its own configuration. Services communicate via HTTP, and async events are handled through Redis-backed job queues (BullMQ).

---

## Services

### Authentication Service (Go)

**Port:** configurable (default `:8080`) | **DB:** MySQL (`authenticationservicedb`)

The backbone of the platform. Handles user identity, JWT-based authentication, and a full **Role-Based Access Control (RBAC)** system.

#### Highlights

**RBAC System (Roles, Permissions, User-Roles)**

A fully normalised three-table RBAC design:
- `permissions` — atomic capabilities (`user_read`, `role_create`, `permission_manage`, etc.)
- `roles` — named collections of permissions (`admin`, `manager`, `user`)
- `permission_role` — many-to-many join with cascade deletes
- `user_roles` — assigns roles to users; supports soft-delete

Seed data ships with all 13 permissions pre-mapped across 3 roles, making the system immediately usable.

**JWT Authentication**

Tokens are signed with HS256 using a configurable secret. Every sensitive route is gated by the `JWTAuthenticate` middleware, which validates the `Authorization: Bearer <token>` header and rejects malformed or expired tokens with actionable error messages.

**Password Security (PBKDF2)**

Passwords are hashed using `PBKDF2-SHA256` with 210,000 iterations and a 16-byte random salt — matching Django's security defaults. The format is `pbkdf2_sha256$iterations$salt$key`, making it portable and auditable.

**Reverse Proxy Utility**

The service includes a generic `ReverseProxyToService` utility for forwarding requests to downstream services — a lightweight API Gateway capability built in.

**Rate Limiting**

All routes are protected by a per-IP rate limiter (`go-chi/httprate`, 5 req/min) at the router level, guarding against brute-force attacks.

**Structured Logging with Zerolog**

Per-request loggers enriched with the `correlation-id` are injected into context via middleware, enabling full request tracing across log files. Log rotation is handled by Lumberjack.

**Generic Middleware Validators**

A compile-time–generic validator (`Validate[T any]()` and `ValidateParams[T any]()`) decodes, unmarshals, and validates request bodies/params against any DTO struct using `go-playground/validator` tags — eliminating boilerplate in every handler.

---

### Hotel Service (Node.js / TypeScript)

**Port:** configurable | **DB:** MySQL (`hotelservicedb`) | **ORM:** Sequelize

Manages the hotel catalogue: CRUD for hotels and room types, with soft-delete patterns throughout.

#### Highlights

**Generic Base Repository (OOP Pattern)**

`BaseRepository<T extends Model>` provides `findById`, `findAll`, `create`, `update`, and `softDelete` using Sequelize generics. Concrete repositories (`HotelRepository`, `RoomTypeRepository`) extend it and override only what they need — keeping data access DRY and type-safe.

**Soft Delete via Paranoid Mode**

Hotels and room types use `deleted_at` timestamps instead of hard deletes, preserving historical integrity. Sequelize's `paranoid: true` option is leveraged on the `RoomType` model; hotels implement a manual soft-delete pattern (`hotel.deletedAt = new Date()`).

**Sequelize Migrations (TypeScript)**

Database schema is managed through TypeScript-based Sequelize CLI migrations — covering table creation, column additions, and index definitions across 5 migration files.

**Correlation ID via AsyncLocalStorage**

The Node.js `AsyncLocalStorage` API propagates the `X-Correlation-Id` header through the entire async call stack without passing it explicitly — enabling correlation-aware logging even in deeply nested async functions.

**Layered Architecture**

Clean separation: `validators → controllers → services → repositories → models`. Zod schemas validate request bodies/params before they reach the controller layer.

---

### Booking Service (Node.js / TypeScript)

**Port:** configurable | **DB:** MySQL (`bookingservicedb`) | **ORM:** Prisma | **Cache:** Redis

The most architecturally complex service. Handles booking creation and confirmation with strong consistency guarantees.

#### Highlights

**Idempotency Key Pattern**

Every booking creation returns a UUID `idempotencyKey`. To confirm a booking, the client must present this key via `POST /booking/confirm/:idempotencyKey`. This two-phase flow (create → confirm) prevents double-bookings and makes the confirmation endpoint **safely retryable**.

The `IdempotencyKey` table stores a `finalized` boolean. Once set to `true`, any re-submission of the same key is rejected with a `400 Bad Request`, making the endpoint idempotent by design.

**Distributed Locking with Redlock**

Booking creation acquires a distributed Redis lock on `booking:{hotelId}` with a 60-second TTL using the **Redlock algorithm** via the `redlock` library. This prevents race conditions when concurrent requests attempt to book the same hotel simultaneously — a critical correctness guarantee in a distributed system.

**Pessimistic Locking with MySQL `FOR UPDATE`**

Booking confirmation executes inside a Prisma transaction. The idempotency key row is fetched using a raw `SELECT ... FOR UPDATE` query, acquiring a **pessimistic row-level lock** in MySQL. This ensures that even if two confirmation requests arrive simultaneously, only one can mutate the row — the other waits or fails cleanly.

**Prisma with MariaDB Adapter**

Uses `@prisma/adapter-mariadb` for a connection-pool-aware MySQL driver, configured with `connectionLimit: 5`. The Prisma schema is cleanly modelled with a one-to-one `Booking ↔ IdempotencyKey` relationship.

**Booking Status State Machine**

Bookings move through a well-defined status enum: `PENDING → CONFIRMED / CANCELLED / IN_PROGRESS / COMPLETED`, enforced at the database level.

**Iterative Schema Evolution**

11 Prisma migration files document the full schema evolution from initial design through multiple revisions — demonstrating a disciplined, production-style database change management process.

---

### Notification Service (Node.js / TypeScript)

**Port:** configurable | **Queue:** BullMQ + Redis | **Email:** Nodemailer (Gmail SMTP)

A dedicated async notification microservice. Decoupled from the rest of the system via a Redis-backed job queue.

#### Highlights

**BullMQ Job Queue Architecture**

Email sends are never performed inline. Producers add jobs to a named BullMQ queue (`queue-mailer`); a Worker independently processes them. This pattern means:
- Booking/Auth services are never blocked by email delivery
- Failed email jobs are retried automatically by BullMQ
- The queue can be monitored, paused, and inspected independently

**Handlebars Email Templates**

Email bodies are rendered from `.hbs` Handlebars templates loaded from disk at runtime. The `renderMailerTemplates(templateId, data)` utility compiles templates with dynamic data injection. A production-quality `welcome.hbs` template (responsive HTML email, gradient header, Airbnb-branded) ships with the service.

**Worker Lifecycle Management**

`setupEmailWorker()` initialises the BullMQ Worker on server start and registers `completed`/`failed` event listeners for structured logging — giving full observability over async email processing.

---

### Review Service (Go)

**Port:** configurable | **DB:** MySQL (`reviewservicedb`)

A lean Go service for hotel reviews. Currently scaffolded with clean architecture ready for full implementation.

#### Highlights

**Clean Go Architecture**

Follows the same layered pattern as the Auth Service: `router → controller → service → repository`. All layers are interface-driven, making them trivially mockable for unit testing.

**Database Migration with Goose**

Schema is managed by `goose` SQL migrations. The review table includes proper indexing on `booking_id`, `user_id`, and `hotel_id` with a `CHECK` constraint enforcing `rating BETWEEN 1 AND 5`.

---

## Key Engineering Implementations

| Concern | Implementation |
|---|---|
| **Authentication** | JWT (HS256), configurable secret, 30-min expiry |
| **Password Hashing** | PBKDF2-SHA256, 210k iterations, random salt |
| **RBAC** | 3-table normalised design (roles, permissions, user_roles) |
| **Idempotency** | UUID idempotency keys, `finalized` flag, two-phase booking |
| **Distributed Locking** | Redlock algorithm over Redis for booking concurrency |
| **Pessimistic DB Locking** | `SELECT ... FOR UPDATE` inside Prisma transactions |
| **Async Messaging** | BullMQ job queues (Redis-backed) for email notifications |
| **Email Templating** | Handlebars templates with dynamic data injection |
| **Rate Limiting** | Per-IP rate limiting at router middleware level |
| **Soft Deletes** | `deleted_at` pattern across all services |
| **Correlation IDs** | Request tracing via `X-Correlation-Id` header (context / AsyncLocalStorage) |
| **Structured Logging** | Zerolog (Go) + Winston (Node.js) with daily log rotation |
| **Validation** | Zod schemas (TypeScript), go-playground/validator (Go), generic middleware wrappers |
| **Generic Repository** | Type-safe `BaseRepository<T>` in Sequelize using TypeScript generics |
| **Reverse Proxy** | Built-in `httputil.ReverseProxy` utility for service forwarding |
| **Schema Migrations** | Goose (Go services), Prisma Migrate + Sequelize CLI (Node services) |

---

## Infrastructure & DevOps

**Per-Service Docker Compose**

Each service ships with its own `docker-compose.yml`, giving developers the ability to spin up only the services they need. Each compose file provisions its own isolated MySQL container on a dedicated port:

| Service | MySQL Host Port |
|---|---|
| Authentication Service | `3308` |
| Booking Service | `3307` |
| Hotel Service | `3306` |
| Review Service | `3309` |

**Root-Level Compose**

A root `docker-compose.yml` + `init.sql` bootstraps a shared MySQL instance for `bookingservicedb` and `hotelservicedb` in one command — ideal for integrated local development.

**Environment-Based Configuration**

Every service exposes a `.env.example` file documenting all required environment variables. Configuration is loaded at startup via `godotenv` (Go) or `dotenv` (Node.js), with sensible defaults for local development.

---

## Tech Stack

| Layer | Technology |
|---|---|
| **Languages** | Go 1.26, TypeScript 5.9 / Node.js |
| **Frameworks** | go-chi (Go), Express 5 (Node.js) |
| **Databases** | MySQL 8.0 |
| **ORMs / Query Builders** | Prisma 7 (Booking), Sequelize 6 (Hotel), `database/sql` (Auth, Review) |
| **Cache / Queue** | Redis 7, BullMQ 5, Redlock |
| **Auth** | JWT (golang-jwt), PBKDF2-SHA256 (crypto/pbkdf2) |
| **Validation** | Zod 4 (TypeScript), go-playground/validator v10 |
| **Logging** | Zerolog + Lumberjack (Go), Winston + DailyRotateFile (Node.js) |
| **Email** | Nodemailer (Gmail SMTP) |
| **Templating** | Handlebars |
| **Migrations** | Goose (Go), Prisma Migrate, Sequelize CLI |
| **Containerisation** | Docker, Docker Compose |

---

## Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.22+
- Node.js 20+ & pnpm
- `goose` CLI (`go install github.com/pressly/goose/v3/cmd/goose@latest`)

### Running a Service

**Authentication Service**
```bash
cd AuthenticationService
cp .env.example .env        # fill in DB credentials
docker-compose up -d        # start MySQL on :3308
make migrate-up url="user:user123@tcp(localhost:3308)/authenticationservicedb"
make run
```

**Booking Service**
```bash
cd BookingService
cp .env.example .env        # fill in DB + Redis config
docker-compose up -d        # start MySQL on :3307 + Redis on :6379
pnpm install
pnpm dev
```

**Hotel Service**
```bash
cd HotelService
cp .env.example .env
docker-compose up -d
pnpm install
pnpm migrate                # runs Sequelize migrations
pnpm dev
```

**Notification Service**
```bash
cd NotificationService
cp .env.example .env        # fill in Redis + SMTP credentials
docker-compose up -d        # starts Redis on :6380
pnpm install
pnpm dev
```

**Review Service**
```bash
cd ReviewService
docker-compose up -d        # starts MySQL on :3309
make migrate-up url="user:user123@tcp(localhost:3309)/reviewservicedb"
make run
```

### Key API Endpoints

| Service | Method | Endpoint | Description |
|---|---|---|---|
| Auth | `POST` | `/v1/user/register` | Register a new user |
| Auth | `POST` | `/v1/user/login` | Login, receive JWT |
| Auth | `GET` | `/v1/user/{id}` | Get user by ID (JWT required) |
| Auth | `POST` | `/v1/roles/` | Create a role |
| Auth | `GET` | `/v1/roles/` | List all roles (filter by `?name=`) |
| Hotel | `POST` | `/api/v1/hotel/` | Create a hotel |
| Hotel | `GET` | `/api/v1/hotel/` | List all hotels |
| Hotel | `GET` | `/api/v1/hotel/:id` | Get hotel by ID |
| Booking | `POST` | `/api/v1/booking/` | Create a booking (returns idempotency key) |
| Booking | `POST` | `/api/v1/booking/confirm/:idempotencyKey` | Confirm a booking |
| Review | `GET` | `/api/v1/reviews/` | Create review (stub) |
