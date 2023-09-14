#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

uniform sampler2D texture0;
uniform vec4 colDiffuse;

out vec4 finalColor;

const vec2 size = vec2(1000, 600);
uniform float samples;
uniform float quality = 2.5;

void main(){
    vec4 sum = vec4(0);
    vec2 sizeFactor = vec2(1)/size*quality;

    vec4 source = texture(texture0, fragTexCoord);

    const int range = 2;

    for (int x = -range; x <= range; x++)
    {
        for (int y = -range; y <= range; y++)
        {
            sum += texture(texture0, fragTexCoord + vec2(x, y)*sizeFactor);
        }
    }

    finalColor = ((sum/(samples*samples)) + source)*colDiffuse;
}