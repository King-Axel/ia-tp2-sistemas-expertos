import ReactECharts from "echarts-for-react";
import { createMembershipSeries } from "../utils/membershipFunctions.js";

const chartColors = ["#71717a", "#a1a1aa", "#d4d4d8", "#60a5fa", "#0070f3"];

function MembershipChart({ variable, value, showSelectedValue = true }) {
  const option = {
    color: chartColors,
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
        color: "#a1a1aa",
        fontSize: 11,
        fontFamily: "Space Grotesk, Inter, sans-serif",
      },
    },
    tooltip: {
      trigger: "axis",
      backgroundColor: "#111111",
      borderColor: "#27272a",
      textStyle: {
        color: "#e4e4e7",
        fontSize: 12,
      },
      valueFormatter: (tooltipValue) => Number(tooltipValue).toFixed(2),
    },
    xAxis: {
      type: "value",
      min: variable.universe[0],
      max: variable.universe[1],
      axisLabel: {
        color: "#71717a",
        fontSize: 11,
        formatter: variable.axisFormatter,
      },
      axisLine: {
        lineStyle: {
          color: "#27272a",
        },
      },
      splitLine: {
        lineStyle: {
          color: "#1f1f1f",
        },
      },
    },
    yAxis: {
      type: "value",
      min: 0,
      max: 1,
      interval: 0.5,
      axisLabel: {
        color: "#71717a",
        fontSize: 11,
      },
      axisLine: {
        lineStyle: {
          color: "#27272a",
        },
      },
      splitLine: {
        lineStyle: {
          color: "#1f1f1f",
        },
      },
    },
    series: createMembershipSeries(variable, value, { showSelectedValue }),
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
