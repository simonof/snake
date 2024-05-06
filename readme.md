# Snake Game

A classic Snake game implemented in Go using the Raylib library.

## Description

This Snake game is a simple implementation of the classic arcade game where the player controls a snake that grows in length as it eats food items. The player must navigate the snake around the game board, avoiding collisions with the walls and the snake's own body. The game ends if the snake collides with a wall or itself.

The game features a scoreboard that keeps track of the player's score, which increases every time the snake eats a food item. Additionally, the snake's speed increases after every 10 food items eaten, providing a gradually increasing level of difficulty.

## How to Play

1. Use the arrow keys (up, down, left, right) to control the direction of the snake.
2. Guide the snake to eat the food items scattered across the game board.
3. Avoid collisions with the walls and the snake's own body.
4. The game ends if the snake collides with a wall or itself.

## Requirements

- Go programming language
- Raylib library
- Raylib Go bindings (`raylib-go`)

## Installation

1. Install Go from the official website: [https://golang.org/](https://golang.org/)
2. Install Raylib by following the installation instructions on the official website: [https://www.raylib.com/](https://www.raylib.com/)
3. Install the Raylib Go bindings (`raylib-go`) using the following command:

   ```bash
   go get -u github.com/gen2brain/raylib-go/raylib
