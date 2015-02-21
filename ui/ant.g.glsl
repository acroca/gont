#version 400
layout(points) in;
layout(triangle_strip, max_vertices = 3) out;

in float _direction[];
in vec3 col[];
out vec3 col2;

void main() {
  col2 = col[0];
  mat4 RotationMatrix = mat4( cos(_direction[0]), -sin(_direction[0]), 0, 0,
                              sin(_direction[0]),  cos(_direction[0]), 0, 0,
                              0,                   0,                  1, 0,
                              0,                   0,                  0, 1 );
  gl_Position = (gl_in[0].gl_Position + vec4(-0.02, -0.01, 0.0, 0.0) * RotationMatrix);
  EmitVertex();
  gl_Position = (gl_in[0].gl_Position + vec4(-0.02,  0.01, 0.0, 0.0) * RotationMatrix);
  EmitVertex();
  gl_Position = (gl_in[0].gl_Position + vec4( 0.02, 0, 0.0, 0.0) * RotationMatrix);
  EmitVertex();

  EndPrimitive();
}
