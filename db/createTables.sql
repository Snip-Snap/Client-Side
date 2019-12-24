CREATE TABLE barber(
    barberID            serial          Primary Key,
    fullName            varchar(30)     not null,
    gender              boolean,
    phoneNumber         varchar(15)     unique not null,
    userName            varchar(30)     not null,
    hashedPassword      varchar         not null
);

CREATE TABLE schedule(
    scheduleID        serial        Primary Key,
    barberID          integer       not null,
    workDate          timestamp     not null,
    startTime         timestamp     not null,
    endTime           timestamp     not null,
    cancelled         boolean       not null,

    Constraint schedule_barberID_fkey Foreign Key (barberID)
        References barber (barberID) Match Simple
        On Update No Action On Delete No Action
);