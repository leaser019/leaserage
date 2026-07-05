# Frontend Routing

Use the smallest file set that proves the user-facing flow.

## Common Paths

- Page/screen: route -> layout/shell -> page/container -> data loader/query -> view components -> tests or screenshot.
- Component behavior: public props -> state/events -> child components -> accessibility states -> unit or interaction tests.
- Form flow: schema/defaults -> inputs -> validation -> submit action/mutation -> loading/error/success states -> tests.
- Client data: query key/cache -> fetcher/client -> transform/selector -> invalidation/refetch -> tests.
- Navigation/routing: route params -> guards/middleware -> links/actions -> redirects/not-found states -> smoke check.
- Styling/layout: design tokens/theme -> component classes/styles -> responsive constraints -> visual check across target breakpoints.
- Assets/media: import/declaration -> optimized render path -> fallback/error state -> build or browser check.
- i18n/content: message key -> default text -> interpolation/plurals -> locale fallback -> snapshot or focused UI check.

## Frontend Safety Checks

- Preserve existing layout scaffolding unless the user explicitly asks to redesign it.
- Check loading, empty, error, disabled, and success states when behavior changes.
- Keep text inside responsive containers; verify long labels and mobile widths when touching layout.
- Prefer existing design-system components, icons, spacing, and state patterns.
- Do not add new dependencies, global styles, or state libraries unless clearly necessary.
- For visual changes, use the cheapest reliable verification: targeted unit/interaction test, typecheck/build, browser smoke check, or screenshot.
