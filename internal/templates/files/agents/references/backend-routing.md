# Backend Routing

Use the smallest file set that proves the flow.

## Common Paths

- HTTP/API: routes -> middleware -> request/validator -> controller/handler -> service/use case -> resource/serializer -> tests.
- Domain/service logic: public method -> collaborators -> data invariants -> unit or focused integration tests.
- Database/data access: model/entity -> query/repository -> migration/schema -> factory/fixtures -> tests.
- Jobs/queues/commands: dispatcher/command -> job handler -> retries/timeouts/idempotency -> worker config -> logs/tests.
- Events/webhooks/integrations: signature/auth -> payload parser -> mapping -> retry/error path -> fixtures or contract tests.
- Config/env/deployment: config loader -> defaults/templates -> docs/install scripts -> dry-run or local smoke check.
- Observability: log/metric/tracing call site -> sensitive-data risk -> sampling/noise risk -> verification query or local output.

## Backend Safety Checks

- Check existing data assumptions before changing casts, enums, constraints, or defaults.
- Treat migrations, backfills, destructive data operations, auth, authorization, payment, and production config as large tasks.
- Prefer dry-run and temp-home/test-database checks before touching real user or production state.
- Keep compatibility unless the user explicitly approves a breaking API or data contract change.
