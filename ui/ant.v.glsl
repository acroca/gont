#version 400

in vec2 position;
in float direction;
in uint kind;

out float _direction;

void main() {
  _direction = direction;
  gl_Position = vec4(position, 0.0, 1.0);
}
