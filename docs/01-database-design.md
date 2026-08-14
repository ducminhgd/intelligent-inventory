# Database design

This design is a compact design based on the requirement, and with assumptions:
1. There is a `users` table, and its primary key uses `UUID` data type.

## Table `manufacturers`

Each record represents for a maker of the vehicles

| Column name  |  Data type  |                    Description                     |
| ------------ | ----------- | -------------------------------------------------- |
| `id`         | SERIAL      | Primary key of the records                         |
| `name`       | TEXT        | The name of the manufacturer                       |
| `created_at` | TIMESTAMPTZ | Timestamp that the record is created               |
| `created_by` | UUID        | UUID of the user, who created it                   |
| `updated_at` | TIMESTAMPTZ | Timestamp of the latest that the record is updated |
| `updated_by` | UUID        | UUID of the user, who updated it                   |
| `deleted_at` | TIMESTAMPTZ | Timestamp that the record is soft-deleted          |
| `deleted_by` | UUID        | UUID of the user, who deleted it                   |

## Table `models`

Each record represents for a model of the vehicles

|    Column name    |  Data type  |                           Description                            |
| ----------------- | ----------- | ---------------------------------------------------------------- |
| `id`              | SERIAL      | Primary key of the records                                       |
| `manufacturer_id` | INT4        | ID of the manufacturer. Foreign key, refer to `manufacturers.id` |
| `name`            | TEXT        | The name of the model                                            |
| `created_at`      | TIMESTAMPTZ | Timestamp that the record is created                             |
| `created_by`      | UUID        | UUID of the user, who created it                                 |
| `updated_at`      | TIMESTAMPTZ | Timestamp of the latest that the record is updated               |
| `updated_by`      | UUID        | UUID of the user, who updated it                                 |
| `deleted_at`      | TIMESTAMPTZ | Timestamp that the record is soft-deleted                        |
| `deleted_by`      | UUID        | UUID of the user, who deleted it                                 |

## Table `stocked_vehicles`

Each record represents for a vehicle in stock

|    Column name    |   Data type   |                     Description                      |
| ----------------- | ------------- | ---------------------------------------------------- |
| `id`              | SERIAL        | Primary key of the records                           |
| `model_id`        | INT4          | ID of the model. Foreign key, refer to `models.id`   |
| `name`            | TEXT          | The name of the model                                |
| `price`           | DECIMAL(16,4) | Price of the vehicle                                 |
| `proposed_action` | TEXT          | Enum of proposed actions, indexed                    |
| `created_at`      | TIMESTAMPTZ   | Timestamp that the record is created, indexed B-Tree |
| `created_by`      | UUID          | UUID of the user, who created it                     |
| `updated_at`      | TIMESTAMPTZ   | Timestamp of the latest that the record is updated   |
| `updated_by`      | UUID          | UUID of the user, who updated it                     |
| `deleted_at`      | TIMESTAMPTZ   | Timestamp that the record is soft-deleted            |
| `deleted_by`      | UUID          | UUID of the user, who deleted it                     |

Values of `proposed_action`:
- `NO_ACTION`: no action is been taken for this vehicle.
- `PRICE_REDUCED`: the price of this vehicle is reduced
- `DESTROY`: the vehicle needs to be destroyed.