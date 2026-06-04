#include "raylib.h"
#include <iostream>
using namespace std;
int main() {
  InitWindow(800, 450, "");
  SetTargetFPS(60);
  Vector2 position = {1.0f, 1.0f};
  cout << position.x << position.y;
  while (!WindowShouldClose()) {
    BeginDrawing();
    ClearBackground(BLACK);
    EndDrawing();
  }

  CloseWindow();
  return 0;
}
