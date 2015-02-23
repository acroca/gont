package ui

const (
	antV = `
		#version 400

		in vec2 position;
		in float direction;

		out float _direction;

		void main() {
		  _direction = direction;
		  gl_Position = vec4(position, 0.0, 1.0);
		}`
	antG = `
		#version 400
		layout(points) in;
		layout(triangle_strip, max_vertices = 5) out;

		in float _direction[];

		out vec2 texCoord;

		void main() {
		  mat4 RotationMatrix = mat4( cos(_direction[0]), -sin(_direction[0]), 0, 0,
		                              sin(_direction[0]),  cos(_direction[0]), 0, 0,
		                              0,                   0,                  1, 0,
		                              0,                   0,                  0, 1 );

		  gl_Position = (gl_in[0].gl_Position + vec4(-0.02, -0.02, 0.0, 0.0) * RotationMatrix);
		  texCoord = vec2(0, 1);
		  EmitVertex();
		  gl_Position = (gl_in[0].gl_Position + vec4(-0.02,  0.02, 0.0, 0.0) * RotationMatrix);
		  texCoord = vec2(1, 1);
		  EmitVertex();
		  gl_Position = (gl_in[0].gl_Position + vec4( 0.02,  0.02, 0.0, 0.0) * RotationMatrix);
		  texCoord = vec2(1, 0);
		  EmitVertex();
		  gl_Position = (gl_in[0].gl_Position + vec4( 0.02, -0.02, 0.0, 0.0) * RotationMatrix);
		  texCoord = vec2(0, 0);
		  EmitVertex();
		  gl_Position = (gl_in[0].gl_Position + vec4(-0.02, -0.02, 0.0, 0.0) * RotationMatrix);
		  texCoord = vec2(0, 1);
		  EmitVertex();

		  EndPrimitive();
		}`
	antF = `
		#version 400
		uniform sampler2D tex;

		in vec2 texCoord;

		out vec4 outColor;

		void main()
		{
		  outColor = vec4(texture(tex, texCoord).rgb, 1);
		}`
)
