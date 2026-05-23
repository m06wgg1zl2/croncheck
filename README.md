# croncheck

A dead-simple cron expression validator and next-run visualizer with timezone-aware output for the terminal.

---

## Installation

```bash
go install github.com/yourname/croncheck@latest
```

Or build from source:

```bash
git clone https://github.com/yourname/croncheck.git && cd croncheck && go build -o croncheck .
```

---

## Usage

Validate a cron expression and preview the next scheduled runs:

```bash
croncheck "0 9 * * MON-FRI"
```

**Output:**
```
✔ Valid cron expression: "0 9 * * MON-FRI"

Next 5 runs (UTC):
  1. Mon, 02 Jun 2025 09:00:00 UTC
  2. Tue, 03 Jun 2025 09:00:00 UTC
  3. Wed, 04 Jun 2025 09:00:00 UTC
  4. Thu, 05 Jun 2025 09:00:00 UTC
  5. Fri, 06 Jun 2025 09:00:00 UTC
```

Specify a timezone with the `--tz` flag:

```bash
croncheck "30 18 * * *" --tz "America/New_York"
```

Show more upcoming runs with `--count`:

```bash
croncheck "*/15 * * * *" --count 10
```

---

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--tz` | `UTC` | Timezone for output (e.g. `Europe/London`) |
| `--count` | `5` | Number of upcoming runs to display |

---

## License

MIT © [yourname](https://github.com/yourname)