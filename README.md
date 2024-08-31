Requirements
It is a CLI-based game, so you need to use the command line to interact with the game. The game should work as follows:

1) When the game starts, it should display a welcome message along with the rules of the game.
2) The computer should randomly select a number between 1 and 100.
3) User should select the difficulty level (easy, medium, hard) which will determine the number of chances they get to guess the number.
4) The user should be able to enter their guess.
5) If the user’s guess is correct, the game should display a congratulatory message along with the number of attempts it took to guess 
   the number.
6) If the user’s guess is incorrect, the game should display a message indicating whether the number is greater or less than the user’s 
   guess.

The game should end when the user guesses the correct number or runs out of chances.

Here is a sample output of the game:
****************************************************************************************************************************************
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess: 50
Incorrect! The number is less than 50.

Enter your guess: 25
Incorrect! The number is greater than 25.

Enter your guess: 35
Incorrect! The number is less than 35.

Enter your guess: 30
Congratulations! You guessed the correct number in 4 attempts.
****************************************************************************************************************************************

How to Run the code

1) Get this code on your computer
2) Install Go Compiler 
3) Go to the file location 
4) Open command
5) $ go build main.Go
6) $ go run main
7) Enjoy the Game

If you don't have go setup in your machine
1) go to https://go.dev/
2) Scroll down you will find the go Compiler
3) copy paste the main.go code to the Compiler
4) hit run
5) Enjoy the Game.
