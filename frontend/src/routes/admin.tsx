import { useEffect, useState } from "react";
import { useUser } from "@clerk/clerk-react";
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

const API_KEY = import.meta.env.VITE_API_KEY;
const API_URL = import.meta.env.VITE_API_URL;
const headers = {
  "Content-Type": "application/json",
  "x-api-key": API_KEY,
};

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

const setTimes = [
  { long: "5 Years", short: "5Yr" },
  { long: "1 Year", short: "1Yr" },
  { long: "6 Month", short: "6M" },
  { long: "1 Month", short: "1M" },
  { long: "1 Week", short: "1W" },
  { long: "1 Day", short: "1D" },
];

const moodOptions = [
  { label: "1", value: 1 },
  { label: "2", value: 2 },
  { label: "3", value: 3 },
  { label: "4", value: 4 },
  { label: "5", value: 5 },
  { label: "6", value: 6 },
  { label: "7", value: 7 },
  { label: "8", value: 8 },
  { label: "9", value: 9 },
  { label: "10", value: 10 },
];

let sampleData = [
  {
    date: "2025-01-01",
    mood: 1,
    notes: "I'm feeling great!",
  },
  {
    date: "2025-01-02",
    mood: 2,
    notes: "I'm feeling good!",
  },
  {
    date: "2025-01-03",
  },
  {
    date: "2025-01-04",
    mood: 4,
    notes: "I'm feeling bad!",
  },

  {
    date: "2025-01-05",
    mood: 5,
    notes: "I'm feeling terrible!",
  },

  {
    date: "2025-01-06",
    mood: 6,
    notes: "I'm feeling terrible!",
  },

  {
    date: "2025-01-07",
    mood: 7,
    notes: "I'm feeling terrible!",
  },

  {
    date: "2025-01-08",
    mood: 8,
    notes: "I'm feeling terrible!",
  },

  {
    date: "2025-01-09",
    mood: 9,
    notes: "I'm feeling terrible!",
  },
];

interface iData {
  labels: string[] | [];
  datasets: {
    label: string;
    data: number[] | [];
    borderColor: string;
    backgroundColor: string;
  }[];
}

