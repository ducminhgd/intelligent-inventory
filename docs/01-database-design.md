# Database design

This design is a compact design based on the requirement, and with assumptions:
1. There is a `users` table, and its primary key uses `UUID` data type.
2. A vehicle is stocked, then a record is created. `created_at` on `stocked_vehicles` is the
   stock-in date used as the basis for the aging computation.
3. The aging threshold (90 days by default) is configurable.

## Table `manufacturers`

Each record represents for a maker of the vehicles

| Column name  |   Data type   |                    Description                     |
| ------------ | ------------- | -------------------------------------------------- |
| `id`         | SERIAL        | Primary key of the records                         |
| `name`       | TEXT NOT NULL | The name of the manufacturer                       |
| `created_at` | TIMESTAMPTZ   | Timestamp that the record is created               |
| `created_by` | UUID          | UUID of the user, who created it                   |
| `updated_at` | TIMESTAMPTZ   | Timestamp of the latest that the record is updated |
| `updated_by` | UUID          | UUID of the user, who updated it                   |
| `deleted_at` | TIMESTAMPTZ   | Timestamp that the record is soft-deleted          |
| `deleted_by` | UUID          | UUID of the user, who deleted it                   |

## Table `models`

Each record represents for a model of the vehicles

|    Column name    |   Data type   |                           Description                            |
| ----------------- | ------------- | ---------------------------------------------------------------- |
| `id`              | SERIAL        | Primary key of the records                                       |
| `manufacturer_id` | INT4          | ID of the manufacturer. Foreign key, refer to `manufacturers.id` |
| `name`            | TEXT NOT NULL | The name of the model                                            |
| `created_at`      | TIMESTAMPTZ   | Timestamp that the record is created                             |
| `created_by`      | UUID          | UUID of the user, who created it                                 |
| `updated_at`      | TIMESTAMPTZ   | Timestamp of the latest that the record is updated               |
| `updated_by`      | UUID          | UUID of the user, who updated it                                 |
| `deleted_at`      | TIMESTAMPTZ   | Timestamp that the record is soft-deleted                        |
| `deleted_by`      | UUID          | UUID of the user, who deleted it                                 |

## Table `stocked_vehicles`

Each record represents for a vehicle in stock

| Column name  |       Data type        |                                    Description                                     |
| ------------ | ---------------------- | ---------------------------------------------------------------------------------- |
| `id`         | SERIAL                 | Primary key of the records                                                         |
| `vin`        | TEXT                   | Vehicle Identification Number. Unique across vehicles                              |
| `model_id`   | INT4                   | ID of the model. Foreign key, refer to `models.id`. Indexed                        |
| `name`       | TEXT NOT NULL          | The name of the vehicle                                                            |
| `price`      | DECIMAL(16,4) NOT NULL | Price of the vehicle                                                               |
| `action`     | TEXT                   | Action/status on the vehicle, constrained by `CHECK`. See values below             |
| `created_at` | TIMESTAMPTZ            | Timestamp the vehicle is stocked (record created). Basis for aging. Indexed B-Tree |
| `created_by` | UUID                   | UUID of the user, who created it                                                   |
| `updated_at` | TIMESTAMPTZ            | Timestamp of the latest that the record is updated                                 |
| `updated_by` | UUID                   | UUID of the user, who updated it                                                   |
| `deleted_at` | TIMESTAMPTZ            | Timestamp that the record is soft-deleted                                          |
| `deleted_by` | UUID                   | UUID of the user, who deleted it                                                   |

Values of `action` (enforced with `CHECK (action IN ('NONE', 'PRICE_REDUCTION_PLANNED', 'PRICE_REDUCED', 'DESTROYED'))`):
- `NONE`: no action has been taken for this vehicle.
- `PRICE_REDUCTION_PLANNED`: a price reduction is planned for this vehicle.
- `PRICE_REDUCED`: the price of this vehicle has been reduced.
- `DESTROYED`: the vehicle has been destroyed.

## Open items (deferred)

- `vin` nullability: nullable keeps POC seeding simple; use `NOT NULL` if every vehicle has one.
- FK `ON DELETE` / `ON UPDATE` behavior on `manufacturer_id` and `model_id`.
- Index on `models.manufacturer_id`.
- `updated_at` maintenance mechanism (trigger vs application).
- `BIGINT` vs `INT4` for surrogate keys.
