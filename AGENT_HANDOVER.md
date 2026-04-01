# Project Status & Handover

## Project Overview
**Name:** Dream Journal
**Architecture:** Local Client-Server using **Wails** (Go backend, Svelte + TailwindCSS v3 frontend).
**Goal:** A Linux desktop application for writing and managing dreams, which also exposes a local REST API (port 8080) for AI/RAG ingestion.
**Database:** Local SQLite (`~/.dream-journal/data.db`).

## Current State (Phase 2: Data Persistence & Logic)
- **Frontend:** Svelte with Tailwind CSS v3 is initialized and running. Basic UI skeleton exists (`App.svelte`).
- **Backend (`app.go`):** SQLite connection is established. Base schema (`dreams` table) is set up with fields for `id`, `title`, `content`, `date`, `moonphase`, and `created_at`.
- **Wails Bindings:** The Wails bridge is being wired up to allow the Svelte frontend to call Go methods (`SaveDream`, `GetDreamsDesc`).

## Active Technical Direction: Astronomical Calculations (Moon Phase)
The user is manually implementing highly accurate mathematical calculations in Go to determine the moon phase natively, avoiding external APIs.

**Current Focus (Go Backend):**
1.  **Time Conversion:** Implementing conversion from UTC (Coordinated Universal Time) to TT (Terrestrial Time) for precise astronomical formulas.
2.  **Leap Seconds Logic:** Writing a `leapseconds` function to support the UTC to TT conversion.
3.  **Dynamic Leap Seconds Table:** The user plans to store the leap seconds data in the SQLite database so it can be dynamically selected and periodically updated/checked.

## Next Steps for the Next Agent
1.  **Assist with UTC to TT Conversion:** Help the user finalize the Go logic for calculating TT, specifically handling the leap seconds offset (TAI - UTC) + 32.184s.
2.  **Database Schema Update:** Assist in modifying the `app.initDB()` function in `app.go` to create a new table for caching leap seconds (e.g., `leap_seconds`).
3.  **Periodic Updates:** Help implement a Go routine or startup check to fetch and update the leap seconds table if it's outdated.
4.  **Frontend Wiring:** Once the Go backend is stable and compiled, guide the user to run `wails generate module` and finalize the Javascript imports in `App.svelte` to display the dreams and their highly accurate moon phases.
5.  **Phase 4 (API):** Eventually, stand up the local HTTP server in Go to expose the `/api/dreams` endpoints.