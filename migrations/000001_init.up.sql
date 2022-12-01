CREATE TABLE users
(
    Id              BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    Login           TEXT UNIQUE NOT NULL,
    Hashed_Password TEXT        NOT NULL
);

CREATE TABLE flights
(
    Id               BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    Airl_Code        TEXT  NOT NULL,
    Flt_Num          TEXT  NOT NULL,
    Flt_Date         TEXT  NOT NULL,
    Orig_IATA        TEXT  NOT NULL,
    Dest_IATA        TEXT  NOT NULL,
    Departure_Time   TEXT  NOT NULL,
    Arrival_Time     TEXT  NOT NULL,
    Total_Cash       FLOAT NOT NULL,
    Correctly_Parsed bool  NOT NULL
);

CREATE TABLE tickets
(
    Id               BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    Flight_Id        INT REFERENCES flights (id) ON DELETE CASCADE NOT NULL,

    Airl_Code        TEXT                                          NOT NULL,
    Flt_Num          TEXT                                          NOT NULL,
    Flt_Date         TEXT                                          NOT NULL,

    Ticket_Code      TEXT                                          NOT NULL,
    Ticket_Capacity  INT                                           NOT NULL,
    Ticket_Type      TEXT                                          NOT NULL,

    Amount           INT                                           NOT NULL,
    Total_Cash       FLOAT                                         NOT NULL,

    Correctly_Parsed bool                                          NOT NULL
);

CREATE TABLE refresh_tokens
(
    Id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    Token TEXT UNIQUE NOT NULL,
    Expiration_Time_Unix BIGINT NOT NULL,
    Created_Time_Unix BIGINT NOT NULL
);