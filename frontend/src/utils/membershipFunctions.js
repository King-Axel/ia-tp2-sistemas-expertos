export function formatTime(minutes) {
  const hours = Math.floor(minutes / 60);
  const remainingMinutes = minutes % 60;

  return `${String(hours).padStart(2, "0")}:${String(remainingMinutes).padStart(2, "0")}`;
}

export function formatValue(value, unit, formatter) {
  if (formatter) {
    return formatter(value);
  }

  return `${value}${unit}`;
}

export function clampValue(value, universe) {
  const [min, max] = universe;

  if (!Number.isFinite(value)) {
    return min;
  }

  return Math.min(Math.max(value, min), max);
}

export function generateMembershipPoints(fuzzySet, universe, setIndex, setsCount) {
  const [start, end] = fuzzySet.range;
  const midpoint = (start + end) / 2;

  if (setIndex === 0) {
    return [
      [universe[0], 1],
      [midpoint, 1],
      [end, 0],
    ];
  }

  if (setIndex === setsCount - 1) {
    return [
      [start, 0],
      [midpoint, 1],
      [universe[1], 1],
    ];
  }

  // Intermediate sets use triangular membership functions.
  return [
    [start, 0],
    [midpoint, 1],
    [end, 0],
  ];
}

export function calculateLeftShoulderDegree(value, range) {
  const [start, end] = range;
  const midpoint = (start + end) / 2;

  if (value < start || value > end) {
    return 0;
  }

  if (value <= midpoint) {
    return 1;
  }

  return (end - value) / (end - midpoint);
}

export function calculateTriangularDegree(value, range) {
  const [start, end] = range;
  const midpoint = (start + end) / 2;

  if (value < start || value > end) {
    return 0;
  }

  if (value === midpoint) {
    return 1;
  }

  if (value < midpoint) {
    return (value - start) / (midpoint - start);
  }

  return (end - value) / (end - midpoint);
}

export function calculateRightShoulderDegree(value, range) {
  const [start, end] = range;
  const midpoint = (start + end) / 2;

  if (value < start || value > end) {
    return 0;
  }

  if (value >= midpoint) {
    return 1;
  }

  return (value - start) / (midpoint - start);
}

export function calculateMembershipDegree(value, fuzzySet, setIndex, setsCount) {
  if (setIndex === 0) {
    return calculateLeftShoulderDegree(value, fuzzySet.range);
  }

  if (setIndex === setsCount - 1) {
    return calculateRightShoulderDegree(value, fuzzySet.range);
  }

  return calculateTriangularDegree(value, fuzzySet.range);
}

export function createSelectedMembershipPoints(variable, value) {
  return variable.sets
    .map((fuzzySet, index) => ({
      name: fuzzySet.name,
      value: [value, calculateMembershipDegree(value, fuzzySet, index, variable.sets.length)],
    }))
    .filter((point) => point.value[1] > 0);
}

export function createMembershipSeries(variable, value, options = {}) {
  const {
    showSelectedValue = true,
    selectedColor = "var(--color-accent)",
    selectedTextColor = "var(--color-text-soft)",
  } = options;
  const membershipSeries = variable.sets.map((fuzzySet, index) => ({
    name: fuzzySet.name,
    type: "line",
    smooth: false,
    symbol: "circle",
    symbolSize: 5,
    showSymbol: false,
    lineStyle: {
      width: 3.5,
    },
    emphasis: {
      focus: "series",
    },
    data: generateMembershipPoints(fuzzySet, variable.universe, index, variable.sets.length),
  }));

  if (!showSelectedValue || value === undefined || value === null) {
    return membershipSeries;
  }

  return [
    ...membershipSeries,
    {
      name: "Valor actual",
      type: "line",
      symbol: "none",
      data: [
        [value, 0],
        [value, 1],
      ],
      lineStyle: {
        color: selectedColor,
        width: 2,
        type: "dashed",
      },
      label: {
        show: true,
        formatter: variable.formatValue ? variable.formatValue(value) : `${value}${variable.unit}`,
        color: selectedTextColor,
        fontSize: 11,
        position: "top",
      },
      tooltip: {
        show: false,
      },
      z: 10,
    },
    {
      name: "Pertenencia actual",
      type: "scatter",
      symbolSize: 9,
      data: createSelectedMembershipPoints(variable, value),
      itemStyle: {
        color: selectedColor,
        borderColor: selectedTextColor,
        borderWidth: 1,
      },
      label: {
        show: true,
        formatter: (params) => Number(params.value[1]).toFixed(2),
        color: selectedTextColor,
        fontSize: 10,
        position: "top",
      },
      tooltip: {
        formatter: (params) =>
          `${params.data.name}<br/>μ = ${Number(params.value[1]).toFixed(2)}`,
      },
      z: 20,
    },
  ];
}
