create table if not exists SensorRaw (
    "Name" varchar(50) not null,
    "Timestamp" bigint not null,
    "Value" decimal not null,
    primary key ("Name", "Timestamp")
);

create table if not exists SensorAverage (
    "Value" decimal not null,
    "Timestamp" bigint not null,
    primary key ("Timestamp")
);



