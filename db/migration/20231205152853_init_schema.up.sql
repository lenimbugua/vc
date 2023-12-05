CREATE TABLE "periods" (
  "id" integer PRIMARY KEY,
  "competition_id" integer NOT NULL,
  "end_time" timestamptz NOT NULL,
  "round_id" bigint NOT NULL,
  "start_time" timestamptz NOT NULL DEFAULT (now())
);
