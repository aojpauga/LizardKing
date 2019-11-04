/**/
/*I left some comments as examples of how to create Zones, rooms and exits*/
/**/
PRAGMA foreign_keys=ON;
BEGIN TRANSACTION;
CREATE TABLE zones (
    id              INTEGER PRIMARY KEY,
    name            TEXT NOT NULL
);
INSERT INTO "zones" VALUES(1,'{ 1 10} Natedog Lizard Den');
/*INSERT INTO "zones" VALUES(1,'{ 1 10} Generic Smurf Village');
INSERT INTO "zones" VALUES(3,'{ 1 20} Copper  Plains of the North');
INSERT INTO "zones" VALUES(6,'{ 5 35} Hatchet New Ofcol');*/
CREATE TABLE rooms (
    id              INTEGER PRIMARY KEY,
    zone_id         INTEGER NOT NULL,
    name            TEXT NOT NULL,
    description     TEXT NOT NULL,

    FOREIGN KEY(zone_id) REFERENCES zones(id)
);
INSERT INTO "rooms" VALUES(101,1,'Starting Room','This is the start room where you begin your adventure!
');
INSERT INTO "rooms" VALUES(102,1,'Cool guy room','This is the Cool Guy Only zone!
');


/*
INSERT INTO "rooms" VALUES(101,1,'Dimly Lit Path','This path is made of crushed granite with curious blue streaks in it.
The path continues north to a village, or south to the west gate.
');
INSERT INTO "rooms" VALUES(102,1,'Dimly Lit Path','This section of the road is badly worn and frequently traversed.  There is
faint singing to the north.  The road back to the forest lies to the south.
');
INSERT INTO "rooms" VALUES(103,1,'Dimly Lit Path','This section of the road is badly worn and frequently traversed.  Singing is
getting stronger to the north.  The road back to the forest lies to the south.
');
INSERT INTO "rooms" VALUES(104,1,'Entrance to Smurf Village','You stand at the entrance to a tiny village.  You can see across the whole
village easily.  Little blue creatures with white hats and boots and little 
fluffy tails scamper about.  They sing continuously.  The singing is getting
on your nerves and you have the urge to squash them.
');
INSERT INTO "rooms" VALUES(105,1,'Smurfy Road','You are way too big to be using this road.  But who cares.  Who''s gonna stop
you?  These little blue smurfs?  The singing is growing intense.  Your ears
are starting to ring.  Smurfy Road continues to the north, west, and east.
There is still time to turn back and go south.
');
INSERT INTO "rooms" VALUES(345,3,'Steep slope','You try to climb down the slope.
>
You slip!
You fall and tumble.
You hit your head HARD.
You die.

You''ve fallen, and you can''t get up!
');*/
CREATE TABLE exits (
    from_room_id    INTEGER NOT NULL,
    to_room_id      INTEGER NOT NULL,
    direction       TEXT NOT NULL CHECK(direction IN ('n','e','s','w','u','d')),
    description     TEXT NOT NULL,

    PRIMARY KEY(from_room_id, direction),
    FOREIGN KEY(from_room_id) REFERENCES rooms(id),
    FOREIGN KEY(to_room_id) REFERENCES rooms(id)
);
INSERT INTO "exits" VALUES(101,102,'n','Door to the Cool Guy Room
');
INSERT INTO "exits" VALUES(102,101,'s','Door to the Start Zone
');
/*
INSERT INTO "exits" VALUES(1001,1006,'e','More of the same.
');
INSERT INTO "exits" VALUES(1001,1002,'s','More of the same.
');*/
CREATE TABLE players (
id INTEGER PRIMARY KEY,
name TEXT NOT NULL,
salt TEXT NOT NULL,
hash TEXT NOT NULL,
UNIQUE(name)
);

CREATE TABLE characters (
id INTEGER PRIMARY KEY,
player_name TEXT NOT NULL,
name TEXT NOT NULL,
class TEXT NOT NULL,
level INTEGER NOT NULL,
FOREIGN KEY(player_name) REFERENCES players(name)
);

COMMIT;
