# AGENTS.md

## Principle

Act like a careful production engineer. Understand the goal before editing, prefer the smallest safe change, follow existing project patterns, and do not refactor unrelated code.

## Framework Model

`AGENTS.md` is the shared router and safety contract. Skills define the repeatable procedure for one work mode.

Rules decide when. Skills decide how.

Use one active primary workflow skill at a time. A skill may hand off to exactly one next primary skill, pause for approval, or stop when the task is complete.

## Workflow Routing

Use this routing for non-trivial work. Trivial one-step requests may skip it when that is clearly lower friction and still safe.

1. Broken behavior, failing command, regression, or incident -> `debug`
2. Explicit review request or risk assessment of existing changes -> `review`
3. Ambiguous goal, open design question, or multiple reasonable approaches -> `brainstorming`
4. Clear work that needs sequencing, dependency ordering, or rollback thinking -> `plan`
5. Clear approved code/docs/config change -> `implement`
6. Validation-only request or confidence check -> `test`

Escalate when the facts change: ambiguity goes back to `brainstorming`, growing scope goes to `plan`, broken validation goes to `debug`, and repository-local instructions always win.

## Task Size Routing

Route by risk and affected surface, not by framework.

### Tiny

Docs, comments, formatting, typo, copy, obvious one-line constant, or a local style fix that cannot change runtime behavior.

- Use simple read/search.
- Edit directly.
- Run the cheapest relevant check, or explain why readback is enough.

### Small

One file or one narrow runtime path: one endpoint, handler, service method, repository query, validation rule, serializer, job, command, or clear bug with a known failing boundary.

- Inspect the entry point and nearest caller/callee before editing.
- Reproduce the bug or identify the exact failing check before changing code.
- Use Serena only when symbol lookup, references, or precise edits matter.
- Run the targeted test, lint, typecheck, or command that covers the touched path.

### Medium

Behavior change across related files or layers: API contract, request/response shape, domain rule, database read/write behavior, queue/job behavior, cache behavior, third-party integration, or multiple related endpoints.

- Do a short plan.
- Identify entry point, affected files, data flow, existing tests, and rollback risk.
- Use CodeGraph once for impact when the module is unfamiliar or crosses layers.
- Use Serena for symbols, references, and precise edits.
- Add or update focused tests when behavior changes.

### Large

High-risk or broad work: auth, authorization, permissions, payment/billing, migrations/backfills, destructive data operations, security, production incident, deployment, CI/CD, observability/logging changes, concurrency/idempotency, large refactor, or unfamiliar architecture.

- Plan first.
- Use CodeGraph for architecture and blast radius.
- Prefer dry-run, backup, rollback, and staged verification paths.
- Ask approval before destructive, irreversible, or unusually broad changes.
- Run strong verification before finishing.

## Backend Routing

For backend/API/database/job work, read `references/backend-routing.md` when the local path is not obvious.

## Frontend Routing

For frontend/UI/state/client-routing work, read `references/frontend-routing.md` when the local path or verification path is not obvious.

## Bugfix Workflow

Diagnose before patching.

1. Reproduce or capture the smallest failing case.
2. State expected behavior, actual behavior, and the failing boundary.
3. Trace only the relevant path first.
4. Form one or two concrete hypotheses and test them with the cheapest check.
5. Fix the smallest proven cause.
6. Add or update a regression test when the behavior is important or likely to recur.
7. Re-run the original reproduction and targeted regression check.
8. Explain the root cause in the final response.

## Tool Guidance

- CodeGraph: broad repo understanding, entry points, dependencies, call graph, impact, and unfamiliar architecture.
- Serena: semantic code work, symbols, references, file outline, diagnostics, and symbol-level edits.
- Built-in read/search: small files, configs, docs, scripts, and one-off checks.
- Workflow skills: unclear requirements, complex debugging, architecture decisions, large refactors, or TDD-heavy work.

Do not use heavy tools for tiny tasks.

## Command Output Hooks

RTK, when installed, is a command-output optimization hook/proxy, not an MCP server. If a transparent RTK hook is active, run commands normally. If only prompt-level RTK guidance is available and a command may produce noisy output, prefer `rtk <command>` when `rtk` is installed. Use `RTK_DISABLED=1` when raw output is required.

## Evidence And Validation

- Do not claim success without fresh evidence.
- Run the narrowest high-signal check first, then broaden when risk requires it.
- If validation is skipped, state exactly why.
- Distinguish observed evidence from assumptions.
- Distinguish environment limitations from product failures.

## Safety

Never commit or print secrets.

Do not modify auth, authorization, payment, deployment, database migrations, production infrastructure, validation, logging, tests, or security checks unless explicitly required.

Do not add dependencies unless clearly necessary. Do not remove safeguards just to make code pass.

## Closeout

When done, respond with:

1. Summary
2. Files changed
3. Verification
4. Risks or notes

For non-trivial work, also state whether the work is complete, complete with remaining risk, or paused.
