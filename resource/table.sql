--drop table customer;
--drop table device;
--drop table location;
--drop table locationrole;
--drop table devicetype;
--drop table sku;

--drop table log_device_fail_sku;

CREATE TABLE IF NOT EXISTS  devicetype
(
  id serial NOT NULL,
  device text,
  createtime timestamp with time zone default now(),
  
  CONSTRAINT devicetype_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

CREATE TABLE IF NOT EXISTS  sku
(
  id serial NOT NULL,
  sku text,
  devicetype int,
  createtime timestamp with time zone default now(),
  
  CONSTRAINT sku_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);


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
  devicetype int,
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


drop table deviceinworld;
CREATE TABLE deviceinworld AS (SELECT 
  location.countrycode,
  count(device.serialnumber) AS total 
FROM 
  public.device, 
  public.location
WHERE 
  device.locationid = location.locationid AND 
  device.isdeleted = '0'
GROUP BY
  location.countrycode
ORDER BY 
  location.countrycode
);

ALTER TABLE deviceinworld ADD COLUMN id serial NOT NULL;
ALTER TABLE deviceinworld ADD PRIMARY KEY (id );

drop table deviceinus;
CREATE table deviceinus AS (SELECT 
  location.city, 
  location.stateprovince,
  count(device.serialnumber) AS total 
FROM 
  public.device, 
  public.location
WHERE 
  device.locationid = location.locationid AND
  location.countrycode ='US' AND
  device.isdeleted = '0'
GROUP BY
  location.city,
  location.stateprovince
ORDER BY 
  location.city
);

ALTER TABLE deviceinus ADD COLUMN id serial NOT NULL;
ALTER TABLE deviceinus ADD COLUMN pb980 int default 0;
ALTER TABLE deviceinus ADD COLUMN scd700 int default 0;
ALTER TABLE deviceinus ADD COLUMN forcetriad int default 0;
ALTER TABLE deviceinus ADD COLUMN cooltip int default 0;
ALTER TABLE deviceinus ADD PRIMARY KEY (id);

drop table deviceindmp;
CREATE table deviceindmp AS (select
	device.locationid,
    device.devicetype, 
	device.maintenanceexpirationdate,
	device.serialnumber,
	device.sku,
	device.actualshipdate,
	location.city, 
    location.stateprovince
	
from 
	device,
	location
where device.devicetype >0 AND
      location.locationid = device.locationid AND
      location.countrycode ='US' AND
      device.isdeleted = '0'

);
ALTER TABLE deviceindmp ADD COLUMN id serial NOT NULL;
ALTER TABLE deviceindmp ADD PRIMARY KEY (id);



  
-- executing
update device set devicetype = 1 where sku in (select sku from sku where devicetype = 1);
update device set devicetype = 2 where sku in (select sku from sku where devicetype = 2);
update device set devicetype = 3 where sku in (select sku from sku where devicetype = 3);
update device set devicetype = 4 where sku in (select sku from sku where devicetype = 4);

update location set city = 'WAITE PARK' where id = 487587;



select location.city, count(device.serialnumber) from device,location where device.devicetype = 4 and device.locationid = location.locationid group by city order by city;
