#version 400
in vec3 col2;
out vec4 outColor;

void main()
{
  outColor = vec4(col2, 1);
}
