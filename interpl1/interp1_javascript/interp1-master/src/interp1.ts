/**
 * Method of interpolation.
 */
export type InterpolationMethod = 'linear' | 'nearest' | 'next' | 'previous';

/**
 * Point data type.
 */
type Point = [number, number];

/**
 * Finds the index of range in which a query value is included in a sorted
 * array with binary search.
 * @param  xs Array sorted in ascending order.
 * @param  xq Query value.
 * @return    Index of range plus percentage to next index. 
 */
function binaryFindIndex(
  xs: ReadonlyArray<number>,
  xq: number,
): number {
  /* Special case of only one element in array. */
  if (xs.length === 1 && xs[0] === xq) return 0;

  /* Determine bounds. */
  let lower: number = 0;
  let upper: number = xs.length - 1;

  /* Find index of range. */
  while (lower < upper) {
    /* Determine test range. */
    const mid: number = Math.floor((lower + upper) / 2);
    const prev: number = xs[mid];
    const next: number = xs[mid + 1];
    if (xq < prev) {
      /* Query value is below range. */
      upper = mid;
    } else if (xq > next) {
      /* Query value is above range. */
      lower = mid + 1;
    } else {
      /* Query value is in range. */
      return mid + (xq - prev) / (next - prev); 
    }
  }

  /* Range not found. */
  return -1;
}

/**
 * Interpolates a value.
 * @param  vs     Array of values to interpolate between.
 * @param  index  Index of new to be interpolated value.
 * @param  method Kind of interpolation. Can be 'linear', 'nearest', 'next' or 'previous'.
 * @return        Interpolated value.
 */
function interpolate(
  vs: ReadonlyArray<number>,
  index: number,
  method: InterpolationMethod,
): number {
  switch (method) {
    case 'nearest': {
      return vs[Math.round(index)]
    }
    case 'next': {
      return vs[Math.ceil(index)];
    }
    case 'previous': {
      return vs[Math.floor(index)];
    }
    case 'linear':
    default: {
      const prev: number = Math.floor(index);
      const next: number = Math.ceil(index);
      const lambda: number = index - prev;
      return (1 - lambda) * vs[prev] + lambda * vs[next];
    }
  }
}

/**
 * Interpolates values linearly in one dimension.
 * @param  xs     Array of independent sample points.
 * @param  vs     Array of dependent values v(x) with length equal to xs.
 * @param  xqs    Array of query points.
 * @param  method Method of interpolation.
 * @return        Interpolated values vq(xq) with length equal to xqs.
 */
export default function interp1(
  xs: ReadonlyArray<number>,
  vs: ReadonlyArray<number>,
  xqs: ReadonlyArray<number>,
  method: InterpolationMethod = 'linear',
): number[] {
  /*
   * Throws an error if number of independent sample points is not equal to
   * the number of dependent values.
   */
  if (xs.length !== vs.length) {
    throw new Error(
      `Arrays of sample points xs and corresponding values vs have to have
      equal length.`
    );
  }

  /* Combine x and v arrays. */
  const zipped: Point[] = xs.map((x, index) => [x, vs[index]]);

  /* Sort points by independent variabel in ascending order. */
  zipped.sort((a, b) => {
    const diff: number = a[0] - b[0];

    /* Check if some x value occurs twice. */
    if (diff === 0) {
      throw new Error(
        'Two sample points have equal value ' + a[0] + '. This is not allowed.'
      );
    }

    return diff;
  });

  /* Extract sorted x and v arrays */
  const sortedX: number[] = [];
  const sortedV: number[] = [];
  for (let i: number = 0; i < zipped.length; i++) {
    const point: Point = zipped[i];
    sortedX.push(point[0]);
    sortedV.push(point[1]);
  }

  /* Interpolate values */
  const yqs: ReadonlyArray<number> = xqs.map(xq => {
    /* Determine index of range of query value. */
    const index: number = binaryFindIndex(sortedX, xq);

    /* Check if value lies in interpolation range. */
    if (index === -1) {
      throw new Error(
        `Query value ` + xq + ` lies outside of range. Extrapolation is not
        supported.`
      );
    }

    /* Interpolate value. */
    return interpolate(sortedV, index, method);
  });

  /* Return result. */
  return yqs.slice();
}
