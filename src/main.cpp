#include "raylib.h"
#include <cmath>
using namespace std;
const int screenWidth = 800;
const int screenHeight = 450;
class Complex {
private:
  float x, y, r, ang;

public:
  Complex(float x, float y) {
    this->x = x;
    this->y = y;
  }
  ~Complex() {}
  float calculateRadius() { return sqrt(x * x + y * y); }
  float calculateAngle() { return atan(y / x); }
  void drawComplexPoint() {
    DrawCircleV({x + screenWidth, y - screenHeight}, 5.0f, RED);
    DrawLineV({(float)screenWidth / 2, (float)screenHeight / 2},
              {x + screenWidth, y - screenHeight}, WHITE);
  }
};

int main() {
  Complex c(10.0f, 20.0f);
  InitWindow(800, 450, "Complex Number Graphing");
  SetTargetFPS(60);
  while (!WindowShouldClose()) {
    BeginDrawing();
    DrawGrid(10, 20.0f);
    c.drawComplexPoint();
    EndDrawing();
  }
  CloseWindow();
}
