package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)

	titleHeight := 200

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Congrats! You created your first window!", 190, int32(titleHeight), 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
