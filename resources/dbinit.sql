BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "log" (
	"uid"	TEXT NOT NULL UNIQUE,
	"contest_id"	TEXT NOT NULL,
	"sta_callsign"	TEXT NOT NULL,
	"dx_callsign"	TEXT NOT NULL,
	"time"	INTEGER NOT NULL,
	"mode"	TEXT NOT NULL,
	"rst_sent"	TEXT NOT NULL,
	"rst_rcvd"	TEXT NOT NULL,
	"exch_sent"	TEXT NOT NULL,
	"exch_rcvd"	TEXT NOT NULL,
	"freq_hz"	INTEGER NOT NULL,
	PRIMARY KEY("uid")
);
CREATE TABLE IF NOT EXISTS "contests" (
	"uid"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"start_time"	INTEGER NOT NULL,
	"end_time"	INTEGER NOT NULL,
	"config"	TEXT NOT NULL,
	"contest_profile"	TEXT NOT NULL,
	"station_callsign"	TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS "meta_data" (
	"key"	TEXT NOT NULL UNIQUE,
	"value"	TEXT NOT NULL,
	PRIMARY KEY("key")
);
COMMIT;
