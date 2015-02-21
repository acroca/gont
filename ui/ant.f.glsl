#version 400
uniform sampler2D tex;

in vec2 texCoord;

out vec4 outColor;

void main()
{
  outColor = vec4(texture(tex, texCoord).rgb, 1);
}
