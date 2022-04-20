-- name: initialize-database
CREATE TABLE IF NOT EXISTS beacons (
    "id"        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"ip"        TEXT NOT NULL,	
	"callback"  TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS listeners (
    "id"        integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"protocol"  TEXT NOT NULL,	
    "port"      INTEGER
);


-- name: get-all-listeners
SELECT * FROM listeners;

--name: sample-data
INSERT INTO beacons (ip, callback) VALUES ("127.0.0.1", "4:43pm PST");
INSERT INTO listeners (protocol, port) VALUES ("ICMP", "N/A");