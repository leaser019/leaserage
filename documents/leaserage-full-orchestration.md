# Leaserage Full Orchestration

This document captures the full Leaserage workflow state machine implied by the installed `AGENTS.md` and workflow skills. Use this when you need the broader routing model, approval gates, validation handoffs, and stop states. For day-to-day usage, the shorter README diagram is enough.

```mermaid
flowchart TD
  A([Start task]) --> SIZE{Tiny, clear, low risk?}
  SIZE -- Yes --> DIRECT([Handle directly])
  SIZE -- No --> ROUTE{Route by primary need}

  ROUTE -->|Broken / failing / regressed / incident| DEBUG[systematic-debugging]
  ROUTE -->|Incoming review feedback| REVIEW_IN[receiving-code-review]
  ROUTE -->|Completed work needs review| REVIEW_OUT[requesting-code-review]
  ROUTE -->|Ambiguous / multiple approaches| BRAIN[brainstorming]
  ROUTE -->|Clear but sequenced / risky / dependency-heavy| PLAN[writing-plans]
  ROUTE -->|Approved plan execution| IMPL[executing-plans / subagent-driven-development]
  ROUTE -->|Behavior change or bugfix| TDD[test-driven-development]
  ROUTE -->|Validation is the main job| TEST[verification-before-completion]

  BRAIN --> BRAIN_GATE{Meaningful design or tradeoff decision?}
  BRAIN_GATE -- Yes --> APPROVAL1([Pause for approval])
  BRAIN_GATE -- No --> BRAIN_NEXT{Next need}
  APPROVAL1 --> BRAIN_NEXT
  BRAIN_NEXT -->|Execution structure needed| PLAN
  BRAIN_NEXT -->|Behavior work is now clear| TDD
  BRAIN_NEXT -->|Plan execution is now clear| IMPL
  BRAIN_NEXT -->|Validation-only follow-up| TEST
  BRAIN_NEXT -->|Direction-only stop| DONE([Finish])

  PLAN --> PLAN_GATE{Meaningful scope, sequence, rollback, or risk decision?}
  PLAN_GATE -- Yes --> APPROVAL2([Pause for approval])
  PLAN_GATE -- No --> PLAN_NEXT{Next need}
  APPROVAL2 --> PLAN_NEXT
  PLAN_NEXT -->|Start plan execution| IMPL
  PLAN_NEXT -->|Start behavior implementation| TDD
  PLAN_NEXT -->|Approach still unresolved| BRAIN
  PLAN_NEXT -->|Plan-only stop| DONE

  IMPL --> IMPL_CHECK{What happened during execution?}
  IMPL_CHECK -->|Small change plus targeted check is enough| DONE
  IMPL_CHECK -->|Needs broader regression confidence| TEST
  IMPL_CHECK -->|Ambiguity appeared| BRAIN
  IMPL_CHECK -->|Scope grew or dependencies changed| PLAN
  IMPL_CHECK -->|Behavior is broken| DEBUG

  TDD --> IMPL_CHECK

  DEBUG --> DEBUG_GATE{Root cause clear and fix materially changes code, config, or workflow?}
  DEBUG_GATE -- Yes --> APPROVAL3([Pause for approval])
  DEBUG_GATE -- No --> DEBUG_NEXT{Fix path}
  APPROVAL3 --> DEBUG_NEXT
  DEBUG_NEXT -->|Actual fix now| TDD
  DEBUG_NEXT -->|Substantial fix path| PLAN
  DEBUG_NEXT -->|Primarily validation gap| TEST
  DEBUG_NEXT -->|Diagnosis-only stop| DONE

  REVIEW_IN --> REVIEW_NEXT{Follow-up needed?}
  REVIEW_OUT --> REVIEW_NEXT
  REVIEW_NEXT -->|No material findings| DONE
  REVIEW_NEXT -->|Defect needs diagnosis| DEBUG
  REVIEW_NEXT -->|Direct fix work| IMPL
  REVIEW_NEXT -->|Validation gap| TEST
  REVIEW_NEXT -->|Broader restructure needed| PLAN

  TEST --> TEST_NEXT{Validation result}
  TEST_NEXT -->|Evidence sufficient| DONE
  TEST_NEXT -->|Validation failed or issue reproduced| DEBUG
  TEST_NEXT -->|Validation-only stop with limits noted| DONE

  DIRECT --> DONE

  classDef startNode fill:#f5f0ff,stroke:#8b6cf0,stroke-width:2px,color:#222;
  classDef decisionNode fill:#ede9fe,stroke:#8b6cf0,stroke-width:2px,color:#222;
  classDef skillNode fill:#f3f0ff,stroke:#8b6cf0,stroke-width:2px,color:#222;
  classDef pauseNode fill:#fff7ed,stroke:#f59e0b,stroke-width:2px,color:#222;
  classDef finishNode fill:#eefcf3,stroke:#22a06b,stroke-width:2px,color:#222;

  class A,DIRECT startNode;
  class SIZE,ROUTE,BRAIN_GATE,BRAIN_NEXT,PLAN_GATE,PLAN_NEXT,IMPL_CHECK,DEBUG_GATE,DEBUG_NEXT,REVIEW_NEXT,TEST_NEXT decisionNode;
  class DEBUG,REVIEW_IN,REVIEW_OUT,BRAIN,PLAN,IMPL,TDD,TEST skillNode;
  class APPROVAL1,APPROVAL2,APPROVAL3 pauseNode;
  class DONE finishNode;
```

