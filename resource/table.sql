--drop table customer;
--drop table device;
--drop table location;
--drop table locationrole;

--drop table log_device_fail_sku;



CREATE TABLE IF NOT EXISTS  customer
(
  id serial NOT NULL,
  customerid text,
  name text,
  phone text,
  fax text,
  lastchangedate text,
  distributorflag text,
  isdeleted text,
  batchnumber text,
  createtime timestamp with time zone,
  
  CONSTRAINT customer_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

CREATE TABLE IF NOT EXISTS  device
(
  id serial NOT NULL,
  locationid text,
  customerid text,
  maintenanceexpirationdate text,
  serialnumber text,
  sku text,
  sourcesystem text,
  installcountrycode text,
  lastchangedate text,
  installationdate text,
  actualshipdate text,
  isdeleted text,
  batchnumber text,
  createtime timestamp with time zone,

  CONSTRAINT device_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);


CREATE TABLE IF NOT EXISTS  location
(
  id serial NOT NULL,
  locationid text,
  addressline1 text,
  addressmodifier2 text,
  addressmodifier3 text,
  addressmodifier4 text,
  city text,
  stateprovince text,
  postalcode text,
  countrycode text,
  addressmodifier1 text,
  lastchangedate text,
  isdeleted text,
  batchnumber text,
  createtime timestamp with time zone,

  CONSTRAINT location_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

CREATE TABLE IF NOT EXISTS  locationrole
(
  id serial NOT NULL,
  customerid text, 
  locationid text,
  locationrole text,
  lastchangedate text,
  isdeleted text,
  batchnumber text,
  createtime timestamp with time zone,

  CONSTRAINT locationrole_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

CREATE TABLE IF NOT EXISTS  log_device_fail_sku
(
  id serial NOT NULL,
  locationid text,
  customerid text,
  maintenanceexpirationdate text,
  serialnumber text,
  sku text,
  sourcesystem text,
  installcountrycode text,
  installationdate text,  
  actualshipdate text,
  isdeleted text,
  lastchangedate text,
  batchnumber text,
  exception text,
  createtime timestamp with time zone,

  CONSTRAINT log_device_fail_sku_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);


CREATE VIEW deviceinus AS

SELECT 
  customer.customerid, 
  location.locationid, 
  device.serialnumber, 
  device.maintenanceexpirationdate, 
  device.sku, 
  device.sourcesystem, 
  device.installcountrycode, 
  location.addressline1, 
  location.addressmodifier2, 
  location.addressmodifier3, 
  location.addressmodifier4, 
  location.addressmodifier1, 
  location.city, 
  location.stateprovince, 
  location.countrycode, 
  locationrole.locationrole
FROM 
  public.customer, 
  public.device, 
  public.location, 
  public.locationrole
WHERE 
  customer.customerid = device.customerid AND
  device.locationid = location.locationid AND
  locationrole.customerid = customer.customerid AND 
  locationrole.locationid = location.locationid AND   
  location.countrycode ='US' 
  
  
CREATE VIEW deviceloc AS
SELECT 
  location.locationid, 
  location.city, 
  device.serialnumber, 
  location.countrycode, 
  location.stateprovince
FROM 
  public.device, 
  public.location
WHERE 
  location.locationid = device.locationid AND
  location.countrycode ='US';
  
CREATE VIEW  devicenumincity AS 
SELECT 
	location.city, 
	count(device.serialnumber) AS total 
FROM 
	public.device, 
  	public.location 
WHERE 
	location.locationid = device.locationid AND 
	location.countrycode ='US' 
GROUP BY 
	location.city
ORDER BY
	city;
	
CREATE VIEW  devicenuminprovince AS 
SELECT 
	location.stateprovince, 
	count(device.serialnumber) AS total 
FROM 
	public.device, 
  	public.location 
WHERE 
	location.locationid = device.locationid AND 
	location.countrycode ='US' 
GROUP BY 
	location.stateprovince
ORDER BY
	stateprovince;
  