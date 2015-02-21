#version 400

in vec2 position;
in float direction;
in vec3 color;
in uint kind;

out vec3 col;
out float _direction;

void main() {
  col = color;
  _direction = direction;
  gl_Position = vec4(position, 0.0, 1.0);
}
