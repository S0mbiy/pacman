# Architecture of pacman program

## Aproach
The approach for this implementation of the game pacman is to divide the game in the following parts:
#### - Characters:
  - Pacman:
    - Inside the Pacman character is the code to take the arrow keys to move the character, also within it is checked if Pacman collides with a coin to get the point and delete the coin from the place. After each movement the points are checked to know if the goal is achived.
  - Ghosts:
    - For the ghosts with just one is enough because each one moves randomly, the position of the ghosts is saved inside another array and checked to show the game, the movements of the ghosts are randomly generated but always checked with the objects map to avoid that a ghost can go through the walls.

#### - Map
  - The map is made with an array, each array element represents a space of the map, with this array we can signilize the elements of the map with the following numbers:
    - Wall: -1
    - Empty space: 0
    - Coin: 1
    - Ghost wall: -2 and -3
  - Also there is a second map to signilize the position of the ghosts and Pacman, making possible to check the position of this characters and if they collide with each other to know if the player has lost when it collides with a ghost.

With this approach the game can be checked each time the user inputs a key to move Pacman or when a ghost moves to update the visual part of the game, even when it is all in console, making a notifier approach, meaning that when a change is made the printer is notified by a message sent through a special channel. In the win or lose part is made with channels to notify that the game has finished.
