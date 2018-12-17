# Golang microservices assignment

This assignment is meant to evaluate the golang proficiency of full-time engineers.
Your code structure should follow microservices best practices and our evaluation will focus primarily on your ability to follow good design principles and less on correctness and completeness of algorithms. During the face to face interview you will have the opportunity to explain your design choices and provide justifications for the eventually omitted parts.

## Evaluation points in order of importance

- use of clean code which is self documenting
- use of packages to achieve separation of concerns
- use of domain driven design
- use of golang idiomatic principles
- use of docker
- tests for business logic
- use of code quality checkers such as linters and build tools
- use of git with appropriate commit messages
- documentation: README and inline code comments

Results: please share a git repository with us containing your implementation.

Level of experience targeted: EXPERT

Avoid using frameworks such as go-kit and go-micro since one of the purposes of the assignment is to evaluate the candidate ability of structuring the solution in their own way.
Try to progress as far as you can in 2 hours.
If you have questions please make some assumptions and collect in writing your interpretations.

Good luck.

Time limitations: time is measured from when we send you the assignment to the final commit time onto your repository.

## Technical test

Given an archive (data.tar), write 2 services. One service should parse the CSV file.
The file is of unknown size, it can contain several millions of records.
The service has limited resources available (e.g. 200MB ram).
While reading the file, it should call a second service that either creates a new record in a database, or updates the existing one.

The end result should be a database containing 100 client records, representing the latest version found in the CSV. Database can be Map.


It's safe to assume the email addresses are all valid, so are the phone numbers.
The phone numbers are all UK numbers.
Choose the approach that you think is best (ie most flexible).

You can use any method of communication between the two services, although the use of gRPC is a bonus.

## Bonus points

- Use gRPC for the communication between the 2 services
- The phone numbers should all be converted to the same format.
  Include the country code (+44) to each one. This can either be done by adding a separate field in the DB, or by appending the country code to the phone number itself.
- Database in docker container
- Docker-compose file

## Note

As mentioned earlier, the services have limited resources available, and the CSV file can be several hundred megabytes (if not gigabites) in size.
This means that you will not able to read the entire file at once.

We are looking for the ingestor (the service reading the CSV) to be written in a way that is easy to reuse, give or take a few customisations.
The services themselves should handle certain signals correctly (e.g. a TERM or KILL signal should result in a graceful shutdown).

**Do not mention 90 percent of everything or 90poe anywhere on the code or repository name.**
