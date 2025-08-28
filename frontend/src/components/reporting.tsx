import { useState } from "react";

const Reporting = () => {
  const [report, setReport] = useState<string>("");

  const handleGenerateReport = async () => {
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
