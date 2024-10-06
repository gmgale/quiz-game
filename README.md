Players can join a game using a code.
Players can enter their names.
The server displays questions and leaderboards.
The game starts when the host clicks "Start Game."
Each question has a time limit.
Scoring is based on correctness and response time.
Questions are imported from a JSON file.
The game runs on a local network and is discoverable by other devices.
Accessible via Safari (web clients).


3.1 Start Game (Host)
The host needs to start the game to begin the quiz. We've added a "Start Game" button in the CreateGameComponent.

Steps for the Host:

Create the Game
Click "Create Game" in the navigation menu.
Click the "Create Game" button.
A game code will be generated and displayed.
Share the Game Code
Provide the game code to players so they can join.
Start the Game
Once all players have joined, click the "Start Game" button.
This will send a request to the backend to start the game session.
All connected players will receive the first question via WebSocket.
4. Playing the Game

4.1 Answering Questions
Players receive questions in real-time and can submit answers.

Steps for Players:

Wait for the Game to Start
After joining, wait for the host to start the game.
A message may indicate that the game is waiting to start.
Answer Questions
When a question appears, select an option by clicking on it.
You will receive immediate feedback on whether your answer was correct.
Proceed to Next Questions
The next question will appear automatically after the time limit or when the host advances the game.
5. Viewing the Leaderboard

At the end of the game, players will be redirected to the leaderboard.

Steps:

The leaderboard component displays the final scores.
Players can see their rank and compare scores.
6. Additional Notes

6.1 Synchronizing Game State
Ensure that the host starts the game only after all players have joined.
Players who try to join after the game has started may receive an error.
6.2 Handling Errors
If a player enters an incorrect game code, display an appropriate error message.
Implement retry mechanisms or prompts to re-enter the game code.
6.3 WebSocket Connections
WebSocket connections are established when players navigate to the game component.
The backend broadcasts questions and game state updates via WebSockets.