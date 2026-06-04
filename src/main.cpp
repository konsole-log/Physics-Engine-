#include "raylib.h"
#include <iostream>
using namespace std;
int main() {
  InitWindow(800, 450, "");
  SetTargetFPS(60);
  Vector2 position = {20.0f, 10.0f};
  cout << position.x << position.y;
  while (!WindowShouldClose()) {
    BeginDrawing();
    ClearBackground(BLACK);
    DrawCircle(position.x, position.y, 20.0f, RED);
    EndDrawing();
  }

  CloseWindow();
  return 0;
}