function Admin() {
  const { user } = useUser();
  const [startDate, setStartDate] = useState<string>(
    new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toISOString().split("T")[0] +
      "00:00:00"
  );
  const [endDate, setEndDate] = useState<string>(
    new Date().toISOString().split("T")[0] + "24:00:00"
  );
  const [mood, setMood] = useState<number>(0);
  const [notes, setNotes] = useState<string>("");
  const [timeFrame, setTimeFrame] = useState<string>("");
  const [data, setData] = useState<iData>({
    labels: [],
    datasets: [
      {
        label: "Mood Ratings",
        data: [],
        borderColor: "rgb(51, 60, 77)",
        backgroundColor: "rgb(51, 60, 77)",
      },
    ],
  });
  console.log(user);
  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(
        `${API_URL}/api/moods?start_date=${startDate}&end_date=${endDate}`,
        { headers }
      );
      const data = await response.json();
      const labels = data.map((item: any) => item.date);
      const moodData = data.map((item: any) => item.mood);
      setData({
        labels,
        datasets: [
          {
            label: "Mood Ratings",
            data: moodData,
            borderColor: "rgb(51, 60, 77)",
            backgroundColor: "rgb(51, 60, 77)",
          },
        ],
      });
    };
    fetchData();
  }, []);

  const handleStartDateChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setStartDate(e.target.value);
  };

  const handleEndDateChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEndDate(e.target.value);
  };

  const handleDateRangeChange = (timeFrame: string) => {
    setTimeFrame(timeFrame);
    if (timeFrame === "5Yr") {
      setStartDate(
        new Date(Date.now() - 5 * 365 * 24 * 60 * 60 * 1000)
          .toISOString()
          .split("T")[0]
      );
      setEndDate(new Date().toISOString().split("T")[0]);
    } else if (timeFrame === "1Yr") {
      setStartDate(
        new Date(Date.now() - 365 * 24 * 60 * 60 * 1000)
          .toISOString()
          .split("T")[0]
      );
      setEndDate(new Date().toISOString().split("T")[0]);
    } else if (timeFrame === "6M") {
      setStartDate(
        new Date(Date.now() - 6 * 30 * 24 * 60 * 60 * 1000)
          .toISOString()
          .split("T")[0]
      );
      setEndDate(new Date().toISOString().split("T")[0]);
    } else if (timeFrame === "1M") {
      setStartDate(
        new Date(Date.now() - 30 * 24 * 60 * 60 * 1000)
          .toISOString()
          .split("T")[0]
      );
      setEndDate(new Date().toISOString().split("T")[0]);
    } else if (timeFrame === "1W") {
      setStartDate(
        new Date(Date.now() - 7 * 24 * 60 * 60 * 1000)
          .toISOString()
          .split("T")[0]
      );
      setEndDate(new Date().toISOString().split("T")[0]);
    } else if (timeFrame === "1D") {
      setStartDate(
        new Date(Date.now() - 24 * 60 * 60 * 1000).toISOString().split("T")[0]
      );
      setEndDate(new Date().toISOString().split("T")[0]);
    }
  };

  const handleMoodChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setMood(Number(e.target.value));
  };

  const handleNotesChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setNotes(e.target.value);
  };

  return (
    <div className="p-6">
      <h3 className="text-2xl font-semibold z-50 mb-3">
        Welcome to your Dashboard
      </h3>
      <div className="md:gap-3 xl:grid xl:grid-cols-[85%_15%]">
        <section className="mb-12 min-w-[50%] md:mr-3">
          <div className="mb-6 flex flex-row gap-2">
            <input
              type="date"
              className="input"
              value={startDate}
              onChange={handleStartDateChange}
            />
            <input
              type="date"
              className="input"
              value={endDate}
              onChange={handleEndDateChange}
            />
            <button className="btn btn-neutral md:mr-12">
              <img src={refresh} alt="refresh" className="w-6" />
            </button>
            <button
              className="btn btn-primary w-[90%] max-w-50 xl hidden xl:inline"
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
              <button
                className={`btn join-item ${
                  timeFrame === time.short ? "btn-neutral" : "btn-primary"
                }`}
                key={time.long}
                onClick={() => handleDateRangeChange(time.short)}
              >
                <span className="hidden lg:inline">{time.long}</span>
                <span className="lg:hidden">{time.short}</span>
              </button>
            ))}
          </div>
          <Line options={options} data={data} />
          <button
            className="btn btn-primary w-[90%] mt-3 max-w-50 xl:hidden"
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
                <h2 className="card-title text-2xl text-primary">N%</h2>
                <p>mood increase from previous month</p>
              </div>
            </div>
            <div className="card bg-base-100 card-md shadow-sm">
              <div className="card-body">
                <h2 className="card-title text-2xl text-neutral">N%</h2>
                <p>omething increase from previous month</p>
              </div>
            </div>
            <div className="card bg-base-100 card-md shadow-sm ">
              <div className="card-body">
                <h2 className="card-title text-2xl text-primary">N%</h2>
                <p>mood increase from previous month</p>
              </div>
            </div>
            <div className="card bg-base-100 card-md shadow-sm ">
              <div className="card-body">
                <h2 className="card-title text-2xl text-neutral">N%</h2>
                <p>omething increase from previous month</p>
              </div>
            </div>
          </div>
        </section>
        <section className="my-12 flex justify-between flex-wrap">
          <button className="btn btn-neutral text-white">
            Generate Report
          </button>
          <div className="flex gap-3">
            <button className="btn btn-primary" disabled={true}>
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
              <select
                defaultValue="Pick a Rating"
                className="select"
                onChange={handleMoodChange}
                value={mood}
              >
                {moodOptions.map((option) => (
                  <option key={option.value} value={option.value}>
                    {option.label}
                  </option>
                ))}
              </select>
            </fieldset>
            <fieldset className="fieldset">
              <legend className="fieldset-legend">Special Notes</legend>
              <textarea
                className="textarea h-24"
                placeholder="Notes Here"
                onChange={handleNotesChange}
                value={notes}
              ></textarea>
              <div className="label">Optional</div>
            </fieldset>
            <button className="btn btn-primary mt-3">Add Mood</button>
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
