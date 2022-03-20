Consists of 2 apps:
1. **NATS Publisher:** A small console application that continuously publishes (at a predefined interval) JSON messages on a
NATS channel. Each JSON message contains random float values for 3 sensors.

2. **NATS Subscriber:** A small console application that subscribes to the NATS channel, calculates the average of the 3 sensor
readings and writes the results (as well as raw data) to a target PostgreSQL database (in 2 separate
tables for raw data and average).

Steps:
- install database
- create new project 
- create database model
- install nats server and go
- create NATS Publisher console app:
  - connect to nats server
  - continously publish a message string
  - continously publish a message that consists of array with 3 objects
- create NATS Subsriber console app:
  - connect to nats server
  - subscribe to message data
  - save raw data to database
  - find average value for 3 sensory readings
  - save result in different table
