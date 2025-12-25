import SummaryCard from "./SummaryCard.tsx";

const Summary = () => {
  return (
      <div className={"card justify-center gap-2"}>
        <span className={"font-body text-3xl"}>Summary</span>
        <div className={"pt-3"}>
          <div className={"flex justify-center gap-2"}>
            <SummaryCard/>
            <SummaryCard/>
          </div>
          <div className={"flex justify-center gap-2 m-2"}>

            <SummaryCard/>
            <SummaryCard/>
          </div>

        </div>
      </div>
  )
}

export default Summary