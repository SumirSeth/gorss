# Goal:
A production-ready, extensible RSS Aggregator API written in Go. It should:

- Accept feed URLs

- Periodically fetch + parse them (RSS/Atom)

- Cache items in memory or DB

- Expose a REST API to serve data to frontend


# Features:

| Stage            | Covered? | Notes                       |
| ---------------- | -------- | --------------------------- |
| Feed collection  | ✅        | API: `/feeds` POST |
| Feed fetching    | ✅        | Go scheduler polls feeds    |
| Parsing XML      | ✅        | `gofeed` handles RSS/Atom   |
| Deduplication    | ✅        | Timestamp or GUID filtering |
| Storage          | ✅        | In-memory or SQLite         |
| API exposure     | ✅        | REST API serves frontend    |
| Frontend reading | ❌        | Use exposed API to build UI |


> [!NOTE]
> This is a work in progress. It is not yet ready for production use.