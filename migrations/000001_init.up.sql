CREATE TABLE users
(
    Id             SERIAL PRIMARY KEY UNIQUE NOT NULL,
    Login          TEXT UNIQUE               NOT NULL,
    Hashed_Password TEXT                      NOT NULL
);

CREATE TABLE flights
(
    Id              SERIAL PRIMARY KEY UNIQUE   NOT NULL,
    Airl_Code        TEXT                        NOT NULL,
    Flt_Num          TEXT                        NOT NULL,
    Flt_Date         TEXT                        NOT NULL,
    Orig_IATA        TEXT                        NOT NULL,
    Dest_IATA        TEXT                        NOT NULL,
    Departure_Time   TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    Arrival_Time      TIMESTAMP                   NOT NULL,
    Total_Cash       FLOAT                       NOT NULL,
    Correctly_Parsed bool                        NOT NULL
);

CREATE TABLE tickets
(
    Id              SERIAL PRIMARY KEY UNIQUE                     NOT NULL,
    Flight_Id        INT REFERENCES flights (id) ON DELETE CASCADE NOT NULL,

    Airl_Code        TEXT                                          NOT NULL,
    Flt_Num          TEXT                                          NOT NULL,
    Flt_Date         TEXT                                          NOT NULL,

    Ticket_Code      TEXT                                          NOT NULL,
    Ticket_Capacity  INT                                           NOT NULL,
    Ticket_Type      TEXT                                          NOT NULL,

    Amount          INT                                           NOT NULL,
    Total_Cash       FLOAT                                         NOT NULL,

    Correctly_Parsed bool                                          NOT NULL
);
