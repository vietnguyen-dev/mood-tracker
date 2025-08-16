import { useState } from "react";
import { createFileRoute } from "@tanstack/react-router";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";
import type { TooltipItem } from "chart.js";
import { Line } from "react-chartjs-2";
import refresh from "../assets/refresh-removebg-preview.png";

export const Route = createFileRoute("/admin")({
  component: Admin,
});
ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

export const options = {
  responsive: true,
  plugins: {
    legend: {
      position: "top" as const,
    },
    tooltip: {
      callbacks: {
        afterBody: function (context: TooltipItem<"line">[]) {
          let index = context[0].dataIndex;
          let dataSet = context[0].dataset;
          console.log(dataSet.data[index]);
          return "text here";
        },
      },
    },
  },
};

const labels = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "november",
  "december",
];

const data = {
  labels,
  datasets: [
    {
      label: "Mood Ratings",
      data: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10],
      borderColor: "rgb(53, 162, 235)",
      backgroundColor: "rgba(53, 162, 235, 0.5)",
    },
  ],
};

const setTimes = [
  { long: "5 Years", short: "5Yr" },
  { long: "1 Year", short: "1Yr" },
  { long: "6 Month", short: "6M" },
  { long: "1 Month", short: "1M" },
  { long: "1 Week", short: "1W" },
  { long: "1 Day", short: "1D" },
];

function Admin() {
  return (
    <div className="p-6">
      <h3 className="text-2xl font-semibold z-50 mb-3">
        Welcome to your Dashboard
      </h3>
      <div className="md:gap-3 xl:grid xl:grid-cols-[85%_15%]">
        <section className="mb-12 min-w-[50%] md:mr-3">
          <div className="mb-6 flex flex-row gap-2">
            <input type="date" className="input" />
            <input type="date" className="input" />
            <button className="btn btn-primary md:mr-12">
              <img src={refresh} alt="refresh" className="w-6" />
            </button>
            <button
              className="btn btn-secondary w-[90%] max-w-50 xl hidden xl:inline"
              onClick={() => {
                const modal = document.getElementById(
                  "new-mood-modal"
                ) as HTMLDialogElement;
                if (document) {
                  modal.showModal();
                }
              }}
            >
              Add Current Mood
            </button>
          </div>
          <div className="join mb-6">
            {setTimes.map((time) => (
              <button className="btn join-item" key={time.long}>
                <span className="hidden lg:inline">{time.long}</span>
                <span className="lg:hidden">{time.short}</span>
              </button>
            ))}
          </div>
          <Line options={options} data={data} />
          <button
            className="btn btn-secondary w-[90%] mt-3 max-w-50 xl:hidden"
            onClick={() => {
              const modal = document.getElementById(
                "new-mood-modal"
              ) as HTMLDialogElement;
              if (document) {
                modal.showModal();
              }
            }}
          >
            Add Current Mood
          </button>
        </section>
        <section className="">
          <div className="grid grid-cols-2 gap-3 xl:grid-cols-1">
            <div className="card bg-base-100 card-md shadow-sm">
              <div className="card-body">
                <h2 className="card-title text-2xl text-secondary">N%</h2>
                <p>mood increase from previous month</p>
              </div>
            </div>
            <div className="card bg-base-100 card-md shadow-sm">
              <div className="card-body">
                <h2 className="card-title text-2xl text-primary">N%</h2>
                <p>omething increase from previous month</p>
              </div>
            </div>
            <div className="card bg-base-100 card-md shadow-sm ">
              <div className="card-body">
                <h2 className="card-title text-2xl text-secondary">N%</h2>
                <p>mood increase from previous month</p>
              </div>
            </div>
            <div className="card bg-base-100 card-md shadow-sm ">
              <div className="card-body">
                <h2 className="card-title text-2xl text-primary">N%</h2>
                <p>omething increase from previous month</p>
              </div>
            </div>
          </div>
        </section>
        <section className="my-12 flex justify-between flex-wrap">
          <button className="btn btn-primary text-white">
            Generate Report
          </button>
          <div className="flex gap-3">
            <button className="btn btn-secondary" disabled={true}>
              Save
            </button>
            <button className="btn btn-accent" disabled={true}>
              X
            </button>
          </div>
        </section>
      </div>
      <dialog id="new-mood-modal" className="modal">
        <div className="modal-box">
          <form method="dialog">
            <button className="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">
              ✕
            </button>
          </form>
          <h3 className="font-bold text-lg">Add New Mood</h3>
          <form>
            <fieldset className="fieldset">
              <legend className="fieldset-legend">Rate Your Mood</legend>
              <select defaultValue="Pick a Rating" className="select">
                <option disabled={true}>Pick a color</option>
                <option>Crimson</option>
                <option>Amber</option>
                <option>Velvet</option>
              </select>
            </fieldset>
            <fieldset className="fieldset">
              <legend className="fieldset-legend">Special Notes</legend>
              <textarea className="textarea h-24" placeholder="Bio"></textarea>
              <div className="label">Optional</div>
            </fieldset>
          </form>
          <p className="py-4">
            Press ESC key, click on the X, or click outside to close.
          </p>
        </div>
        {/* This form and button create the backdrop that closes the modal when clicked */}
        <form method="dialog" className="modal-backdrop">
          <button>close</button>
        </form>
      </dialog>
    </div>
  );
}
