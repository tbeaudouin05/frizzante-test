

---

## Project Structure

This is a Frizzante application with a Svelte 5 front-end (SSR/CSR) and a Go back-end.

```
hello-frizzante/
├── app/                              # Svelte 5 front-end (SSR bundle)
│   ├── app.server.ts|.svelte         # SSR entry
│   ├── app.client.ts|.svelte         # CSR entry
│   ├── exports.server.ts             # Exports server-side Svelte views to Frizzante
│   ├── exports.client.ts             # Code-split client imports for hydration
│   ├── index.html                    # Dev server HTML (Vite)
│   ├── lib/
│   │   ├── components/               # UI components
│   │   │   ├── ui/                   # shadcn/ui for Svelte components
│   │   │   │   ├── button/
│   │   │   │   ├── input/
│   │   │   │   ├── card/
│   │   │   │   ├── checkbox/
│   │   │   │   ├── calendar/         # Installed via shadcn CLI
│   │   │   │   └── select/           # Installed via shadcn CLI
│   │   │   └── icons/ …
│   │   ├── scripts/core/             # Frizzante helpers (e.g., action, href)
│   │   └── views/                    # Svelte pages rendered by Go handlers
│   │       ├── Welcome.svelte
│   │       ├── Todos.svelte
│   │       └── Lessons.svelte        # Calendar + shadcn Select time picker
│   ├── package.json / pnpm-lock.yaml # Front-end deps (pnpm)
│   └── vite.config.ts / svelte.config.js / tailwind.config.ts
│
├── lib/                              # Go server-side (Frizzante)
│   ├── core/                         # Frizzante core (send/receive/view, server, etc.)
│   ├── routes/handlers/              # HTTP handlers
│   │   ├── welcome/
│   │   ├── todos/
│   │   └── lessons/                  # Lessons view + book/cancel handlers
│   └── session/memory/               # In-memory session store and types
│       ├── new.go                    # Session initializer (pre-populates Todos, sets Lessons = [])
│       └── types.go                  # Session, Todo, Lesson types
│
├── main.go                           # Wires routes to handlers and SSR renderer
└── README.md
```

## Common Frizzante Commands

Use these frequently while working on the project:

- `frizzante --configure`
  - Initializes/updates Frizzante configuration for your environment.

- `frizzante --dev`
  - Runs the Go server with live-reload and starts Vite for the Svelte app (SSR + HMR).

- Front-end build (from `app/`):
  - `pnpm vite build`
  - Produces `app/dist` assets for SSR.

- Package built assets into the Go app (from project root):
  - `frizzante --package`
  - Copies `app/dist` into Go embeds at `lib/core/view/ssr` so the server serves the latest bundle.

- Production build:
  - `frizzante --build`
  - Creates a standalone binary at `.gen/bin/app`.

## Tech Stack

- **Frizzante (Go)**: server, routing, SSR rendering, session storage.
- **Svelte 5**: views with `$state`, `$derived` and Svelte transitions.
- **Vite**: dev server and build pipeline for the app bundle.
- **Tailwind CSS + shadcn/ui (Svelte)**: headless UI primitives + styled components.

## Views and Routing

- Views are Svelte files under `app/lib/views/` and exported in:
  - `app/exports.server.ts` (server-side components)
  - `app/exports.client.ts` (client-side lazy imports for hydration)

- Routes are registered in `main.go` as `route.Route{ Pattern, Handler }`.
  - Examples:
    - `GET /` and `GET /welcome` → `welcome.View`
    - `GET /todos` plus actions `/check|/uncheck|/add|/remove` → Todos handlers
    - `GET /lessons` → `lessons.View`
    - `GET /lessons/book` → `lessons.Book` (adds a booking to session)
    - `GET /lessons/cancel` → `lessons.Cancel` (removes by index)

## Lessons Page (example)

- UI: `app/lib/views/Lessons.svelte`
  - Date via `Calendar` (shadcn calendar built on `bits-ui`).
  - Time via shadcn `Select` dropdown.
  - Student name via shadcn `Input`.
  - Form uses Frizzante `action("/lessons/book")` to submit `student`, `date` (ISO), and `time`.
  - Upcoming lessons list with Cancel action.

- Server: `lib/routes/handlers/lessons/`
  - `view.go` reads session and returns `Lessons` to the view.
  - `book.go` appends a new `Lesson{ Student, Date, Time }` to session.
  - `cancel.go` removes a lesson by index.
  - Session schema in `lib/session/memory/types.go`.

## Using shadcn/ui (Svelte)

The project ships with shadcn components installed under `app/lib/components/ui/`.
To add another component, use the shadcn CLI inside `app/` (pnpm required):

```sh
pnpm dlx shadcn-svelte@latest add <component>
# e.g. calendar, select, dialog, sheet, alert-dialog, etc.
```

The CLI will scaffold files like `app/lib/components/ui/<component>/` and update dependencies.

## Development Workflow

- **Dev (SSR + Vite)**
  - Back-end auto-reloads and Vite serves the front-end:
    ```sh
    frizzante --dev
    ```

- **When you change front-end code**
  - Vite hot-reloads in dev.
  - For a production-like run, build and package:
    ```sh
    # from app/
    pnpm vite build
    # from project root
    frizzante --package
    frizzante --dev
    ```

## Notes

- This codebase uses Svelte 5 `$state` and `$derived`, which require Svelte 5 configuration (already set up in `svelte.config.js`).
- Frizzante actions/helpers live under `app/lib/scripts/core/` (e.g., `action.ts`, `href.ts`).
- UI is Tailwind-based; styles are configured in `tailwind.config.ts` and imported via Vite.
