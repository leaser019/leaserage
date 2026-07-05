# AGENTS.md

## Principle

Act like a careful production engineer. Understand the goal before editing. Prefer the smallest safe change. Follow existing project patterns. Do not refactor unrelated code.

## Task routing

### Tiny

Typo, text, CSS, one-line obvious fix.

- Use simple read/search.
- Edit directly.
- Run the cheapest relevant check.

### Small

One component, one endpoint, one clear bug.

- Inspect nearby files first.
- Use Serena only when symbol lookup or references matter.
- Run targeted lint/test/typecheck if available.

### Medium

Behavior change or multiple related files.

- Do a short plan.
- Use CodeGraph once for impact/blast-radius if the module is unfamiliar.
- Use Serena for symbols/references/precise edits.
- Add or update tests when behavior changes.

### Large

Auth, authorization, payment, database, deployment, CI/CD, production bug, large refactor, or unfamiliar architecture.

- Plan first.
- Use CodeGraph for architecture and impact.
- Use Serena for implementation.
- Run strong verification before finishing.

## Tool guidance

Use tools intentionally.

- CodeGraph: broad repo understanding, entry points, dependencies, call graph, blast radius.
- Serena: semantic code work, symbols, references, file outline, diagnostics, symbol-level edits.
- Built-in read/search: small files, configs, docs, scripts, one-off checks.
- Workflow skills: unclear requirements, complex debugging, architecture decisions, large refactors, or TDD-heavy work.

Do not use heavy tools for tiny tasks.

## Repo orientation

Before broad exploration, check these if they exist:

- `docs/agent/repo-map.md`
- `docs/agent/architecture.md`
- `docs/agent/commands.md`
- `docs/agent/testing.md`

Use them for orientation only. Verify against source before editing.

## Workflow

For medium/large tasks:

1. Understand:
   - identify entry point
   - identify affected files
   - identify main symbols/functions/classes
   - identify relevant tests/checks
   - identify main risk

2. Edit:
   - make the smallest safe change
   - prefer symbol-level edits
   - re-read changed sections
   - run verification
   - review diff

Stop exploring once the safe implementation path is clear.

## Safety

Never commit or print secrets.

Do not modify auth, authorization, payment, deployment, database migrations, production infrastructure, validation, logging, tests, or security checks unless explicitly required.

Do not add dependencies unless clearly necessary.

Do not remove safeguards just to make code pass.

## Verification

Use existing project commands from package scripts, Makefile, README, or CI config.

Prefer cheapest reliable check first:

1. syntax/type check for tiny changes
2. targeted lint/test for small changes
3. related package test/build for medium changes
4. full lint, typecheck, test, build for large/risky changes

If a check cannot be run, explain why.

## Terminal output

For large output:

- focus on the first real error
- keep only relevant stack traces
- do not paste huge logs
- rerun targeted checks before full checks

## Closeout

When done, respond with:

1. Summary
2. Files changed
3. Verification
4. Risks or notes
