create table client (
	clientID 			serial			Primary Key,
	userName			varchar(25)		not null,
	hashedPassword 		varchar(255)	not null,
	firstName 			varchar(35)		not null,
	lastName 			varchar(35)		not null,
	phoneNumber 		varchar(15)		unique not null,
	gender 				varchar(1)
);
create table shop (
	shopID 				serial Primary Key,
	streetAddr 			VARCHAR(50) not null,
	state 				VARCHAR(50) not null,
	areaCode 			VARCHAR(50) not null,
	city 				VARCHAR(50) not null,
	country 			VARCHAR(50) not null

);

create table barber(
	barberID 			serial			Primary Key,
	shopID				integer			not null,
	userName			varchar(25)		not null,
	hashedPassword 		varchar(255)	not null,
	firstName 			varchar(35)		not null,
	lastName 			varchar(35)		not null,
	phoneNumber 		varchar(15)		unique not null,
	gender 				varchar(1),
	DOB 				date 			not null,
	hireDate 			date			not null,
	dismissDate 		date,
	seatNum 			integer			not null,

    Constraint barberID_shopID_fkey Foreign Key (shopID)
        References shop (shopID) Match Simple
        On Update No Action On Delete No Action
);



create table appointment (
	apptID              serial          Primary Key,
	clientID 			integer			not null,
	barberID 			integer			not null,
	clientCancelled		boolean			not null,
	paymentType			varchar(10)		not null,
	barberCancelled		boolean			not null,
	apptDate			date			not null,
	startTime			time			not null,
	endTime				time			not null,
	
    Constraint appointment_clientID_fkey Foreign Key (clientID)
        References client (clientID) Match Simple
        On Update No Action On Delete No Action,
    
    Constraint appointment_barberID_fkey Foreign Key (barberID)
        References barber (barberID) Match Simple
        On Update No Action On Delete No Action
);


create table service (
	serviceID			serial			Primary Key,
	serviceName 		varchar(255) 	not null,
	serviceDescription	varchar(255)	not null,
	price 				numeric(4,2) 	not null,
	customDuration 		integer 		not null
);

create table review (
	reviewID			serial			Primary Key,
	apptID 				integer			not null,
	rating 				integer 		not null,
	comment 			varchar(255) 	not null,

	Constraint rating Check(rating <= 5 and rating >= 0),
	
    Constraint review_apptID_fkey Foreign Key (apptID)
        References appointment (apptID) Match Simple
        On Update CASCADE On Delete CASCADE
);

create table schedule (
	scheduleID			serial			Primary Key,
	barberID 			integer			not null,
	workDate 			date			not null,
	startTime 			time 			not null,
	endTime 			time			not null,

    Constraint schedule_barberID_fkey Foreign Key (barberID)
        References barber (barberID) Match Simple
        On Update No Action On Delete No Action
);
create table supplier (
	supplierID 		serial			PRIMARY KEY,
	supName 		VARCHAR(50)		not null,
	streetAddr 		VARCHAR(50)		not null,
	state 			VARCHAR(50) 	not null,
	areaCode 		VARCHAR(50)		not null,
	phoneNum 		VARCHAR(50)		not null


);
create table product (
	productID 		serial 			PRIMARY KEY,
	supplierID   	integer 		not null,
	product_Name 	VARCHAR(50)		not null,
	manuFacturer 	VARCHAR(50) 	not null,
	weight 			DECIMAL(4,2)	not null,
	price 			NUMERIC(4,2)	not null,
	quantity 		integer,

	Constraint supplier_supplierID_fkey Foreign Key (supplierID)
        References supplier (supplierID) Match Simple
        On Update No Action On Delete No Action
);

create table cli_order (
	orderID				serial			Primary Key,
	clientID 			integer			not null,
	paymentType 		varchar(10)		not null,
	purchaseTime 		timestamp 		not null,

    Constraint  client_clientID_fkey Foreign Key (clientID)
        References client (clientID) Match Simple
        On Update No Action On Delete No Action
);


create table composed_of (
	orderID        integer         not null,
	productID      integer         not null,
	price           NUMERIC(5,2)    not null,
	quantity        integer         not null,

    Constraint product_productID_fkey Foreign Key (productID)
        References product (productID) Match Simple
        On Update No Action On Delete No Action,
	
	Constraint cli_order_orderID_fkey Foreign Key (orderID)
        References cli_order (orderID) Match Simple
        On Update No Action On Delete No Action,
	
	Constraint composedOf_pk Primary Key (orderID, productID)


);

create table supply_order (
	supplyOrderID serial Primary Key,
	shopID INT not null,
	purchaseTIme VARCHAR(50),
	
	Constraint supply_order_shopID_fkey Foreign Key (shopID)
		REFERENCES shop(shopID) Match Simple 
		On Update No Action ON Delete No Action
);
create table receives (
	supplierID 		integer			not null, 
	supplyOrderID 	integer			not null,

	Constraint supplier_supplierID_fkey Foreign Key (supplierID)
        References supplier (supplierID) Match Simple
        On Update No Action On Delete No Action,

	Constraint supply_order_supplyOrderID_fkey Foreign Key (supplyOrderID)
        References supply_order (supplyOrderID) Match Simple
        On Update No Action On Delete No Action,
	
	Constraint receives_id PRIMARY KEY (supplyOrderID, supplierID)
	
);




create table usedIn (
	productID INT not null,
	apptID INT not null,
	
	Constraint usedIn_usedIn_pk_fkey Foreign Key (apptID)
		REFERENCES appointment(apptID) Match Simple 
		On Update CASCADE ON Delete CASCADE,
	
	Constraint usedIn_pk Primary Key (productID, apptID)

);

create table Includes (
	apptID			integer not null,
	serviceID		integer not null,
	price 			numeric(4,2) not null,
	

	Constraint Includes_serviceID_pk_fkey Foreign Key (serviceID)
		REFERENCES service(serviceID) Match Simple 
		On Update CASCADE ON Delete CASCADE,

	Constraint Includes_apptID_pk_fkey Foreign Key (apptID)
		REFERENCES appointment(apptID) Match Simple 
		On Update CASCADE ON Delete CASCADE,

	Constraint includes_pk Primary Key (serviceID, apptID)

);

ALTER TABLE barber
ADD CONSTRAINT shop_shopID_fk
FOREIGN KEY (shopID)
REFERENCES shop(shopID)
ON UPDATE No Action ON DELETE No Action;