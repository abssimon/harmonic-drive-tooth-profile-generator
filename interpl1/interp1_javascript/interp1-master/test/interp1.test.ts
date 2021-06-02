import interp1 from "../src/interp1"

describe('interp1', () => {
  it('interp1 exists', () => {
    expect(interp1).toBeTruthy();
  });

  it('throws an error if the number of sample points does not equal the number of values', () => {
    expect(() => {
      interp1([1, 2, 3], [1, 2, 3, 4], [2, 2.5, 3]);
    }).toThrowError();
  });

  it('throws an error if two values of independent variable occur more than once', () => {
    expect(() => {
      interp1([5, 5, 8], [8, -3, 0], [6]);
    }).toThrowError();
  });

  it('throws an error if query points lie outside of sample points range', () => {
    expect(() => {
      interp1([-5, 0, 7], [3, 9, -8], [-6]);
    }).toThrowError();

    expect(() => {
      interp1([2, 4, 5], [-2, 4, -3], [7]);
    }).toThrowError();
  });

  it('returns an empty array for empty query points', () => {
    expect(interp1([], [], [])).toEqual([]);

    expect(interp1([7, 8, 9], [-5, 9, 81], [])).toEqual([]);
  });

  it('returns exact values', () => {
    expect(interp1([3], [7], [3])).toEqual([7]);
    
    expect(interp1(
      [-8, 0, 5],
      [1, 92, 4],
      [-8, 5, 0],
    )).toEqual([1, 4, 92]);
  });

  it('interpolates values linearly', () => {
    expect(interp1(
      [-7, -3, 0, 4, 9],
      [0, 8, -4, -2, 3],
      [-7, -5, -3, -1.5, 0, 2, 4, 6.5, 9],
      'linear',
    )).toEqual([0, 4, 8, 2, -4, -3, -2, 0.5, 3]);
  });

  it('interpolates values with nearest neighbours', () => {
    expect(interp1(
      [-5, 6, 8],
      [0, 4, -2],
      [-5, 0, 6, 6.99, 7, 8],
      'nearest',
    )).toEqual([0, 0, 4, 4, -2, -2]);
  });

  it('interpolates values with previous neighbours', () => {
    expect(interp1(
      [-9, -7, 0, 1, 2],
      [0, 7, 3, -4, -2],
      [-9, -8, -7.1, -7, -1, 0, 1, 1.75, 2],
      'previous',
    )).toEqual([0, 0, 0, 7, 7, 3, -4, -4, -2]);
  });

  it('interpolates values with next neighbours', () => {
    expect(interp1(
      [-4, -1, 0, 4, 7, 9],
      [-9, 0, 3, -4, 2, -1],
      [-4, -3.9, -1, -0.6, 0, 3, 4, 6, 7, 8, 9],
      'next',
    )).toEqual([-9, 0, 0, 3, 3, -4, -4, 2, 2, -1, -1]);
  });

  it('interpolates values linearly by default', () => {
    const xs: number[] = [-3, 0, 4, 5];
    const vs: number[] = [9, -2, 0, 3];
    const xqs: number[] = [-3, -1.5, 0, 2, 4, 5];
    expect(interp1(xs, vs, xqs)).toEqual(interp1(xs, vs, xqs));
  });

  it('does not require the sample points to be in order', () => {
    expect(interp1(
      [9, -2, 0, 3],
      [8, 0, 3, 1],
      [-2, 0, 1.5, 3],
      'linear',
    )).toEqual([0, 3, 2, 1]);
  })
});
