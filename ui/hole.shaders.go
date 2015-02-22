package ui

const (
	holeV = `
    #version 400

    in vec2 position;

    void main() {
      gl_Position = vec4(position, 0.0, 1.0);
    }`
	holeG = `
    #version 400
    layout(points) in;
    layout(line_strip, max_vertices = 11) out;

    void main() {
			int lines = 10;
			float step = 6.38 / lines;
			for(float i = 0; i <= 6.38 ; i+=step){
	      gl_Position = (gl_in[0].gl_Position + vec4(0.02 * cos(i), 0.02 * sin(i), 0.0, 0.0));
	      EmitVertex();
			}
      EndPrimitive();
    }`
	holeF = `
    #version 400

    out vec4 outColor;

    void main()
    {
      outColor = vec4(0.6, 0.1, 0.1, 1);
    }`
)
