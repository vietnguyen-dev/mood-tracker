import { useState } from "react";

const API_KEY = import.meta.env.VITE_API_KEY;
const API_URL = import.meta.env.VITE_API_URL;

interface iReportData {
  data: [];
}

const Reporting: React.FC<iReportData> = ({ data }) => {
  console.log(data);
  const [report, setReport] = useState<string>("dfsdf");

  const handleGenerateReport = async () => {
    const response = await fetch(
      `${API_URL}/api/generate-report?question=write a report about my mood based on the following data`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "x-api-key": API_KEY,
        },
        body: JSON.stringify({
          moodData: data,
        }),
      }
    );
    const report = await response.json();
    console.log(report);
    setReport("Report generated");
  };

  const cancelReport = () => {
    setReport("");
  };

  return (
    <>
      <section className="my-12 flex justify-between flex-wrap">
        <button
          className="btn btn-neutral text-white"
          onClick={handleGenerateReport}
        >
          Generate Report
        </button>
        <div className="flex gap-3">
          <button className="btn btn-primary" disabled={report.length === 0}>
            Save
          </button>
          <button
            className="btn btn-accent"
            disabled={report.length === 0}
            onClick={cancelReport}
          >
            X
          </button>
        </div>
      </section>
      {report.length > 0 && <div className="mt-4">{report}</div>}
    </>
  );
};

export default Reporting;
