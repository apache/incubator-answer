import Color from 'color';

/**
 * Bootstrap Color Weight:
 * $blue-100: tint-color($blue, 80%) !default;
 * $blue-200: tint-color($blue, 60%) !default;
 * $blue-300: tint-color($blue, 40%) !default;
 * $blue-400: tint-color($blue, 20%) !default;
 * $blue-500: $blue !default;
 * $blue-600: shade-color($blue, 20%) !default;
 * $blue-700: shade-color($blue, 40%) !default;
 * $blue-800: shade-color($blue, 60%) !default;
 * $blue-900: shade-color($blue, 80%) !default;
 */

/**
 *  The `weight` parameter in `Color`:
 *    1. Must use decimals rather than percentages. eg: color.mix(Color("blue"), 0.6)
 *    2. The value is the difference between `1 - $weight` in `bootstrap`.
 *      eg: color.mix(Color("blue"), 0.6) === shade-color($blue, 40%) !default
 */

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
