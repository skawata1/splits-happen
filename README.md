# Design and Development Challenge – Bowling Score
This script will accept a string of bowling scores as a parameter and output the total score.
If no bowling scores are inputted then it will generate a random bowling game and calculate the total score. 


## Usage
As this is a golang application, a standalone binary with all its dependencies self contained is generated from the build and can be executed on any machine with the same operating system. Executables for windows and linux 64 bit version will be included. 
A bowling score can input using the "bowlingInput" param for the application

The script will interpret the following from the provided string:
*   An "X" represents a strike. It will add 10 plus the next two rolls to the total score.
*   A "/" represents a spare. Paired with the preceding roll, it will add 10 plus the next roll to the total score.
*   A "-" represents a miss. It adds nothing to the total score.
*   Any number represents that number of pins knocked over. It will be added to the total score unless superceded by anything described above.

Example uses of the script are as follows:
*   Calculating a score with all strikes:
```
$ ./splits-happen -bowlingInput="XXXXXXXXXXXX"
Bowling Scores Inputted
[X X X X X X X X X X X X] Total: 300
```
*   Calculating a score with mixed rolls:
```
$ ./splits-happen -bowlingInput="X7/9-X-88/-6XXX81"
Bowling Scores Inputted
[X 7 / 9 - X - 8 8 / - 6 X X X 8 1] Total: 167
```
*   Not providing the "bowlingInput" parameter or by specifying its value as "random", the bowling scores will be randomly generated:
```
$ ./splits-happen
Bowling Scores Randomly Generated
[6 / X 5 4 4 2 X 6 / 5 4 X 4 / - -] Total:  128