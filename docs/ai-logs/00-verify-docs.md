# Verify docs

**Date:** 2026-08-15
**Reviewer:** principal-swe (agent)
**Reviewed:** `README.md` (Detail requirements), `docs/00-architectural-design.md`, `docs/01-database-design.md`
**Method:** `/grill-me` — one question at a time, each with a recommended answer.

> Scope note: the architectural diagram (`docs/media/architectural-design.png`) could not be
> rendered in this session ("unsupported image"), so the architecture was reviewed from the doc
> text only. The diagram should still be manually checked against the text.

---

## Ground rule established up front

**The docs describe the *full production system*; the code is a POC subset.**
The HA topology (Redis, replication, K8s) is the north-star, not what the POC builds.
→ Follow-up action: add an explicit "POC boundary" marker to the docs so the full-system
design and the POC implementation don't silently diverge.

---

## Conversation & decisions

### Q1 — Architecture scope
- **Asked:** the doc specifies full HA (Redis, source/replica, K8s) for a POC with three modest requirements. Intent?
- **Decision:** "The designs are for the full system, the code will be a POC."
- **Outcome:** no change to the architecture; keep the full-system topology. Document the POC boundary.

### Q2 — "Real-time" vs 5-min cache TTL
- **Asked:** requirement says "real-time overview", but the cache serves reads with a 1–5 min TTL, so a manager can see stale stock.
- **Decision:** "For API get list / get list method, we don't cache. We cache when get one record, and invalidate cache on update."
- **Outcome:** list endpoints are **not** cached; single-record GETs are cache-aside and invalidated on write.
- **Doc fix:** the "Caching Strategy" section currently reads as if *all* reads are cache-aside.
  State explicitly: *list reads bypass cache; single-record reads are cached (TTL 1–5 min) and invalidated on any write.*

### Q3 — API response envelope vs REST semantics
- **Asked:** the `{error:int, message, data}` envelope (error=0 success) contradicts "HTTP status is the status".
- **Decision:** "The response I mentioned is the body of the response, we don't override the HTTP status code."
- **Outcome:** HTTP status codes are authoritative; the envelope is the body shape only. Keep it.
- **Non-blocking note:** an integer `error: 0` on success is redundant next to the status code, and the
  user's own REST rule #13 prefers a nested `{error:{code, message, details}}` with a *string* code.
  Two sources of truth can drift. Left as-is per user's call.

### Q4 — Aging computation column (`created_at` vs a stock-in date)
- **Asked:** ">90 days in inventory" has no arrival-date column; `created_at` is overloaded. Backfill/import would corrupt age.
- **Decision:** "Each vehicle is stocked then a record is created; removal = soft-delete. Use `created_at` for aging. The threshold must be configurable."
- **Outcome:** `created_at` is the stock-in date; aging = `created_at < now() - threshold`.
- **Doc fix (important):** `created_at` is described as "timestamp the record is created" — nothing says it
  is *the* date aging depends on. Write down the invariant: **record creation == vehicle stocked**, or a
  future engineer breaks the aging logic. Also: make the 90-day threshold a **config value**, not a SQL literal.

### Q5 — Vehicle status / "in inventory" semantics
- **Asked:** "sold" is a business event, "deleted" is administrative; does the full system need to distinguish them?
- **Decision:** "Soft-delete only (as-is)."
- **Outcome:** out-of-stock == soft-delete. No status column.
- **Non-blocking note:** sold vehicles are indistinguishable from deleted ones, so time-to-sell / sales
  counts won't be queryable later without a migration. Accepted for the POC.

### Q6 — `stocked_vehicles.name` denormalization
- **Asked:** `name` is described as "the name of the model", duplicating `models.name` via `model_id`.
- **Decision:** "My bad, it is the name of the vehicle, not model."
- **Outcome:** no schema change — it's a per-vehicle display name, not a redundant model copy.
- **Doc fix:** correct the column description to "the name of the vehicle" (copy-paste from `models`).
  Open question for later: is `name` required (NOT NULL) and/or unique?

### Q7 — `proposed_action` values & semantics (two rounds)
- **Asked (a):** enum mixes future intent (`NO_ACTION`) with completed states (`PRICE_REDUCED`, `DESTROY`).
- **Decision (a):** "Model as completed status."
- **Asked (b):** completed-only can't express the requirement's literal example "Price Reduction Planned".
- **Decision (b):** "Cover both planned + done."
- **Outcome:** enum covers the full lifecycle — `NONE`, `PRICE_REDUCTION_PLANNED`, `PRICE_REDUCED`, `DESTROYED`.
- **Doc fix:** rename the column (`proposed_action` no longer fits a "planned + done" enum — suggest
  `action` or `action_status`) and enforce with a `CHECK` constraint.

### Q8 — Vehicle identity (VIN)
- **Asked:** two identical units are only distinguishable by surrogate `id`; a dealership keys on VIN.
- **Decision:** "Add VIN column."
- **Outcome:** add `vin` (UNIQUE). Decide nullability: nullable keeps POC seeding simple; NOT NULL if every row has one.

### Q9 — Constraints & indexes
- **Asked:** doc lists columns/types only — no NOT NULL, ON DELETE, CHECK, or FK indexes.
- **Decision:** "Minimal subset."
- **Outcome (document now):**
  - index on `stocked_vehicles.model_id` (FK),
  - `CHECK` on the action enum,
  - `NOT NULL` on `name` / `price`.
- **Deferred to later:** FK `ON DELETE RESTRICT`, index on `models.manufacturer_id`,
  `updated_at` maintenance mechanism, `BIGINT` vs `INT4` for surrogate keys.

---

## Resolved design (summary)

| Concern | Resolution |
|---|---|
| Scope | Docs = full system; code = POC. Add POC-boundary marker. |
| Caching | List uncached; single-record cached + invalidated on write. |
| API | Status codes authoritative; `{error, message, data}` body envelope. |
| Aging basis | `created_at` = stock-in date; threshold configurable. |
| Out-of-stock | Soft-delete only (no status column). |
| `stocked_vehicles.name` | Vehicle name (doc typo to fix). |
| Action field | `NONE` / `PRICE_REDUCTION_PLANNED` / `PRICE_REDUCED` / `DESTROYED`; rename column; add CHECK. |
| Vehicle identity | Add `vin` (UNIQUE). |
| Constraints | Minimal now: FK index (`model_id`), CHECK (action), NOT NULL (`name`, `price`). |

---

## Non-blocking notes (for the user to fold in later)

1. **Typos:** "Softare" → "Software"; "Reponse" → "Response".
2. **`price DECIMAL(16,4)`** has no currency column — state the single-currency assumption.
3. **"Persistent cache" Redis (optional)** overlaps with the "Cache" Redis — clarify its distinct purpose or drop it.
4. **`created_at` / `updated_at`** should be `NOT NULL DEFAULT NOW()`; `updated_at` maintenance (trigger vs app) unspecified.
5. **Verify the architecture diagram** against the text — it couldn't be rendered during this review.
