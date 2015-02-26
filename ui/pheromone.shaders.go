package ui

var (
	pheromoneShaders = &Shader{
		Vertex: `
	    #version 400

	    in vec2 position;
	    in float intensity;

			out float _intensity;

	    void main() {
				_intensity = intensity;
	      gl_Position = vec4(position, 0.0, 1.0);
	    }`,
		Geometry: `
	    #version 400
	    layout(points) in;
	    layout(triangle_strip, max_vertices = 14) out;

			in float _intensity[];

			out float intensity;

	    void main() {
				intensity = _intensity[0];

				float PI = 3.14159265358979323846264;
				float step = 1.0 / 4;
				float radius = 0.005;

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
			in float intensity;

	    out vec4 outColor;

	    void main()
	    {
	      outColor = vec4(0, 0, 1, intensity);
	    }`,
	}
)
