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
    layout(triangle_strip, max_vertices = 5) out;

    void main() {
      gl_Position = (gl_in[0].gl_Position + vec4(-0.02, -0.02, 0.0, 0.0));
      EmitVertex();
      gl_Position = (gl_in[0].gl_Position + vec4(-0.02,  0.02, 0.0, 0.0));
      EmitVertex();
      gl_Position = (gl_in[0].gl_Position + vec4( 0.02,  0.02, 0.0, 0.0));
      EmitVertex();
      gl_Position = (gl_in[0].gl_Position + vec4( 0.02, -0.02, 0.0, 0.0));
      EmitVertex();
      gl_Position = (gl_in[0].gl_Position + vec4(-0.02, -0.02, 0.0, 0.0));
      EmitVertex();

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