## Routing Basis

- `systematic-debugging`: broken behavior, failing command, regression, unexpected output, or incident.
- `receiving-code-review`: respond to incoming review feedback.
- `requesting-code-review`: inspect completed work for bugs, regressions, missing validation, and risk.
- `brainstorming`: unclear goal, unresolved design, tradeoff, or multiple viable approaches.
- `writing-plans`: clear work that needs sequencing, dependency ordering, checkpoints, or rollback thinking.
- `executing-plans`: execute an approved written plan inline.
- `subagent-driven-development`: execute an approved written plan with task-focused subagents.
- `test-driven-development`: implement behavior changes or bugfixes test-first.
- `verification-before-completion`: validation-only request, confidence check, or regression proof.
- direct handling: tiny low-risk task where workflow routing would add noise.

## Approval Gates

Pause for approval before:

- meaningful design or tradeoff decisions after brainstorming
- substantial execution plans with risk, sequencing, or rollback choices
- destructive, irreversible, broad, or production-sensitive changes
- systematic-debugging fixes that materially change code, config, workflow, data, auth, payment, or deployment behavior

Approval gates are conditional. Do not pause for every small edit, but do not skip approval for irreversible or high-risk actions.

## Handoff Rules

- ambiguity during execution -> `brainstorming`
- scope growth or dependency ordering -> `writing-plans`
- failed validation or reproduced breakage -> `systematic-debugging`
- direct implementation that needs more proof -> `verification-before-completion`
- review finding with root cause unknown -> `systematic-debugging`
- review finding with obvious narrow fix -> `test-driven-development` or `executing-plans`

## Evidence Rules

- Do not claim success without fresh evidence.
- Run the narrowest high-signal check first.
- Broaden validation when risk, blast radius, or user-facing behavior requires it.
- If validation is skipped, state exactly why.
- Distinguish observed evidence from assumptions and environment limits.

## Stop States

Finish only when the current workflow has a clear stop state:

- direct task completed and read back or checked
- systematic-debugging root cause explained and fix verified, or diagnosis-only stop requested
- review findings reported with file/line evidence
- writing-plans delivered as a handoff artifact
- implementation verified with relevant checks
- verification-before-completion workflow reports evidence, failures, or explicit limits
