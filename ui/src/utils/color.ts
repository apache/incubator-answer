import Color from 'color';

const WHITE = Color('#fff');
const BLACK = Color('#000');

export const mixColour = (baseColor, opColor, weight) => {
  return Color(baseColor).mix(Color(opColor), weight);
};

export const tintColor = (color, weight) => {
  return mixColour(WHITE, color, weight);
};

export const shadeColor = (color, weight) => {
  return mixColour(BLACK, color, weight);
};

export const shiftColor = (color, weight) => {
  if (weight > 0) {
    return shadeColor(color, weight);
  }
  return tintColor(color, weight);
};
