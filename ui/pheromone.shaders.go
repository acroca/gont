package ui

const (
	pheromoneV = `
    #version 400

    in vec2 position;
    in float intensity;

		out float _intensity;

    void main() {
			_intensity = intensity;
      gl_Position = vec4(position, 0.0, 1.0);
    }`
	pheromoneG = `
    #version 400
    layout(points) in;
    layout(triangle_strip, max_vertices = 4) out;

		in float _intensity[];

		out float intensity;

    void main() {
			intensity = _intensity[0];

      gl_Position = (gl_in[0].gl_Position + vec4(-0.004, -0.004, 0.0, 0.0));
      EmitVertex();
      gl_Position = (gl_in[0].gl_Position + vec4(-0.004,  0.004, 0.0, 0.0));
      EmitVertex();
      gl_Position = (gl_in[0].gl_Position + vec4( 0.004,  0, 0.0, 0.0));
      EmitVertex();
			gl_Position = (gl_in[0].gl_Position + vec4(-0.004, -0.004, 0.0, 0.0));
      EmitVertex();

      EndPrimitive();
    }`
	pheromoneF = `
    #version 400
		in float intensity;

    out vec4 outColor;

    void main()
    {
      outColor = vec4(0, 0, intensity, 1);
    }`
)
