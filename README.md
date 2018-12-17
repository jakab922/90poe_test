# Description of the solution

## Folder structure

The source files live under the "github.com/jakab922/phone_data/" folder.
The "server" folder contains the code for the server and similarly the "client" folder for the client. Also there is a utils folder which holds the shared code.
We have a Dockerfile for both of them with the appropriate suffix.
I wanted to create a helm chart for this but ran out of time.

## About the client

So the client reads at most a 100 lines(the batch size) in one go from the csv file and creates a slice of PhoneData objects from it. It also cleans the
phone number in the process. By cleaning I mean it removes all non-numeric characters but keeps the + sign. Then it expects the phone number to be in
either "0123456789" or "+44123456789" format and transforms the number to the latter. It serializes the slice of PhoneData objects and tries to send
them to the server and that's it.

## About the server

It listens on the "/store" endpoint for json data that can be parsed into a slice of PhoneData which in turn can be stored in a psql database. Also we check if the table we want to persist the data to exists.

## About the helm chart and connecting the pieces

I wanted to create a helm chart with the psql db and the server as a deployment and the client as a job, but I ran out of time so I skipped that. Also I haven't done the GRPC service since I've never used it before and it would have been too much new information.

How to run this? 

I have created a file called env_file on the top level of the folder structure. Open 2 terminals and source the env_file. Compile the client and the server and run first the server in one terminal and a client in an another one. Also I'm assuming you have access to a running psql database with the setting defined in the env_file. If not please install psql and modify the env_file accordingly.

.
