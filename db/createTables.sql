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
--YEEET
create table cut(
    cutID            serial          Primary Key,
    hairCutStyle     varchar(30),
    cost             integer       not null,
    defaulDuration   integer       not null,
    customDuration   integer       not null
);
create table appointment(
    apptID            serial          Primary Key,
    startTime         time         not null,
    isCancelled       boolean        not null,
    paymentType       varchar(30)    not null,
    apptDate          date           not null
);
create table review(
    reviewID            serial          Primary Key,
    clientID            integer         not null,
    barberID            integer         not null,
    --apptID              integer         not null,
    comments            varchar(128)     not null,
    rating              smallint        not null,

    Constraint review_clientID_fkey 
        Foreign Key (clientID) 
        References  client (clientID)
        Match Simple
        On Update No Action On Delete No Action,

    Constraint review_barberID_fkey 
        Foreign Key (barberID) 
        References  barber (barberID)
        Match Simple
        On Update No Action On Delete No Action

    -- Constraint review_apptID_fkey 
    --     Foreign Key (apptID) 
    --     References  appointment (apptID)
    --     Match Simple
    --     On Update No Action On Delete No Action
);
