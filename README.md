
## Programming practice

The purpose of this assignment is to assess your comfort with basic programming constructs.

### Technical Requirements

You may only use standard libraries.

### Assignment

The task is to write a program that selects a winner from a list of participants in a tournament. 
The list contains the results of 3 different competitions, where participants can participate in one or more of the competitions. 
You should read the results from a text file and return the winner of the entire tournament. 
To win, a participant must have participated in all 3 competitions and have the fastest total competition time (the sum of the time for each competition). 
Try to solve the problem in an object-oriented manner and as if it were part of a larger program.

Each line in the file represents a time for a participant in a competition and consists of:
Participant's name
Unique ID
Start time
End time
Competition type (race type)
These are separated by commas and no spaces, as shown in the example below:
Steve Carell, 3100693, 14:07:10, 14:14:05, 1000m

There are 3 different types of competitions: 1000m, eggRace, and sackRace. 
All participants can participate in all competitions, but some participants have only participated in some of the competitions. 
If a participant is in multiple competitions, they appear on multiple lines in the file, with the same name and ID, as shown in the example below:
Steve Carell, 3100693, 14:07:10, 14:14:05, 1000m
Steve Carell, 3100693, 12:16:11, 12:19:08, egg Race

You want to return the name and ID of the winner along with their average time. 
If there are multiple people with the same time, return all winners.

Don't forget to validate the different values and provide useful error messages when something goes wrong in the program. 
If reading a line fails, log the information about the line and continue reading the file. 
It should be possible to read a file with errors and still get a winner.

There is an example file that you can use, race-results.txt. 
It contains various errors that could occur in a file, but it should still return a winner. 
You should be able to read other files with the same format as race-results.txt, the file is just an example.

### Assumptions:

You cannot assume that each line is correctly formatted. 
Return a clear error message when it is misformatted. 
However, you can assume that none of the values contain a comma.

Validate the values according to the following rules:

A name contains only letters and spaces.
An ID contains only numbers.
Time is written in the format HH:mm:ss.
There are three types of competitions: 1000m, egg Race, sackRace.
Keep in mind that there may be additional validations you want to perform in your program that are not mentioned above. 
For example, if two participants have different names but the same ID, there is an error in the file, and the program should return an error message.


