import ReactECharts from "echarts-for-react";
import { createMembershipSeries } from "../utils/membershipFunctions.js";

function getDesignToken(tokenName) {
  if (typeof window === "undefined") {
    return `var(${tokenName})`;
  }

  return getComputedStyle(document.documentElement).getPropertyValue(tokenName).trim() || `var(${tokenName})`;
}

function MembershipChart({ variable, value, showSelectedValue = true }) {
  const colors = {
    chart: [
      getDesignToken("--color-chart-1"),
      getDesignToken("--color-chart-2"),
      getDesignToken("--color-chart-3"),
      getDesignToken("--color-chart-4"),
      getDesignToken("--color-chart-5"),
    ],
    text: getDesignToken("--color-text-soft"),
    muted: getDesignToken("--color-text-muted"),
    subtle: getDesignToken("--color-text-subtle"),
    border: getDesignToken("--color-border"),
    grid: getDesignToken("--color-chart-grid"),
    surface: getDesignToken("--color-surface"),
    accent: getDesignToken("--color-accent"),
  };

  const option = {
    color: colors.chart,
    animation: false,
    backgroundColor: "transparent",
    grid: {
      left: 42,
      right: 18,
      top: 38,
      bottom: 34,
    },
    legend: {
      top: 0,
      left: 0,
      itemWidth: 12,
      itemHeight: 8,
      textStyle: {
        color: colors.muted,
        fontSize: 11,
        fontFamily: "Space Grotesk, Inter, sans-serif",
      },
    },
    tooltip: {
      trigger: "axis",
      backgroundColor: colors.surface,
      borderColor: colors.border,
      textStyle: {
        color: colors.text,
        fontSize: 12,
      },
      valueFormatter: (tooltipValue) => Number(tooltipValue).toFixed(2),
    },
    xAxis: {
      type: "value",
      min: variable.universe[0],
      max: variable.universe[1],
      axisLabel: {
        color: colors.subtle,
        fontSize: 11,
        formatter: variable.axisFormatter,
      },
      axisLine: {
        lineStyle: {
          color: colors.border,
        },
      },
      splitLine: {
        lineStyle: {
          color: colors.grid,
        },
      },
    },
    yAxis: {
      type: "value",
      min: 0,
      max: 1,
      interval: 0.5,
      axisLabel: {
        color: colors.subtle,
        fontSize: 11,
      },
      axisLine: {
        lineStyle: {
          color: colors.border,
        },
      },
      splitLine: {
        lineStyle: {
          color: colors.grid,
        },
      },
    },
    series: createMembershipSeries(variable, value, {
      showSelectedValue,
      selectedColor: colors.accent,
      selectedTextColor: colors.text,
    }),
  };

  return (
    <ReactECharts
      option={option}
      notMerge
      lazyUpdate
      style={{ height: 260, width: "100%" }}
    />
  );
}

export default MembershipChart;
