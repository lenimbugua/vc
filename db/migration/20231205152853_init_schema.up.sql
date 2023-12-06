CREATE TABLE "periods"(
  "id" serial PRIMARY KEY,
  "competition_id" integer NOT NULL,
  "end_time" timestamptz NOT NULL,
  "round_id" bigint NOT NULL,
  "start_time" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "markets" (
  "id" serial PRIMARY KEY,
  "name" varchar(50) 
  "market_id" integer NOT NULL
);

INSERT INTO markets (name, market_id)
VALUES
    ("Correct Score (Full Time)", 1),
    ("Half-Time Score", 2),
    ("Match Result", 3),
    ("Half-Time Result", 4),
    ("Double Chance", 5),
    ("Double Chance (Half-Time)", 6),
    ("Over/Under 1.5", 7),
    ("Over/Under 2.5", 8),
    ("Over/Under 3.5", 9),
    ("Handicap -1",10),
    ("Handicap -2", 11),
    ("Half Time / Full Time", 12),
    ("Total Goals", 13),
    ("Goal:Goal Full Time", 14),
    ("Goal:Goal Half Time",15),
    ("1X2 and Over/Under 1.5", 16),
    ("1X2 and Over/Under 2.5", 17),
    ("1X2 and Over/Under 3.5", 18),
    ("1X2 and Over/Under 4.5", 19),
    ("1X2 and Over/Under 5.5", 20),
    ("1X2 and Goal/No Goal", 21),
    ("Team 1 Over/Under 1.5", 22),
    ("Team 2 Over/Under 1.5", 23),
    ("Team 1 Goal/No Goal", 24),
    ("Team 2 Goal/No Goal", 25),
    ("Total Goals Odd/Even", 26),
    ("Time of First Goal", 27),
    ("First Team to Score",28),
    ("Multi-Goals", 29),
    ("First Player to Score", 30),
    ("Penalty in Match", 31);
  