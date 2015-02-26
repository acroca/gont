package ui

var (
	holeShaders = &Shader{
		Vertex: `
	    #version 400

	    in vec2 position;

	    void main() {
	      gl_Position = vec4(position, 0.0, 1.0);
	    }`,
		Geometry: `
	    #version 400
	    layout(points) in;
	    layout(triangle_strip, max_vertices = 14) out;

	    void main() {
				float PI = 3.14159265358979323846264;
				float step = 1.0 / 6;
				float radius = 0.02;

				gl_Position = (gl_in[0].gl_Position + vec4(radius, 0, 0, 0));
				EmitVertex();

				for(float i = step; i < 1.0 ; i+=step){
		      gl_Position = (gl_in[0].gl_Position + vec4(radius * cos(i * PI), radius * sin(i * PI), 0.0, 0.0));
		      EmitVertex();
		      gl_Position = (gl_in[0].gl_Position + vec4(radius * cos(i * PI), -radius * sin(i * PI), 0.0, 0.0));
		      EmitVertex();
				}

				gl_Position = (gl_in[0].gl_Position + vec4(-radius, 0, 0, 0));
				EmitVertex();

	      EndPrimitive();
	    }`,
		Fragment: `
	    #version 400

	    out vec4 outColor;

	    void main()
	    {
	      outColor = vec4(0.6, 0.1, 0.1, 1);
	    }`,
	}
)
