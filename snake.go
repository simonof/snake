package main

import (
	"fmt"
	"math/rand"
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
	gridSize     = 20
)

type Snake struct {
	segments     []raylib.Vector2
	direction    raylib.Vector2
	moveCounter  int
	speedCounter int
	speed        int
}

type Food struct {
	position raylib.Vector2
}

func main() {
	// Initialize Raylib
	raylib.InitWindow(screenWidth, screenHeight, "Snake Game")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	snake := Snake{
		segments: []raylib.Vector2{{10, 10}, {9, 10}, {8, 10}},
		direction: raylib.Vector2{
			X: 1,
			Y: 0,
		},
		speed: 10,
	}

	food := Food{}

	score := 0

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for !raylib.WindowShouldClose() {
		// Update
		snake.moveCounter++
		if snake.moveCounter >= snake.speed {
			snake.moveCounter = 0

			// Move the snake
			// Add new head segment
			newHead := raylib.Vector2{
				X: snake.segments[0].X + snake.direction.X,
				Y: snake.segments[0].Y + snake.direction.Y,
			}
			snake.segments = append([]raylib.Vector2{newHead}, snake.segments...)

			// Check for collision with food
			if newHead.X == food.position.X && newHead.Y == food.position.Y {
				// Spawn new food
				food.position = raylib.Vector2{X: float32(rand.Int31n(int32(screenWidth / gridSize))), Y: float32(rand.Int31n(int32(screenHeight / gridSize)))}
				score++

				// Increase speed after every 10 food items
				if score%10 == 0 {
					snake.speed -= 2
					if snake.speed < 2 {
						snake.speed = 2
					}
				}
			} else {
				// Remove tail segment
				snake.segments = snake.segments[:len(snake.segments)-1]
			}

			// Check for collision with walls
			if snake.segments[0].X < 0 || snake.segments[0].X >= screenWidth/gridSize ||
				snake.segments[0].Y < 0 || snake.segments[0].Y >= screenHeight/gridSize {
				raylib.CloseWindow()
			}

			// Check for collision with self
			for i := 1; i < len(snake.segments); i++ {
				if snake.segments[0].X == snake.segments[i].X && snake.segments[0].Y == snake.segments[i].Y {
					raylib.CloseWindow()
				}
			}
		}

		// Check for input
		if raylib.IsKeyPressed(raylib.KeyRight) && snake.direction.X != -1 {
			snake.direction = raylib.Vector2{X: 1, Y: 0}
		}
		if raylib.IsKeyPressed(raylib.KeyLeft) && snake.direction.X != 1 {
			snake.direction = raylib.Vector2{X: -1, Y: 0}
		}
		if raylib.IsKeyPressed(raylib.KeyUp) && snake.direction.Y != 1 {
			snake.direction = raylib.Vector2{X: 0, Y: -1}
		}
		if raylib.IsKeyPressed(raylib.KeyDown) && snake.direction.Y != -1 {
			snake.direction = raylib.Vector2{X: 0, Y: 1}
		}

		// Begin Drawing
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RayWhite)

		// Draw snake
		for _, segment := range snake.segments {
			raylib.DrawRectangle(int32(segment.X*gridSize), int32(segment.Y*gridSize), gridSize, gridSize, raylib.Black)
		}

		// Draw food
		raylib.DrawRectangle(int32(food.position.X*gridSize), int32(food.position.Y*gridSize), gridSize, gridSize, raylib.Red)

		// Draw score
		raylib.DrawText(fmt.Sprintf("Score: %d", score), 10, 10, 20, raylib.Black)

		raylib.EndDrawing()
		// End Drawing
	}
}
