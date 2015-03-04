package ui

var (
	antShaders = &Shader{
		Vertex: `
			#version 400

			in vec2 position;
			in float direction;
			in float redAnt;

			out float _direction;
			out float _redAnt;

			void main() {
			  _direction = direction;
			  _redAnt = redAnt;
			  gl_Position = vec4(position, 0.0, 1.0);
			}`,
		Geometry: `
			#version 400
			layout(points) in;
			layout(triangle_strip, max_vertices = 5) out;

			in float _direction[];
			in float _redAnt[];

			out float redAnt;
			out vec2 texCoord;

			void main() {
				redAnt = _redAnt[0];
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
			}`,
		Fragment: `
			#version 400
			uniform sampler2D tex;

			in float redAnt;
			in vec2 texCoord;

			out vec4 outColor;

			void main()
			{
			  outColor = texture(tex, texCoord).rgba;
				if(redAnt > 0.5){
					outColor.r = (2.0 + outColor.r) / 3;
					outColor.g = outColor.g / 20;
					outColor.b = outColor.b / 20;
				}
			}`,
	}
)
