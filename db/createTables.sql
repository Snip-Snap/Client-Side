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

create table client(
    clientID            serial          Primary Key,
    fullName            varchar(30)     not null,
    gender              boolean,
    phoneNumber         varchar(15)     unique not null,
    userName            varchar(30)     not null,
    hashedPassword      varchar         not null
);

create table review(
    reviewID            serial          Primary Key,
    clientID            integer         not null,
    barberID            integer         not null,
    apptID              integer         not null,
    comments            varchar(128)     not null,
    rating              smallint        not null,

    Constraint review_clientID_fkey 
        Foreign Key (clientID) 
        References client (clientID)
        Match Simple
        On Update No Action On Delete No Action,

    Constraint review_barberID_fkey 
        Foreign Key (barberID) 
        References barber (barberID)
        Match Simple
        On Update No Action On Delete No Action
);